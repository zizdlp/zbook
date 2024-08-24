"use client";

import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import React, { useState } from "react";
import { MdOutlineVisibilityOff, MdVisibility } from "react-icons/md";

export default function UpdateVisibleButton({
  username,
  repo_name,
  repo_username,
  is_visible,
}: {
  username: string;
  repo_name: string;
  repo_username: string;
  is_visible: boolean;
}) {
  const [isVisible, setIsVisible] = useState(is_visible);

  async function updateFollowStatus() {
    if (isVisible) {
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_REPO_VISIBILITY,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          repo_name: decodeURIComponent(repo_name),
          repo_username: repo_username,
          username: username,
        },
      }).then((data: any) => {
        setIsVisible(!isVisible);
      });
    } else {
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_REPO_VISIBILITY,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          repo_name: decodeURIComponent(repo_name),
          repo_username: repo_username,
          username: username,
        },
      }).then((data: any) => {
        setIsVisible(!isVisible);
      });
    }
  }
  return (
    <div
      className="bg-green-500 flex space-x-2 items-center dark:bg-green-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-green-600 dark:hover:bg-green-500"
      onClick={() => updateFollowStatus()}
    >
      {isVisible ? <MdVisibility /> : <MdOutlineVisibilityOff />}

      <span className="flex-1 whitespace-nowrap">
        {isVisible ? "取消" : "允许"}
      </span>
    </div>
  );
}
