import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  CreateChangeRequestRequestSchema,
  GetChangeRequestAuditItemsRequestSchema,
  ListChangeRequestExecutionsRequestSchema,
  ListChangeRequestWorkItemsRequestSchema,
  CancelChangeRequestRequestSchema,
  CloseChangeRequestRequestSchema
} from '../operations/appstack/changeRequests.js';

// Export all appstack change requests tools
export const getAppStackChangeRequestTools = () => [
  {
    name: 'create_change_request',
    description: 'Create a change request',
    inputSchema: zodToJsonSchema(CreateChangeRequestRequestSchema),
  },
  {
    name: 'get_change_request_audit_items',
    description: 'Get audit items for a change request',
    inputSchema: zodToJsonSchema(GetChangeRequestAuditItemsRequestSchema),
  },
  {
    name: 'list_change_request_executions',
    description: 'List change request executions',
    inputSchema: zodToJsonSchema(ListChangeRequestExecutionsRequestSchema),
  },
  {
    name: 'list_change_request_work_items',
    description: 'List work items for a change request',
    inputSchema: zodToJsonSchema(ListChangeRequestWorkItemsRequestSchema),
  },
  {
    name: 'cancel_change_request',
    description: 'Cancel a change request',
    inputSchema: zodToJsonSchema(CancelChangeRequestRequestSchema),
  },
  {
    name: 'close_change_request',
    description: 'Close a change request',
    inputSchema: zodToJsonSchema(CloseChangeRequestRequestSchema),
  }
];