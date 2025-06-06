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
 * 基于结构化参数生成流水线YAML（不创建流水线）
 * @param options 结构化的流水线配置选项
 * @returns 生成的YAML字符串
 */
export async function generatePipelineYamlFunc(
  options: {
    // 技术栈信息（必须明确指定）
    buildLanguage: 'java' | 'nodejs' | 'python' | 'go' | 'dotnet';
    buildTool: 'maven' | 'gradle' | 'npm' | 'yarn' | 'pip' | 'go' | 'dotnet';
    deployTarget?: 'vm' | 'k8s' | 'none';
    
    // 代码仓库配置
    repoUrl?: string;
    branch?: string;
    serviceName?: string;
    serviceConnectionId?: string;
    
    // 技术版本配置
    jdkVersion?: string;        // Java版本，如 "1.8", "11", "17"
    mavenVersion?: string;      // Maven版本，如 "3.6.3"
    nodeVersion?: string;       // Node.js版本，如 "18.12", "20.x"
    pythonVersion?: string;     // Python版本，如 "3.7", "3.12"
    goVersion?: string;         // Go版本，如 "1.21"
    
    // 构建配置
    buildCommand?: string;      // 自定义构建命令
    testCommand?: string;       // 自定义测试命令
    
    // 制品上传配置
    uploadType?: 'packages' | 'flowPublic';
    packagesServiceConnection?: string;
    artifactName?: string;
    artifactVersion?: string;
    packagesRepoId?: string;
    includePathInArtifact?: boolean;
    
    // VM部署配置（当deployTarget为'vm'时）
    machineGroupId?: string;
    executeUser?: string;
    artifactDownloadPath?: string;
    deployCommand?: string;
    pauseStrategy?: 'firstBatchPause' | 'noPause' | 'eachBatchPause';
    batchNumber?: number;
    
    // Kubernetes部署配置（当deployTarget为'k8s'时）
    kubernetesClusterId?: string;
    kubectlVersion?: string;
    namespace?: string;
    yamlPath?: string;
    dockerImage?: string;
  }
): Promise<string> {
  // 准备变量，确保版本号有双引号
  const variables: TemplateVariables = {
    // 基础配置
    ...(options.repoUrl && { repoUrl: options.repoUrl }),
    ...(options.branch && { branch: options.branch }),
    ...(options.serviceName && { serviceName: options.serviceName }),
    ...(options.serviceConnectionId && { serviceConnectionId: options.serviceConnectionId }),
    ...(options.packagesServiceConnection && { packagesServiceConnection: options.packagesServiceConnection }),
    ...(options.machineGroupId && { machineGroupId: options.machineGroupId }),
    ...(options.namespace && { namespace: options.namespace }),
    ...(options.dockerImage && { dockerImage: options.dockerImage }),
    
    // 版本相关（确保双引号）
    ...(options.jdkVersion && { jdkVersion: `"${options.jdkVersion}"` }),
    ...(options.mavenVersion && { mavenVersion: `"${options.mavenVersion}"` }),
    ...(options.nodeVersion && { nodeVersion: `"${options.nodeVersion}"` }),
    ...(options.pythonVersion && { pythonVersion: `"${options.pythonVersion}"` }),
    ...(options.goVersion && { goVersion: `"${options.goVersion}"` }),
    ...(options.kubectlVersion && { kubectlVersion: `"${options.kubectlVersion}"` }),
    
    // 构建物上传相关
    ...(options.uploadType && { uploadType: options.uploadType }),
    ...(options.artifactName && { artifactName: options.artifactName }),
    ...(options.artifactVersion && { artifactVersion: options.artifactVersion }),
    ...(options.packagesRepoId && { packagesRepoId: options.packagesRepoId }),
    ...(options.includePathInArtifact !== undefined && { includePathInArtifact: options.includePathInArtifact }),
    
    // 部署相关
    ...(options.executeUser && { executeUser: options.executeUser }),
    ...(options.artifactDownloadPath && { artifactDownloadPath: options.artifactDownloadPath }),
    ...(options.kubernetesClusterId && { kubernetesClusterId: options.kubernetesClusterId }),
    ...(options.yamlPath && { yamlPath: options.yamlPath }),
    
    // 命令
    ...(options.buildCommand && { buildCommand: options.buildCommand }),
    ...(options.testCommand && { testCommand: options.testCommand }),
    ...(options.deployCommand && { deployCommand: options.deployCommand }),
  };
  
  // 转换为模块化流水线选项
  const deployTargets = options.deployTarget ? [options.deployTarget] : [];
  
  // 使用模块化架构生成YAML
  return generateModularPipeline({
    keywords: [options.buildLanguage, options.buildTool],
    buildLanguages: [options.buildLanguage],
    buildTools: [options.buildTool],
    deployTargets: deployTargets,
    uploadType: options.uploadType || 'packages',
    variables: variables
  });
}

