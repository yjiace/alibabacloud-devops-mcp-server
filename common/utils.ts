import { getUserAgent } from "universal-user-agent";
import { createYunxiaoError } from "./errors.js";
import { VERSION } from "./version.js";

const DEFAULT_YUNXIAO_API_BASE_URL = "https://openapi-rdc.aliyuncs.com";

/**
 * Get the Yunxiao API base URL from environment variables or use the default
 * @returns The Yunxiao API base URL
 */
export function getYunxiaoApiBaseUrl(): string {
  return process.env.YUNXIAO_API_BASE_URL || DEFAULT_YUNXIAO_API_BASE_URL;
}

type RequestOptions = {
  method?: string;
  body?: unknown;
  headers?: Record<string, string>;
}

export function debug(message: string, data?: unknown): void {
  if (data !== undefined) {
    console.error(`[DEBUG] ${message}`, typeof data === 'object' ? JSON.stringify(data, null, 2) : data);
  } else {
    console.error(`[DEBUG] ${message}`);
  }
}

async function parseResponseBody(response: Response): Promise<unknown> {
  const contentType = response.headers.get("content-type");
  if (contentType?.includes("application/json")) {
    return response.json();
  }
  return response.text();
}

export function buildUrl(baseUrl: string, params: Record<string, string | number | undefined>): string {
  // Handle baseUrl that doesn't have protocol
  const fullBaseUrl = baseUrl.startsWith('http')
    ? baseUrl
    : `${getYunxiaoApiBaseUrl()}${baseUrl.startsWith('/') ? baseUrl : `/${baseUrl}`}`;

  try {
    const url = new URL(fullBaseUrl);

    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        url.searchParams.append(key, value.toString());
      }
    });

    const result = url.toString();
    console.error(`[DEBUG] Final URL: ${result}`);

    // If we started with a relative URL, return just the path portion
    if (!baseUrl.startsWith('http')) {
      // Extract the path and query string from the full URL
      const urlObj = new URL(result);
      return urlObj.pathname + urlObj.search;
    }

    return result;
  } catch (error) {
    console.error(`[ERROR] Failed to build URL: ${error}`);

    // Fallback: manually append query parameters
    let urlWithParams = baseUrl;
    const queryParts: string[] = [];

    Object.entries(params).forEach(([key, value]) => {
      if (value !== undefined) {
        queryParts.push(`${encodeURIComponent(key)}=${encodeURIComponent(value.toString())}`);
      }
    });

    if (queryParts.length > 0) {
      urlWithParams += (urlWithParams.includes('?') ? '&' : '?') + queryParts.join('&');
    }

    console.error(`[DEBUG] Fallback URL: ${urlWithParams}`);
    return urlWithParams;
  }
}

const USER_AGENT = `modelcontextprotocol/servers/alibabacloud-devops-mcp-server/v${VERSION} ${getUserAgent()}`;

export async function yunxiaoRequest(
  urlPath: string,
  options: RequestOptions = {},
): Promise<unknown> {
  // Check if the URL is already a full URL or a path
  let url = urlPath.startsWith("http") ? urlPath : `${getYunxiaoApiBaseUrl()}${urlPath.startsWith("/") ? urlPath : `/${urlPath}`}`;
  const requestHeaders: Record<string, string> = {
    "Accept": "application/json",
    "Content-Type": "application/json",
    "User-Agent": USER_AGENT,
    ...options.headers,
  };

  if (process.env.YUNXIAO_ACCESS_TOKEN) {
    requestHeaders["x-yunxiao-token"] = process.env.YUNXIAO_ACCESS_TOKEN;
  }

  debug(`Request: ${options.method} ${url}`);
  debug(`Headers:`, requestHeaders);

  const response = await fetch(url, {
    method : options.method || "GET",
    headers: requestHeaders,
    body: options.body ? JSON.stringify(options.body) : undefined,
  } as RequestInit);

  const responseBody = await parseResponseBody(response);
  debug(`Response Body:`, responseBody)

  if (!response.ok) {
    throw createYunxiaoError(response.status, responseBody);
  }

  return responseBody;
}

export function pathEscape(filePath: string): string {
  // 先使用encodeURIComponent进行编码
  let encoded = encodeURIComponent(filePath);

  // 将编码后的%2F（/的编码）替换回/
  encoded = encoded.replace(/%2F/gi, "/");

  return encoded;
}

/**
 * Handle repository ID encoding
 * @param repositoryId Repository ID which may contain unencoded slash
 * @returns Properly encoded repository ID
 */
export function handleRepositoryIdEncoding(repositoryId: string): string {
  let encodedRepoId = repositoryId;

  // Automatically handle unencoded slashes in repositoryId
  if (repositoryId.includes("/")) {
    // Found unencoded slash, automatically URL encode it
    const parts = repositoryId.split("/", 2);
    if (parts.length === 2) {
      const encodedRepoName = encodeURIComponent(parts[1]);
      // Remove + signs from encoding (spaces are encoded as +, but we need %20)
      const formattedEncodedName = encodedRepoName.replace(/\+/g, "%20");
      encodedRepoId = `${parts[0]}%2F${formattedEncodedName}`;
    }
  }

  return encodedRepoId;
}

/**
 * Converts a floating point number to an integer string (removes decimal point and decimal part)
 * Used primarily for handling numeric IDs that might come as floats from JSON parsing
 * @param value Value to convert
 * @returns Integer string representation
 */
export function floatToIntString(value: any): string {
  // 如果传入的是字符串，先尝试转为浮点数
  if (typeof value === 'string') {
    const floatValue = parseFloat(value);
    if (!isNaN(floatValue)) {
      value = floatValue;
    } else {
      return value; // 如果转换失败，返回原字符串
    }
  }

  // 处理浮点数
  if (typeof value === 'number') {
    const intValue = Math.floor(value + 0.5); // 四舍五入转整数
    return intValue.toString();
  }

  // 处理其他情况，直接转字符串
  return String(value);
}


/**
 * 将各种时间格式转换为毫秒时间戳
 * 支持：
 * - 已有时间戳（number）直接返回
 * - Date对象转换为时间戳
 * - ISO格式日期字符串 (如: '2023-01-01T00:00:00Z')
 * - 日期字符串 (如: '2023-01-01')
 *
 * @param time 时间输入
 * @returns 毫秒时间戳
 */
export function convertToTimestamp(time: number | string | Date): number {
  if (typeof time === 'number') {
    // 如果已经是数字，假设已是时间戳
    return time;
  } else if (time instanceof Date) {
    // 如果是Date对象，转换为时间戳
    return time.getTime();
  } else if (typeof time === 'string') {
    // 尝试解析日期字符串
    const date = new Date(time);
    if (!isNaN(date.getTime())) {
      return date.getTime();
    }
  }

  // 无法转换时返回原值（如果是数字）或当前时间戳
  return typeof time === 'number' ? time : Date.now();
}
