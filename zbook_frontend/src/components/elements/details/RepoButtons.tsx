"use client";
import { ListDataType } from "@/fetchs/model";
import { OperationContext } from "@/providers/OperationProvider";
import { useContext } from "react";
import { BiEdit } from "react-icons/bi";
import { MdDeleteOutline } from "react-icons/md";

export default function RepoButtons({
  username,
  reponame,
  authname,
  listType,
}: {
  username: string;
  reponame: string;
  authname: string;
  listType: ListDataType;
}) {
  const {
    setUpdateRepoOpen,
    setDeleteRepoOpen,
    setOperationUsername,
    setOperationRepoName,
  } = useContext(OperationContext);
  if (authname != username) {
    return <></>;
  }
  return (
    <>
      {listType === ListDataType.LIST_USER_REPO && (
        <BiEdit
          onClick={() => {
            setOperationUsername(username);
            setOperationRepoName(decodeURIComponent(reponame));
            setUpdateRepoOpen(true);
          }}
          className="p-1 w-7 h-7 cursor-pointer text-gray-500 border dark:border-0 border-gray-200 rounded dark:bg-[#263142] hover:bg-sky-500 hover:text-white dark:hover:bg-gray-900 dark:text-gray-400"
        />
      )}
      {(listType === ListDataType.LIST_USER_REPO ||
        listType === ListDataType.LIST_ADMIN_REPO) && (
        <MdDeleteOutline
          onClick={() => {
            setOperationRepoName(decodeURIComponent(reponame));
            setOperationUsername(username);
            setDeleteRepoOpen(true);
          }}
          className="p-1 w-7 h-7 cursor-pointer text-gray-500 border border-gray-200 dark:border-0 rounded dark:bg-[#263142] hover:bg-red-500 hover:text-white dark:hover:bg-gray-900 dark:text-gray-400"
        />
      )}
    </>
  );
}
