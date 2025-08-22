import { z } from 'zod';
import { yunxiaoRequest, buildUrl } from '../../common/utils.js';

// Define the referenced schemas based on their usage
// Since we can't find the full definitions, we'll create simplified versions
const AppSchema = z.object({
  name: z.string().optional().describe("应用名称"),
}).describe("应用");

const OrchestrationRevisionSchema = z.object({
  author: z.string().optional().describe("作者"),
  commitTime: z.number().optional().describe("提交时间"),
  message: z.string().optional().describe("提交信息"),
  sha: z.string().optional().describe("提交SHA"),
}).describe("编排版本");

const AppOrchestrationSchema = z.object({
  app: AppSchema.optional().describe("应用"),
  creatorId: z.string().optional().describe("创建人ID"),
  description: z.string().optional().describe("描述"),
  format: z.enum(["MANIFEST", "HELM_CHARTS"]).optional().describe("格式"),
  gmtCreate: z.string().optional().describe("创建时间"),
  gmtModified: z.string().optional().describe("修改时间"),
  modifierId: z.string().optional().describe("修改人ID"),
  name: z.string().optional().describe("名称"),
  revision: OrchestrationRevisionSchema.optional().describe("版本"),
  sn: z.string().optional().describe("序列号"),
  storageType: z.enum(["BUILTIN", "EXTERNAL"]).optional().describe("存储类型"),
  suitableResourceTypes: z.array(z.enum(["KUBERNETES", "HOST", "BePending"])).optional().describe("适用资源类型"),
  type: z.string().optional().describe("类型"),
}).describe("应用编排");

const AppBuiltInOrchestrationSchema = AppOrchestrationSchema.extend({
  componentList: z.array(z.object({}).passthrough()).optional().describe("组件列表"),
  groupNameMap: z.record(z.string()).optional().describe("组名映射"),
  labelList: z.array(z.object({}).passthrough()).optional().describe("标签列表"),
  labelPolicy: z.string().optional().describe("标签策略"),
  placeholderList: z.array(z.object({}).passthrough()).optional().describe("占位符列表"),
  syncSourceTemplate: z.object({}).passthrough().optional().describe("同步源模板"),
}).describe("应用内置编排");

const AppExternalCodeupOrchestrationSchema = AppOrchestrationSchema.extend({
  datasource: z.object({}).passthrough().optional().describe("Codeup数据源"),
}).describe("外置Codeup应用编排");

const AppExternalGitlabOrchestrationSchema = AppOrchestrationSchema.extend({
  datasource: z.object({}).passthrough().optional().describe("Gitlab数据源"),
}).describe("外置GitLab应用编排");

// Schema for the GetLatestOrchestration API
export const GetLatestOrchestrationRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  appName: z.string().describe("应用名"),
  envName: z.string().describe("环境名"),
});

export const GetLatestOrchestrationResponseSchema = z.union([
  AppBuiltInOrchestrationSchema,
  AppExternalCodeupOrchestrationSchema,
  AppExternalGitlabOrchestrationSchema
]);

// Schema for the ListAppOrchestration API
export const ListAppOrchestrationRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  appName: z.string().describe("应用名"),
});

export const ListAppOrchestrationResponseSchema = z.object({
  current: z.number().describe("当前页数"),
  pageSize: z.number().describe("每页大小"),
  pages: z.number().describe("总页数"),
  records: z.array(z.union([
    AppBuiltInOrchestrationSchema,
    AppExternalCodeupOrchestrationSchema,
    AppExternalGitlabOrchestrationSchema
  ])).describe("数据列表"),
  total: z.number().describe("总数"),
});

// Schema for the CreateAppOrchestration API
export const CreateAppOrchestrationRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  appName: z.string().describe("应用名"),
  templateId: z.string().describe("编排模板 ID"),
  templateType: z.enum(["PREDEFINED", "ORG"]).describe("编排模板类型"),
});

export const CreateAppOrchestrationResponseSchema = z.union([
  AppBuiltInOrchestrationSchema,
  AppExternalCodeupOrchestrationSchema,
  AppExternalGitlabOrchestrationSchema
]);

// Schema for the DeleteAppOrchestration API
export const DeleteAppOrchestrationRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  appName: z.string().describe("应用名"),
  sn: z.string().describe("应用编排唯一序列号"),
});

export const DeleteAppOrchestrationResponseSchema = z.boolean().describe("调用是否成功");

// Schema for the GetAppOrchestration API
export const GetAppOrchestrationRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  appName: z.string().describe("应用名"),
  sn: z.string().describe("应用编排唯一序列号，未指定 tag 和 sha 时将查找最新版本"),
  tagName: z.string().optional().describe("编排 tag"),
  sha: z.string().optional().describe("编排 commit sha"),
});

export const GetAppOrchestrationResponseSchema = z.union([
  AppBuiltInOrchestrationSchema,
  AppExternalCodeupOrchestrationSchema,
  AppExternalGitlabOrchestrationSchema
]);

