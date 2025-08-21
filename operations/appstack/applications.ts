import { z } from 'zod';
import { yunxiaoRequest, buildUrl } from '../../common/utils.js';
import { YunxiaoError } from '../../common/errors.js';
import {debug} from "util";

// Schema for the ListApplications API
export const ListApplicationsRequestSchema = z.object({
  organizationId: z.string().describe("组织id"),
  pagination: z.enum(['keyset']).default('keyset').describe("分页模式参数，目前只支持键集分页 keyset 模式"),
  perPage: z.number().optional().describe("分页尺寸参数，决定一页最多返回多少对象"),
  orderBy: z.string().optional().default("id").describe("分页排序属性，决定根据何种属性进行记录排序；推荐在实现严格遍历时，使用 id 属性"),
  sort: z.enum(['asc', 'desc']).optional().default('asc').describe("分页排序为升降序，asc 为升序，desc 为降序；推荐在实现严格遍历时，使用升序"),
  nextToken: z.string().optional().describe("分页 token，获取第一页数据时无需传入，否则需要传入前一页查询结果中的 nextToken 字段"),
});

export const ListApplicationsResponseSchema = z.object({
  data: z.array(z.object({
    appTemplateDisplayName: z.string().nullable().optional().describe("应用模版展示名称"),
    appTemplateName: z.string().nullable().optional().describe("应用模版名称"),
    creatorId: z.string().optional().describe("应用创建者id"),
    description: z.string().optional().describe("应用描述"),
    gmtCreate: z.string().optional().describe("创建时间"),
    name: z.string().optional().describe("应用名"),
  })),
  nextToken: z.string().optional(),
});

export type ListApplicationsRequest = z.infer<typeof ListApplicationsRequestSchema>;
export type ListApplicationsResponse = z.infer<typeof ListApplicationsResponseSchema>;

/**
 * List applications in an organization with pagination
 * 
 * @param params - The request parameters
 * @returns The list of applications
 */
export async function listApplications(params: ListApplicationsRequest): Promise<ListApplicationsResponse> {
  const { organizationId, ...queryParams } = params;
  
  // Build query string properly
  const query: Record<string, string | number> = {};
  if (queryParams.pagination) query.pagination = queryParams.pagination;
  if (queryParams.perPage) query.perPage = queryParams.perPage;
  if (queryParams.orderBy) query.orderBy = queryParams.orderBy;
  if (queryParams.sort) query.sort = queryParams.sort;
  if (queryParams.nextToken) query.nextToken = queryParams.nextToken;
  
  try {
    // Build the full URL with query parameters
    const url = buildUrl(`/oapi/v1/appstack/organizations/${organizationId}/apps:search`, query);
    
    const response = await yunxiaoRequest(
      url,
      {
        method: 'GET',
      }
    );
    return ListApplicationsResponseSchema.parse(response);
  } catch (error) {
    throw error;
  }
}