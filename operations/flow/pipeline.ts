import { z } from "zod";
import * as utils from "../../common/utils.js";

// 定义流水线配置Schema
export const PipelineConfigSourceSchema = z.object({
  data: z.object({
    branch: z.string().nullable().optional().describe("Default branch"),
    cloneDepth: z.number().int().nullable().optional().describe("Clone depth"),
    credentialId: z.number().int().nullable().optional().describe("Credential ID"),
    credentialLabel: z.string().nullable().optional().describe("Credential label"),
    credentialType: z.string().nullable().optional().describe("Credential type"),
    events: z.array(z.string()).nullable().optional().describe("Trigger events"),
    isBranchMode: z.boolean().nullable().optional().describe("Whether branch mode is enabled"),
    isCloneDepth: z.boolean().nullable().optional().describe("Whether clone depth is enabled"),
    isSubmodule: z.boolean().nullable().optional().describe("Whether submodule is enabled"),
    isTrigger: z.boolean().nullable().optional().describe("Whether commit trigger is enabled"),
    label: z.string().nullable().optional().describe("Display name"),
    namespace: z.string().nullable().optional().describe("Namespace"),
    repo: z.string().nullable().optional().describe("Repository URL"),
    serviceConnectionId: z.number().int().nullable().optional().describe("Service connection ID"),
    triggerFilter: z.string().nullable().optional().describe("Trigger filter condition"),
    webhook: z.string().nullable().optional().describe("Webhook URL"),
  }),
  sign: z.string().nullable().optional().describe("Code source identifier"),
  type: z.string().nullable().optional().describe("Code source type"),
});

export const PipelineTagSchema = z.object({
  id: z.number().int().nullable().optional().describe("Tag ID"),
  name: z.string().nullable().optional().describe("Tag name"),
});

export const PipelineConfigSchema = z.object({
  flow: z.string().nullable().optional().describe("Flow configuration"),
  settings: z.string().nullable().optional().describe("Pipeline settings"),
  sources: z.array(PipelineConfigSourceSchema).nullable().optional().describe("Code source configurations"),
});

// 定义流水线详情Schema
export const PipelineDetailSchema = z.object({
  createTime: z.number().int().nullable().optional().describe("Creation time"),
  creatorAccountId: z.string().nullable().optional().describe("Creator account ID"),
  envId: z.number().int().nullable().optional().describe("Environment ID: 0-Daily 1-Pre-release 2-Production"),
  envName: z.string().nullable().optional().describe("Environment name"),
  groupId: z.number().int().nullable().optional().describe("Pipeline group ID"),
  modifierAccountId: z.string().nullable().optional().describe("Last modifier account ID"),
  name: z.string().nullable().optional().describe("Pipeline name"),
  pipelineConfig: PipelineConfigSchema.nullable().optional().describe("Pipeline configuration"),
  tagList: z.array(PipelineTagSchema).nullable().optional().describe("Tag list"),
  updateTime: z.number().int().nullable().optional().describe("Update time"),
});

// 定义流水线列表项Schema
export const PipelineListItemSchema = z.object({
  createAccountId: z.string().nullable().optional().describe("Creator account ID"),
  createTime: z.number().int().nullable().optional().describe("Creation time"),
  pipelineId: z.number().int().nullable().optional().describe("Pipeline ID"),
  pipelineName: z.string().nullable().optional().describe("Pipeline name"),
});

// 类型导出
export type PipelineDetail = z.infer<typeof PipelineDetailSchema>;
export type PipelineListItem = z.infer<typeof PipelineListItemSchema>;

// 获取流水线详情的参数Schema
export const GetPipelineSchema = z.object({
  organizationId: z.string().describe("Organization ID, can be found in the basic information page of the organization admin console"),
  pipelineId: z.string().describe("Pipeline ID"),
});

