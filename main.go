package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-devops-mcp-server/tools/codeup"
	"github.com/aliyun/alibaba-devops-mcp-server/tools/flow"
	"github.com/aliyun/alibaba-devops-mcp-server/tools/projex"
	"github.com/aliyun/alibaba-devops-mcp-server/utils"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

var (
	Version = utils.Version
)

func newMCPServer() *server.MCPServer {
	return server.NewMCPServer(
		"alibaba-devops-mcp-server",
		Version,
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)
}

func addTools(s *server.MCPServer) {
	s.AddTool(codeup.CreateChangeRequestTool, codeup.CreateChangeRequestFunc)

	s.AddTool(codeup.ListRepositoriesTool, codeup.ListRepositoriesFunc)

	s.AddTool(codeup.GetRepositoryTool, codeup.GetRepositoryFunc)

	s.AddTool(codeup.ListChangeRequestsTool, codeup.ListChangeRequestsFunc)

	s.AddTool(codeup.CreateChangeRequestCommentTool, codeup.CreateChangeRequestCommentFunc)

	s.AddTool(codeup.GetCompareTool, codeup.GetCompareFunc)

	s.AddTool(codeup.ListChangeRequestPatchSetsTool, codeup.ListChangeRequestPatchSetsFunc)

	s.AddTool(codeup.GetChangeRequestTool, codeup.GetChangeRequestFunc)

	s.AddTool(codeup.GetBranchTool, codeup.GetBranchFunc)

	s.AddTool(codeup.ListBranchesTool, codeup.ListBranchesFunc)

	s.AddTool(codeup.CreateBranchTool, codeup.CreateBranchFunc)

	s.AddTool(codeup.DeleteBranchTool, codeup.DeleteBranchFunc)

	s.AddTool(codeup.CreateFileTool, codeup.CreateFileFunc)

	s.AddTool(codeup.DeleteFileTool, codeup.DeleteFileFunc)

	s.AddTool(codeup.ListFilesTool, codeup.ListFilesFunc)

	s.AddTool(codeup.GetFileBlobsTool, codeup.GetFileBlobsFunc)

	s.AddTool(codeup.UpdateFileTool, codeup.UpdateFileFunc)

	s.AddTool(projex.GetWorkItemTool, projex.GetWorkItemFunc)

	s.AddTool(projex.SearchProjectsTool, projex.SearchProjectsFunc)

	s.AddTool(projex.GetProjectTool, projex.GetProjectFunc)

	//s.AddTool(projex.ListSprintsTool, projex.ListSprintsFunc)

	s.AddTool(projex.GetSprintTool, projex.GetSprintFunc)

	//s.AddTool(projex.SearchWorkitemsTool, projex.SearchWorkitemsFunc)

	s.AddTool(flow.ListPipelinesTool, flow.ListPipelinesFunc)

	s.AddTool(flow.GetPipelineTool, flow.GetPipelineFunc)

	s.AddTool(flow.CreatePipelineTool, flow.CreatePipelineFunc)

	s.AddTool(flow.GetPipelineRunTool, flow.GetPipelineRunFunc)

	s.AddTool(flow.ListPipelineRunsTool, flow.ListPipelineRunsFunc)

	s.AddTool(flow.GetLatestPipelineRunTool, flow.GetLatestPipelineRunFunc)
}

func run(transport, addr string) error {
	s := newMCPServer()
	addTools(s)

	switch transport {
	case "stdio":
		if err := server.ServeStdio(s); err != nil {
			if err == context.Canceled {
				return nil
			}
			return err
		}
	case "sse":
		log.Printf("Yunxiao MCP Server (SSE) listening on %s", addr)
		srv := server.NewSSEServer(s, server.WithSSEEndpoint("http://"+addr))
		if err := srv.Start(addr); err != nil {
			if err == context.Canceled {
				return nil
			}
			return fmt.Errorf("服务器错误: %v", err)
		}
	default:
		return fmt.Errorf(
			"invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}
	return nil
}

func main() {
	accessToken := flag.String("token", "", "YUNXIAO_ACCESS_TOKEN")
	apiBase := flag.String("api-base", "", "https://openapi-rdc.aliyuncs.com")
	showVersion := flag.Bool("version", false, "Show version information")
	transport := flag.String("transport", "stdio", "Transport type (stdio or sse)")
	addr := flag.String("sse-address", "localhost:8000", "The host and port to start the sse server on")
	flag.Parse()

	if *showVersion {
		fmt.Printf("alibaba-devops-mcp-server\n")
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	if *accessToken == "" {
		*accessToken = os.Getenv("YUNXIAO_ACCESS_TOKEN")
	}

	if *accessToken == "" {
		log.Fatal("Must provide Yunxiao access token, either through the --token parameter or the YUNXIAO_ACCESS_TOKEN environment variable.")
	}

	if *apiBase != "" {
		utils.DefaultYunxiaoUrl = *apiBase
	}

	utils.SetYunxiaoAccessToken(*accessToken)

	// Start the server
	if err := run(*transport, *addr); err != nil {
		log.Fatalf("Run Server error: %v\n", err)
	}
}
