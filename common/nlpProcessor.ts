/**
 * 自然语言处理模块
 * 用于从用户描述中提取流水线相关的关键信息
 * 基于云效Flow官方文档优化
 */

import { TemplateVariables } from './pipelineTemplates.js';

// 解析结果接口
export interface ParsedPipelineInfo {
  // 识别的关键词
  detectedKeywords: string[];
  
  // 流水线类型
  pipelineTypes: string[];
  
  // 编程语言
  programmingLanguages: string[];
  
  // 构建工具
  buildTools: string[];
  
  // 部署目标
  deployTargets: string[];
  
  // 提取的变量
  variables: TemplateVariables;
  
  // 匹配的模板（保留兼容性，但不再使用）
  matchedTemplates: any[];
}

// 关键词映射定义（基于云效官方文档扩展）
const KEYWORD_MAPPINGS = {
  // 流水线类型
  pipelineTypes: {
    '构建': ['build', '构建', '编译', 'compile'],
    '部署': ['deploy', '部署', '发布', 'release', '上线'],
    '测试': ['test', '测试', '单元测试', 'unit test'],
    '发布': ['release', '发布', '上线', 'publish']
  },
  
  // 编程语言（基于官方步骤文档）
  programmingLanguages: {
    'java': ['java', 'jar', 'maven', 'gradle', 'spring', 'springboot', 'javabuild'],
    'nodejs': ['node', 'nodejs', 'node.js', 'npm', 'yarn', 'javascript', 'js', 'express', 'koa', 'nodebuild'],
    'python': ['python', 'py', 'pip', 'django', 'flask', 'fastapi', 'pytest', 'pythonbuild'],
    'go': ['go', 'golang', 'gin', 'golangbuild'],
    'php': ['php', 'composer', 'laravel', 'phpbuild'],
    'csharp': ['c#', 'csharp', 'dotnet', '.net', 'nuget', 'dotnetcore', 'dotnetcorebuild'],
    'cpp': ['c++', 'cpp', 'cmake', 'gcc', 'gccbuild'],
    'rust': ['rust', 'cargo', 'rustbuild'],
    'ruby': ['ruby', 'gem', 'rails', 'rubybuild'],
    'asp': ['asp', 'aspnet', 'aspnetbuild']
  },
  
  // 构建工具（基于官方步骤文档）
  buildTools: {
    'maven': ['maven', 'mvn', 'pom.xml'],
    'gradle': ['gradle', 'build.gradle'],
    'npm': ['npm', 'package.json'],
    'yarn': ['yarn'],
    'pip': ['pip', 'requirements.txt', 'setup.py'],
    'go': ['go build', 'go mod'],
    'docker': ['docker', 'dockerfile', '容器', 'container', 'acrdockerbuild', 'privateregistrydockerbuild'],
    'webpack': ['webpack', 'babel', 'rollup'],
    'dotnet': ['dotnet', 'nuget'],
    'composer': ['composer'],
    'cargo': ['cargo'],
    'gem': ['gem'],
    'custom': ['custom', '自定义', 'customenvironmentbuild']
  },
  
  // 部署目标（基于官方组件文档）
  deployTargets: {
    'kubernetes': ['kubernetes', 'k8s', '容器编排', 'pod', 'deployment', 'service', 'kubernetesdeploy', 'kubernetesbatchdeploy', 'kubernetesbluegreendeploy'],
    'docker': ['docker', '容器', 'container', 'vmdockerdeploy'],
    'vm': ['vm', '虚拟机', '主机', 'host', 'server', '服务器', 'vmdeploy'],
    'ecs': ['ecs', 'ecsappdeploy'],
    'sae': ['sae', 'serverless', 'saedeploy', 'saejobupdate'],
    'edas': ['edas', 'edasecsdeploy', 'edaskubernetesdeploy'],
    'ess': ['ess', 'essecideploy', 'essecsdeploy'],
    'appstack': ['appstack', 'appstackflowdeploy'],
    'asm': ['asm', 'asmbluegreendeploy'],
    'cloud': ['云', 'cloud', 'aws', 'azure', 'gcp']
  },
  
  // 版本相关（基于官方文档）
  versions: {
    'jdk': {
      '1.6': ['jdk1.6', 'java1.6'],
      '1.7': ['jdk1.7', 'java1.7'],
      '1.8': ['jdk1.8', 'java1.8', 'java8'],
      '11': ['jdk11', 'java11'],
      '17': ['jdk17', 'java17'],
      '21': ['jdk21', 'java21']
    },
    'maven': {
      '3.6.3': ['maven3.6.3', 'mvn3.6.3'],
      '3.8.4': ['maven3.8.4', 'mvn3.8.4'],
      '3.9.3': ['maven3.9.3', 'mvn3.9.3']
    },
    'node': {
      '16.8': ['node16', 'nodejs16'],
      '18.12': ['node18', 'nodejs18'],
      '20': ['node20', 'nodejs20']
    },
    'python': {
      '3.9': ['python3.9', 'py3.9'],
      '3.12': ['python3.12', 'py3.12']
    },
    'go': {
      '1.19.x': ['go1.19', 'golang1.19'],
      '1.20.x': ['go1.20', 'golang1.20'],
      '1.21.x': ['go1.21', 'golang1.21']
    }
  }
};

