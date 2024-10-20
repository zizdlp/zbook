"use client";
import React, { useState } from "react";
import { CiImageOff, CiImageOn } from "react-icons/ci"; // 确保你已安装 react-icons 并导入 CiImageOff

const ImageWithFallback = ({
  src,
  alt,
  show_caption,
}: {
  src: string;
  alt: string;
  show_caption: boolean;
}) => {
  const [imgLoaded, setImgLoaded] = useState(false);
  const [imgError, setImgError] = useState(false);
  return (
    <>
      {!imgLoaded && !imgError && (
        <CiImageOn className="w-full my-[1.25em] h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 animate-pulse text-slate-500 dark:text-slate-400" />
      )}
      {imgError && (
        <CiImageOff className="w-full my-[1.25em] h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 text-slate-500 dark:text-slate-400" />
      )}
      <div className="not-prose">
        <img
          className="w-full rounded-md border"
          src={src}
          alt={alt}
          style={{ display: imgLoaded && !imgError ? "block" : "none" }}
          onLoad={() => setImgLoaded(true)}
          onError={() => setImgError(true)}
        />
        {alt != "" && show_caption && (
          <p className="font-medium text-sm italic text-slate-500 mb-[1.25em] font-mono text-center dark:text-slate-400">
            {alt}
          </p>
        )}
      </div>
    </>
  );
};

export default ImageWithFallback;
