"use client";

import React, { useState, useEffect, useRef } from "react";
import { useTranslations } from "next-intl";
import {
  ListCommentNotificationInfo,
  ListFollowerNotificationInfo,
  ListRepoNotificationInfo,
  ListSystemNotificationInfo,
} from "@/fetchs/model";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import LoadingNotificationList from "./LoadingNotificationList";
import NoItemFound from "../NoItemFound";
import InfCard from "./InfCard";
import NotiComponent from "./NotiComponent";
import { FetchError } from "@/fetchs/util";

export default function ListNotifications({
  notificationType,
}: {
  notificationType: string;
}) {
  const t = useTranslations("Notifications");
  const [currentPage, setCurrentPage] = useState(1);
  const isFetchingData = useRef(false);
  const [hasMore, setHasMore] = useState(true);
  const [listModelInfo, setListModelInfo] = useState<
    Array<
      | ListFollowerNotificationInfo
      | ListCommentNotificationInfo
      | ListRepoNotificationInfo
      | ListSystemNotificationInfo
    >
  >([]);

  useEffect(() => {
    async function fetchMoreSystemNotification() {
      if (isFetchingData.current) {
        return;
      }
      isFetchingData.current = true;
      try {
        let data = [];
        let values = {
          page_id: currentPage,
          page_size: 5,
        };
        switch (notificationType) {
          case "systemNotification":
            data = await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.LIST_SYSTEM_NOTIFICATION,
              tags: [],
              xforward: "",
              agent: "",
              values: values,
            });
            if (data.error) {
              throw new FetchError(data.message, data.status);
            }
            break;
          case "followerNotification":
            data = await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.LIST_FOLLOWER_NOTIFICATION,
              tags: [],
              xforward: "",
              agent: "",
              values: values,
            });
            if (data.error) {
              throw new FetchError(data.message, data.status);
            }
            break;
          case "repoNotification":
            data = await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.LIST_REPO_NOTIFICATION,
              tags: [],
              xforward: "",
              agent: "",
              values: values,
            });
            if (data.error) {
              throw new FetchError(data.message, data.status);
            }
            break;
          case "commentNotification":
            data = await fetchServerWithAuthWrapper({
              endpoint:
                FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_NOTIFICATION,
              tags: [],
              xforward: "",
              agent: "",
              values: values,
            });
            if (data.error) {
              throw new FetchError(data.message, data.status);
            }
            break;
          default:
            break;
        }
        if (data.notifications != undefined && data.notifications.length > 0) {
          setListModelInfo((prevState) => [
            ...prevState,
            ...data.notifications,
          ]);
        } else {
          setHasMore(false);
        }
        isFetchingData.current = false;
      } catch (error) {
        setHasMore(false);
        isFetchingData.current = false;
      }
    }
    if (!isFetchingData.current) {
      fetchMoreSystemNotification();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentPage]);

  return (
    <>
      {listModelInfo.map((model: any, index) => (
        <InfCard
          key={index}
          isLast={index === listModelInfo.length - 1}
          newLimit={() => {
            if (isFetchingData.current || !hasMore) {
              return;
            } else {
              setCurrentPage(currentPage + 1);
            }
          }}
        >
          <NotiComponent ListModelInfo={model} noti_type={notificationType} />
        </InfCard>
      ))}
      {hasMore && <LoadingNotificationList itemCount={3} />}
      {!hasMore && currentPage == 1 && listModelInfo.length == 0 && (
        <NoItemFound title={t("NoNotifications")} />
      )}
    </>
  );
}
