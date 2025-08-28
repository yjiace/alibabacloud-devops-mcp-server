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
    name: 'create_appstack_change_request',
    description: '[application delivery] Create a change request',
    inputSchema: zodToJsonSchema(CreateChangeRequestRequestSchema),
  },
  {
    name: 'get_change_request_audit_items',
    description: '[application delivery] Get audit items for a change request',
    inputSchema: zodToJsonSchema(GetChangeRequestAuditItemsRequestSchema),
  },
  {
    name: 'list_change_request_executions',
    description: '[application delivery] List change request executions',
    inputSchema: zodToJsonSchema(ListChangeRequestExecutionsRequestSchema),
  },
  {
    name: 'list_change_request_work_items',
    description: '[application delivery] List work items for a change request',
    inputSchema: zodToJsonSchema(ListChangeRequestWorkItemsRequestSchema),
  },
  {
    name: 'cancel_change_request',
    description: '[application delivery] Cancel a change request',
    inputSchema: zodToJsonSchema(CancelChangeRequestRequestSchema),
  },
  {
    name: 'close_change_request',
    description: '[application delivery] Close a change request',
    inputSchema: zodToJsonSchema(CloseChangeRequestRequestSchema),
  }
];