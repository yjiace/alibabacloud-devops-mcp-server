import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  GetEnvVariableGroupsRequestSchema,
  CreateVariableGroupRequestSchema,
  DeleteVariableGroupRequestSchema,
  GetVariableGroupRequestSchema,
  UpdateVariableGroupRequestSchema,
  GetAppVariableGroupsRequestSchema,
  GetAppVariableGroupsRevisionRequestSchema
} from '../operations/appstack/variableGroups.js';

// Export all appstack variable groups tools
export const getAppStackVariableGroupTools = () => [
  {
    name: 'get_env_variable_groups',
    description: 'Get variable groups for an environment',
    inputSchema: zodToJsonSchema(GetEnvVariableGroupsRequestSchema),
  },
  {
    name: 'create_variable_group',
    description: 'Create a variable group',
    inputSchema: zodToJsonSchema(CreateVariableGroupRequestSchema),
  },
  {
    name: 'delete_variable_group',
    description: 'Delete a variable group',
    inputSchema: zodToJsonSchema(DeleteVariableGroupRequestSchema),
  },
  {
    name: 'get_variable_group',
    description: 'Get a variable group',
    inputSchema: zodToJsonSchema(GetVariableGroupRequestSchema),
  },
  {
    name: 'update_variable_group',
    description: 'Update a variable group',
    inputSchema: zodToJsonSchema(UpdateVariableGroupRequestSchema),
  },
  {
    name: 'get_app_variable_groups',
    description: 'Get variable groups for an application',
    inputSchema: zodToJsonSchema(GetAppVariableGroupsRequestSchema),
  },
  {
    name: 'get_app_variable_groups_revision',
    description: 'Get the revision of variable groups for an application',
    inputSchema: zodToJsonSchema(GetAppVariableGroupsRevisionRequestSchema),
  }
];