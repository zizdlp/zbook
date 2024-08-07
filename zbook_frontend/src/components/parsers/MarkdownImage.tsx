import Image from "next/image";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { CiImageOff } from "react-icons/ci";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";

class LRUCache {
  private capacity: number;
  private cache: Map<string, { value: string; timestamp: number }>;

  constructor(capacity: number) {
    this.capacity = capacity;
    this.cache = new Map();
  }

  get(key: string): string | undefined {
    if (!this.cache.has(key)) return undefined;

    const { value, timestamp } = this.cache.get(key)!;
    // Check if the cached value is expired (more than 60 seconds)
    if (Date.now() - timestamp > 60 * 1000) {
      this.cache.delete(key);
      return undefined;
    }

    // Update the timestamp to indicate recent use
    this.cache.delete(key);
    this.cache.set(key, { value, timestamp: Date.now() });

    return value;
  }

  put(key: string, value: string): void {
    // Remove the least recently used item if cache is full
    if (this.cache.size === this.capacity) {
      const leastUsedKey = this.cache.keys().next().value;
      this.cache.delete(leastUsedKey);
    }

    // Add new item with current timestamp
    this.cache.set(key, { value, timestamp: Date.now() });
  }
}
const lruCache = new LRUCache(300); // 设定容量为300
export default async function MarkdownImage({
  path,
  username,
  repo_name,
}: {
  path: string;
  username: string;
  repo_name: string;
}) {
  try {
    // 检查内存缓存
    if (lruCache.get(path)) {
      return renderImage(lruCache.get(path)!);
    }
    const xforward = headers().get("x-forwarded-for") ?? "";
    const agent = headers().get("User-Agent") ?? "";
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_MARKDOWN_IMAGE,
      xforward: xforward,
      agent: agent,
      tags: [],
      values: {
        file_path: decodeURIComponent(path),
        username: username,
        repo_name: decodeURIComponent(repo_name),
      },
    });
    if (data.error) {
      throw new FetchError(data.message, data.status);
    }

    const pathExtension = path.slice(-4).toLowerCase();
    const fourPathExtension = path.slice(-5).toLowerCase();

    if (isSupportedImageExtension(pathExtension, fourPathExtension)) {
      const file_type = getFileType(pathExtension, fourPathExtension);
      const avatar = `data:image/${file_type};base64,${data.file}`;
      lruCache.put(path, avatar);
      return renderImage(avatar);
    } else {
      logger.error(`Unsupported image file type:${pathExtension}`);
      throw new Error();
    }
  } catch (error) {
    return (
      <CiImageOff className="rounded-md my-[1.25em] w-full h-96 py-40 bg-gray-200 dark:bg-gray-700/75 text-slate-500 dark:text-slate-400" />
    );
  }
}

// 辅助函数，用于渲染图片
function renderImage(imageSrc: string) {
  return (
    <Image
      src={imageSrc}
      className="rounded-md w-full my-[2em]"
      alt="image"
      width={256}
      height={256}
    />
  );
}

// 辅助函数，用于检查支持的图片后缀
function isSupportedImageExtension(
  pathExtension: string,
  fourPathExtension: string
) {
  const imageExtensions = [
    ".svg",
    ".png",
    ".jpg",
    ".jpeg",
    ".gif",
    ".bmp",
    ".webp",
  ];

  return (
    imageExtensions.includes(pathExtension) ||
    imageExtensions.includes(fourPathExtension)
  );
}

// 辅助函数，用于获取文件类型
function getFileType(pathExtension: string, fourPathExtension: string) {
  if (pathExtension === ".png") {
    return "png";
  } else if (pathExtension === ".svg") {
    return "svg+xml";
  } else if (pathExtension === ".jpg" || fourPathExtension === ".jpeg") {
    return "jpeg";
  } else if (pathExtension === ".gif") {
    return "gif";
  } else if (pathExtension === ".bmp") {
    return "bmp";
  } else if (fourPathExtension === ".webp") {
    return "webp";
  }

  // 默认返回 jpg，可以根据需要调整
  return "jpeg";
}
