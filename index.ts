#!/usr/bin/env node
import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { SSEServerTransport } from "@modelcontextprotocol/sdk/server/sse.js";
import {
    CallToolRequestSchema,
    ListToolsRequestSchema,
} from "@modelcontextprotocol/sdk/types.js";
import { z } from 'zod';
import { zodToJsonSchema } from 'zod-to-json-schema';

import * as branches from './operations/codeup/branches.js';
import * as files from './operations/codeup/files.js';
import * as repositories from './operations/codeup/repositories.js';
import * as changeRequests from './operations/codeup/changeRequests.js';
import * as changeRequestComments from './operations/codeup/changeRequestComments.js';
import * as organization from './operations/organization/organization.js';
import * as members from './operations/organization/members.js';
import * as project from './operations/projex/project.js';
import * as workitem from './operations/projex/workitem.js';
import * as sprint from './operations/projex/sprint.js';
import * as compare from './operations/codeup/compare.js'
import * as pipeline from './operations/flow/pipeline.js'
import * as pipelineJob from './operations/flow/pipelineJob.js'
import * as serviceConnection from './operations/flow/serviceConnection.js'
import * as packageRepositories from './operations/packages/repositories.js'
import * as artifacts from './operations/packages/artifacts.js'
import {
    isYunxiaoError,
    YunxiaoError,
    YunxiaoValidationError
} from "./common/errors.js";
import { VERSION } from "./common/version.js";
import {config} from "dotenv";
import * as types from "./common/types.js";
import { getAllTools } from "./tool-registry/index.js";
import { handleToolRequest } from "./tool-handlers/index.js";

const server = new Server(
    {
        name: "alibabacloud-devops-mcp-server",
        version: VERSION,
    },
    {
        capabilities: {
            tools: {},
        },
    }
);

function formatYunxiaoError(error: YunxiaoError): string {
    let message = `Yunxiao API Error: ${error.message}`;

    if (error instanceof YunxiaoValidationError) {
        message = `Parameter validation failed: ${error.message}`;
        if (error.response) {
            message += `\n errorMessage: ${JSON.stringify(error.response, null, 2)}`;
        }
        // 添加常见参数错误的提示
        if (error.message.includes('name')) {
            message += `\n Suggestion: Please check whether the pipeline name meets the requirements.`;
        }
        if (error.message.includes('content') || error.message.includes('yaml')) {
            message += `\n Suggestion: Please check whether the generated YAML format is correct.`;
        }
    } else {
        // 处理通用的Yunxiao错误
        message = `Yunxiao API error (${error.status}): ${error.message}`;
        if (error.response) {
            const response = error.response as any;
            if (response.errorCode) {
                message += `\n errorCode: ${response.errorCode}`;
            }
            if (response.errorMessage && response.errorMessage !== error.message) {
                message += `\n errorMessage: ${response.errorMessage}`;
            }
            if (response.data && typeof response.data === 'object') {
                message += `\n data: ${JSON.stringify(response.data, null, 2)}`;
            }
        }
        
        // 根据状态码提供通用建议
        switch (error.status) {
            case 400:
                message += `\n Suggestion: Please check whether the request parameters are correct, especially whether all required parameters have been provided.`;
                break;
            case 500:
                message += `\n Suggestion: Internal server error. Please try again later or contact technical support.`;
                break;
            case 502:
            case 503:
            case 504:
                message += `\n Suggestion: The service is temporarily unavailable. Please try again later.`;
                break;
        }
    }

    return message;
}

server.setRequestHandler(ListToolsRequestSchema, async () => {
    return {
        tools: getAllTools(),
    };
});

server.setRequestHandler(CallToolRequestSchema, async (request) => {
    try {
        if (!request.params.arguments) {
            throw new Error("Arguments are required");
        }

        // Delegate to our modular tool handler
        return await handleToolRequest(request);
    } catch (error) {
        if (error instanceof z.ZodError) {
            throw new Error(`Invalid input: ${JSON.stringify(error.errors)}`);
        }
        if (isYunxiaoError(error)) {
            throw new Error(formatYunxiaoError(error));
        }
        throw error;
    }
});

config();

// Check if we should run in SSE mode
const useSSE = process.argv.includes('--sse') || process.env.MCP_TRANSPORT === 'sse';

async function runServer() {
    if (useSSE) {
        // Import express only when needed for SSE mode
        const { default: express } = await import('express');
        const app: any = express();
        const port = process.env.PORT || 3000;
        
        // Store sessions
        const sessions: Record<string, { transport: SSEServerTransport; server: Server }> = {};
        
        // SSE endpoint - handles initial connection
        app.get('/sse', async (req: any, res: any) => {
            // In SSE mode, we can use console.log for debugging since it doesn't interfere with the protocol
            console.log(`New SSE connection from ${req.ip}`);
            
            // Create transport with endpoint for POST messages
            const sseTransport = new SSEServerTransport('/messages', res);
            const sessionId = sseTransport.sessionId;
            
            if (sessionId) {
                sessions[sessionId] = { transport: sseTransport, server };
            }
            
            try {
                await server.connect(sseTransport);
                // In SSE mode, console.error is acceptable for status messages
                console.error(`Yunxiao MCP Server connected via SSE with session ${sessionId}`);
            } catch (error) {
                console.error("Failed to start SSE server:", error);
                res.status(500).send("Server error");
            }
        });
        
        // POST endpoint - handles incoming messages
        app.use(express.json({ limit: '10mb' })); // Add JSON body parser
        app.post('/messages', async (req: any, res: any) => {
            const sessionId = req.query.sessionId as string;
            const session = sessions[sessionId];
            
            if (!session) {
                res.status(404).send("Session not found");
                return;
            }
            
            try {
                await session.transport.handlePostMessage(req, res, req.body);
            } catch (error) {
                console.error("Error handling POST message:", error);
                res.status(500).send("Server error");
            }
        });
        
        // Start server
        const serverInstance: any = app.listen(port, () => {
            console.log(`Yunxiao MCP Server running in SSE mode on port ${port}`);
            console.log(`Connect via SSE at http://localhost:${port}/sse`);
            console.log(`Send messages to http://localhost:${port}/messages?sessionId=<session-id>`);
        });
        
        // Handle graceful shutdown
        process.on('SIGINT', () => {
            console.log('Shutting down SSE server...');
            serverInstance.close(() => {
                console.log('Server closed.');
                process.exit(0);
            });
        });
    } else {
        // Stdio mode (default)
        // In stdio mode, we must avoid console.log/console.error as they interfere with the JSON-RPC protocol
        const transport = new StdioServerTransport();
        await server.connect(transport);
        // Don't output anything to stdout/stderr in stdio mode - only JSON-RPC messages should go through the transport
    }
}

runServer().catch((error) => {
    // Only output error to stderr in SSE mode, not in stdio mode
    if (useSSE) {
        console.error("Fatal error in main():", error);
    }
    process.exit(1);
});