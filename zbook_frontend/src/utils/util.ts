import { toast } from "react-toastify";
import { logger } from "./logger";
import { BsFillBookmarkCheckFill } from "react-icons/bs";
import { TiWarning } from "react-icons/ti";
import { FaInfoCircle } from "react-icons/fa";
import { MdError, MdTipsAndUpdates } from "react-icons/md";
import { IconType } from "react-icons/lib";
export function fileToBase64(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      const result = reader.result as string; // 将结果断言为 string 类型
      resolve(result.split(",")[1]); // 获取base64字符串
    };
    reader.onerror = reject;
    reader.readAsDataURL(file);
  });
}

export function getCurrentDateTime() {
  const now = new Date();

  const year = now.getFullYear();
  const month = padZero(now.getMonth() + 1); // months are zero-based
  const day = padZero(now.getDate());
  const hours = padZero(now.getHours());
  const minutes = padZero(now.getMinutes());
  const seconds = padZero(now.getSeconds());

  return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}`;
}

function padZero(num: number) {
  return num < 10 ? `0${num}` : num;
}

export function isContains(str: string | undefined, substr: string) {
  if (str == undefined) {
    return false;
  }
  return str.indexOf(substr) >= 0;
}

export function joinUrlParts(...parts: string[]): string {
  // 使用 map 方法对每个部分进行处理
  const processedParts = parts.map((part) =>
    part.trim().replace(/^\/|\/$/g, "")
  ); // 去除每个部分的前导和尾随斜杠

  // 使用 join 方法将处理后的部分拼接成一个字符串
  const url = processedParts.join("/");
  return url;
}

export function generateRandomString(length: number): string {
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  let result = "";
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  return result;
}

export function isPrefixWithOutLocal(
  pathname: string,
  locale: string,
  prefix: string,
  relative_path: string
): boolean {
  let folderPath = prefix + relative_path;
  const folderPathSegments = folderPath
    .split("/")
    .filter((segment) => segment !== "");
  const pathnameSegments = decodeURIComponent(pathname)
    .split("/")
    .filter((segment) => segment !== "");
  if (pathnameSegments.length < folderPathSegments.length) {
    // 如果 pathname 的段数比 folderPath 的段数少，或者对应的段不一致，则不是前缀
    return false;
  }
  for (let i = 0; i < folderPathSegments.length; i++) {
    if (folderPathSegments[i] !== pathnameSegments[i]) {
      return false;
    }
  }
  return true;
}
export function isPrefix(
  pathname: string,
  locale: string,
  prefix: string,
  relative_path: string
): boolean {
  if (isPrefixWithOutLocal(pathname, locale, prefix, relative_path)) {
    return true;
  }
  let folderPath = "/" + locale + prefix + relative_path;
  const folderPathSegments = folderPath
    .split("/")
    .filter((segment) => segment !== "");
  const pathnameSegments = decodeURIComponent(pathname)
    .split("/")
    .filter((segment) => segment !== "");
  if (pathnameSegments.length < folderPathSegments.length) {
    // 如果 pathname 的段数比 folderPath 的段数少，或者对应的段不一致，则不是前缀
    return false;
  }
  for (let i = 0; i < folderPathSegments.length; i++) {
    if (folderPathSegments[i] !== pathnameSegments[i]) {
      return false;
    }
  }
  return true;
}

export function isSameUrl(
  pathname: string,
  locale: string,
  prefix: string,
  relative_path: string
) {
  return (
    decodeURIComponent(pathname) === "/" + locale + prefix + relative_path ||
    decodeURIComponent(pathname) === prefix + relative_path
  );
}
export function toastInfo(id: any, message: any) {
  toast.update(id, {
    render: message,
    type: "success",
    isLoading: false,
    autoClose: 500,
  });
}
export function toastError(id: any, message: any) {
  toast.update(id, {
    render: message,
    type: "error",
    isLoading: false,
    autoClose: 1500,
  });
}
export function parseUserAgent(userAgent: string) {
  let uaData = {
    platform: "Unknown",
    os: "Unknown",
    browser: "Unknown",
    browserVersion: "Unknown",
    engine: "Unknown",
    engineVersion: "Unknown",
  };
  // 确保 userAgent 是一个非空字符串
  if (!userAgent || typeof userAgent !== "string") {
    return uaData;
  }

  try {
    // Extract platform and OS
    const platformMatch = userAgent.match(/\(([^)]+)\)/);
    if (platformMatch) {
      const platformParts = platformMatch[1].split(";");
      if (platformParts.length > 0) {
        uaData.platform = platformParts[0].trim();
      }
      if (platformParts.length > 1) {
        uaData.os = platformParts[1].trim();
      }
    }

    // Extract browser and version
    const browserRegexes = [
      { regex: /(Chrome)\/([0-9.]+)/, name: "Chrome" },
      { regex: /(Safari)\/([0-9.]+)/, name: "Safari" },
      { regex: /(Firefox)\/([0-9.]+)/, name: "Firefox" },
      { regex: /(Edg)\/([0-9.]+)/, name: "Edge" },
      { regex: /(OPR)\/([0-9.]+)/, name: "Opera" },
      { regex: /(MSIE) ([0-9.]+)/, name: "Internet Explorer" },
      { regex: /(Trident)\/.*rv:([0-9.]+)/, name: "Internet Explorer" },
    ];

    for (const { regex, name } of browserRegexes) {
      const match = userAgent.match(regex);
      if (match && match[2]) {
        uaData.browser = name;
        uaData.browserVersion = match[2];
        break;
      }
    }

    // Extract rendering engine and version
    const engineRegexes = [
      { regex: /(AppleWebKit)\/([0-9.]+)/, name: "AppleWebKit" },
      { regex: /(Gecko)\/([0-9.]+)/, name: "Gecko" },
      { regex: /(KHTML)\/([0-9.]+)/, name: "KHTML" },
      { regex: /(Trident)\/([0-9.]+)/, name: "Trident" },
    ];

    for (const { regex, name } of engineRegexes) {
      const match = userAgent.match(regex);
      if (match && match[2]) {
        uaData.engine = name;
        uaData.engineVersion = match[2];
        break;
      }
    }
  } catch (error) {
    logger.error(`Error parsing userAgent:${error}`);
    return uaData;
  }
  return uaData;
}

export function getLayoutForLocale(jsonConfig: any, locale: any) {
  const layout = jsonConfig.layout;

  // 尝试获取指定语言的布局
  let langLayout = layout[locale];

  // 如果指定语言的布局不存在，尝试获取 "default" 布局
  if (!langLayout) {
    langLayout = layout["default"];
  }

  // 如果 "default" 布局也不存在，返回一个空布局
  if (!langLayout) {
    return [];
  }

  return langLayout;
}
const bgColorsMap = new Map<string, string>([
  ["note", "cyan"],
  ["warning", "yellow"],
  ["info", "green"],
  ["tip", "lime"],
  ["error", "red"],
]);

const iconTypes = {
  note: BsFillBookmarkCheckFill,
  warning: TiWarning,
  info: FaInfoCircle,
  tip: MdTipsAndUpdates,
  error: MdError,
};

export function getAdmonitionBackground(admonitionType: string): string {
  const color = bgColorsMap.get(admonitionType);
  if (!color) {
    console.warn(`Unknown admonition type: ${admonitionType}`);
    return "bg-gray-600/75 dark:bg-gray-500/25"; // 默认颜色
  }
  return `bg-${color}-600/75 dark:bg-${color}-500/25`;
}

type ParentType = "note" | "warning" | "info" | "tip" | "error";

export function getClientAdmonitionType(admtype: string): {
  bg: string;
  Icon: IconType;
} {
  let type: ParentType = "note"; // 默认类型为 "note"
  if (admtype == "note") type = "note";
  else if (admtype == "warning") type = "warning";
  else if (admtype == "info") type = "info";
  else if (admtype == "tip") type = "tip";
  else if (admtype == "error") type = "error";

  return {
    bg: getAdmonitionBackground(type),
    Icon: iconTypes[type],
  };
}
