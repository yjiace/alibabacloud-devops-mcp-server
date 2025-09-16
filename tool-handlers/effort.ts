import {
  ListCurrentUserEffortRecordsSchema,
  ListEffortRecordsSchema,
  CreateEffortRecordSchema,
  ListEstimatedEffortsSchema,
  CreateEstimatedEffortSchema,
  UpdateEffortRecordSchema,
  UpdateEstimatedEffortSchema
} from "../common/types.js";
import {
  listCurrentUserEffortRecords,
  listEffortRecords,
  createEffortRecord,
  listEstimatedEfforts,
  createEstimatedEffort,
  updateEffortRecord,
  updateEstimatedEffort
} from "../operations/projex/effort.js";

export async function listCurrentUserEffortRecordsHandler(
  params: unknown
) {
  const validatedParams = ListCurrentUserEffortRecordsSchema.parse(params);
  return listCurrentUserEffortRecords(validatedParams);
}

export async function listEffortRecordsHandler(
  params: unknown
) {
  const validatedParams = ListEffortRecordsSchema.parse(params);
  return listEffortRecords(validatedParams);
}

export async function createEffortRecordHandler(
  params: unknown
) {
  const validatedParams = CreateEffortRecordSchema.parse(params);
  return createEffortRecord(validatedParams);
}

export async function listEstimatedEffortsHandler(
  params: unknown
) {
  const validatedParams = ListEstimatedEffortsSchema.parse(params);
  return listEstimatedEfforts(validatedParams);
}

export async function createEstimatedEffortHandler(
  params: unknown
) {
  const validatedParams = CreateEstimatedEffortSchema.parse(params);
  return createEstimatedEffort(validatedParams);
}

export async function updateEffortRecordHandler(
  params: unknown
) {
  const validatedParams = UpdateEffortRecordSchema.parse(params);
  return updateEffortRecord(validatedParams);
}

export async function updateEstimatedEffortHandler(
  params: unknown
) {
  const validatedParams = UpdateEstimatedEffortSchema.parse(params);
  return updateEstimatedEffort(validatedParams);
}