"use client";
import { OperationContext } from "@/providers/OperationProvider";
import { useContext } from "react";
import { useTranslations } from "next-intl";
import { MdDeleteOutline } from "react-icons/md";
export default function DeleteButton({
  comment_id,
  username,
  repo_name,
  dataType,
}: {
  comment_id: number;
  username: string;
  repo_name: string;
  dataType: string;
}) {
  const t = useTranslations("Dialog");
  const {
    setDeleteRepoOpen,
    setDeleteCommentOpen,
    setDeleteUserOpen,
    setOperationCommentID,
    setOperationRepoName,
    setOperationUsername,
  } = useContext(OperationContext);
  if (dataType == "repo") {
    return (
      <div
        className="bg-sky-500 dark:bg-sky-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-sky-600 dark:hover:bg-sky-500"
        onClick={() => {
          setOperationUsername(username);
          setOperationRepoName(decodeURIComponent(repo_name));
          setDeleteRepoOpen(true);
        }}
      >
        <span className="flex-1 whitespace-nowrap">{t("Delete")}</span>
      </div>
    );
  } else if (dataType == "user") {
    return (
      <div
        className="bg-sky-500 dark:bg-sky-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-sky-600 dark:hover:bg-sky-500"
        onClick={() => {
          setOperationUsername(username);
          setDeleteUserOpen(true);
        }}
      >
        <span className="flex-1 whitespace-nowrap">{t("Delete")}</span>
      </div>
    );
  } else if (dataType == "comment") {
    return (
      <MdDeleteOutline
        onClick={() => {
          setOperationCommentID(comment_id);
          setDeleteCommentOpen(true);
        }}
        className="p-1 w-7 h-7 cursor-pointer text-gray-500 border border-gray-200 dark:border-0 rounded dark:bg-[#263142] hover:bg-red-500 hover:text-white dark:hover:bg-gray-900 dark:text-gray-400"
      />
    );
  }
}