/**
 * 基于结构化参数创建流水线
 * @param organizationId 组织ID
 * @param options 结构化的流水线配置选项
 * @returns 创建结果，包含流水线ID和生成的YAML
 */
export async function createPipelineWithOptionsFunc(
  organizationId: string,
  options: {
    // 基础信息
    name: string;
    
    // 技术栈信息（必须明确指定）
    buildLanguage: 'java' | 'nodejs' | 'python' | 'go' | 'dotnet';
    buildTool: 'maven' | 'gradle' | 'npm' | 'yarn' | 'pip' | 'go' | 'dotnet';
    deployTarget?: 'vm' | 'k8s' | 'none';
    
    // 代码仓库配置（大模型应该从IDE上下文中获取）
    repoUrl?: string;
    branch?: string;
    serviceName?: string;
    serviceConnectionId?: string;
    
    // 技术版本配置
    jdkVersion?: string;        // Java版本，如 "1.8", "11", "17"
    mavenVersion?: string;      // Maven版本，如 "3.6.3"
    nodeVersion?: string;       // Node.js版本，如 "18.12", "20.x"
    pythonVersion?: string;     // Python版本，如 "3.7", "3.12"
    goVersion?: string;         // Go版本，如 "1.21"
    
    // 构建配置
    buildCommand?: string;      // 自定义构建命令
    testCommand?: string;       // 自定义测试命令
    
    // 制品上传配置
    uploadType?: 'packages' | 'flowPublic';
    packagesServiceConnection?: string;
    artifactName?: string;
    artifactVersion?: string;
    packagesRepoId?: string;
    includePathInArtifact?: boolean;
    
    // VM部署配置（当deployTarget为'vm'时）
    machineGroupId?: string;
    executeUser?: string;
    artifactDownloadPath?: string;
    deployCommand?: string;
    pauseStrategy?: 'firstBatchPause' | 'noPause' | 'eachBatchPause';
    batchNumber?: number;
    
    // Kubernetes部署配置（当deployTarget为'k8s'时）
    kubernetesClusterId?: string;
    kubectlVersion?: string;
    namespace?: string;
    yamlPath?: string;
    dockerImage?: string;
  }
): Promise<{
  pipelineId: number;
  generatedYaml: string;
  usedTemplate: string;
}> {
  // 获取默认服务连接ID（如果用户没有明确指定）
  let defaultServiceConnectionId: string | null = null;
  const hasServiceConnectionId = options.serviceConnectionId;
  if (!hasServiceConnectionId) {
    defaultServiceConnectionId = await getDefaultServiceConnectionId(organizationId);
  }
  
  // 获取默认Packages服务连接ID（如果用户没有明确指定且需要packages上传）
  let defaultPackagesServiceConnectionId: string | null = null;
  const hasPackagesServiceConnectionId = options.packagesServiceConnection;
  const needsPackagesUpload = !options.uploadType || options.uploadType === 'packages';
  if (!hasPackagesServiceConnectionId && needsPackagesUpload) {
    defaultPackagesServiceConnectionId = await getDefaultPackagesServiceConnectionId(organizationId);
  }
  
  // 获取默认主机组ID（如果用户没有明确指定且需要VM部署）
  let defaultMachineGroupId: string | null = null;
  const hasMachineGroupId = options.machineGroupId;
  const needsVMDeploy = options.deployTarget === 'vm';
  if (!hasMachineGroupId && needsVMDeploy) {
    defaultMachineGroupId = await getDefaultHostGroupId(organizationId);
  }
  
  // 准备模块化流水线生成的变量
  const finalVariables: TemplateVariables = {
    // 基础配置（直接使用用户提供的值）
    ...(options.repoUrl && { repoUrl: options.repoUrl }),
    ...(options.branch && { branch: options.branch }),
    ...(options.serviceName && { serviceName: options.serviceName }),
    
    // 使用获取到的默认服务连接ID
    ...(defaultServiceConnectionId && !hasServiceConnectionId && { serviceConnectionId: defaultServiceConnectionId }),
    
    // 使用获取到的默认Packages服务连接ID
    ...(defaultPackagesServiceConnectionId && !hasPackagesServiceConnectionId && { packagesServiceConnection: defaultPackagesServiceConnectionId }),
    
    // 使用获取到的默认主机组ID
    ...(defaultMachineGroupId && !hasMachineGroupId && { machineGroupId: defaultMachineGroupId }),
    
    // 用户明确指定的值优先级最高
    ...(options.serviceConnectionId && { serviceConnectionId: options.serviceConnectionId }),
    ...(options.packagesServiceConnection && { packagesServiceConnection: options.packagesServiceConnection }),
    ...(options.machineGroupId && { machineGroupId: options.machineGroupId }),
    ...(options.namespace && { namespace: options.namespace }),
    ...(options.dockerImage && { dockerImage: options.dockerImage }),
    
    // 版本相关（确保双引号）
    ...(options.jdkVersion && { jdkVersion: `"${options.jdkVersion}"` }),
    ...(options.mavenVersion && { mavenVersion: `"${options.mavenVersion}"` }),
    ...(options.nodeVersion && { nodeVersion: `"${options.nodeVersion}"` }),
    ...(options.pythonVersion && { pythonVersion: `"${options.pythonVersion}"` }),
    ...(options.goVersion && { goVersion: `"${options.goVersion}"` }),
    ...(options.kubectlVersion && { kubectlVersion: `"${options.kubectlVersion}"` }),
    
    // 构建物上传相关
    ...(options.uploadType && { uploadType: options.uploadType }),
    ...(options.artifactName && { artifactName: options.artifactName }),
    ...(options.artifactVersion && { artifactVersion: options.artifactVersion }),
    ...(options.packagesRepoId && { packagesRepoId: options.packagesRepoId }),
    ...(options.includePathInArtifact !== undefined && { includePathInArtifact: options.includePathInArtifact }),
    
    // 部署相关
    ...(options.executeUser && { executeUser: options.executeUser }),
    ...(options.artifactDownloadPath && { artifactDownloadPath: options.artifactDownloadPath }),
    ...(options.kubernetesClusterId && { kubernetesClusterId: options.kubernetesClusterId }),
    ...(options.yamlPath && { yamlPath: options.yamlPath }),
    
    // 命令
    ...(options.buildCommand && { buildCommand: options.buildCommand }),
    ...(options.testCommand && { testCommand: options.testCommand }),
    ...(options.deployCommand && { deployCommand: options.deployCommand }),
  };
  
  console.log('🔍 [DEBUG] finalVariables:', JSON.stringify(finalVariables, null, 2));
  
  // 转换为模块化流水线选项
  const deployTargets = options.deployTarget ? [options.deployTarget] : [];
  
  // 使用模块化架构生成YAML
  const generatedYaml = generateModularPipeline({
    keywords: [options.buildLanguage, options.buildTool],
    buildLanguages: [options.buildLanguage],
    buildTools: [options.buildTool],
    deployTargets: deployTargets,
    uploadType: options.uploadType || 'packages',
    variables: finalVariables
  });
  
  console.log('生成的YAML:', generatedYaml);
  
  // 创建流水线
  const pipelineId = await createPipelineFunc(organizationId, options.name, generatedYaml);
  
  return {
    pipelineId,
    generatedYaml,
    usedTemplate: '模块化流水线'
  };
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


