import { 
  getMachineDeployLog,
  addHostListToHostGroup,
  addHostListToDeployGroup,
  deleteHostListFromDeployGroup,
  deleteHostListFromHostGroup,
  getDeployGroup,
  listResourceInstances,
  getResourceInstance,
  updateResourceInstance,
  GetMachineDeployLogRequestSchema,
  AddHostListToHostGroupRequestSchema,
  AddHostListToDeployGroupRequestSchema,
  DeleteHostListFromDeployGroupRequestSchema,
  DeleteHostListFromHostGroupRequestSchema,
  GetDeployGroupRequestSchema,
  ListResourceInstancesRequestSchema,
  GetResourceInstanceRequestSchema,
  UpdateResourceInstanceRequestSchema
} from '../operations/appstack/deploymentResources.js';

/**
 * Handle the appstack deployment resources tool requests
 * 
 * @param request - The tool request
 * @returns The tool response or null if not handled
 */
export async function handleAppStackDeploymentResourceTools(request: any) {
  switch (request.params.name) {
    case 'get_machine_deploy_log':
      const getLogParams = GetMachineDeployLogRequestSchema.parse(request.params.arguments);
      const getLogResult = await getMachineDeployLog(getLogParams);
      return {
        content: [{ type: "text", text: JSON.stringify(getLogResult, null, 2) }],
      };
      
    case 'add_host_list_to_host_group':
      const addHostGroupParams = AddHostListToHostGroupRequestSchema.parse(request.params.arguments);
      const addHostGroupResult = await addHostListToHostGroup(addHostGroupParams);
      return {
        content: [{ type: "text", text: JSON.stringify(addHostGroupResult, null, 2) }],
      };
      
    case 'add_host_list_to_deploy_group':
      const addDeployGroupParams = AddHostListToDeployGroupRequestSchema.parse(request.params.arguments);
      const addDeployGroupResult = await addHostListToDeployGroup(addDeployGroupParams);
      return {
        content: [{ type: "text", text: JSON.stringify(addDeployGroupResult, null, 2) }],
      };
      
    case 'delete_host_list_from_deploy_group':
      const deleteDeployGroupParams = DeleteHostListFromDeployGroupRequestSchema.parse(request.params.arguments);
      const deleteDeployGroupResult = await deleteHostListFromDeployGroup(deleteDeployGroupParams);
      return {
        content: [{ type: "text", text: JSON.stringify(deleteDeployGroupResult, null, 2) }],
      };
      
    case 'delete_host_list_from_host_group':
      const deleteHostGroupParams = DeleteHostListFromHostGroupRequestSchema.parse(request.params.arguments);
      const deleteHostGroupResult = await deleteHostListFromHostGroup(deleteHostGroupParams);
      return {
        content: [{ type: "text", text: JSON.stringify(deleteHostGroupResult, null, 2) }],
      };
      
    case 'get_deploy_group':
      const getDeployGroupParams = GetDeployGroupRequestSchema.parse(request.params.arguments);
      const getDeployGroupResult = await getDeployGroup(getDeployGroupParams);
      return {
        content: [{ type: "text", text: JSON.stringify(getDeployGroupResult, null, 2) }],
      };
      
    case 'list_resource_instances':
      const listInstancesParams = ListResourceInstancesRequestSchema.parse(request.params.arguments);
      const listInstancesResult = await listResourceInstances(listInstancesParams);
      return {
        content: [{ type: "text", text: JSON.stringify(listInstancesResult, null, 2) }],
      };
      
    case 'get_resource_instance':
      const getInstanceParams = GetResourceInstanceRequestSchema.parse(request.params.arguments);
      const getInstanceResult = await getResourceInstance(getInstanceParams);
      return {
        content: [{ type: "text", text: JSON.stringify(getInstanceResult, null, 2) }],
      };
      
    case 'update_resource_instance':
      const updateInstanceParams = UpdateResourceInstanceRequestSchema.parse(request.params.arguments);
      const updateInstanceResult = await updateResourceInstance(updateInstanceParams);
      return {
        content: [{ type: "text", text: JSON.stringify(updateInstanceResult, null, 2) }],
      };
      
    default:
      return null;
  }
}