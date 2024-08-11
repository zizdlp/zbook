"use client";

import { MdSearch } from "react-icons/md";
import { useSearchParams } from "next/navigation";
import { usePathname, useRouter } from "@/navigation";
import { ListDataType } from "@/fetchs/model";
import { useTranslations } from "next-intl";
import { IoMdPersonAdd } from "react-icons/io";
import { useContext } from "react";
import { SearchDialogContext } from "@/providers/SearchDialogProvider";
import { OperationContext } from "@/providers/OperationProvider";
import { SearchType } from "@/utils/const_value";

export default function SearchList({
  listType,
  username,
  repo_name,
}: {
  listType: ListDataType;
  username: string;
  repo_name: string;
}) {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();
  const t = useTranslations("DataList");
  const { setSearchDialogOpen, setSearchType } =
    useContext(SearchDialogContext);
  const { setOperationUsername, setOperationRepoName } =
    useContext(OperationContext);
  function handleSearch(term: string) {
    if (searchParams) {
      const params = new URLSearchParams(searchParams);
      if (term) {
        params.set("query", term);
      } else {
        params.delete("query");
      }
      replace(`${pathname}?${params.toString()}`);
    }
  }

  let placeholder = "";
  switch (listType) {
    case ListDataType.LIST_USER_REPO:
      placeholder = t("SearchListUserRepoTip");
      break;
    case ListDataType.LIST_USER_FAVORITE:
      placeholder = t("SearchListUserFavoriteTip");
      break;
    case ListDataType.LIST_PUBLIC_REPO:
      placeholder = t("SearchPublicRepoTip");
      break;
    case ListDataType.LIST_USER_FOLLOWER:
      placeholder = t("SearchListUserFollowerTip");
      break;
    case ListDataType.LIST_USER_FOLLOWING:
      placeholder = t("SearchListUserFollowingTip");
      break;

    case ListDataType.LIST_ADMIN_COMMENT:
      placeholder = t("SearchListAdminCommentTip");
      break;
    case ListDataType.LIST_ADMIN_COMMENT_REPORT:
      placeholder = t("SearchListAdminCommentReportTip");
      break;
    case ListDataType.LIST_ADMIN_SESSION:
      placeholder = t("SearchListAdminSessionTip");
      break;
    case ListDataType.LIST_ADMIN_REPO:
      placeholder = t("SearchListAdminRepoTip");
      break;
    case ListDataType.LIST_ADMIN_USER:
      placeholder = t("SearchListAdminUserTip");
      break;
    case ListDataType.LIST_REPO_VISI:
      placeholder = t("SearchListRepoUserTip");
      break;
    default:
      break;
      throw new Error("Unsupported oauth-party type");
  }
  return (
    <div className="relative flex w-full mb-4">
      <input
        type={"text"}
        placeholder={placeholder}
        defaultValue={searchParams?.get("query")?.toString()}
        onChange={(e) => {
          handleSearch(e.target.value);
        }}
        className={`py-3 h-12 pl-10 rounded-md border border-slate-300/75 dark:border-0  dark:text-slate-400 grow  dark:bg-slate-800
          placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base 
           ${"focus:outline-0"}`}
      />
      <MdSearch className="absolute left-3 top-1/2 h-[24px] w-[24px] -translate-y-1/2 text-gray-500 peer-focus:text-gray-900" />
      {listType == ListDataType.LIST_REPO_VISI && (
        <IoMdPersonAdd
          className="absolute right-3 top-1/2 h-[24px] w-[24px] -translate-y-1/2 text-gray-500 peer-focus:text-gray-900 cursor-pointer hover:text-sky-600 dark:hover:text-sky-400"
          onClick={() => {
            setOperationRepoName(decodeURIComponent(repo_name));
            setOperationUsername(username);
            setSearchType(SearchType.VISI_USER); //搜索仓库可见用户
            setSearchDialogOpen(true);
          }}
        />
      )}
    </div>
  );
}
