package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

var (
	yunxiaoAccessToken string
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
	return ioutil.ReadAll(g.Response.Body)
}

func NewYunxiaoClient(method, urlString string, opts ...Option) *YunxiaoClient {
	urlString = DefaultYunxiaoUrl + urlString
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
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
	_payload, _ := json.Marshal(g.Payload)
	req, err := http.NewRequest(g.Method, g.Url, bytes.NewReader(_payload))
	if err != nil {
		return nil, NewNetworkError(0, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "mcp-yunxiao "+Version+" Go/"+runtime.GOOS+"/"+runtime.GOARCH+"/"+runtime.Version())

	accessToken := GetYunxiaoAccessToken()

	req.Header.Set("x-yunxiao-token", accessToken)

	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return g, NewNetworkError(resp.StatusCode, err)
	}

	g.Response = resp

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

	if err = json.Unmarshal(body, object); err != nil {
		bodyStr := string(body)
		errorMessage := fmt.Sprintf("Failed to parse body response: %v body: %s", err, bodyStr)
		return mcp.NewToolResultError(errorMessage), NewNetworkError(g.Response.StatusCode, errors.New(errorMessage))
	}

	result, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to format result response: %s", err.Error())),
			NewNetworkError(g.Response.StatusCode, err)
	}

	return mcp.NewToolResultText(string(result)), nil
}
