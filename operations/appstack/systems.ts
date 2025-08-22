import { z } from 'zod';
import { yunxiaoRequest, buildUrl } from '../../common/utils.js';

// Schema for the ListSystems API
export const ListSystemsRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  current: z.number().default(1).describe("当前页号（从 1 开始，默认取 1）"),
  pageSize: z.number().default(10).describe("分页记录数（默认 10 条）"),
});

export const SystemSchema = z.object({
  creatorId: z.string().optional().describe("系统创建者 ID"),
  description: z.string().optional().describe("系统描述"),
  gmtCreate: z.string().optional().describe("使用本地时区呈现的系统创建时间"),
  name: z.string().optional().describe("系统的唯一名，仅允许包含小写字母、中划线和数字，且开头、结尾均为小写字母或数字"),
  ownerId: z.string().optional(),
});

export const ListSystemsResponseSchema = z.object({
  current: z.number().describe("当前页数"),
  pageSize: z.number().describe("每页大小"),
  pages: z.number().describe("总页数"),
  records: z.array(SystemSchema).describe("数据列表"),
  total: z.number().describe("总数"),
});

// Schema for the CreateSystem API
export const CreateSystemRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  name: z.string().describe("系统名"),
  description: z.string().optional().describe("系统描述"),
  ownerId: z.string().optional().describe("应用 owner ID"),
});

export const CreateSystemResponseSchema = SystemSchema;

// Schema for the UpdateSystem API
export const UpdateSystemRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  description: z.string().optional().describe("系统描述"),
  ownerId: z.string().optional().describe("系统 owner ID"),
});

export const UpdateSystemResponseSchema = SystemSchema;

// Schema for the ListAttachedApps API
export const ListAttachedAppsRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  current: z.number().default(1).describe("当前页号（从 1 开始，默认取 1）"),
  pageSize: z.number().default(10).describe("分页记录数（默认 10 条）"),
});

export const AppWithSourcesVOSchema = z.object({
  // Based on the reference in the swagger, we'll define a basic structure
  // You may need to update this with the actual fields from AppWithSourcesVO
  name: z.string().optional().describe("应用名"),
  description: z.string().optional().describe("应用描述"),
});

export const ListAttachedAppsResponseSchema = z.object({
  current: z.number().describe("当前页数"),
  pageSize: z.number().describe("每页大小"),
  pages: z.number().describe("总页数"),
  records: z.array(AppWithSourcesVOSchema).describe("数据列表"),
  total: z.number().describe("总数"),
});

// Schema for the AttachApps API
export const AttachAppsRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  appNames: z.array(z.string()).describe("待关联的应用名列表"),
});

export const AttachAppsResponseSchema = z.boolean().describe("调用是否成功");

// Schema for the DetachApps API
export const DetachAppsRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  appNames: z.array(z.string()).describe("待解除关联的应用名列表"),
});

export const DetachAppsResponseSchema = z.boolean().describe("调用是否成功");

// Schema for the ListSystemMembers API
export const ListSystemMembersRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  current: z.number().default(1).describe("当前页号（从 1 开始，默认取 1）"),
  pageSize: z.number().default(10).describe("分页记录数（默认 10 条）"),
});

export const MemberVOSchema = z.object({
  // Based on the reference in the swagger, we'll define a basic structure
  // You may need to update this with the actual fields from MemberVO
  id: z.string().optional().describe("成员ID"),
  name: z.string().optional().describe("成员名称"),
  type: z.string().optional().describe("成员类型"),
});

export const ListSystemMembersResponseSchema = z.object({
  current: z.number().describe("当前页数"),
  pageSize: z.number().describe("每页大小"),
  pages: z.number().describe("总页数"),
  records: z.array(MemberVOSchema).describe("数据列表"),
  total: z.number().describe("总数"),
});

// Schema for the CreateSystemMembers API
export const ResourcePlayerRequestSchema = z.object({
  // Based on the reference in the swagger, we'll define a basic structure
  // You may need to update this with the actual fields from ResourcePlayerRequest
  id: z.string().optional().describe("玩家ID"),
  type: z.string().optional().describe("玩家类型"),
});

export const AddMembersRequestSchema = z.object({
  appNames: z.array(z.string()).describe("添加应用列表"),
  playerList: z.array(ResourcePlayerRequestSchema).describe("添加的成员列表"),
  roleNames: z.array(z.string()).describe("添加成员的角色列表"),
});

export const CreateSystemMembersRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  members: AddMembersRequestSchema,
});

export const CreateSystemMembersResponseSchema = z.boolean().describe("调用是否成功");

// Schema for the UpdateSystemMember API
export const UpdateMemberRequestSchema = z.object({
  player: ResourcePlayerRequestSchema,
  roleNames: z.array(z.string()).describe("更新成员的角色列表"),
});

export const UpdateSystemMemberRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  member: UpdateMemberRequestSchema,
});

export const UpdateSystemMemberResponseSchema = z.boolean().describe("调用是否成功");

// Schema for the DeleteSystemMember API
export const DeleteSystemMemberRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统"),
  subjectType: z.string().describe("成员类型"),
  subjectId: z.string().describe("成员id"),
});

export const DeleteSystemMemberResponseSchema = z.boolean().describe("调用是否成功");

