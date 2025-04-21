package main

import (
	"devops.aliyun.com/mcp-yunxiao/tools/codeup"
	"devops.aliyun.com/mcp-yunxiao/tools/flow"
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

	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
