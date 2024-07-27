"use client";

import React, { useContext } from "react";
import { useTranslations } from "next-intl";
import {
  SearchDialogContext,
  SearchType,
} from "@/providers/SearchDialogProvider";
import { OperationContext } from "@/providers/OperationProvider";
export default function SearchUserButton({ username }: { username: string }) {
  const { searchDialogOpen, setSearchDialogOpen, setSearchType } =
    useContext(SearchDialogContext);
  const { setOperationRepoID, setOperationUsername } =
    useContext(OperationContext);
  const t = useTranslations("SideBar");
  return (
    <button
      onClick={() => {
        setOperationUsername(username);
        setSearchType(SearchType.USER_DOCUMENT); //搜索当前用户的仓库
        setSearchDialogOpen(!searchDialogOpen);
      }}
      type="button"
      className="flex w-full items-center text-sm leading-6 text-slate-400 rounded-md ring-none border-[0.1rem] border-slate-200 dark:border-0
                py-1.5 pl-2 pr-3 dark:bg-slate-800 dark:highlight-white/5 dark:hover:bg-slate-700 h-12"
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