// Schema for the UpdateAppOrchestration API
export const UpdateAppOrchestrationRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  appName: z.string().describe("应用名"),
  sn: z.string().describe("应用编排唯一序列号"),
  branchName: z.enum(["master"]).optional().describe("本次提交的编排分支，不填写则使用默认主干"),
  commitMessage: z.string().optional().describe("本次提交的描述信息"),
  description: z.string().optional().describe("编排描述"),
  fromRevisionSha: z.string().optional().describe("本次提交的基线版本 SHA 值"),
  name: z.string().describe("编排名"),
  spec: z.object({}).passthrough().optional().describe("编排规范"),
});

export const UpdateAppOrchestrationResponseSchema = z.union([
  AppBuiltInOrchestrationSchema,
  AppExternalCodeupOrchestrationSchema,
  AppExternalGitlabOrchestrationSchema
]);

export type GetLatestOrchestrationRequest = z.infer<typeof GetLatestOrchestrationRequestSchema>;
export type GetLatestOrchestrationResponse = z.infer<typeof GetLatestOrchestrationResponseSchema>;
export type ListAppOrchestrationRequest = z.infer<typeof ListAppOrchestrationRequestSchema>;
export type ListAppOrchestrationResponse = z.infer<typeof ListAppOrchestrationResponseSchema>;
export type CreateAppOrchestrationRequest = z.infer<typeof CreateAppOrchestrationRequestSchema>;
export type CreateAppOrchestrationResponse = z.infer<typeof CreateAppOrchestrationResponseSchema>;
export type DeleteAppOrchestrationRequest = z.infer<typeof DeleteAppOrchestrationRequestSchema>;
export type DeleteAppOrchestrationResponse = z.infer<typeof DeleteAppOrchestrationResponseSchema>;
export type GetAppOrchestrationRequest = z.infer<typeof GetAppOrchestrationRequestSchema>;
export type GetAppOrchestrationResponse = z.infer<typeof GetAppOrchestrationResponseSchema>;
export type UpdateAppOrchestrationRequest = z.infer<typeof UpdateAppOrchestrationRequestSchema>;
export type UpdateAppOrchestrationResponse = z.infer<typeof UpdateAppOrchestrationResponseSchema>;

/**
 * Get the latest orchestration for an environment
 * 
 * @param params - The request parameters
 * @returns The latest orchestration
 */
export async function getLatestOrchestration(params: GetLatestOrchestrationRequest): Promise<GetLatestOrchestrationResponse> {
  const { organizationId, appName, envName } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/apps/${appName}/envs/${envName}/orchestration:latestAvailable`,
      {
        method: 'GET',
      }
    );
    return GetLatestOrchestrationResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * List application orchestrations
 * 
 * @param params - The request parameters
 * @returns The list of application orchestrations
 */
export async function listAppOrchestration(params: ListAppOrchestrationRequest): Promise<ListAppOrchestrationResponse> {
  const { organizationId, appName } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/apps/${appName}/orchestrations`,
      {
        method: 'GET',
      }
    );
    return ListAppOrchestrationResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Create an application orchestration
 * 
 * @param params - The request parameters
 * @returns The created application orchestration
 */
export async function createAppOrchestration(params: CreateAppOrchestrationRequest): Promise<CreateAppOrchestrationResponse> {
  const { organizationId, appName, ...body } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/apps/${appName}/orchestrations`,
      {
        method: 'POST',
        body: body,
      }
    );
    return CreateAppOrchestrationResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Delete an application orchestration
 * 
 * @param params - The request parameters
 * @returns Whether the deletion was successful
 */
export async function deleteAppOrchestration(params: DeleteAppOrchestrationRequest): Promise<DeleteAppOrchestrationResponse> {
  const { organizationId, appName, sn } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/apps/${appName}/orchestrations/${sn}`,
      {
        method: 'DELETE',
      }
    );
    return DeleteAppOrchestrationResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Get an application orchestration
 * 
 * @param params - The request parameters
 * @returns The application orchestration
 */
export async function getAppOrchestration(params: GetAppOrchestrationRequest): Promise<GetAppOrchestrationResponse> {
  const { organizationId, appName, sn, tagName, sha } = params;
  
  // Build query string properly
  const query: Record<string, string | number> = {};
  if (tagName) query.tagName = tagName;
  if (sha) query.sha = sha;
  
  try {
    const url = buildUrl(`/oapi/v1/appstack/organizations/${organizationId}/apps/${appName}/orchestrations/${sn}`, query);
    
    const response = await yunxiaoRequest(
      url,
      {
        method: 'GET',
      }
    );
    return GetAppOrchestrationResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Update an application orchestration
 * 
 * @param params - The request parameters
 * @returns The updated application orchestration
 */
export async function updateAppOrchestration(params: UpdateAppOrchestrationRequest): Promise<UpdateAppOrchestrationResponse> {
  const { organizationId, appName, sn, ...body } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/apps/${appName}/orchestrations/${sn}`,
      {
        method: 'PUT',
        body: body,
      }
    );
    return UpdateAppOrchestrationResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}