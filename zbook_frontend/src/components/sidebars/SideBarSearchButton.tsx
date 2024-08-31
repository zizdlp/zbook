"use client";

import React, { useContext } from "react";
import { useTranslations } from "next-intl";
import { SearchDialogContext } from "@/providers/SearchDialogProvider";
import { OperationContext } from "@/providers/OperationProvider";
import { SearchType } from "@/utils/const_value";
export default function SideBarSearchButton({
  username,
  repo_name,
  searchType,
}: {
  username: string;
  repo_name: string;
  searchType: SearchType;
}) {
  const { searchDialogOpen, setSearchDialogOpen, setSearchType } =
    useContext(SearchDialogContext);
  const { setOperationRepoName, setOperationUsername } =
    useContext(OperationContext);
  const t = useTranslations("SideBar");
  return (
    <div className="sticky top-0 pointer-events-none z-50 px-4">
      <div className="h-10 bg-white dark:bg-slate-900"></div>
      <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
        <button
          onClick={() => {
            setOperationUsername(username);
            setOperationRepoName(decodeURIComponent(repo_name));
            setSearchType(searchType);
            setSearchDialogOpen(!searchDialogOpen);
          }}
          type="button"
          className="flex w-full items-center text-sm leading-6 text-slate-400 rounded-md ring-none border border-slate-300/75 dark:border-0 
                py-3 px-3 dark:bg-slate-800 dark:highlight-white/5 dark:hover:bg-slate-700 h-12"
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
      </div>
      <div className="h-4 bg-gradient-to-b from-white dark:from-slate-900"></div>
    </div>
  );
}
