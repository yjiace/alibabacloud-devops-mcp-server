package main

import (
	"devops.aliyun.com/mcp-yunxiao/tools/codeup"
	"devops.aliyun.com/mcp-yunxiao/tools/projex"
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Yunxiao mcp server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	s.AddTool(codeup.CreateChangeRequestTool, codeup.CreateChangeRequestFunc)

	s.AddTool(codeup.ListRepositoriesTool, codeup.ListRepositoriesFunc)

	s.AddTool(codeup.GetRepositoryTool, codeup.GetRepositoryFunc)

	s.AddTool(codeup.ListChangeRequestsTool, codeup.ListChangeRequestsFunc)

	s.AddTool(codeup.CreateChangeRequestCommentTool, codeup.CreateChangeRequestCommentFunc)

	s.AddTool(codeup.GetCompareTool, codeup.GetCompareFunc)

	s.AddTool(codeup.ListChangeRequestPatchSetsTool, codeup.ListChangeRequestPatchSetsFunc)

	s.AddTool(codeup.GetChangeRequestTool, codeup.GetChangeRequestFunc)

	s.AddTool(projex.GetWorkItemTool, projex.GetWorkItemFunc)

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
