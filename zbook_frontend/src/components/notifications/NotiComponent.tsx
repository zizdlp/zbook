"use client";
import { useContext } from "react";
import { NotiDialogContext } from "@/providers/NotiDialogProvider";
import { useTranslations } from "next-intl";
import AvatarImageClient from "../AvatarImageClient";
import {
  ListCommentNotificationInfo,
  ListFollowerNotificationInfo,
  ListRepoNotificationInfo,
  ListSystemNotificationInfo,
} from "@/fetchs/model";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import TimeElement from "../TimeElement";
import NotificationLink from "./NotificationLink";
interface NotiComponentProps {
  ListModelInfo:
    | ListFollowerNotificationInfo
    | ListCommentNotificationInfo
    | ListRepoNotificationInfo
    | ListSystemNotificationInfo;
  noti_type: string;
}

export default function NotiComponent(props: NotiComponentProps) {
  const t = useTranslations("Notifications");
  const {
    setNotiDialogOpen,
    setMutationReadNotification,
    mutationReadNotification,
  } = useContext(NotiDialogContext);
  let specificModelInfo;
  let username: string = "";
  let mainText: string = "";
  let subText: string = "";
  let redirect_url: string = "";
  let createdInfo: string = "";
  switch (props.noti_type) {
    case "repoNotification":
      specificModelInfo = props.ListModelInfo as ListRepoNotificationInfo;
      username = specificModelInfo.username;
      mainText = specificModelInfo.username;
      subText = t("NewRepoCreated");
      redirect_url =
        "/workspace/" +
        specificModelInfo.username +
        "/o/" +
        specificModelInfo.repo_name;
      break;
    case "commentNotification":
      specificModelInfo = props.ListModelInfo as ListCommentNotificationInfo;
      username = specificModelInfo.username;
      mainText = specificModelInfo.username + t("NewCommented");
      subText = specificModelInfo.comment_content;
      redirect_url =
        "/workspace/" +
        specificModelInfo.repo_username +
        "/o/" +
        specificModelInfo.repo_name +
        "/" +
        specificModelInfo.relative_path;
      break;
    case "followerNotification":
      specificModelInfo = props.ListModelInfo as ListFollowerNotificationInfo;
      username = specificModelInfo.username;
      mainText = username + t("NewFollower");
      subText = specificModelInfo.email;
      redirect_url = "/workspace/" + specificModelInfo.username;

      break;
    case "systemNotification":
      specificModelInfo = props.ListModelInfo as ListSystemNotificationInfo;
      username = "admin";
      mainText = specificModelInfo.title;
      subText = specificModelInfo.contents;
      redirect_url = specificModelInfo.redirect_url ?? "";
      break;
    default:
      specificModelInfo = null;
      break;
  }

  async function setNotiRead() {
    if (props.noti_type === "repoNotification") {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.MARK_REPO_NOTI_READED,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          noti_id: props.ListModelInfo.noti_id,
        },
      });
      if (!data.error) {
        props.ListModelInfo.readed = true;
      }
    } else if (props.noti_type === "commentNotification") {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.MARK_COMMENT_NOTI_READED,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          noti_id: props.ListModelInfo.noti_id,
        },
      });
      if (!data.error) {
        props.ListModelInfo.readed = true;
      }
    } else if (props.noti_type === "followerNotification") {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.MARK_FOLLOWER_NOTI_READED,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          noti_id: props.ListModelInfo.noti_id,
        },
      });
      if (!data.error) {
        props.ListModelInfo.readed = true;
      }
    } else if (props.noti_type === "systemNotification") {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.MARK_SYSTEM_NOTI_READED,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          noti_id: props.ListModelInfo.noti_id,
        },
      });
      if (!data.error) {
        props.ListModelInfo.readed = true;
      }
    }
    setMutationReadNotification(!mutationReadNotification);
  }
  return (
    <NotificationLink
      redirect_url={redirect_url}
      onClickFunc={() => {
        if (redirect_url != "") {
          setNotiDialogOpen(false);
        }
        setNotiRead();
      }}
    >
      <div className="flex items-center justify-center">
        <div
          className={`flex-shrink-0 ${
            props.ListModelInfo.readed ? "" : "bg-green-600"
          } ml-4 w-2 h-2 md:w-2 md:h-2 rounded-full`}
        />
        <div className="flex-shrink-0">
          <AvatarImageClient
            username={username}
            className="w-8 h-8 md:w-12 md:h-12 ml-2 mr-4 rounded-full shadow-lg"
          />
        </div>
        <div className="flex flex-col md:py-5 py-3 text-xs md:text-sm">
          <strong className="">{mainText}</strong>
          <p className="md:font-medium">{subText}</p>
        </div>
      </div>
      <div className="px-4 flex-shrink-0">
        <span className="text-sm">
          {createdInfo}
          <TimeElement timeInfo={props.ListModelInfo.created_at} />
        </span>
      </div>
    </NotificationLink>
  );
}
