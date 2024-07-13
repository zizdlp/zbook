import React from "react";
import { ReactNode } from "react";

export default function SideBarLiContent({
  children,
  isSelected,
}: {
  isSelected: boolean;
  children: ReactNode;
}) {
  return (
    <div
      className={`flex items-center p-3 text-base font-bold text-gray-900 rounded-md  border-[0.01rem] dark:border-0 
    hover:bg-sky-400 hover:text-white dark:hover:bg-sky-900/75 dark:text-white cursor-pointer ${
      isSelected
        ? "bg-sky-500/75 text-white dark:bg-sky-800"
        : "bg-gray-50 dark:bg-slate-800/50"
    }`}
    >
      {children}
    </div>
  );
}