// 获取流水线列表的参数Schema
export const ListPipelinesSchema = z.object({
  organizationId: z.string().describe("Organization ID, can be found in the basic information page of the organization admin console"),
  createStartTime: z.number().int().optional().describe("Creation start time in milliseconds timestamp format (e.g., 1729178040000). For filtering pipelines created after this time."),
  createEndTime: z.number().int().optional().describe("Creation end time in milliseconds timestamp format (e.g., 1729178040000). For filtering pipelines created before this time."),
  executeStartTime: z.number().int().optional().describe("Execution start time in milliseconds timestamp format (e.g., 1729178040000). For filtering pipelines executed after this time."),
  executeEndTime: z.number().int().optional().describe("Execution end time in milliseconds timestamp format (e.g., 1729178040000). For filtering pipelines executed before this time."),
  pipelineName: z.string().optional().describe("Pipeline name for filtering"),
  statusList: z.string().optional().describe("Pipeline status list, comma separated (SUCCESS,RUNNING,FAIL,CANCELED,WAITING)"),
  perPage: z.number().int().min(1).max(30).default(10).optional().describe("Number of items per page, default 10, max 30"),
  page: z.number().int().min(1).default(1).optional().describe("Page number, default 1"),
});

// 导出参数类型
export type ListPipelinesOptions = z.infer<typeof ListPipelinesSchema>;

/**
 * 获取流水线详情
 * @param organizationId 组织ID
 * @param pipelineId 流水线ID
 * @returns 流水线详情
 */
export async function getPipelineFunc(
  organizationId: string,
  pipelineId: string
): Promise<PipelineDetail> {
  const url = `/oapi/v1/flow/organizations/${organizationId}/pipelines/${pipelineId}`;

  const response = await utils.yunxiaoRequest(url, {
    method: "GET",
  });

  return PipelineDetailSchema.parse(response);
}

/**
 * 获取流水线列表
 * @param organizationId 组织ID
 * @param options 查询选项
 * @returns 流水线列表
 */
export async function listPipelinesFunc(
  organizationId: string,
  options?: Omit<ListPipelinesOptions, 'organizationId'>
): Promise<{
  items: PipelineListItem[],
  pagination: {
    nextPage: number | null,
    page: number,
    perPage: number,
    prevPage: number | null,
    total: number,
    totalPages: number
  }
}> {
  const baseUrl = `/oapi/v1/flow/organizations/${organizationId}/pipelines`;
  
  // 构建查询参数
  const queryParams: Record<string, string | number | undefined> = {};
  
  // 处理时间戳参数
  // 如果传入的是日期字符串或Date对象，自动转换为毫秒时间戳
  if (options?.createStartTime !== undefined) {
    queryParams.createStartTime = utils.convertToTimestamp(options.createStartTime);
  }
  
  if (options?.createEndTime !== undefined) {
    queryParams.createEndTime = utils.convertToTimestamp(options.createEndTime);
  }
  
  if (options?.executeStartTime !== undefined) {
    queryParams.executeStartTime = utils.convertToTimestamp(options.executeStartTime);
  }
  
  if (options?.executeEndTime !== undefined) {
    queryParams.executeEndTime = utils.convertToTimestamp(options.executeEndTime);
  }
  
  if (options?.pipelineName !== undefined) {
    queryParams.pipelineName = options.pipelineName;
  }
  
  if (options?.statusList !== undefined) {
    queryParams.statusList = options.statusList;
  }
  
  if (options?.perPage !== undefined) {
    queryParams.perPage = options.perPage;
  }
  
  if (options?.page !== undefined) {
    queryParams.page = options.page;
  }

  // 使用buildUrl函数构建带有查询参数的URL
  const url = utils.buildUrl(baseUrl, queryParams);
  
  const response = await utils.yunxiaoRequest(url, {
    method: "GET",
  });

  // 处理响应头中的分页信息
  const pagination = {
    nextPage: null as number | null,
    page: 1,
    perPage: 10,
    prevPage: null as number | null,
    total: 0,
    totalPages: 0
  };

  // 如果响应中包含数组，则对每个对象进行解析
  let items: PipelineListItem[] = [];
  if (Array.isArray(response)) {
    items = response.map(item => PipelineListItemSchema.parse(item));
  }

  return {
    items,
    pagination
  };
}

