"use client";
import React, { useState } from "react";
import { CiImageOff, CiImageOn } from "react-icons/ci"; // 确保你已安装 react-icons 并导入 CiImageOff

const ImageWithFallback = ({ src, alt }: { src: string; alt: string }) => {
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
      <img
        className="w-full rounded-md border"
        src={src}
        alt={alt}
        style={{ display: imgLoaded && !imgError ? "block" : "none" }}
        onLoad={() => setImgLoaded(true)}
        onError={() => setImgError(true)}
      />
    </>
  );
};

export default ImageWithFallback;
