import { z } from "zod";
import { yunxiaoRequest } from "../../common/utils.js";
import {
  EffortRecordSchema,
  EstimatedEffortSchema,
  IdentifierDTOSchema,
  ListCurrentUserEffortRecordsSchema,
  ListEffortRecordsSchema,
  CreateEffortRecordSchema,
  ListEstimatedEffortsSchema,
  CreateEstimatedEffortSchema,
  UpdateEffortRecordSchema,
  UpdateEstimatedEffortSchema
} from "./types.js";

// List current user effort records
export async function listCurrentUserEffortRecords(
  params: z.infer<typeof ListCurrentUserEffortRecordsSchema>
) {
  const validatedParams = ListCurrentUserEffortRecordsSchema.parse(params);
  
  const url = `/oapi/v1/projex/organizations/${validatedParams.organizationId}/effortRecords`;
  const queryParams = {
    startDate: validatedParams.startDate,
    endDate: validatedParams.endDate
  };
  
  const response = await yunxiaoRequest(`${url}?startDate=${queryParams.startDate}&endDate=${queryParams.endDate}`, {
    method: "GET"
  });

  return z.array(EffortRecordSchema).parse(response);
}

// List effort records
export async function listEffortRecords(
  params: z.infer<typeof ListEffortRecordsSchema>
) {
  const validatedParams = ListEffortRecordsSchema.parse(params);
  
  const url = `/oapi/v1/projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/effortRecords`;
  
  const response = await yunxiaoRequest(url, {
    method: "GET"
  });

  return z.array(EffortRecordSchema).parse(response);
}

// Create effort record
export async function createEffortRecord(
  params: z.infer<typeof CreateEffortRecordSchema>
) {
  const { request, ...validatedParams } = CreateEffortRecordSchema.parse(params);
  
  const url = `/oapi/v1/projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/effortRecords`;
  
  const response = await yunxiaoRequest(url, {
    method: "POST",
    body: request
  });

  return IdentifierDTOSchema.parse(response);
}

// List estimated efforts
export async function listEstimatedEfforts(
  params: z.infer<typeof ListEstimatedEffortsSchema>
) {
  const validatedParams = ListEstimatedEffortsSchema.parse(params);
  
  const url = `/oapi/v1/projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/estimatedEfforts`;
  
  const response = await yunxiaoRequest(url, {
    method: "GET"
  });

  return z.array(EstimatedEffortSchema).parse(response);
}

// Create estimated effort
export async function createEstimatedEffort(
  params: z.infer<typeof CreateEstimatedEffortSchema>
) {
  const { request, ...validatedParams } = CreateEstimatedEffortSchema.parse(params);
  
  const url = `/oapi/v1/projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/estimatedEfforts`;
  
  const response = await yunxiaoRequest(url, {
    method: "POST",
    body: request
  });

  return IdentifierDTOSchema.parse(response);
}

// Update effort record
export async function updateEffortRecord(
  params: z.infer<typeof UpdateEffortRecordSchema>
) {
  const { request, ...validatedParams } = UpdateEffortRecordSchema.parse(params);
  
  const url = `/oapi/v1/projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.workitemId}/effortRecords/${validatedParams.id}`;
  
  const response = await yunxiaoRequest(url, {
    method: "PUT",
    body: request
  });

  return response;
}

// Update estimated effort
export async function updateEstimatedEffort(
  params: z.infer<typeof UpdateEstimatedEffortSchema>
) {
  const { request, ...validatedParams } = UpdateEstimatedEffortSchema.parse(params);
  
  const url = `/oapi/v1/projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.workitemId}/estimatedEfforts/${validatedParams.id}`;
  
  const response = await yunxiaoRequest(url, {
    method: "PUT",
    body: request
  });

  return response;
}