// 定义创建流水线运行的参数Schema
export const CreatePipelineRunSchema = z.object({
  organizationId: z.string().describe("Organization ID, can be found in the basic information page of the organization admin console"),
  pipelineId: z.string().describe("Pipeline ID to run"),
  params: z.string().optional().describe("Pipeline run parameters in JSON string format. Can include: branchModeBranchs(array), envs(object), runningBranchs(object), runningTags(object), runningPipelineArtifacts(object), runningAcrArtifacts(object), runningPackagesArtifacts(object), comment(string), needCreateBranch(boolean), releaseBranch(string)"),
  
  // 添加自然语言相关参数，这些参数将被转换为params参数的内容
  description: z.string().optional().describe("Natural language description of how to run the pipeline, e.g. 'Run pipeline using branch mode with branches main and develop'"),
  branches: z.array(z.string()).optional().describe("Branches to use in branch mode or specific branches for repositories"),
  branchMode: z.boolean().optional().describe("Whether to run in branch mode"),
  releaseBranch: z.string().optional().describe("Specific release branch to use"),
  createReleaseBranch: z.boolean().optional().describe("Whether to create a release branch"),
  environmentVariables: z.record(z.string()).optional().describe("Environment variables for the pipeline run"),
  repositories: z.array(z.object({
    url: z.string().describe("Repository URL"),
    branch: z.string().optional().describe("Branch to use for this repository"),
    tag: z.string().optional().describe("Tag to use for this repository")
  })).optional().describe("Specific repository configurations")
});

// 导出类型
export type CreatePipelineRunOptions = z.infer<typeof CreatePipelineRunSchema>;

/**
 * 运行流水线
 * @param organizationId 组织ID
 * @param pipelineId 流水线ID
 * @param options 运行选项，可以是直接的JSON字符串或者自然语言描述的选项
 * @returns 流水线运行ID
 */
