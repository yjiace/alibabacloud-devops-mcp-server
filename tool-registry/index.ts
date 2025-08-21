import { getCodeManagementTools } from './code-management.js';
import { getOrganizationTools } from './organization.js';
import { getProjectManagementTools } from './project-management.js';
import { getPipelineTools } from './pipeline.js';
import { getPackageManagementTools } from './packages.js';
import { getServiceConnectionTools } from './service-connections.js';

export const getAllTools = () => [
  ...getCodeManagementTools(),
  ...getOrganizationTools(),
  ...getProjectManagementTools(),
  ...getPipelineTools(),
  ...getPackageManagementTools(),
  ...getServiceConnectionTools(),
];