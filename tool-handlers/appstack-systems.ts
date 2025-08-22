import { z } from 'zod';
import { 
  listSystems,
  createSystem,
  updateSystem,
  listAttachedApps,
  attachApps,
  detachApps,
  listSystemMembers,
  createSystemMembers,
  updateSystemMember,
  deleteSystemMember,
  ListSystemsRequestSchema,
  CreateSystemRequestSchema,
  UpdateSystemRequestSchema,
  ListAttachedAppsRequestSchema,
  AttachAppsRequestSchema,
  DetachAppsRequestSchema,
  ListSystemMembersRequestSchema,
  CreateSystemMembersRequestSchema,
  UpdateSystemMemberRequestSchema,
  DeleteSystemMemberRequestSchema
} from '../operations/appstack/systems.js';

/**
 * Handle the appstack system tool requests
 * 
 * @param request - The tool request
 * @returns The tool response or null if not handled
 */
export async function handleAppStackSystemTools(request: any) {
  switch (request.params.name) {
    case 'list_systems':
      const listSystemsParams = ListSystemsRequestSchema.parse(request.params.arguments);
      const listSystemsResult = await listSystems(listSystemsParams);
      return {
        content: [{ type: "text", text: JSON.stringify(listSystemsResult, null, 2) }],
      };
      
    case 'create_system':
      const createSystemParams = CreateSystemRequestSchema.parse(request.params.arguments);
      const createSystemResult = await createSystem(createSystemParams);
      return {
        content: [{ type: "text", text: JSON.stringify(createSystemResult, null, 2) }],
      };
      
    case 'update_system':
      const updateSystemParams = UpdateSystemRequestSchema.parse(request.params.arguments);
      const updateSystemResult = await updateSystem(updateSystemParams);
      return {
        content: [{ type: "text", text: JSON.stringify(updateSystemResult, null, 2) }],
      };
      
    case 'list_attached_apps':
      const listAttachedAppsParams = ListAttachedAppsRequestSchema.parse(request.params.arguments);
      const listAttachedAppsResult = await listAttachedApps(listAttachedAppsParams);
      return {
        content: [{ type: "text", text: JSON.stringify(listAttachedAppsResult, null, 2) }],
      };
      
    case 'attach_apps':
      const attachAppsParams = AttachAppsRequestSchema.parse(request.params.arguments);
      const attachAppsResult = await attachApps(attachAppsParams);
      return {
        content: [{ type: "text", text: JSON.stringify(attachAppsResult, null, 2) }],
      };
      
    case 'detach_apps':
      const detachAppsParams = DetachAppsRequestSchema.parse(request.params.arguments);
      const detachAppsResult = await detachApps(detachAppsParams);
      return {
        content: [{ type: "text", text: JSON.stringify(detachAppsResult, null, 2) }],
      };
      
    case 'list_system_members':
      const listSystemMembersParams = ListSystemMembersRequestSchema.parse(request.params.arguments);
      const listSystemMembersResult = await listSystemMembers(listSystemMembersParams);
      return {
        content: [{ type: "text", text: JSON.stringify(listSystemMembersResult, null, 2) }],
      };
      
    case 'create_system_members':
      const createSystemMembersParams = CreateSystemMembersRequestSchema.parse(request.params.arguments);
      const createSystemMembersResult = await createSystemMembers(createSystemMembersParams);
      return {
        content: [{ type: "text", text: JSON.stringify(createSystemMembersResult, null, 2) }],
      };
      
    case 'update_system_member':
      const updateSystemMemberParams = UpdateSystemMemberRequestSchema.parse(request.params.arguments);
      const updateSystemMemberResult = await updateSystemMember(updateSystemMemberParams);
      return {
        content: [{ type: "text", text: JSON.stringify(updateSystemMemberResult, null, 2) }],
      };
      
    case 'delete_system_member':
      const deleteSystemMemberParams = DeleteSystemMemberRequestSchema.parse(request.params.arguments);
      const deleteSystemMemberResult = await deleteSystemMember(deleteSystemMemberParams);
      return {
        content: [{ type: "text", text: JSON.stringify(deleteSystemMemberResult, null, 2) }],
      };
      
    default:
      return null;
  }
}