export async function createPipelineRunFunc(
  organizationId: string,
  pipelineId: string,
  options?: Partial<Omit<CreatePipelineRunOptions, 'organizationId' | 'pipelineId'>>
): Promise<number> {
  const url = `/oapi/v1/flow/organizations/${organizationId}/pipelines/${pipelineId}/runs`;
  
  // 如果用户已经提供了格式化的params，直接使用
  if (options?.params) {
    const body = {
      params: options.params
    };

    const response = await utils.yunxiaoRequest(url, {
      method: "POST",
      body: body,
    });

    return Number(response);
  }

  // 否则，基于用户提供的自然语言参数构建params
  const paramsObject: Record<string, any> = {};

  // 处理分支模式相关参数
  if (options?.branchMode && options?.branches && options.branches.length > 0) {
    paramsObject.branchModeBranchs = options.branches;
  }
  
  // 处理Release分支相关参数
  if (options?.createReleaseBranch !== undefined) {
    paramsObject.needCreateBranch = options.createReleaseBranch;
  }
  
  if (options?.releaseBranch) {
    paramsObject.releaseBranch = options.releaseBranch;
  }
  
  // 处理环境变量
  if (options?.environmentVariables && Object.keys(options.environmentVariables).length > 0) {
    paramsObject.envs = options.environmentVariables;
  }
  
  // 处理特定仓库配置
  if (options?.repositories && options.repositories.length > 0) {
    // 初始化runningBranchs和runningTags对象
    const runningBranchs: Record<string, string> = {};
    const runningTags: Record<string, string> = {};
    
    // 填充分支和标签信息
    options.repositories.forEach(repo => {
      if (repo.branch) {
        runningBranchs[repo.url] = repo.branch;
      }
      if (repo.tag) {
        runningTags[repo.url] = repo.tag;
      }
    });
    
    // 只有在有内容时才添加到params对象
    if (Object.keys(runningBranchs).length > 0) {
      paramsObject.runningBranchs = runningBranchs;
    }
    
    if (Object.keys(runningTags).length > 0) {
      paramsObject.runningTags = runningTags;
    }
  }
  
  // 如果有自然语言描述，尝试解析它
  if (options?.description) {
    // 此处可以添加更复杂的自然语言处理逻辑
    // 当前实现是简单的关键词匹配
    const description = options.description.toLowerCase();
    
    // 检测分支模式
    if ((description.includes('branch mode') || description.includes('分支模式')) && 
        !paramsObject.branchModeBranchs && 
        options?.branches?.length) {
      paramsObject.branchModeBranchs = options.branches;
    }
    
    // 检测是否需要创建release分支
    if ((description.includes('create release') || description.includes('创建release')) && 
        paramsObject.needCreateBranch === undefined) {
      paramsObject.needCreateBranch = true;
    }
    
    // 如果提到特定release分支但没有指定
    if ((description.includes('release branch') || description.includes('release分支')) && 
        !paramsObject.releaseBranch && 
        options?.branches?.length) {
      // 假设第一个分支就是release分支
      paramsObject.releaseBranch = options.branches[0];
    }
  }
  
  // 构建请求体
  const body: Record<string, any> = {};
  
  // 只有在有参数时才添加params字段
  if (Object.keys(paramsObject).length > 0) {
    body.params = JSON.stringify(paramsObject);
  }
  
  const response = await utils.yunxiaoRequest(url, {
    method: "POST",
    body: body,
  });
  
  // API返回的是一个数字(流水线运行ID)
  return Number(response);
}

// 定义流水线运行的参数Schema
export const GetLatestPipelineRunSchema = z.object({
  organizationId: z.string().describe("Organization ID, can be found in the basic information page of the organization admin console"),
  pipelineId: z.string().describe("Pipeline ID to get the latest run information"),
});

// 定义流水线运行结果类型
export const PipelineRunActionSchema = z.object({
  data: z.string().optional().describe("Action data"),
  disable: z.boolean().optional().describe("Whether the action is disabled"),
  displayType: z.string().optional().describe("Display type of the action"),
  name: z.string().optional().describe("Action name"),
  order: z.number().int().optional().describe("Order of the action"),
  params: z.record(z.string()).optional().describe("Action parameters"),
  title: z.string().optional().describe("Action title"),
  type: z.string().optional().describe("Action type")
});

export const PipelineRunJobSchema = z.object({
  actions: z.array(PipelineRunActionSchema).optional().describe("Available actions for the job"),
  endTime: z.number().int().optional().describe("End time of the job"),
  id: z.number().int().optional().describe("Job ID"),
  jobSign: z.string().optional().describe("Job unique identifier"),
  result: z.string().optional().describe("Job result data in JSON string format"),
  name: z.string().optional().describe("Job name"),
  params: z.string().optional().describe("Job parameters in JSON string format"),
  startTime: z.number().int().optional().describe("Start time of the job"),
  status: z.string().optional().describe("Job status: FAIL, SUCCESS, RUNNING")
});

export const PipelineStageInfoSchema = z.object({
  endTime: z.number().int().optional().describe("End time of the stage"),
  jobs: z.array(PipelineRunJobSchema).optional().describe("Jobs in this stage"),
  name: z.string().optional().describe("Stage name"),
  startTime: z.number().int().optional().describe("Start time of the stage"),
  status: z.string().optional().describe("Stage status: FAIL, SUCCESS, RUNNING")
});

export const PipelineStageSchema = z.object({
  index: z.string().optional().describe("Stage index"),
  name: z.string().optional().describe("Stage name"),
  stageInfo: PipelineStageInfoSchema.optional().describe("Stage detailed information")
});

