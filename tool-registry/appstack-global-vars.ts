import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  CreateGlobalVarRequestSchema,
  DeleteGlobalVarRequestSchema,
  GetGlobalVarRequestSchema,
  UpdateGlobalVarRequestSchema,
  ListGlobalVarsRequestSchema
} from '../operations/appstack/globalVars.js';

// Export all appstack global variables tools
export const getAppStackGlobalVarTools = () => [
  {
    name: 'create_global_var',
    description: 'Create a global variable group',
    inputSchema: zodToJsonSchema(CreateGlobalVarRequestSchema),
  },
  {
    name: 'delete_global_var',
    description: 'Delete a global variable group',
    inputSchema: zodToJsonSchema(DeleteGlobalVarRequestSchema),
  },
  {
    name: 'get_global_var',
    description: 'Get a global variable group',
    inputSchema: zodToJsonSchema(GetGlobalVarRequestSchema),
  },
  {
    name: 'update_global_var',
    description: 'Update a global variable group',
    inputSchema: zodToJsonSchema(UpdateGlobalVarRequestSchema),
  },
  {
    name: 'list_global_vars',
    description: 'List global variable groups',
    inputSchema: zodToJsonSchema(ListGlobalVarsRequestSchema),
  }
];