import { z } from 'zod';
import { listApplications, ListApplicationsRequestSchema, ListApplicationsResponseSchema } from '../operations/appstack/applications.js';

// Tool schema for ListApplications
export const ListApplicationsToolSchema = z.object({
  name: z.literal('list_applications'),
  description: z.literal('List applications in an organization with pagination'),
  inputSchema: ListApplicationsRequestSchema,
});

/**
 * Handle the list_applications tool request
 * 
 * @param request - The tool request
 * @returns The tool response or null if not handled
 */
export async function handleAppStackTools(request: any) {
  if (request.params.name === 'list_applications') {
    const params = ListApplicationsRequestSchema.parse(request.params.arguments);
    const result = await listApplications(params);
    return {
      content: [{ type: "text", text: JSON.stringify(result, null, 2) }],
    };
  }
  
  return null;
}