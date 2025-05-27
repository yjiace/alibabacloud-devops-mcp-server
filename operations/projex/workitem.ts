import { z } from "zod";
import { yunxiaoRequest } from "../../common/utils.js";
import {
  WorkItemSchema,
  FilterConditionSchema,
  ConditionsSchema
} from "../../common/types.js";

export async function getWorkItemFunc(
  organizationId: string,
  workItemId: string
): Promise<z.infer<typeof WorkItemSchema>> {
  const url = `/oapi/v1/projex/organizations/${organizationId}/workitems/${workItemId}`;

  const response = await yunxiaoRequest(url, {
    method: "GET",
  });

  return WorkItemSchema.parse(response);
}

export async function searchWorkitemsFunc(
  organizationId: string,
  category: string,
  spaceId: string,
  subject?: string,
  status?: string,
  createdAfter?: string,
  createdBefore?: string,
  updatedAfter?: string,
  updatedBefore?: string,
  creator?: string,
  assignedTo?: string,
  advancedConditions?: string,
  orderBy: string = "gmtCreate", // Possible values: gmtCreate, subject, status, priority, assignedTo
  includeDetails: boolean = false // 新增参数：是否自动补充缺失的description等详细信息
): Promise<z.infer<typeof WorkItemSchema>[]> {
  const url = `/oapi/v1/projex/organizations/${organizationId}/workitems:search`;

  // Prepare payload
  const payload: Record<string, any> = {
    category: category,
    spaceId: spaceId,
  };

  // Process condition parameters
  const conditions = buildWorkitemConditions({
    subject,
    status,
    createdAfter,
    createdBefore,
    updatedAfter,
    updatedBefore,
    creator,
    assignedTo,
    advancedConditions
  });
  
  if (conditions) {
    payload.conditions = conditions;
  }

  // Add orderBy parameter
  payload.orderBy = orderBy;

  const response = await yunxiaoRequest(url, {
    method: "POST",
    body: payload,
  });

  // Ensure response is an array
  if (!Array.isArray(response)) {
    return [];
  }

  // Parse each work item object first
  const workItems = response.map(workitem => WorkItemSchema.parse(workitem));

  // 如果需要补充详细信息，使用分批并发方式获取
  if (includeDetails) {
    // 找出description为null的工作项，同时确保id是有效的字符串
    const itemsNeedingDetails = workItems.filter(item => 
      typeof item.id === 'string' && item.id.length > 0 && 
      (item.description === null || item.description === undefined)
    );

    if (itemsNeedingDetails.length > 0) {
      // 分批并发获取详情
      const descriptionMap = await batchGetWorkItemDetails(organizationId, itemsNeedingDetails);

      // 更新workItems中的description
      return workItems.map(item => {
        if (typeof item.id === 'string' && descriptionMap.has(item.id)) {
          return {
            ...item,
            description: descriptionMap.get(item.id) || item.description
          };
        }
        return item;
      });
    }
  }

  return workItems;
}

// 分批并发获取工作项详情
async function batchGetWorkItemDetails(
  organizationId: string, 
  workItems: z.infer<typeof WorkItemSchema>[],
  batchSize: number = 10,  // 每批处理10个
  maxItems: number = 100   // 最多处理100个
): Promise<Map<string, string | null>> {
  const descriptionMap = new Map<string, string | null>();
  
  // 限制处理数量
  const limitedItems = workItems.slice(0, maxItems);
  
  if (workItems.length > maxItems) {
    console.log(`工作项数量较多，只处理前 ${maxItems} 个`);
  }

  // 分批处理
  for (let i = 0; i < limitedItems.length; i += batchSize) {
    const batch = limitedItems.slice(i, i + batchSize);
    
    // 批次内并发执行
    const batchResults = await Promise.allSettled(
      batch.map(async (item) => {
        // 再次检查item.id是否为有效字符串
        if (typeof item.id !== 'string' || item.id.length === 0) {
          return { 
            id: item.id || 'unknown', 
            description: null,
            success: false 
          };
        }
        
        const itemId: string = item.id;
        
        try {
          const detailedItem = await getWorkItemFunc(organizationId, itemId);
          return { 
            id: itemId, 
            description: detailedItem.description,
            success: true 
          };
        } catch (error) {
          console.warn(`获取工作项 ${itemId} 详情失败:`, error);
          return { 
            id: itemId, 
            description: null,
            success: false 
          };
        }
      })
    );

    // 处理批次结果
    batchResults.forEach((result) => {
      if (result.status === 'fulfilled') {
        // 确保description类型正确，将undefined转换为null
        const description = result.value.description === undefined ? null : result.value.description;
        descriptionMap.set(result.value.id, description);
      }
    });
  }
  return descriptionMap;
}

// Build work item search conditions
function buildWorkitemConditions(args: {
  subject?: string;
  status?: string;
  createdAfter?: string;
  createdBefore?: string;
  updatedAfter?: string;
  updatedBefore?: string;
  creator?: string;
  assignedTo?: string;
  advancedConditions?: string;
}): string | undefined {
  // If advanced conditions are provided directly, use them preferentially
  if (args.advancedConditions) {
    return args.advancedConditions;
  }

  // Build condition group
  const filterConditions: z.infer<typeof FilterConditionSchema>[] = [];

  // Process title
  if (args.subject) {
    filterConditions.push({
      className: "string",
      fieldIdentifier: "subject",
      format: "input",
      operator: "CONTAINS",
      toValue: null,
      value: [args.subject],
    });
  }

  // Process status
  if (args.status) {
    const statusValues = args.status.split(",");
    const values = statusValues.map(v => v.trim());

    filterConditions.push({
      className: "status",
      fieldIdentifier: "status",
      format: "list",
      operator: "CONTAINS",
      toValue: null,
      value: values,
    });
  }

  // Process creation time range
  if (args.createdAfter) {
    const createdBefore = args.createdBefore ? `${args.createdBefore} 23:59:59` : null;

    filterConditions.push({
      className: "dateTime",
      fieldIdentifier: "gmtCreate",
      format: "input",
      operator: "BETWEEN",
      toValue: createdBefore,
      value: [`${args.createdAfter} 00:00:00`],
    });
  }

  //process updated time range
  if (args.updatedAfter) {
    const updatedBefore = args.updatedBefore ? `${args.updatedBefore} 23:59:59` : null;

    filterConditions.push({
      className: "dateTime",
      fieldIdentifier: "gmtModified",
      format: "input",
      operator: "BETWEEN",
      toValue: updatedBefore,
      value: [`${args.updatedAfter} 00:00:00`],
    })
  }

  // Process creator
  if (args.creator) {
    const creatorValues = args.creator.split(",");
    const values = creatorValues.map(v => v.trim());

    filterConditions.push({
      className: "user",
      fieldIdentifier: "creator",
      format: "list",
      operator: "CONTAINS",
      toValue: null,
      value: values,
    });
  }

  // Process assignee
  if (args.assignedTo) {
    const assignedToValues = args.assignedTo.split(",");
    const values = assignedToValues.map(v => v.trim());

    filterConditions.push({
      className: "user",
      fieldIdentifier: "assignedTo",
      format: "list",
      operator: "CONTAINS",
      toValue: null,
      value: values,
    });
  }

  // If there are no conditions, return undefined
  if (filterConditions.length === 0) {
    return undefined;
  }

  // Build complete condition object
  const conditions: z.infer<typeof ConditionsSchema> = {
    conditionGroups: [filterConditions],
  };

  // Serialize to JSON
  return JSON.stringify(conditions);
} 