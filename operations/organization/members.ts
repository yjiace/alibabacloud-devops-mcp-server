import {OrganizationMembersSchema, OrganizationMembers} from '../../common/types.js';
import {buildUrl, yunxiaoRequest} from "../../common/utils.js";

/**
 * 查询组织成员列表
 * @param organizationId 组织ID
 * @param page 当前页，默认1
 * @param perPage 每页数据条数，默认100
 * @returns 组织成员列表
 */
export const getOrganizationMembersFunc = async (
  organizationId: string,
  page: number = 1,
  perPage: number = 100
): Promise<OrganizationMembers> => {

  const url = `/oapi/v1/platform/organizations/${organizationId}/members`;

  const params = {
    page: page,
    perPage: perPage
  };
  const urlWithParams = buildUrl(url, params);

  const response = await yunxiaoRequest(urlWithParams, {
    method: "GET",
  });

  // 验证响应数据结构
  return OrganizationMembersSchema.parse(response);
};