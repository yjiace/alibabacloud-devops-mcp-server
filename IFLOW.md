# IFLOW.md

This file provides guidance to iFlow Cli when working with code in this repository.

## Common Commands

### Building
- `npm run build` - Compiles TypeScript code to JavaScript in the `dist` directory
- `npm run watch` - Continuously watches and compiles TypeScript files

### Running
- `npm start` - Runs the compiled server from `dist/index.js` in stdio mode
- `npm run start:sse` - Runs the compiled server from `dist/index.js` in SSE mode

### Development
- Use `npm run watch` during development for automatic recompilation
- The server entry point is `index.ts` which exports all functionality as an MCP server
- To run in SSE mode during development: `node dist/index.js --sse`
- To run with specific toolsets: `node dist/index.js --toolsets=code-management,project-management`

## Architecture Overview

This is an MCP (Model Context Protocol) server implementation that provides AI assistants with the ability to interact with Alibaba Cloud DevOps (Yunxiao) platform. 

The server is structured into several modules:

1. **Core Entry Point** (`index.ts`):
   - Initializes the MCP server
   - Registers available tools based on enabled toolsets
   - Handles tool requests and maps them to appropriate functions
   - Supports both stdio and SSE transports

2. **Operations Modules** (in `operations/` directory):
   - `codeup/` - Contains functions for code repository operations (branches, files, repositories, change requests)
   - `flow/` - Contains functions for pipeline operations and service connections
   - `organization/` - Contains functions for organization and member management
   - `packages/` - Contains functions for package/artifact repository operations
   - `projex/` - Contains functions for project and work item management

3. **Common Modules** (in `common/` directory):
   - `types.ts` - Defines all Zod schemas for input validation
   - `errors.ts` - Custom error handling for Yunxiao API responses
   - `version.ts` - Version information for the server
   - `toolsets.ts` - Toolset definitions and configuration
   - `toolsetManager.ts` - Toolset management implementation

The server implements a standard MCP server pattern where:
1. Tools are registered with their schemas in the ListTools handler based on enabled toolsets
2. Actual tool execution happens in the CallTool handler, routed to appropriate toolset handlers
3. Each operation has a dedicated function file that makes API calls to Yunxiao
4. All inputs are validated using Zod schemas before processing

The server exposes dozens of tools covering:
- Code repository management (branches, files, repositories)
- Code review operations (change requests, comments)
- Project management (projects, work items, work item types)
- Pipeline management (pipelines, runs, jobs)
- Package repository management (artifacts, repositories)
- Organization management (members, departments, roles)
- Service connections management

## Toolsets

The server now supports toolsets, allowing you to enable only the tools you need. This can reduce the number of tools presented to the AI assistant and improve performance.

Available toolsets:
- `code-management`: Code repository management tools (includes commit management tools)
- `organization-management`: Organization management tools
- `project-management`: Project management tools (includes effort management tools)
- `pipeline-management`: Pipeline management tools (includes service connections, resource member, and VM deploy order tools)
- `packages-management`: Package repository management tools
- `application-delivery`: Application delivery tools

To use toolsets, you can specify them via command line arguments or environment variables:

1. Via command line argument:
```bash
npm start -- --toolsets=code-management,project-management
```

2. Via environment variable:
```bash
MCP_TOOLSETS=code-management,project-management npm start
```

If no toolsets are specified, all tools will be enabled by default.

## SSE Mode

The server can run in SSE (Server-Sent Events) mode, which allows it to be accessed over HTTP instead of stdio. This is useful when deploying the server as a remote service.

To run in SSE mode:
1. Use `npm run start:sse` or `node dist/index.js --sse`
2. The server will start an HTTP server on port 3000 (configurable with PORT environment variable)
3. Clients can connect via SSE at `http://localhost:3000/sse`
4. Messages are sent to `http://localhost:3000/messages?sessionId=<session-id>`

In SSE mode, the server maintains sessions for each connected client, allowing for proper request/response correlation.