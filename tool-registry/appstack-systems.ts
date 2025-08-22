import { z } from 'zod';
import { zodToJsonSchema } from 'zod-to-json-schema';
import { 
  ListSystemsRequestSchema,
  CreateSystemRequestSchema,
  UpdateSystemRequestSchema,
  ListAttachedAppsRequestSchema,
  AttachAppsRequestSchema,
  DetachAppsRequestSchema,
  ListSystemMembersRequestSchema,
  CreateSystemMembersRequestSchema,
  UpdateSystemMemberRequestSchema,
  DeleteSystemMemberRequestSchema
} from '../operations/appstack/systems.js';

// Export all appstack system tools
export const getAppStackSystemTools = () => [
  {
    name: 'list_systems',
    description: '分页查找系统',
    inputSchema: zodToJsonSchema(ListSystemsRequestSchema),
  },
  {
    name: 'create_system',
    description: '创建系统',
    inputSchema: zodToJsonSchema(CreateSystemRequestSchema),
  },
  {
    name: 'update_system',
    description: '更新系统信息',
    inputSchema: zodToJsonSchema(UpdateSystemRequestSchema),
  },
  {
    name: 'list_attached_apps',
    description: '查找系统所关联的应用列表',
    inputSchema: zodToJsonSchema(ListAttachedAppsRequestSchema),
  },
  {
    name: 'attach_apps',
    description: '将应用关联到系统',
    inputSchema: zodToJsonSchema(AttachAppsRequestSchema),
  },
  {
    name: 'detach_apps',
    description: '解除应用与系统的关联',
    inputSchema: zodToJsonSchema(DetachAppsRequestSchema),
  },
  {
    name: 'list_system_members',
    description: '查找系统成员列表',
    inputSchema: zodToJsonSchema(ListSystemMembersRequestSchema),
  },
  {
    name: 'create_system_members',
    description: '添加系统成员',
    inputSchema: zodToJsonSchema(CreateSystemMembersRequestSchema),
  },
  {
    name: 'update_system_member',
    description: '更新系统成员角色',
    inputSchema: zodToJsonSchema(UpdateSystemMemberRequestSchema),
  },
  {
    name: 'delete_system_member',
    description: '删除系统成员',
    inputSchema: zodToJsonSchema(DeleteSystemMemberRequestSchema),
  }
];