// URL 正则表达式
const URL_REGEX = /(https?:\/\/|git@|ssh:\/\/)[^:\s]+\.git/g;

// 分支名正则表达式
const BRANCH_REGEX = /(?:分支|branch)[：:\s]*([a-zA-Z0-9\-_\/]+)/gi;

// 服务名提取正则表达式
const SERVICE_NAME_REGEX = /(?:项目名称是|项目名称|项目叫|项目|应用名称是|应用名称|应用叫|应用|服务名称是|服务名称|服务叫|服务|service|project|app)[：:\s是]*([a-zA-Z0-9\-_]+)/gi;

// 版本号提取正则表达式
const VERSION_REGEX = /(?:jdk|java|node|python|go|golang|maven)\s*([0-9\.x]+)/gi;

// 机器组ID提取正则表达式
const MACHINE_GROUP_REGEX = /(?:机器组|machinegroup|machine group)[：:\s]*([a-zA-Z0-9\-_]+)/gi;

// 命名空间提取正则表达式
const NAMESPACE_REGEX = /(?:namespace|命名空间)[：:\s]*([a-zA-Z0-9\-_]+)/gi;

/**
 * 从文本中提取关键词
 */
function extractKeywords(text: string): string[] {
  // 转换为小写并分词
  const words = text.toLowerCase()
    .replace(/[，。！？；：""''（）【】\[\]{}]/g, ' ')
    .split(/\s+/)
    .filter(word => word.length > 1);
  
  return [...new Set(words)];
}

/**
 * 根据关键词映射识别类型
 */
function identifyTypes(keywords: string[], mappings: Record<string, string[]>): string[] {
  const identified: string[] = [];
  
  for (const [type, typeKeywords] of Object.entries(mappings)) {
    const hasMatch = typeKeywords.some(keyword => 
      keywords.some(k => k.includes(keyword) || keyword.includes(k))
    );
    
    if (hasMatch) {
      identified.push(type);
    }
  }
  
  return identified;
}

/**
 * 提取版本信息
 */
function extractVersions(description: string): Record<string, string> {
  const versions: Record<string, string> = {};
  
  // 提取版本号
  const versionMatches = [...description.matchAll(VERSION_REGEX)];
  versionMatches.forEach(match => {
    const tool = match[0].split(/\s+/)[0].toLowerCase();
    const version = match[1];
    
    if (tool.includes('jdk') || tool.includes('java')) {
      versions.jdkVersion = version;
    } else if (tool.includes('node')) {
      versions.nodeVersion = version;
    } else if (tool.includes('python')) {
      versions.pythonVersion = version;
    } else if (tool.includes('go')) {
      versions.goVersion = version + '.x';
    } else if (tool.includes('maven')) {
      versions.mavenVersion = version;
    }
  });
  
  return versions;
}

/**
 * 从描述中提取变量信息
 */
