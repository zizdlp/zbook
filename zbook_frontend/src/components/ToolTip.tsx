"use client";
import React, { ReactNode, useState } from "react";

export default function ToolTip({
  message,
  children,
}: {
  message: string;
  children: ReactNode;
}) {
  const [show, setShow] = useState(false);
  return (
    <div className="relative flex flex-col items-center group">
      <div
        className="flex justify-center"
        onMouseEnter={() => setShow(true)}
        onMouseLeave={() => setShow(false)}
      >
        {children}
      </div>
      {show && (
        <div
          className={`absolute whitespace-nowrap bottom-full flex flex-col items-center group-hover:flex`}
        >
          <span className="relative z-10 p-2 text-xs leading-none text-white whitespace-no-wrap bg-gray-600 shadow-lg rounded-md">
            {message}
          </span>
          <div className="w-3 h-3 -mt-2 rotate-45 bg-gray-600" />
        </div>
      )}
    </div>
  );
}
