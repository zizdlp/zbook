"use client";
import { OperationContext } from "@/providers/OperationProvider";
import { useContext } from "react";
import { BiEdit } from "react-icons/bi";
import { MdDeleteOutline } from "react-icons/md";

export default function RepoButtons({
  username,
  reponame,
  authname,
  repo_id,
}: {
  username: string;
  reponame: string;
  authname: string;
  repo_id: number;
}) {
  const {
    setUpdateRepoOpen,
    setDeleteRepoOpen,
    setOperationRepoID,
    setOperationUsername,
    setOperationRepoName,
  } = useContext(OperationContext);
  if (authname != username) {
    console.log("authname:", authname, username);
    return <></>;
  }
  return (
    <>
      <BiEdit
        onClick={() => {
          setOperationRepoID(repo_id);
          setUpdateRepoOpen(true);
        }}
        className="p-1 w-7 h-7 cursor-pointer text-gray-500 bg-gray-200 rounded dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-900 dark:text-gray-400"
      />
      <MdDeleteOutline
        onClick={() => {
          setOperationRepoName(reponame);
          setOperationUsername(username);
          setDeleteRepoOpen(true);
        }}
        className="p-1 w-7 h-7 cursor-pointer text-gray-500 bg-gray-200 rounded dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-900 dark:text-gray-400"
      />
    </>
  );
}
