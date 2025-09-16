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
import {
  listCurrentUserEffortRecordsHandler,
  listEffortRecordsHandler,
  createEffortRecordHandler,
  listEstimatedEffortsHandler,
  createEstimatedEffortHandler,
  updateEffortRecordHandler,
  updateEstimatedEffortHandler
} from './effort.js';

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
    // Effort handlers
    async (request: any) => {
      if (request.params.name === 'list_current_user_effort_records') {
        return listCurrentUserEffortRecordsHandler(request.params.arguments);
      }
      return null;
    },
    async (request: any) => {
      if (request.params.name === 'list_effort_records') {
        return listEffortRecordsHandler(request.params.arguments);
      }
      return null;
    },
    async (request: any) => {
      if (request.params.name === 'create_effort_record') {
        return createEffortRecordHandler(request.params.arguments);
      }
      return null;
    },
    async (request: any) => {
      if (request.params.name === 'list_estimated_efforts') {
        return listEstimatedEffortsHandler(request.params.arguments);
      }
      return null;
    },
    async (request: any) => {
      if (request.params.name === 'create_estimated_effort') {
        return createEstimatedEffortHandler(request.params.arguments);
      }
      return null;
    },
    async (request: any) => {
      if (request.params.name === 'update_effort_record') {
        return updateEffortRecordHandler(request.params.arguments);
      }
      return null;
    },
    async (request: any) => {
      if (request.params.name === 'update_estimated_effort') {
        return updateEstimatedEffortHandler(request.params.arguments);
      }
      return null;
    }
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