function extractVariables(description: string): TemplateVariables {
  const variables: TemplateVariables = {};
  
  // 提取仓库URL
  const urlMatches = description.match(URL_REGEX);
  if (urlMatches && urlMatches.length > 0) {
    variables.repoUrl = urlMatches[0];
  }
  
  // 提取分支名
  const branchMatches = [...description.matchAll(BRANCH_REGEX)];
  if (branchMatches.length > 0) {
    variables.branch = branchMatches[0][1];
  }
  
  // 提取服务名
  const serviceMatches = [...description.matchAll(SERVICE_NAME_REGEX)];
  if (serviceMatches.length > 0) {
    variables.serviceName = serviceMatches[0][1];
  }
  
  // 提取机器组ID
  const machineGroupMatches = [...description.matchAll(MACHINE_GROUP_REGEX)];
  if (machineGroupMatches.length > 0) {
    variables.machineGroupId = machineGroupMatches[0][1];
  }
  
  // 提取命名空间
  const namespaceMatches = [...description.matchAll(NAMESPACE_REGEX)];
  if (namespaceMatches.length > 0) {
    variables.namespace = namespaceMatches[0][1];
  }
  
  // 提取版本信息
  const versions = extractVersions(description);
  Object.assign(variables, versions);
  
  // 根据描述推断其他变量
  const lowerDesc = description.toLowerCase();
  
  // 推断构建命令
  if (lowerDesc.includes('maven') || lowerDesc.includes('mvn')) {
    if (lowerDesc.includes('跳过测试') || lowerDesc.includes('skip test')) {
      variables.buildCommand = 'mvn clean package -Dmaven.test.skip=true';
    } else {
      variables.buildCommand = 'mvn clean package';
    }
  } else if (lowerDesc.includes('npm')) {
    if (lowerDesc.includes('build')) {
      variables.buildCommand = 'npm run build';
    }
  } else if (lowerDesc.includes('yarn')) {
    variables.buildCommand = 'yarn build';
  } else if (lowerDesc.includes('go')) {
    variables.buildCommand = 'go build -o app main.go';
  } else if (lowerDesc.includes('gradle')) {
    variables.buildCommand = 'gradle build';
  }
  
  // 推断测试命令
  if (lowerDesc.includes('maven') || lowerDesc.includes('java')) {
    variables.testCommand = 'mvn test';
  } else if (lowerDesc.includes('npm')) {
    variables.testCommand = 'npm test';
  } else if (lowerDesc.includes('yarn')) {
    variables.testCommand = 'yarn test';
  } else if (lowerDesc.includes('python')) {
    variables.testCommand = 'python -m pytest';
  } else if (lowerDesc.includes('go')) {
    variables.testCommand = 'go test ./...';
  } else if (lowerDesc.includes('gradle')) {
    variables.testCommand = 'gradle test';
  }
  
  // 推断构建产物路径
  if (lowerDesc.includes('java') || lowerDesc.includes('maven') || lowerDesc.includes('gradle')) {
    variables.artifactPath = 'target/';
  } else if (lowerDesc.includes('nodejs') || lowerDesc.includes('npm') || lowerDesc.includes('yarn')) {
    variables.artifactPath = 'dist/';
  } else if (lowerDesc.includes('python')) {
    variables.artifactPath = 'dist/';
  } else if (lowerDesc.includes('go')) {
    variables.artifactPath = 'app';
  }
  
  // 推断部署目标
  if (lowerDesc.includes('kubernetes') || lowerDesc.includes('k8s')) {
    variables.deployTarget = 'kubernetes';
  } else if (lowerDesc.includes('主机') || lowerDesc.includes('vm') || lowerDesc.includes('host')) {
    variables.deployTarget = 'vm';
  } else if (lowerDesc.includes('docker') && !lowerDesc.includes('kubernetes')) {
    variables.deployTarget = 'docker';
  } else if (lowerDesc.includes('sae') || lowerDesc.includes('serverless')) {
    variables.deployTarget = 'sae';
  } else if (lowerDesc.includes('ecs')) {
    variables.deployTarget = 'ecs';
  }
  
  // 推断上传类型
  if (lowerDesc.includes('packages') || lowerDesc.includes('制品仓库')) {
    variables.uploadType = 'packages';
  } else {
    variables.uploadType = 'flowPublic';
  }
  
  // 推断执行用户
  if (lowerDesc.includes('admin') || lowerDesc.includes('管理员')) {
    variables.executeUser = 'admin';
  } else {
    variables.executeUser = 'root';
  }
  
  return variables;
}

/**
 * 解析用户的自然语言描述
 */
export function parseUserDescription(description: string): ParsedPipelineInfo {
  // 提取关键词
  const detectedKeywords = extractKeywords(description);
  
  // 识别各种类型
  const pipelineTypes = identifyTypes(detectedKeywords, KEYWORD_MAPPINGS.pipelineTypes);
  const programmingLanguages = identifyTypes(detectedKeywords, KEYWORD_MAPPINGS.programmingLanguages);
  const buildTools = identifyTypes(detectedKeywords, KEYWORD_MAPPINGS.buildTools);
  const deployTargets = identifyTypes(detectedKeywords, KEYWORD_MAPPINGS.deployTargets);
  
  // 提取变量
  const variables = extractVariables(description);
  
  const result: ParsedPipelineInfo = {
    detectedKeywords,
    pipelineTypes,
    programmingLanguages,
    buildTools,
    deployTargets,
    variables,
    matchedTemplates: [], // 不再需要预选模板，模块化架构会自动匹配
  };
  
  return result;
}

/**
 * 生成流水线名称建议
 */
export function generatePipelineName(parsedInfo: ParsedPipelineInfo): string {
  const { programmingLanguages, pipelineTypes, variables } = parsedInfo;
  
  let name = '';
  
  // 添加服务名
  if (variables.serviceName) {
    name = variables.serviceName;
  } else if (programmingLanguages.length > 0) {
    name = programmingLanguages[0];
  } else {
    name = '应用';
  }
  
  // 添加流水线类型
  if (pipelineTypes.length > 0) {
    const type = pipelineTypes.includes('部署') ? '部署' : 
                 pipelineTypes.includes('构建') ? '构建' : 
                 pipelineTypes[0];
    name += type + '流水线';
  } else {
    name += '流水线';
  }
  
  return name;
} 