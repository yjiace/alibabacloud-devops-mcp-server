package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"runtime"
	"strings"
)

var (
	yunxiaoAccessToken string
	Debug              = false // 默认禁用调试输出
)

const (
	DefaultYunxiaoUrl = "https://openapi-rdc.aliyuncs.com"
	Version           = "0.0.1"
)

type Option func(client *YunxiaoClient)

type YunxiaoClient struct {
	Url       string
	Method    string
	Payload   interface{}
	Headers   map[string]string
	Response  *http.Response
	parsedUrl *url.URL
	Query     map[string]string
}

func (g *YunxiaoClient) GetRespBody() ([]byte, error) {
	return io.ReadAll(g.Response.Body)
}

func NewYunxiaoClient(method, urlString string, opts ...Option) *YunxiaoClient {
	// 添加调试输出
	if Debug {
		fmt.Fprintf(os.Stderr, "YunxiaoClient 初始化:\n")
		fmt.Fprintf(os.Stderr, "  输入URL: %s\n", urlString)
	}

	// 检查urlString是否已经包含基础URL
	var fullUrl string
	if strings.HasPrefix(urlString, "http") {
		// 如果urlString已经是完整URL，直接使用
		fullUrl = urlString
		if Debug {
			fmt.Fprintf(os.Stderr, "  URL已包含http前缀，使用原始URL\n")
		}
	} else {
		// 否则添加默认基础URL
		fullUrl = DefaultYunxiaoUrl + urlString
		if Debug {
			fmt.Fprintf(os.Stderr, "  添加基础URL: %s + %s = %s\n", DefaultYunxiaoUrl, urlString, fullUrl)
		}
	}

	parsedUrl, err := url.Parse(fullUrl)
	if err != nil {
		panic(err)
	}

	if Debug {
		fmt.Fprintf(os.Stderr, "  解析后URL: %s\n", parsedUrl.String())
		fmt.Fprintf(os.Stderr, "    Scheme: %s\n", parsedUrl.Scheme)
		fmt.Fprintf(os.Stderr, "    Host: %s\n", parsedUrl.Host)
		fmt.Fprintf(os.Stderr, "    Path: %s\n", parsedUrl.Path)
	}

	client := &YunxiaoClient{
		Method:    method,
		Url:       parsedUrl.String(),
		parsedUrl: parsedUrl,
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (g *YunxiaoClient) Do() (*YunxiaoClient, error) {
	g.Response = nil

	// 打印原始URL信息
	if Debug {
		fmt.Fprintf(os.Stderr, "原始URL: %s\n", g.Url)
		fmt.Fprintf(os.Stderr, "解析后的URL: %v\n", g.parsedUrl)
	}

	// 创建请求时，不为GET请求添加请求体
	var reqBody io.Reader
	if g.Method != "GET" && g.Payload != nil {
		_payload, _ := json.Marshal(g.Payload)
		reqBody = bytes.NewReader(_payload)
		if Debug {
			fmt.Fprintf(os.Stderr, "添加请求体: %s\n", string(_payload))
		}
	} else if Debug {
		fmt.Fprintf(os.Stderr, "GET请求，不添加请求体\n")
	}

	req, err := http.NewRequest(g.Method, g.Url, reqBody)
	if err != nil {
		return nil, NewNetworkError(0, err)
	}

	// 只设置最基本的请求头，与直接HTTP请求相同
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "mcp-yunxiao "+Version+" Go/"+runtime.GOOS+"/"+runtime.GOARCH+"/"+runtime.Version())

	accessToken := GetYunxiaoAccessToken()
	if accessToken == "" {
		return nil, NewNetworkError(0, errors.New("YUNXIAO_ACCESS_TOKEN is not set"))
	}

	req.Header.Set("x-yunxiao-token", accessToken)

	// 应用自定义头部
	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	if Debug {
		fmt.Fprintf(os.Stderr, "===== YunxiaoClient 请求详情 =====\n")
		fmt.Fprintf(os.Stderr, "原始URL构建: 基础URL=%s + 路径=%s\n", DefaultYunxiaoUrl, g.Url[len(DefaultYunxiaoUrl):])

		reqDump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error dumping request: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "%s\n", string(reqDump))
		}
	}

	// 使用默认的HTTP客户端配置
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return g, NewNetworkError(0, err)
	}

	g.Response = resp

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return g, NewNetworkError(resp.StatusCode, err)
	}
	// 重新设置响应体，因为已经读取了它
	resp.Body = io.NopCloser(bytes.NewBuffer(body))

	if Debug {
		fmt.Fprintf(os.Stderr, "响应体长度: %d 字节\n", len(body))
		if len(body) > 0 {
			if len(body) > 1000 {
				fmt.Fprintf(os.Stderr, "响应体内容 (前1000字节): %s...\n", string(body[:1000]))
			} else {
				fmt.Fprintf(os.Stderr, "响应体内容: %s\n", string(body))
			}
		}
	}

	if resp.StatusCode >= 400 {
		return g, NewNetworkError(resp.StatusCode, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body)))
	}

	return g, nil
}

func GetYunxiaoAccessToken() string {
	if yunxiaoAccessToken != "" {
		return yunxiaoAccessToken
	}
	return os.Getenv("YUNXIAO_ACCESS_TOKEN")
}

func WithPayload(payload interface{}) Option {
	return func(client *YunxiaoClient) {
		client.Payload = payload
	}
}

func (g *YunxiaoClient) HandleMCPResult(object any) (*mcp.CallToolResult, error) {
	_, err := g.Do()
	if err != nil {
		return mcp.NewToolResultText(err.Error()), err
	}

	if object == nil {
		return mcp.NewToolResultText("Operation completed successfully"), nil
	}

	body, err := g.GetRespBody()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to read response body: %s", err.Error())),
			NewNetworkError(g.Response.StatusCode, err)
	}

	// 尝试解析 JSON，如果失败则直接返回原始内容
	if err = json.Unmarshal(body, object); err != nil {
		if Debug {
			fmt.Fprintf(os.Stderr, "JSON parsing failed: %v, returning raw body\n", err)
		}
		return mcp.NewToolResultText(string(body)), nil
	}

	result, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to format result response: %s", err.Error())),
			NewNetworkError(g.Response.StatusCode, err)
	}

	return mcp.NewToolResultText(string(result)), nil
}
