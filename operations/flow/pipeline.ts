import * as utils from "../../common/utils.js";
import {
  PipelineDetailSchema,
  PipelineDetail,
  ListPipelinesOptions,
  PipelineListItemSchema,
  PipelineListItem,
  CreatePipelineOptions,
  CreatePipelineFromDescriptionOptions,
  CreatePipelineRunOptions,
  PipelineRunSchema,
  PipelineRun,
  PipelineRunListItemSchema,
  PipelineRunListItem,
  ListPipelineRunsOptions
} from "../../common/types.js";
import { parseUserDescription, generatePipelineName } from "../../common/nlpProcessor.js";
import { TemplateVariables } from "../../common/pipelineTemplates.js";
import { generateModularPipeline } from "../../common/modularTemplates.js";
import { getCurrentOrganizationInfoFunc } from "../organization/organization.js";
import { listRepositoriesFunc } from "../codeup/repositories.js";
import { listServiceConnectionsFunc } from "./serviceConnection.js";
import { listHostGroupsFunc } from "./hostGroup.js";

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

  const url = utils.buildUrl(baseUrl, queryParams);
  
  const response = await utils.yunxiaoRequest(url, {
    method: "GET",
  });

  const pagination = {
    nextPage: null as number | null,
    page: 1,
    perPage: 10,
    prevPage: null as number | null,
    total: 0,
    totalPages: 0
  };

  let items: PipelineListItem[] = [];
  if (Array.isArray(response)) {
    items = response.map(item => PipelineListItemSchema.parse(item));
  }

  return {
    items,
    pagination
  };
}

/**
 * 智能查询流水线列表，能够解析自然语言中的时间表达
 * @param organizationId 组织ID
 * @param timeReference 自然语言时间引用，如"今天"、"本周"、"上个月"
 * @param options 其他查询选项
 * @returns 流水线列表
 */
export async function smartListPipelinesFunc(
  organizationId: string,
  timeReference?: string,
  options?: Omit<ListPipelinesOptions, 'organizationId' | 'executeStartTime' | 'executeEndTime'>
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
  // 解析时间引用获取开始和结束时间戳
  const { startTime, endTime } = utils.parseDateReference(timeReference);
  
  // 合并选项
  const fullOptions: Omit<ListPipelinesOptions, 'organizationId'> = {
    ...options,
    executeStartTime: startTime,
    executeEndTime: endTime
  };
  
  return listPipelinesFunc(organizationId, fullOptions);
}

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

  const body: Record<string, any> = {};

  if (Object.keys(paramsObject).length > 0) {
    body.params = JSON.stringify(paramsObject);
  }
  
  const response = await utils.yunxiaoRequest(url, {
    method: "POST",
    body: body,
  });

  return Number(response);
}

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
    queryParams.endTme = utils.convertToTimestamp(options.endTime);
  }
  
  if (options?.status !== undefined) {
    queryParams.status = options.status;
  }
  
  if (options?.triggerMode !== undefined) {
    queryParams.triggerMode = options.triggerMode;
  }

  const url = utils.buildUrl(baseUrl, queryParams);
  
  const response = await utils.yunxiaoRequest(url, {
    method: "GET",
  });

  const pagination = {
    nextPage: null as number | null,
    page: options?.page ?? 1,
    perPage: options?.perPage ?? 10,
    prevPage: null as number | null,
    total: 0,
    totalPages: 0
  };

  let items: PipelineRunListItem[] = [];
  if (Array.isArray(response)) {
    items = response.map(item => PipelineRunListItemSchema.parse(item));
  }

  return {
    items,
    pagination
  };
}

/**
 * 创建流水线
 * @param organizationId 组织ID
 * @param name 流水线名称
 * @param content 流水线YAML描述
 * @returns 流水线ID
 */
export async function createPipelineFunc(
  organizationId: string,
  name: string,
  content: string
): Promise<number> {
  const url = `/oapi/v1/flow/organizations/${organizationId}/pipelines`;
  
  const body = {
    name: name,
    content: content
  };

  const response = await utils.yunxiaoRequest(url, {
    method: "POST",
    body: body,
  });

  return Number(response);
}

/**
 * 基于自然语言描述智能创建流水线
 * @param organizationId 组织ID
 * @param description 自然语言描述流水线需求
 * @param options 可选的额外配置
 * @returns 创建结果，包含流水线ID、生成的YAML
 */