export const PipelineRunSourceSchema = z.object({
  data: z.record(z.any()).optional().describe("Source configuration data"),
  name: z.string().optional().describe("Source name"),
  sign: z.string().optional().describe("Source identifier"),
  type: z.string().optional().describe("Source type")
});

export const PipelineRunGlobalParamSchema = z.object({
  encrypted: z.boolean().optional().describe("Whether the parameter is encrypted"),
  key: z.string().optional().describe("Parameter key"),
  value: z.string().optional().describe("Parameter value")
});

export const PipelineRunGroupSchema = z.object({
  id: z.number().int().optional().describe("Group ID"),
  name: z.string().optional().describe("Group name")
});

export const PipelineRunSchema = z.object({
  createTime: z.number().int().optional().describe("Creation time of the run"),
  creatorAccountId: z.string().optional().describe("Creator account ID"),
  globalParams: z.array(PipelineRunGlobalParamSchema).optional().describe("Global parameters"),
  groups: z.array(PipelineRunGroupSchema).optional().describe("Pipeline groups"),
  modifierAccountId: z.string().optional().describe("Last modifier account ID"),
  pipelineId: z.number().int().optional().describe("Pipeline ID"),
  pipelineRunId: z.number().int().optional().describe("Pipeline run ID"),
  sources: z.array(PipelineRunSourceSchema).optional().describe("Code sources used in this run"),
  stageGroup: z.array(z.string()).optional().describe("Stage groups"),
  stages: z.array(PipelineStageSchema).optional().describe("Pipeline stages"),
  status: z.string().optional().describe("Pipeline run status: FAIL, SUCCESS, RUNNING"),
  triggerMode: z.number().int().optional().describe("Trigger mode: 1-Manual, 2-Scheduled, 3-Code commit, 5-Pipeline, 6-Webhook"),
  updateTime: z.number().int().optional().describe("Last update time")
});

// 导出类型
export type PipelineRun = z.infer<typeof PipelineRunSchema>;

/**
 * 获取最近一次流水线运行信息
 * @param organizationId 组织ID
 * @param pipelineId 流水线ID
 * @returns 最近一次流水线运行信息
 */
export async function getLatestPipelineRunFunc(
  organizationId: string,
  pipelineId: string
): Promise<PipelineRun> {
  const url = `/oapi/v1/flow/organizations/${organizationId}/pipelines/${pipelineId}/runs/latestPipelineRun`;
  
  const response = await utils.yunxiaoRequest(url, {
    method: "GET",
  });

  return PipelineRunSchema.parse(response);
}

// 定义获取特定流水线运行实例的参数Schema
export const GetPipelineRunSchema = z.object({
  organizationId: z.string().describe("Organization ID, can be found in the basic information page of the organization admin console"),
  pipelineId: z.string().describe("Pipeline ID"),
  pipelineRunId: z.string().describe("Pipeline run ID to retrieve details for"),
});

// 导出类型
export type GetPipelineRunOptions = z.infer<typeof GetPipelineRunSchema>;

/**
 * 获取特定流水线运行实例
 * @param organizationId 组织ID
 * @param pipelineId 流水线ID
 * @param pipelineRunId 流水线运行ID
 * @returns 流水线运行实例信息
 */
export async function getPipelineRunFunc(
  organizationId: string,
  pipelineId: string,
  pipelineRunId: string
): Promise<PipelineRun> {
  const url = `/oapi/v1/flow/organizations/${organizationId}/pipelines/${pipelineId}/runs/${pipelineRunId}`;
  
  const response = await utils.yunxiaoRequest(url, {
    method: "GET",
  });

  return PipelineRunSchema.parse(response);
}

