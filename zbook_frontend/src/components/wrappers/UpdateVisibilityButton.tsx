"use client";

import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import React, { useState } from "react";
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
      className="bg-green-500 dark:bg-green-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-green-600 dark:hover:bg-green-500"
      onClick={() => updateFollowStatus()}
    >
      <span className="flex-1 whitespace-nowrap">
        {isVisible ? "取消" : "添加"}
      </span>
    </div>
  );
}
