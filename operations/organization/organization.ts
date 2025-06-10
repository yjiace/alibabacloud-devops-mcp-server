import { z } from "zod";
import {buildUrl, yunxiaoRequest} from "../../common/utils.js";
import {
  CurrentOrganizationInfoSchema,
  UserOrganizationsInfoSchema,
  CurrentUserSchema,
  OrganizationDepartmentsSchema,
  DepartmentInfoSchema,
} from "../../common/types.js";

export async function getCurrentOrganizationInfoFunc(
): Promise<z.infer<typeof CurrentOrganizationInfoSchema>> {
  const url = "/oapi/v1/platform/user";

  const response = await yunxiaoRequest(url, {
    method: "GET",
  });

  const responseData = response as { 
    lastOrganization?: string;
    id?: string; 
    name?: string;
  };

  const mappedResponse = {
    lastOrganization: responseData.lastOrganization, // Organization ID
    userId: responseData.id,                         // Map API's "id" to userId
    userName: responseData.name                      // Map API's "name" to userName
  };

  return CurrentOrganizationInfoSchema.parse(mappedResponse);
}

export async function getUserOrganizationsFunc(
): Promise<z.infer<typeof UserOrganizationsInfoSchema>> {
  const url = "/oapi/v1/platform/organizations";

  const response = await yunxiaoRequest(url, {
    method: "GET",
  });

  if (!Array.isArray(response)) {
    return [];
  }

  return UserOrganizationsInfoSchema.parse(response);
}

export async function getOrganizationDepartmentsFunc(
    organizationId: string,
    parentId?: string
): Promise<z.infer<typeof OrganizationDepartmentsSchema>> {
  const baseUrl = `/oapi/v1/platform/organizations/${organizationId}/departments`;

  const params: Record<string, string | undefined> = {};
  if (parentId) {
    params.parentId = parentId;
  }

  const url = buildUrl(baseUrl, params);

  const response = await yunxiaoRequest(url, {
    method: "GET"
  });

  return OrganizationDepartmentsSchema.parse(response);
}

export async function getOrganizationDepartmentInfoFunc(
  organizationId: string,
  id: string
): Promise<z.infer<typeof DepartmentInfoSchema>> {
  const url = `/oapi/v1/platform/organizations/${organizationId}/departments/${id}`;

  const response = await yunxiaoRequest(url, {
    method: "GET",
  });

  return DepartmentInfoSchema.parse(response);
}

export async function getOrganizationDepartmentAncestorsFunc(
    organizationId: string,
    id: string): Promise<z.infer<typeof OrganizationDepartmentsSchema>>  {
  const url = `/oapi/v1/platform/organizations/${organizationId}/departments/${id}/ancestors`;
  const response = await yunxiaoRequest(url, {
    method: "GET",
  })
  return OrganizationDepartmentsSchema.parse(response);
};

export async function getCurrentUserFunc(): Promise<z.infer<typeof CurrentUserSchema>> {
  const url = "/oapi/v1/platform/user";

  const response = await yunxiaoRequest(url, {
    method: "GET",
  });

  return CurrentUserSchema.parse(response);
} 