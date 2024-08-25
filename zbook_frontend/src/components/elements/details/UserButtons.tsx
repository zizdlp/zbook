"use client";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { FetchError } from "@/fetchs/util";
import { OperationContext } from "@/providers/OperationProvider";
import { useContext } from "react";
import { MdDeleteOutline, MdLockOpen, MdLockPerson } from "react-icons/md";

export default function UserButtons({
  username,
  is_blocked,
}: {
  username: string;
  is_blocked: boolean;
}) {
  const { setOperationUsername, setDeleteUserOpen } =
    useContext(OperationContext);
  async function actionUpdateUserBlock() {
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_USER_BLOCK,
        xforward: "xforward",
        agent: "",
        tags: [],
        values: {
          username: username,
          blocked: !is_blocked,
        },
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
    } catch (error) {
      let e = error as FetchError;
    }
    await refreshPage(`/workspace/${username}/admin_users`, true, false);
  }
  return (
    <>
      {is_blocked && (
        <MdLockPerson
          onClick={async () => {
            await actionUpdateUserBlock();
          }}
          className="p-1 w-7 h-7 cursor-pointer text-gray-500 border dark:border-0 border-gray-200 rounded dark:bg-[#263142] hover:bg-sky-500 hover:text-white dark:hover:bg-gray-900 dark:text-gray-400"
        />
      )}
      {!is_blocked && (
        <MdLockOpen
          onClick={async () => {
            await actionUpdateUserBlock();
          }}
          className="p-1 w-7 h-7 cursor-pointer text-gray-500 border dark:border-0 border-gray-200 rounded dark:bg-[#263142] hover:bg-sky-500 hover:text-white dark:hover:bg-gray-900 dark:text-gray-400"
        />
      )}

      <MdDeleteOutline
        onClick={() => {
          setOperationUsername(username);
          setDeleteUserOpen(true);
        }}
        className="p-1 w-7 h-7 cursor-pointer text-gray-500 border border-gray-200 dark:border-0 rounded dark:bg-[#263142] hover:bg-red-500 hover:text-white dark:hover:bg-gray-900 dark:text-gray-400"
      />
    </>
  );
}
