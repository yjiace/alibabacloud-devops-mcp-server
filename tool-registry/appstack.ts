import { z } from 'zod';
import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  ListApplicationsRequestSchema,
  GetApplicationRequestSchema,
  CreateApplicationRequestSchema,
  UpdateApplicationRequestSchema
} from '../operations/appstack/applications.js';

// Export all appstack tools
export const getAppStackTools = () => [
  {
    name: 'list_applications',
    description: 'List applications in an organization with pagination',
    inputSchema: zodToJsonSchema(ListApplicationsRequestSchema),
  },
  {
    name: 'get_application',
    description: 'Get application details by name',
    inputSchema: zodToJsonSchema(GetApplicationRequestSchema),
  },
  {
    name: 'create_application',
    description: 'Create a new application',
    inputSchema: zodToJsonSchema(CreateApplicationRequestSchema),
  },
  {
    name: 'update_application',
    description: 'Update an existing application',
    inputSchema: zodToJsonSchema(UpdateApplicationRequestSchema),
  }
];