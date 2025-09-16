import { z } from "zod";
import { handleApiRequest } from "../../common/utils.js";
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
  
  return handleApiRequest({
    path: `projex/organizations/${validatedParams.organizationId}/effortRecords`,
    method: "GET",
    queryParams: {
      startDate: validatedParams.startDate,
      endDate: validatedParams.endDate
    },
    responseSchema: z.array(EffortRecordSchema)
  });
}

// List effort records
export async function listEffortRecords(
  params: z.infer<typeof ListEffortRecordsSchema>
) {
  const validatedParams = ListEffortRecordsSchema.parse(params);
  
  return handleApiRequest({
    path: `projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/effortRecords`,
    method: "GET",
    responseSchema: z.array(EffortRecordSchema)
  });
}

// Create effort record
export async function createEffortRecord(
  params: z.infer<typeof CreateEffortRecordSchema>
) {
  const { request, ...validatedParams } = CreateEffortRecordSchema.parse(params);
  
  return handleApiRequest({
    path: `projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/effortRecords`,
    method: "POST",
    body: request,
    responseSchema: IdentifierDTOSchema
  });
}

// List estimated efforts
export async function listEstimatedEfforts(
  params: z.infer<typeof ListEstimatedEffortsSchema>
) {
  const validatedParams = ListEstimatedEffortsSchema.parse(params);
  
  return handleApiRequest({
    path: `projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/estimatedEfforts`,
    method: "GET",
    responseSchema: z.array(EstimatedEffortSchema)
  });
}

// Create estimated effort
export async function createEstimatedEffort(
  params: z.infer<typeof CreateEstimatedEffortSchema>
) {
  const { request, ...validatedParams } = CreateEstimatedEffortSchema.parse(params);
  
  return handleApiRequest({
    path: `projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.id}/estimatedEfforts`,
    method: "POST",
    body: request,
    responseSchema: IdentifierDTOSchema
  });
}

// Update effort record
export async function updateEffortRecord(
  params: z.infer<typeof UpdateEffortRecordSchema>
) {
  const { request, ...validatedParams } = UpdateEffortRecordSchema.parse(params);
  
  return handleApiRequest({
    path: `projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.workitemId}/effortRecords/${validatedParams.id}`,
    method: "PUT",
    body: request,
    responseSchema: z.any()
  });
}

// Update estimated effort
export async function updateEstimatedEffort(
  params: z.infer<typeof UpdateEstimatedEffortSchema>
) {
  const { request, ...validatedParams } = UpdateEstimatedEffortSchema.parse(params);
  
  return handleApiRequest({
    path: `projex/organizations/${validatedParams.organizationId}/workitems/${validatedParams.workitemId}/estimatedEfforts/${validatedParams.id}`,
    method: "PUT",
    body: request,
    responseSchema: z.any()
  });
}