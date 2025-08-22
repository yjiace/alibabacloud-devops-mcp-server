import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  GetLatestOrchestrationRequestSchema,
  ListAppOrchestrationRequestSchema,
  CreateAppOrchestrationRequestSchema,
  DeleteAppOrchestrationRequestSchema,
  GetAppOrchestrationRequestSchema,
  UpdateAppOrchestrationRequestSchema
} from '../operations/appstack/appOrchestrations.js';

// Export all appstack application orchestrations tools
export const getAppStackOrchestrationTools = () => [
  {
    name: 'get_latest_orchestration',
    description: 'Get the latest orchestration for an environment',
    inputSchema: zodToJsonSchema(GetLatestOrchestrationRequestSchema),
  },
  {
    name: 'list_app_orchestration',
    description: 'List application orchestrations',
    inputSchema: zodToJsonSchema(ListAppOrchestrationRequestSchema),
  },
  {
    name: 'create_app_orchestration',
    description: 'Create an application orchestration',
    inputSchema: zodToJsonSchema(CreateAppOrchestrationRequestSchema),
  },
  {
    name: 'delete_app_orchestration',
    description: 'Delete an application orchestration',
    inputSchema: zodToJsonSchema(DeleteAppOrchestrationRequestSchema),
  },
  {
    name: 'get_app_orchestration',
    description: 'Get an application orchestration',
    inputSchema: zodToJsonSchema(GetAppOrchestrationRequestSchema),
  },
  {
    name: 'update_app_orchestration',
    description: 'Update an application orchestration',
    inputSchema: zodToJsonSchema(UpdateAppOrchestrationRequestSchema),
  }
];