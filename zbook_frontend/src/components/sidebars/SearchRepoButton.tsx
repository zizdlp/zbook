"use client";

import React, { useContext } from "react";
import { SearchDialogContext } from "@/providers/SearchDialogProvider";
import { OperationContext } from "@/providers/OperationProvider";
import { useTranslations } from "next-intl";
export default function SearchRepoButton({ repo_id }: { repo_id: number }) {
  const { searchDialogOpen, setSearchDialogOpen, setSearchType } =
    useContext(SearchDialogContext);
  const { setOperationRepoID } = useContext(OperationContext);
  const t = useTranslations("SideBar");
  return (
    <button
      onClick={() => {
        setOperationRepoID(repo_id);
        setSearchType(4); //仅展示搜索当前仓库
        setSearchDialogOpen(!searchDialogOpen);
      }}
      type="button"
      className="flex w-full items-center text-sm leading-6 text-slate-400 rounded-md ring-1 
                ring-slate-900/10 shadow-sm py-1.5 pl-2 pr-3 hover:ring-slate-300 dark:hover:ring-slate-700 
                dark:bg-slate-800 dark:highlight-white/5 dark:hover:bg-slate-700 h-10"
    >
      <svg
        width="24"
        height="24"
        fill="none"
        aria-hidden="true"
        className="mr-3 flex-none"
      >
        <path
          d="m19 19-3.5-3.5"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        ></path>
        <circle
          cx="11"
          cy="11"
          r="6"
          stroke="currentColor"
          strokeWidth="2"
          strokeLinecap="round"
          strokeLinejoin="round"
        ></circle>
      </svg>
      {t("Search")}
      <span className="ml-auto pl-3 flex-none text-xs font-semibold">
        {t("SearchShortcuts")}
      </span>
    </button>
  );
}
