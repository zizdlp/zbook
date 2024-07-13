"use client";
import { OperationContext } from "@/providers/OperationProvider";
import { useContext } from "react";
import { useTranslations } from "next-intl";
export default function DeleteButton({
  id,
  username,
  dataType,
}: {
  id: number;
  username: string;
  dataType: string;
}) {
  const t = useTranslations("Dialog");
  const {
    setDeleteRepoOpen,
    setDeleteCommentOpen,
    setDeleteUserOpen,
    setOperationCommentID,
    setOperationRepoID,
    setOperationUsername,
  } = useContext(OperationContext);
  if (dataType == "repo") {
    return (
      <div
        className="bg-sky-500 dark:bg-sky-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-sky-600 dark:hover:bg-sky-500"
        onClick={() => {
          setOperationRepoID(id);
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
      <div
        className="bg-sky-500 dark:bg-sky-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-sky-600 dark:hover:bg-sky-500"
        onClick={() => {
          setOperationCommentID(id);
          setDeleteCommentOpen(true);
        }}
      >
        <span className="flex-1 whitespace-nowrap">{t("Delete")}</span>
      </div>
    );
  }
}