export async function createPipelineFromDescriptionFunc(
  organizationId: string,
  description: string,
  options?: Partial<Omit<CreatePipelineFromDescriptionOptions, 'organizationId' | 'description'>>
): Promise<{
  pipelineId: number;
  generatedYaml: string;
}> {
  // 解析用户描述
  const parsedInfo = parseUserDescription(description);
  
  // 获取当前代码库上下文信息（如果用户没有明确指定的话）
  let repoContext: { repoUrl?: string; branch?: string; serviceName?: string; } | null = null;
  
  // 检查用户是否已经在描述中或参数中提供了代码库信息
  const hasRepoInfo = parsedInfo.variables.repoUrl || options?.repoUrl;
  const hasServiceName = parsedInfo.variables.serviceName || options?.name;
  const hasBranch = parsedInfo.variables.branch || options?.branch;
  
  // 如果用户没有提供代码库基础信息，尝试自动获取
  if (!hasRepoInfo) {
    try {
      repoContext = await getCurrentRepositoryContext(organizationId);
    } catch (error) {
      console.error('无法获取当前代码库上下文，使用默认值:', error);
    }
  }
  
  // 获取默认服务连接ID（如果用户没有明确指定）
  let defaultServiceConnectionId: string | null = null;
  const hasServiceConnectionId = parsedInfo.variables.serviceConnectionId || options?.serviceConnectionId;
  if (!hasServiceConnectionId) {
    defaultServiceConnectionId = await getDefaultServiceConnectionId(organizationId);
  }
  
  // 获取默认Packages服务连接ID（如果用户没有明确指定且需要packages上传）
  let defaultPackagesServiceConnectionId: string | null = null;
  const hasPackagesServiceConnectionId = options?.packagesServiceConnection;
  const needsPackagesUpload = !options?.uploadType || options?.uploadType === 'packages';
  if (!hasPackagesServiceConnectionId && needsPackagesUpload) {
    defaultPackagesServiceConnectionId = await getDefaultPackagesServiceConnectionId(organizationId);
  }
  
  // 获取默认主机组ID（如果用户没有明确指定且模板包含VM部署）
  let defaultMachineGroupId: string | null = null;
  const hasMachineGroupId = parsedInfo.variables.machineGroupId || options?.machineGroupId;
  const needsVMDeploy = parsedInfo.deployTargets.some(target => 
    ['vm', 'host', '主机'].includes(target.toLowerCase())
  );
  if (!hasMachineGroupId && needsVMDeploy) {
    defaultMachineGroupId = await getDefaultHostGroupId(organizationId);
  }
  
  // 准备模块化流水线生成的变量
  const finalVariables: TemplateVariables = {
    // 使用解析出的变量作为基础
    ...parsedInfo.variables,
    
    // 使用代码库上下文作为智能默认值（只在用户未明确指定时使用）
    ...(repoContext && !hasRepoInfo && { repoUrl: repoContext.repoUrl }),
    ...(repoContext && !hasBranch && { branch: repoContext.branch }),
    ...(repoContext && !hasServiceName && { serviceName: repoContext.serviceName }),
    
    // 使用获取到的默认服务连接ID
    ...(defaultServiceConnectionId && !hasServiceConnectionId && { serviceConnectionId: defaultServiceConnectionId }),
    
    // 使用获取到的默认Packages服务连接ID
    ...(defaultPackagesServiceConnectionId && !hasPackagesServiceConnectionId && { packagesServiceConnection: defaultPackagesServiceConnectionId }),
    
    // 使用获取到的默认主机组ID
    ...(defaultMachineGroupId && !hasMachineGroupId && { machineGroupId: defaultMachineGroupId }),
    
    // 基础配置覆盖（用户明确指定的值优先级最高）
    ...(options?.repoUrl && { repoUrl: options.repoUrl }),
    ...(options?.branch && { branch: options.branch }),
    ...(options?.serviceConnectionId && { serviceConnectionId: options.serviceConnectionId }),
    ...(options?.packagesServiceConnection && { packagesServiceConnection: options.packagesServiceConnection }),
    ...(options?.machineGroupId && { machineGroupId: options.machineGroupId }),
    ...(options?.namespace && { namespace: options.namespace }),
    ...(options?.dockerImage && { dockerImage: options.dockerImage }),
    
    // 版本相关覆盖
    ...(options?.jdkVersion && { jdkVersion: options.jdkVersion }),
    ...(options?.mavenVersion && { mavenVersion: options.mavenVersion }),
    ...(options?.nodeVersion && { nodeVersion: options.nodeVersion }),
    ...(options?.pythonVersion && { pythonVersion: options.pythonVersion }),
    ...(options?.goVersion && { goVersion: options.goVersion }),
    
    // 构建物上传相关覆盖
    ...(options?.uploadType && { uploadType: options.uploadType }),
    ...(options?.artifactName && { artifactName: options.artifactName }),
    ...(options?.artifactVersion && { artifactVersion: options.artifactVersion }),
    ...(options?.packagesRepoId && { packagesRepoId: options.packagesRepoId }),
    ...(options?.includePathInArtifact !== undefined && { includePathInArtifact: options.includePathInArtifact }),
    
    // 部署相关覆盖
    ...(options?.executeUser && { executeUser: options.executeUser }),
    ...(options?.artifactDownloadPath && { artifactDownloadPath: options.artifactDownloadPath }),
    
    // 命令覆盖
    ...(options?.buildCommand && { buildCommand: options.buildCommand }),
    ...(options?.testCommand && { testCommand: options.testCommand }),
    ...(options?.deployCommand && { deployCommand: options.deployCommand }),
  };
  
  // 如果没有明确的serviceName但有repoUrl，从repoUrl解析serviceName
  if (!finalVariables.serviceName && finalVariables.repoUrl) {
    let match;
    // 处理git@格式：git@codeup.aliyun.com:org/repo.git
    if (finalVariables.repoUrl.includes('@') && finalVariables.repoUrl.includes(':')) {
      match = finalVariables.repoUrl.match(/:([^\/]+)\/([^\/]+?)(?:\.git)?$/);
      if (match) {
        finalVariables.serviceName = match[2]; // 返回repo名称
      }
    } 
    // 处理https格式：https://codeup.aliyun.com/org/repo.git
    else {
      match = finalVariables.repoUrl.match(/\/([^\/]+)(?:\.git)?$/);
      if (match) {
        finalVariables.serviceName = match[1].replace('.git', '');
      }
    }
  }
  
  console.log('🔍 [DEBUG] finalVariables:', JSON.stringify(finalVariables, null, 2));
  
  // 生成流水线名称
  const pipelineName = options?.name || generatePipelineName(parsedInfo);
  
  // 使用模块化架构生成YAML
  const generatedYaml = generateModularPipeline({
    keywords: parsedInfo.detectedKeywords,
    buildLanguages: parsedInfo.programmingLanguages,
    buildTools: parsedInfo.buildTools,
    deployTargets: parsedInfo.deployTargets,
    uploadType: options?.uploadType || 'packages',
    variables: finalVariables
  });
  console.log('生成的YAML:', generatedYaml);
  // 创建流水线
  const pipelineId = await createPipelineFunc(organizationId, pipelineName, generatedYaml);
  
  return {
    pipelineId,
    generatedYaml
  };
}

