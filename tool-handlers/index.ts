import { handleCodeManagementTools } from './code-management.js';
import { handleOrganizationTools } from './organization.js';
import { handleProjectManagementTools } from './project-management.js';
import { handlePipelineTools } from './pipeline.js';
import { handlePackageManagementTools } from './packages.js';
import { handleServiceConnectionTools } from './service-connections.js';
import { handleAppStackTools } from './appstack.js';
import { handleAppStackTagTools } from './appstack-tags.js';
import { handleAppStackTemplateTools } from './appstack-templates.js';
import { handleAppStackGlobalVarTools } from './appstack-global-vars.js';
import { handleAppStackVariableGroupTools } from './appstack-variable-groups.js';
import { handleAppStackOrchestrationTools } from './appstack-orchestrations.js';
import { handleAppStackChangeRequestTools } from './appstack-change-requests.js';
import { handleAppStackDeploymentResourceTools } from './appstack-deployment-resources.js';
import { handleAppStackChangeOrderTools } from './appstack-change-orders.js';
import { handleEffortTools } from './effort.js';
import { handleResourceMemberTools } from './resourceMember.js';
import { handleTagTools } from './tag.js';
import { handleVMDeployOrderTools } from './vmDeployOrder.js';

export const handleToolRequest = async (request: any) => {
  // Try each handler in sequence until one returns a result
  const handlers = [
    handleCodeManagementTools,
    handleOrganizationTools,
    handleProjectManagementTools,
    handlePipelineTools,
    handlePackageManagementTools,
    handleServiceConnectionTools,
    handleAppStackTools,
    handleAppStackTagTools,
    handleAppStackTemplateTools,
    handleAppStackGlobalVarTools,
    handleAppStackVariableGroupTools,
    handleAppStackOrchestrationTools,
    handleAppStackChangeRequestTools,
    handleAppStackDeploymentResourceTools,
    handleAppStackChangeOrderTools,
    handleEffortTools,
    handleResourceMemberTools,
    handleTagTools,
    handleVMDeployOrderTools
  ];

  for (const handler of handlers) {
    const result = await handler(request);
    if (result !== null) {
      return result;
    }
  }

  // If no handler matched, throw an error
  throw new Error(`Unknown tool: ${request.params.name}`);
};