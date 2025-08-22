import { z } from 'zod';
import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  ListSystemAllReleaseWorkflowsRequestSchema,
  CreateSystemReleaseWorkflowsRequestSchema,
  UpdateSystemReleaseStageRequestSchema,
  ExecuteReleaseStageRequestSchema
} from '../operations/appstack/releaseWorkflows.js';

// Export all appstack release workflow tools
export const getAppStackReleaseWorkflowTools = () => [
  {
    name: 'list_system_all_release_workflows',
    description: '查找系统下所有的发布流程',
    inputSchema: zodToJsonSchema(ListSystemAllReleaseWorkflowsRequestSchema),
  },
  {
    name: 'create_system_release_workflows',
    description: '创建发布流程',
    inputSchema: zodToJsonSchema(CreateSystemReleaseWorkflowsRequestSchema),
  },
  {
    name: 'update_system_release_stage',
    description: '更新发布阶段',
    inputSchema: zodToJsonSchema(UpdateSystemReleaseStageRequestSchema),
  },
  {
    name: 'execute_release_stage',
    description: '执行发布阶段流水线',
    inputSchema: zodToJsonSchema(ExecuteReleaseStageRequestSchema),
  }
];