/**
 * 获取用户当前的默认代码库信息（仅用于代码源配置）
 * @param organizationId 组织ID
 * @returns 默认代码库信息
 */
async function getCurrentRepositoryContext(organizationId: string): Promise<{
  repoUrl?: string;
  branch?: string;
  serviceName?: string;
} | null> {
  try {
    // 获取用户最近的代码库列表（按最新活动排序，取第一个）
    const repositories = await listRepositoriesFunc(
      organizationId,
      1, // page
      1, // perPage - 只取第一个
      'last_activity_at', // orderBy - 按最新活动排序
      'desc' // sort - 降序，最新的在前面
    );
    
    if (repositories && repositories.length > 0) {
      const repo = repositories[0];
      
      // 确保repoUrl有.git后缀
      let repoUrl = repo.webUrl;
      if (repoUrl && !repoUrl.endsWith('.git')) {
        repoUrl = `${repoUrl}.git`;
      }
      
      // 使用repo.name作为serviceName，这个值应该就是仓库名称
      const serviceName = repo.name || 'my-app';
      
      return {
        repoUrl: repoUrl || `https://codeup.aliyun.com/${organizationId}/${serviceName}.git`,
        branch: 'master', // 默认分支，实际应该从repo信息中获取
        serviceName: serviceName
      };
    }
    
    return null;
  } catch (error) {
    console.error('获取当前代码库上下文失败:', error);
    return null;
  }
}

/**
 * 获取默认的服务连接ID（用于代码源配置）
 * @param organizationId 组织ID
 * @returns 服务连接ID
 */
async function getDefaultServiceConnectionId(organizationId: string): Promise<string | null> {
  try {
    // 获取Codeup类型的服务连接（代码源最常用）
    const serviceConnections = await listServiceConnectionsFunc(organizationId, 'codeup');
    if (serviceConnections && serviceConnections.length > 0) {
      // 优先使用UUID，如果没有UUID则使用ID转字符串
      return serviceConnections[0].uuid || null;
    }
    return null;
  } catch (error) {
    console.error('获取Codeup服务连接失败:', error);
    return null;
  }
}

/**
 * 获取默认的Packages服务连接ID（用于制品上传配置）
 * @param organizationId 组织ID
 * @returns Packages服务连接ID
 */
async function getDefaultPackagesServiceConnectionId(organizationId: string): Promise<string | null> {
  try {
    // 获取packages类型的服务连接
    const serviceConnections = await listServiceConnectionsFunc(organizationId, 'packages');
    if (serviceConnections && serviceConnections.length > 0) {
      // 优先使用UUID，如果没有UUID则使用ID转字符串
      return serviceConnections[0].uuid || null;
    }
    return null;
  } catch (error) {
    console.error('获取Packages服务连接失败:', error);
    return null;
  }
}

/**
 * 获取默认的主机组ID（用于VM部署配置）
 * 注意：由于主机组API只返回数字ID而不是UUID，这个函数暂时不使用
 * 用户需要在描述中明确指定主机组UUID
 * @param organizationId 组织ID
 * @returns null（暂不自动获取）
 */
async function getDefaultHostGroupId(organizationId: string): Promise<string | null> {
  // 暂时不自动获取主机组，因为API只返回数字ID，无法在流水线中使用
  // 用户需要在options中明确提供machineGroupId（UUID格式）
  return null;
}