export type ListSystemsRequest = z.infer<typeof ListSystemsRequestSchema>;
export type ListSystemsResponse = z.infer<typeof ListSystemsResponseSchema>;
export type CreateSystemRequest = z.infer<typeof CreateSystemRequestSchema>;
export type CreateSystemResponse = z.infer<typeof CreateSystemResponseSchema>;
export type UpdateSystemRequest = z.infer<typeof UpdateSystemRequestSchema>;
export type UpdateSystemResponse = z.infer<typeof UpdateSystemResponseSchema>;
export type ListAttachedAppsRequest = z.infer<typeof ListAttachedAppsRequestSchema>;
export type ListAttachedAppsResponse = z.infer<typeof ListAttachedAppsResponseSchema>;
export type AttachAppsRequest = z.infer<typeof AttachAppsRequestSchema>;
export type AttachAppsResponse = z.infer<typeof AttachAppsResponseSchema>;
export type DetachAppsRequest = z.infer<typeof DetachAppsRequestSchema>;
export type DetachAppsResponse = z.infer<typeof DetachAppsResponseSchema>;
export type ListSystemMembersRequest = z.infer<typeof ListSystemMembersRequestSchema>;
export type ListSystemMembersResponse = z.infer<typeof ListSystemMembersResponseSchema>;
export type CreateSystemMembersRequest = z.infer<typeof CreateSystemMembersRequestSchema>;
export type CreateSystemMembersResponse = z.infer<typeof CreateSystemMembersResponseSchema>;
export type UpdateSystemMemberRequest = z.infer<typeof UpdateSystemMemberRequestSchema>;
export type UpdateSystemMemberResponse = z.infer<typeof UpdateSystemMemberResponseSchema>;
export type DeleteSystemMemberRequest = z.infer<typeof DeleteSystemMemberRequestSchema>;
export type DeleteSystemMemberResponse = z.infer<typeof DeleteSystemMemberResponseSchema>;

/**
 * List systems in an organization with pagination
 * 
 * @param params - The request parameters
 * @returns The list of systems
 */
export async function listSystems(params: ListSystemsRequest): Promise<ListSystemsResponse> {
  const { organizationId, current, pageSize } = params;
  
  try {
    const url = buildUrl(`/oapi/v1/appstack/organizations/${organizationId}/systems`, {
      current,
      pageSize
    });
    
    const response = await yunxiaoRequest(url, {
      method: 'GET',
    });
    return ListSystemsResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Create a new system
 * 
 * @param params - The request parameters
 * @returns The created system details
 */
export async function createSystem(params: CreateSystemRequest): Promise<CreateSystemResponse> {
  const { organizationId, ...body } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems`,
      {
        method: 'POST',
        body: body,
      }
    );
    return CreateSystemResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Update an existing system
 * 
 * @param params - The request parameters
 * @returns The updated system details
 */
export async function updateSystem(params: UpdateSystemRequest): Promise<UpdateSystemResponse> {
  const { organizationId, systemName, ...body } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}`,
      {
        method: 'PUT',
        body: body,
      }
    );
    return UpdateSystemResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * List attached apps for a system
 * 
 * @param params - The request parameters
 * @returns The list of attached apps
 */
export async function listAttachedApps(params: ListAttachedAppsRequest): Promise<ListAttachedAppsResponse> {
  const { organizationId, systemName, current, pageSize } = params;
  
  try {
    const url = buildUrl(`/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/apps`, {
      current,
      pageSize
    });
    
    const response = await yunxiaoRequest(url, {
      method: 'GET',
    });
    return ListAttachedAppsResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Attach apps to a system
 * 
 * @param params - The request parameters
 * @returns Whether the operation was successful
 */
export async function attachApps(params: AttachAppsRequest): Promise<AttachAppsResponse> {
  const { organizationId, systemName, appNames } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/apps:attach`,
      {
        method: 'POST',
        body: appNames,
      }
    );
    return AttachAppsResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Detach apps from a system
 * 
 * @param params - The request parameters
 * @returns Whether the operation was successful
 */
export async function detachApps(params: DetachAppsRequest): Promise<DetachAppsResponse> {
  const { organizationId, systemName, appNames } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/apps:detach`,
      {
        method: 'POST',
        body: appNames,
      }
    );
    return DetachAppsResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * List system members
 * 
 * @param params - The request parameters
 * @returns The list of system members
 */
export async function listSystemMembers(params: ListSystemMembersRequest): Promise<ListSystemMembersResponse> {
  const { organizationId, systemName, current, pageSize } = params;
  
  try {
    const url = buildUrl(`/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/members`, {
      current,
      pageSize
    });
    
    const response = await yunxiaoRequest(url, {
      method: 'GET',
    });
    return ListSystemMembersResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Add members to a system
 * 
 * @param params - The request parameters
 * @returns Whether the operation was successful
 */
export async function createSystemMembers(params: CreateSystemMembersRequest): Promise<CreateSystemMembersResponse> {
  const { organizationId, systemName, members } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/members`,
      {
        method: 'POST',
        body: members,
      }
    );
    return CreateSystemMembersResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Update a system member
 * 
 * @param params - The request parameters
 * @returns Whether the operation was successful
 */
export async function updateSystemMember(params: UpdateSystemMemberRequest): Promise<UpdateSystemMemberResponse> {
  const { organizationId, systemName, member } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/members`,
      {
        method: 'PUT',
        body: member,
      }
    );
    return UpdateSystemMemberResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Delete a system member
 * 
 * @param params - The request parameters
 * @returns Whether the operation was successful
 */
export async function deleteSystemMember(params: DeleteSystemMemberRequest): Promise<DeleteSystemMemberResponse> {
  const { organizationId, systemName, subjectType, subjectId } = params;
  
  try {
    const url = buildUrl(`/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/members`, {
      subjectType,
      subjectId
    });
    
    const response = await yunxiaoRequest(url, {
      method: 'DELETE',
    });
    return DeleteSystemMemberResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}