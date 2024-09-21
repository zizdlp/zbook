"use client";
import DialogComponent from "@/components/DialogComponent";
import { NotiDialogContext } from "@/providers/NotiDialogProvider";
import { Tab } from "@headlessui/react";
import { useTranslations } from "next-intl";
function classNames(...classes: string[]): string {
  return classes.filter(Boolean).join(" ");
}
import React, { useContext, useEffect, useState } from "react";
import ListNotifications from "./ListNotifications";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { useSession } from "next-auth/react";

export default function NotificationDialog() {
  const t = useTranslations("Notifications");
  const { notiDialogOpen, setNotiDialogOpen, mutationReadNotification } =
    useContext(NotiDialogContext);
  const [systemUnreadedCount, setSystemUnreadedCount] = useState(0);
  const [commentUnreadedCount, setCommentUnreadedCount] = useState(0);
  const [repoUnreadedCount, setRepoUnreadedCount] = useState(0);
  const [followerUnreadedCount, setFollowerUnreadedCount] = useState(0);
  const [username, setUsername] = useState("");
  const { data, status } = useSession();
  useEffect(() => {
    if (data?.username) {
      setUsername(data.username);
    }
  }, [data]);
  let categories = [
    t("SystemNotifications"),
    t("FollowerNotifications"),
    t("RepoNotifications"),
    t("CommentNotifications"),
  ];
  useEffect(() => {
    const fetchData = async () => {
      if (username) {
        const [systemData, commentData, repoData, followerData] =
          await Promise.all([
            await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.GET_LIST_SYSTEM_NOTIFICATION_UNREADED_COUNT,
              xforward: "",
              agent: "",
              tags: [],
              values: {},
            }),
            await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_NOTIFICATION_UNREADED_COUNT,
              xforward: "",
              agent: "",
              tags: [],
              values: {},
            }),
            await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.GET_LIST_REPO_NOTIFICATION_UNREADED_COUNT,
              xforward: "",
              agent: "",
              tags: [],
              values: {},
            }),
            await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.GET_LIST_FOLLOWER_NOTIFICATION_UNREADED_COUNT,
              xforward: "",
              agent: "",
              tags: [],
              values: {},
            }),
          ]);
        if (systemData.count) {
          setSystemUnreadedCount(systemData.count);
        } else {
          setSystemUnreadedCount(0);
        }
        if (commentData.count) {
          setCommentUnreadedCount(commentData.count);
        } else {
          setCommentUnreadedCount(0);
        }
        if (repoData.count) {
          setRepoUnreadedCount(repoData.count);
        } else {
          setRepoUnreadedCount(0);
        }
        if (followerData.count) {
          setFollowerUnreadedCount(followerData.count);
        } else {
          setFollowerUnreadedCount(0);
        }
      }
    };
    fetchData();
  }, [username, mutationReadNotification]);
  function GetCount(index: number) {
    return index == 0
      ? systemUnreadedCount
      : index == 1
        ? followerUnreadedCount
        : index == 2
          ? repoUnreadedCount
          : commentUnreadedCount;
  }
  return (
    <DialogComponent
      showDialog={notiDialogOpen}
      setShowDialog={setNotiDialogOpen}
    >
      <header
        className="justify-center px-4 py-4 overflow-x-auto relative flex flex-row items-center scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
      >
        <p className="text-lg md:text-xl py-2 text-slate-800 dark:text-slate-100">
          {t("Notifications")}
        </p>
      </header>
      <div className="px-2 md:px-4 pb-2 md:pb-4">
        <Tab.Group>
          <Tab.List className="flex space-x-1 rounded-lg md:rounded-xl bg-gray-200/75 dark:bg-gray-900/75 p-1">
            {categories.map((category, index) => (
              <Tab
                key={index}
                className={({ selected }) =>
                  classNames(
                    "w-full rounded-lg py-2.5 text-xs md:text-sm font-medium leading-3 md:leading-5 focus:outline-none",
                    selected
                      ? "bg-white dark:bg-slate-700/75 shadow text-slate-700 dark:text-slate-200"
                      : "text-slate-500 hover:bg-gray-100/25 dark:hover:bg-gray-900/25 hover:text-slate-700 dark:hover:text-slate-200"
                  )
                }
              >
                <div className="flex items-center justify-center">
                  <div
                    className={`flex-shrink-0 bg-green-600
                      w-2 h-2 md:w-2 md:h-2 mr-1 rounded-full ${
                        GetCount(index) == 0 && "hidden"
                      }`}
                  />
                  {category}
                </div>
              </Tab>
            ))}
          </Tab.List>

          <Tab.Panels
            className="my-2
          rounded-lg md:rounded-xl bg-gray-200/75 dark:bg-gray-900/75 px-3 py-1
          overflow-y-auto md:h-96 h-[30rem] scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
          >
            <Tab.Panel key={0} className={"h-full"}>
              <ListNotifications notificationType="systemNotification" />
            </Tab.Panel>
            <Tab.Panel key={1} className={"h-full"}>
              <ListNotifications notificationType="followerNotification" />
            </Tab.Panel>
            <Tab.Panel key={2} className={"h-full"}>
              <ListNotifications notificationType="repoNotification" />
            </Tab.Panel>
            <Tab.Panel key={3} className={"h-full"}>
              <ListNotifications notificationType="commentNotification" />
            </Tab.Panel>
          </Tab.Panels>
        </Tab.Group>
      </div>
    </DialogComponent>
  );
}
