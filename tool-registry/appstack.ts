import { z } from 'zod';
import { zodToJsonSchema } from 'zod-to-json-schema';
import { ListApplicationsRequestSchema } from '../operations/appstack/applications.js';

// Export all appstack tools
export const getAppStackTools = () => [
  {
    name: 'list_applications',
    description: 'List applications in an organization with pagination',
    inputSchema: zodToJsonSchema(ListApplicationsRequestSchema),
  }
];