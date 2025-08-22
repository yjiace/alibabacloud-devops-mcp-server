import { z } from 'zod';
import { yunxiaoRequest, buildUrl } from '../../common/utils.js';

// Schema for Label
export const LabelSchema = z.object({
  // Based on the reference in the swagger, we'll define a basic structure
  // You may need to update this with the actual fields from Label
  key: z.string().optional().describe("标签键"),
  value: z.string().optional().describe("标签值"),
});

// Schema for VariableGroup
export const VariableGroupSchema = z.object({
  // Based on the reference in the swagger, we'll define a basic structure
  // You may need to update this with the actual fields from VariableGroup
  name: z.string().optional().describe("变量组名称"),
  variables: z.record(z.string()).optional().describe("变量映射"),
});

// Schema for FlowV1Pipeline
export const FlowV1PipelineSchema = z.object({
  // Based on the reference in the swagger, we'll define a basic structure
  // You may need to update this with the actual fields from FlowV1Pipeline
  id: z.number().optional().describe("流水线ID"),
  name: z.string().optional().describe("流水线名称"),
});

// Schema for FlowV2Pipeline
export const FlowV2PipelineSchema = z.object({
  // Based on the reference in the swagger, we'll define a basic structure
  // You may need to update this with the actual fields from FlowV2Pipeline
  id: z.number().optional().describe("流水线ID"),
  name: z.string().optional().describe("流水线名称"),
});

// Schema for ReleaseStage
export const ReleaseStageSchema = z.object({
  appName: z.string().optional().describe("应用名"),
  labels: z.array(LabelSchema).describe("标签列表"),
  name: z.string().optional().describe("名称"),
  order: z.string().optional().describe("阶段顺序"),
  pipeline: z.union([FlowV1PipelineSchema, FlowV2PipelineSchema]).optional(),
  releaseWorkflowSn: z.string().optional().describe("所属的流程sn"),
  sn: z.string().optional().describe("唯一序列号"),
  variableGroups: z.array(VariableGroupSchema).describe("变量组列表"),
});

// Schema for ReleaseWorkflow
export const ReleaseWorkflowSchema = z.object({
  appName: z.string().optional().describe("应用名"),
  name: z.string().optional().describe("名称"),
  note: z.string().optional().describe("公告"),
  order: z.string().optional().describe("排序"),
  releaseStages: z.array(ReleaseStageSchema).describe("研发阶段列表"),
  sn: z.string().optional().describe("唯一标识sn"),
  type: z.enum(["CR", "APP_RELEASE"]).optional().describe("研发流程类型"),
});

// Schema for ListSystemAllReleaseWorkflows API
export const ListSystemAllReleaseWorkflowsRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
});

export const ListSystemAllReleaseWorkflowsResponseSchema = z.array(ReleaseWorkflowSchema);

// Schema for CreateSystemReleaseWorkflows API
export const CreateReleaseWorkflowRequestSchema = z.object({
  name: z.string(),
  note: z.string().optional(),
  templateSn: z.string().optional(),
});

export const CreateSystemReleaseWorkflowsRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  workflow: CreateReleaseWorkflowRequestSchema,
});

export const CreateSystemReleaseWorkflowsResponseSchema = ReleaseWorkflowSchema;

// Schema for UpdateSystemReleaseStage API
export const UpdateReleaseStageRequestSchema = z.object({
  commitMessage: z.string().optional().describe("版本commit信息"),
  labelList: z.array(LabelSchema).describe("标签列表"),
  name: z.string().optional().describe("名称"),
  order: z.string().optional().describe("排序"),
  variableGroups: z.array(VariableGroupSchema).describe("变量组列表"),
});

export const UpdateSystemReleaseStageRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  releaseWorkflowSn: z.string().describe("发布流程SN"),
  releaseStageSn: z.string().describe("发布阶段SN"),
  stage: UpdateReleaseStageRequestSchema,
});

export const UpdateSystemReleaseStageResponseSchema = ReleaseStageSchema;

