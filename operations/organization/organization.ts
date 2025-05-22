import { z } from "zod";
import { yunxiaoRequest } from "../../common/utils.js";
import { CurrentOrganizationInfoSchema, UserOrganizationsInfoSchema } from "../../common/types.js";

// Function implementations
export async function getCurrentOrganizationInfoFunc(
): Promise<z.infer<typeof CurrentOrganizationInfoSchema>> {
  const url = "/oapi/v1/platform/user";

  const response = await yunxiaoRequest(url, {
    method: "GET",
  });

  // Type assertion to ensure TypeScript understands the response structure
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

  // Ensure response is an array
  if (!Array.isArray(response)) {
    return [];
  }

  return UserOrganizationsInfoSchema.parse(response);
} 