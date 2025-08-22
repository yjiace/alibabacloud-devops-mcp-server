import { 
  createChangeRequest,
  getChangeRequestAuditItems,
  listChangeRequestExecutions,
  listChangeRequestWorkItems,
  cancelChangeRequest,
  closeChangeRequest,
  CreateChangeRequestRequestSchema,
  GetChangeRequestAuditItemsRequestSchema,
  ListChangeRequestExecutionsRequestSchema,
  ListChangeRequestWorkItemsRequestSchema,
  CancelChangeRequestRequestSchema,
  CloseChangeRequestRequestSchema
} from '../operations/appstack/changeRequests.js';

/**
 * Handle the appstack change requests tool requests
 * 
 * @param request - The tool request
 * @returns The tool response or null if not handled
 */
export async function handleAppStackChangeRequestTools(request: any) {
  switch (request.params.name) {
    case 'create_change_request':
      const createParams = CreateChangeRequestRequestSchema.parse(request.params.arguments);
      const createResult = await createChangeRequest(createParams);
      return {
        content: [{ type: "text", text: JSON.stringify(createResult, null, 2) }],
      };
      
    case 'get_change_request_audit_items':
      const getAuditParams = GetChangeRequestAuditItemsRequestSchema.parse(request.params.arguments);
      const getAuditResult = await getChangeRequestAuditItems(getAuditParams);
      return {
        content: [{ type: "text", text: JSON.stringify(getAuditResult, null, 2) }],
      };
      
    case 'list_change_request_executions':
      const listExecParams = ListChangeRequestExecutionsRequestSchema.parse(request.params.arguments);
      const listExecResult = await listChangeRequestExecutions(listExecParams);
      return {
        content: [{ type: "text", text: JSON.stringify(listExecResult, null, 2) }],
      };
      
    case 'list_change_request_work_items':
      const listWorkParams = ListChangeRequestWorkItemsRequestSchema.parse(request.params.arguments);
      const listWorkResult = await listChangeRequestWorkItems(listWorkParams);
      return {
        content: [{ type: "text", text: JSON.stringify(listWorkResult, null, 2) }],
      };
      
    case 'cancel_change_request':
      const cancelParams = CancelChangeRequestRequestSchema.parse(request.params.arguments);
      const cancelResult = await cancelChangeRequest(cancelParams);
      return {
        content: [{ type: "text", text: JSON.stringify(cancelResult, null, 2) }],
      };
      
    case 'close_change_request':
      const closeParams = CloseChangeRequestRequestSchema.parse(request.params.arguments);
      const closeResult = await closeChangeRequest(closeParams);
      return {
        content: [{ type: "text", text: JSON.stringify(closeResult, null, 2) }],
      };
      
    default:
      return null;
  }
}