"use client";

import Image from "next/image";
import { useEffect, useState } from "react";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { CiImageOff, CiImageOn } from "react-icons/ci";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";

export default function MarkdownImageClient({
  path,
  username,
  repo_name,
}: {
  path: string;
  username: string;
  repo_name: string;
}) {
  const [imageSrc, setImageSrc] = useState<string>();
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    async function fetchImage() {
      try {
        const data = await fetchServerWithAuthWrapper({
          endpoint: FetchServerWithAuthWrapperEndPoint.GET_MARKDOWN_IMAGE,
          xforward: "",
          agent: "",
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

          setImageSrc(avatar);
        } else {
          logger.error(`Unsupported image file type:${pathExtension}`);
          throw new Error("Unsupported image format");
        }
      } catch (error) {
        logger.error(`Failed to fetch markdown image: ${error}`);
        setImageSrc(undefined);
      } finally {
        setLoading(false);
      }
    }

    fetchImage();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  if (loading) {
    return (
      <CiImageOn className="w-full my-[1.25em] h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 animate-pulse text-slate-500 dark:text-slate-400" />
    );
  }

  if (imageSrc != undefined) {
    return (
      <Image
        src={imageSrc}
        className="rounded-md w-full my-[2em]"
        alt="image"
        width={256}
        height={256}
      />
    );
  } else {
    return (
      <CiImageOff className="rounded-md my-[1.25em] w-full h-96 py-40 bg-gray-200 dark:bg-gray-700/75 text-slate-500 dark:text-slate-400" />
    );
  }
}

// 辅助函数：检查支持的图片后缀
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

// 辅助函数：获取文件类型
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

  return "jpeg"; // 默认返回 jpg
}
