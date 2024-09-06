"use client";
import React, { useState } from "react";
import { FaVideo, FaVideoSlash } from "react-icons/fa";

const VideoWithFallBack = ({ src, alt }: { src: string; alt: string }) => {
  const [imgLoaded, setImgLoaded] = useState(false);
  const [imgError, setImgError] = useState(false);

  return (
    <>
      {!imgLoaded && !imgError && (
        <FaVideo className="w-full my-[1.25em] h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 animate-pulse text-slate-500 dark:text-slate-400" />
      )}
      {imgError && (
        <FaVideoSlash className="w-full my-[1.25em] h-96 rounded-md py-40 bg-gray-200 dark:bg-gray-700/75 text-slate-500 dark:text-slate-400" />
      )}
      <iframe
        className="w-full embed-video my-[1.25em] rounded-md"
        src={src}
        style={{ display: imgLoaded && !imgError ? "block" : "none" }}
        onLoad={() => setImgLoaded(true)}
        onError={() => setImgError(true)}
      />
    </>
  );
};

export default VideoWithFallBack;
