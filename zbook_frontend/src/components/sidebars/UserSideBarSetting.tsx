"use client";
import React, { useContext } from "react";
import { OperationContext } from "@/providers/OperationProvider";
import SideBarLiContent from "./SideBarLiContent";
import {
  MdNoteAdd,
  MdEditNotifications,
  MdInsertInvitation,
  MdOutlineInsertInvitation,
} from "react-icons/md";
import ShowComponent from "../ShowComponent";
import { FaUserEdit } from "react-icons/fa";
import { SlUserFollowing, SlUserUnfollow } from "react-icons/sl";
import { useTranslations } from "next-intl";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { toast } from "react-toastify";
import { FetchError } from "@/fetchs/util";
export default function UserSideBarSetting({
  username,
  authname,
  authrole,
  follow_status,
}: {
  username: string;
  authname: string;
  authrole: string;
  follow_status: boolean;
}) {
  const t = useTranslations("SideBar");
  async function updateFollowStatus() {
    const id = toast(t("UpdatingFollowStatus"), {
      type: "info",
      isLoading: true,
    });
    try {
      if (follow_status) {
        const data = await fetchServerWithAuthWrapper({
          endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_FOLLOW,
          xforward: "",
          agent: "",
          tags: [],
          values: {
            username: username,
          },
        });
        if (data.error) {
          throw new FetchError(data.message, data.status);
        }
      } else {
        const data = await fetchServerWithAuthWrapper({
          endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_FOLLOW,
          xforward: "",
          agent: "",
          tags: [],
          values: {
            username: username,
          },
        });
        if (data.error) {
          throw new FetchError(data.message, data.status);
        }
      }
      toast.update(id, {
        render: t("UpdateSuccessful"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
      refreshPage("/workspace/[username]", true, false);
    } catch (error) {
      toast.update(id, {
        render: t("UpdateFailed"),
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }

  const {
    setCreateRepoOpen,
    setCreateSystemNotificationOpen,
    setCreateInvitationOpen,
  } = useContext(OperationContext);

  let isCurrentUser = username === authname;
  let isAdmin = authrole === "admin";
  const { updateUserOpen, setUpdateUserOpen } = useContext(OperationContext);

  return (
    <ul className="my-4 space-y-3">
      <ShowComponent show={isCurrentUser}>
        <li onClick={() => setUpdateUserOpen(!updateUserOpen)}>
          <SideBarLiContent isSelected={false}>
            <FaUserEdit className="text-sky-600 dark:text-sky-400 w-6 h-6" />
            <span className="flex-1 ms-3 whitespace-nowrap">{t("Edit")}</span>
          </SideBarLiContent>
        </li>
      </ShowComponent>

      <ShowComponent show={!isCurrentUser}>
        <li onClick={() => updateFollowStatus()}>
          <SideBarLiContent isSelected={false}>
            {follow_status ? (
              <>
                <SlUserUnfollow className="text-sky-600 dark:text-sky-400 w-6 h-6" />
                <span className="flex-1 ms-3 whitespace-nowrap">
                  {t("UnFollow")}
                </span>
              </>
            ) : (
              <>
                <SlUserFollowing className="text-sky-700 dark:text-sky-400 w-6 h-6" />
                <span className="flex-1 ms-3 whitespace-nowrap">
                  {t("Follow")}
                </span>
              </>
            )}
          </SideBarLiContent>
        </li>
      </ShowComponent>

      <ShowComponent show={isCurrentUser}>
        <li onClick={() => setCreateRepoOpen(true)}>
          <SideBarLiContent isSelected={false}>
            <MdNoteAdd className="text-amber-600 dark:text-amber-400 w-6 h-6" />
            <span className="flex-1 ms-3 whitespace-nowrap">
              {t("NewRepo")}
            </span>
          </SideBarLiContent>
        </li>
      </ShowComponent>
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li onClick={() => setCreateSystemNotificationOpen(true)}>
          <SideBarLiContent isSelected={false}>
            <MdEditNotifications className="text-indigo-600 dark:text-indigo-400 w-6 h-6" />
            <span className="flex-1 ms-3 whitespace-nowrap">
              {t("NewNotification")}
            </span>
          </SideBarLiContent>
        </li>
      </ShowComponent>
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li onClick={() => setCreateInvitationOpen(true)}>
          <SideBarLiContent isSelected={false}>
            <MdOutlineInsertInvitation className="text-teal-600 dark:text-teal-400 w-6 h-6" />
            <span className="flex-1 ms-3 whitespace-nowrap">
              {t("NewInvitation")}
            </span>
          </SideBarLiContent>
        </li>
      </ShowComponent>
    </ul>
  );
}
