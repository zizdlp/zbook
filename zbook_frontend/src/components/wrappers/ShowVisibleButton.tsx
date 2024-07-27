"use client";
import { OperationContext } from "@/providers/OperationProvider";
import { SearchDialogContext } from "@/providers/SearchDialogProvider";
import { useContext } from "react";
import { MdOutlineVisibility } from "react-icons/md";
import { useTranslations } from "next-intl";
import RepoSideBarButton from "../sidebars/RepoSideBarButton";
import { SearchType } from "@/utils/const_value";
export default function ShowVisibleButton({
  repo_id,
  visibility_level,
}: {
  repo_id: number;
  visibility_level: string;
}) {
  const t = useTranslations("SideBar");
  const { searchDialogOpen, setSearchDialogOpen, setSearchType } =
    useContext(SearchDialogContext);
  const { setOperationRepoID } = useContext(OperationContext);
  return (
    <RepoSideBarButton
      onclick={() => {
        setOperationRepoID(repo_id);
        setSearchType(SearchType.VISI_USER); //搜索仓库可见用户
        setSearchDialogOpen(true);
      }}
      url="#"
      title={t("VisibleTo")}
      className="group-hover:bg-indigo-500"
    >
      <MdOutlineVisibility
        className={`h-4 w-4 group-hover:text-white text-gray-700 dark:text-gray-400`}
      />
    </RepoSideBarButton>
  );
}