// Schema for ExecuteReleaseStage API
export const ExecuteReleaseStagePipelineNgRequestSchema = z.object({
  appReleaseSn: z.string().optional().describe("发布sn"),
  params: z.record(z.any()).optional().describe("运行流水线参数"),
});

export const ExecuteReleaseStageRequestSchema = z.object({
  organizationId: z.string().describe("组织ID"),
  systemName: z.string().describe("系统名"),
  releaseWorkflowSn: z.string().describe("发布流程SN"),
  releaseStageSn: z.string().describe("发布阶段SN"),
  execution: ExecuteReleaseStagePipelineNgRequestSchema,
});

export const ExecutePipelineResultSchema = z.object({
  object: z.number().describe("流水线对象"),
  pipelineId: z.number().describe("流水线ID"),
  pipelineRunId: z.number().describe("流水线运行ID"),
});

export const ExecuteReleaseStageResponseSchema = ExecutePipelineResultSchema;

export type ListSystemAllReleaseWorkflowsRequest = z.infer<typeof ListSystemAllReleaseWorkflowsRequestSchema>;
export type ListSystemAllReleaseWorkflowsResponse = z.infer<typeof ListSystemAllReleaseWorkflowsResponseSchema>;
export type CreateSystemReleaseWorkflowsRequest = z.infer<typeof CreateSystemReleaseWorkflowsRequestSchema>;
export type CreateSystemReleaseWorkflowsResponse = z.infer<typeof CreateSystemReleaseWorkflowsResponseSchema>;
export type UpdateSystemReleaseStageRequest = z.infer<typeof UpdateSystemReleaseStageRequestSchema>;
export type UpdateSystemReleaseStageResponse = z.infer<typeof UpdateSystemReleaseStageResponseSchema>;
export type ExecuteReleaseStageRequest = z.infer<typeof ExecuteReleaseStageRequestSchema>;
export type ExecuteReleaseStageResponse = z.infer<typeof ExecuteReleaseStageResponseSchema>;

/**
 * List all release workflows for a system
 * 
 * @param params - The request parameters
 * @returns The list of release workflows
 */
export async function listSystemAllReleaseWorkflows(params: ListSystemAllReleaseWorkflowsRequest): Promise<ListSystemAllReleaseWorkflowsResponse> {
  const { organizationId, systemName } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/releaseWorkflows`,
      {
        method: 'GET',
      }
    );
    return ListSystemAllReleaseWorkflowsResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Create a release workflow for a system
 * 
 * @param params - The request parameters
 * @returns The created release workflow
 */
export async function createSystemReleaseWorkflows(params: CreateSystemReleaseWorkflowsRequest): Promise<CreateSystemReleaseWorkflowsResponse> {
  const { organizationId, systemName, workflow } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/releaseWorkflows`,
      {
        method: 'POST',
        body: workflow,
      }
    );
    return CreateSystemReleaseWorkflowsResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Update a release stage
 * 
 * @param params - The request parameters
 * @returns The updated release stage
 */
export async function updateSystemReleaseStage(params: UpdateSystemReleaseStageRequest): Promise<UpdateSystemReleaseStageResponse> {
  const { organizationId, systemName, releaseWorkflowSn, releaseStageSn, stage } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/releaseWorkflows/${releaseWorkflowSn}/releaseStages/${releaseStageSn}`,
      {
        method: 'PUT',
        body: stage,
      }
    );
    return UpdateSystemReleaseStageResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}

/**
 * Execute a release stage
 * 
 * @param params - The request parameters
 * @returns The execution result
 */
export async function executeReleaseStage(params: ExecuteReleaseStageRequest): Promise<ExecuteReleaseStageResponse> {
  const { organizationId, systemName, releaseWorkflowSn, releaseStageSn, execution } = params;
  
  try {
    const response = await yunxiaoRequest(
      `/oapi/v1/appstack/organizations/${organizationId}/systems/${systemName}/releaseWorkflows/${releaseWorkflowSn}/releaseStages/${releaseStageSn}:execute`,
      {
        method: 'POST',
        body: execution,
      }
    );
    return ExecuteReleaseStageResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}