import { z } from 'zod';
import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  CreateAppTagRequestSchema,
  UpdateAppTagRequestSchema,
  DeleteAppTagRequestSchema,
  SearchAppTagRequestSchema,
  UpdateAppTagBindRequestSchema
} from '../operations/appstack/appTags.js';

// Export all appstack tag tools
export const getAppStackTagTools = () => [
  {
    name: 'create_app_tag',
    description: '[application delivery] Create an application tag',
    inputSchema: zodToJsonSchema(CreateAppTagRequestSchema),
  },
  {
    name: 'update_app_tag',
    description: '[application delivery] Update an application tag',
    inputSchema: zodToJsonSchema(UpdateAppTagRequestSchema),
  },
  {
    name: 'delete_app_tag',
    description: '[application delivery] Delete an application tag',
    inputSchema: zodToJsonSchema(DeleteAppTagRequestSchema),
  },
  {
    name: 'search_app_tags',
    description: '[application delivery] Search application tags',
    inputSchema: zodToJsonSchema(SearchAppTagRequestSchema),
  },
  {
    name: 'update_app_tag_bind',
    description: '[application delivery] Update application tag bindings',
    inputSchema: zodToJsonSchema(UpdateAppTagBindRequestSchema),
  }
];