// 定义流水线运行实例列表项Schema
export const PipelineRunListItemSchema = z.object({
  creatorAccountId: z.string().optional().describe("Creator account ID"),
  endTime: z.number().int().optional().describe("End time of the run"),
  pipelineId: z.number().int().optional().describe("Pipeline ID"),
  pipelineRunId: z.number().int().optional().describe("Pipeline run ID"),
  startTime: z.number().int().optional().describe("Start time of the run"),
  triggerMode: z.number().int().optional().describe("Trigger mode: 1-Manual, 2-Scheduled, 3-Code commit, 5-Pipeline, 6-Webhook")
});

// 定义获取流水线运行实例列表的参数Schema
export const ListPipelineRunsSchema = z.object({
  organizationId: z.string().describe("Organization ID, can be found in the basic information page of the organization admin console"),
  pipelineId: z.string().describe("Pipeline ID to list runs for"),
  perPage: z.number().int().min(1).max(30).default(10).optional().describe("Number of items per page, default 10, max 30"),
  page: z.number().int().min(1).default(1).optional().describe("Page number, default 1"),
  startTime: z.number().int().optional().describe("Execution start time filter in milliseconds timestamp format"),
  endTime: z.number().int().optional().describe("Execution end time filter in milliseconds timestamp format"),
  status: z.string().optional().describe("Run status filter: FAIL, SUCCESS, or RUNNING"),
  triggerMode: z.number().int().optional().describe("Trigger mode filter: 1-Manual, 2-Scheduled, 3-Code commit, 5-Pipeline, 6-Webhook")
});

// 导出类型
export type PipelineRunListItem = z.infer<typeof PipelineRunListItemSchema>;
export type ListPipelineRunsOptions = z.infer<typeof ListPipelineRunsSchema>;

/**
 * 获取流水线运行实例列表
 * @param organizationId 组织ID
 * @param pipelineId 流水线ID
 * @param options 查询选项
 * @returns 流水线运行实例列表和分页信息
 */
export async function listPipelineRunsFunc(
  organizationId: string,
  pipelineId: string,
  options?: Omit<ListPipelineRunsOptions, 'organizationId' | 'pipelineId'>
): Promise<{
  items: PipelineRunListItem[],
  pagination: {
    nextPage: number | null,
    page: number,
    perPage: number,
    prevPage: number | null,
    total: number,
    totalPages: number
  }
}> {
  const baseUrl = `/oapi/v1/flow/organizations/${organizationId}/pipelines/${pipelineId}/runs`;
  
  // 构建查询参数
  const queryParams: Record<string, string | number | undefined> = {};
  
  if (options?.perPage !== undefined) {
    queryParams.perPage = options.perPage;
  }
  
  if (options?.page !== undefined) {
    queryParams.page = options.page;
  }
  
  if (options?.startTime !== undefined) {
    queryParams.startTime = utils.convertToTimestamp(options.startTime);
  }
  
  if (options?.endTime !== undefined) {
    // 注意文档中参数名称是endTme(没有i)，这里进行修正
    queryParams.endTme = utils.convertToTimestamp(options.endTime);
  }
  
  if (options?.status !== undefined) {
    queryParams.status = options.status;
  }
  
  if (options?.triggerMode !== undefined) {
    queryParams.triggerMode = options.triggerMode;
  }

  // 使用buildUrl函数构建带有查询参数的URL
  const url = utils.buildUrl(baseUrl, queryParams);
  
  const response = await utils.yunxiaoRequest(url, {
    method: "GET",
  });

  // 处理响应头中的分页信息
  const pagination = {
    nextPage: null as number | null,
    page: options?.page ?? 1,
    perPage: options?.perPage ?? 10,
    prevPage: null as number | null,
    total: 0,
    totalPages: 0
  };

  // 如果响应中包含数组，则对每个对象进行解析
  let items: PipelineRunListItem[] = [];
  if (Array.isArray(response)) {
    items = response.map(item => PipelineRunListItemSchema.parse(item));
  }

  return {
    items,
    pagination
  };
}


