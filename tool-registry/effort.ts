import {
  ListCurrentUserEffortRecordsSchema,
  ListEffortRecordsSchema,
  CreateEffortRecordSchema,
  ListEstimatedEffortsSchema,
  CreateEstimatedEffortSchema,
  UpdateEffortRecordSchema,
  UpdateEstimatedEffortSchema
} from "../common/types.js";

export const listCurrentUserEffortRecordsTool = {
  name: "list_current_user_effort_records",
  description: "获取用户的实际工时明细，结束时间和开始时间的间隔不能大于6个月",
  parameters: ListCurrentUserEffortRecordsSchema,
};

export const listEffortRecordsTool = {
  name: "list_effort_records",
  description: "获取实际工时明细",
  parameters: ListEffortRecordsSchema,
};

export const createEffortRecordTool = {
  name: "create_effort_record",
  description: "登记实际工时",
  parameters: CreateEffortRecordSchema,
};

export const listEstimatedEffortsTool = {
  name: "list_estimated_efforts",
  description: "获取预计工时明细",
  parameters: ListEstimatedEffortsSchema,
};

export const createEstimatedEffortTool = {
  name: "create_estimated_effort",
  description: "登记预计工时",
  parameters: CreateEstimatedEffortSchema,
};

export const updateEffortRecordTool = {
  name: "update_effort_record",
  description: "更新登记实际工时",
  parameters: UpdateEffortRecordSchema,
};

export const updateEstimatedEffortTool = {
  name: "update_estimated_effort",
  description: "更新登记预计工时",
  parameters: UpdateEstimatedEffortSchema,
};

export const effortTools = [
  listCurrentUserEffortRecordsTool,
  listEffortRecordsTool,
  createEffortRecordTool,
  listEstimatedEffortsTool,
  createEstimatedEffortTool,
  updateEffortRecordTool,
  updateEstimatedEffortTool,
];