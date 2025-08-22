import { z } from 'zod';
import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  SearchAppTemplatesRequestSchema
} from '../operations/appstack/appTemplates.js';

// Export all appstack template tools
export const getAppStackTemplateTools = () => [
  {
    name: 'search_app_templates',
    description: 'Search application templates',
    inputSchema: zodToJsonSchema(SearchAppTemplatesRequestSchema),
  }
];