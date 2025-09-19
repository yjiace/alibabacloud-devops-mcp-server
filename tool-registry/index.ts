import { getCodeManagementTools } from './code-management.js';
import { getOrganizationTools } from './organization.js';
import { getProjectManagementTools } from './project-management.js';
import { getPipelineTools } from './pipeline.js';
import { getPackageManagementTools } from './packages.js';
import { getServiceConnectionTools } from './service-connections.js';
import { getAppStackTools } from './appstack.js';
import { getAppStackTagTools } from './appstack-tags.js';
import { getAppStackTemplateTools } from './appstack-templates.js';
import { getAppStackGlobalVarTools } from './appstack-global-vars.js';
import { getAppStackVariableGroupTools } from './appstack-variable-groups.js';
import { getAppStackOrchestrationTools } from './appstack-orchestrations.js';
import { getAppStackChangeRequestTools } from './appstack-change-requests.js';
import { getAppStackDeploymentResourceTools } from './appstack-deployment-resources.js';
import { getAppStackChangeOrderTools } from './appstack-change-orders.js';
import { getEffortTools } from './effort.js';
import { getResourceMemberTools } from './resourceMember.js';
import { getVMDeployOrderTools } from './vmDeployOrder.js';
import { getCommitTools } from './commit.js';

export const getAllTools = () => [
  ...getCodeManagementTools(),
  ...getOrganizationTools(),
  ...getProjectManagementTools(),
  ...getPipelineTools(),
  ...getPackageManagementTools(),
  ...getServiceConnectionTools(),
  ...getAppStackTools(),
  ...getAppStackTagTools(),
  ...getAppStackTemplateTools(),
  ...getAppStackGlobalVarTools(),
  ...getAppStackVariableGroupTools(),
  ...getAppStackOrchestrationTools(),
  ...getAppStackChangeRequestTools(),
  ...getAppStackDeploymentResourceTools(),
  ...getAppStackChangeOrderTools(),
  ...getEffortTools(),
  ...getResourceMemberTools(),
  ...getVMDeployOrderTools(),
  ...getCommitTools()
];