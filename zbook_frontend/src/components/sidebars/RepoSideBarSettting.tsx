"use client";
import { toast } from "react-toastify";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { useTranslations } from "next-intl";
import { FetchError } from "@/fetchs/util";
import { Anchor } from "@/types/interface";
import RepoSideBarSettingItem from "./RepoSideBarSetttingItem";
import { ThemeColor } from "../TableOfContent";
import { usePathname } from "@/navigation";
export default function RepoSideBarSetting({
  username,
  reponame,
  authname,
  anchors,
  visibility_level,
  theme_color,
  first_path,
}: {
  username: string;
  reponame: string;
  authname: string;
  anchors: Anchor[];
  visibility_level: string;
  theme_color: ThemeColor;
  first_path: string;
}) {
  const pathname = usePathname();
  let page_type = "";
  if (pathname !== undefined) {
    const regex = new RegExp(
      `^\/([^\/]+)?\/?workspace\/${username}\/(?:[^\/]*)\/(?:[^\/]*)\/([^\/]*)`
    );
    const matches = pathname.match(regex);
    if (matches) {
      page_type = matches[2] || "";
    }
  }

  const t = useTranslations("SideBar");
  async function manualSyncRepoHandler() {
    const id = toast(t("SynchronizingRepository"), {
      type: "info",
      isLoading: true,
    });
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.MANUAL_SYNC_REPO,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          repo_name: decodeURIComponent(reponame),
          username: username,
        },
        timeout: 300000, //600s
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      toast.update(id, {
        render: t("RepositorySynchronizedSuccessfully"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
      refreshPage("#", true, false);
    } catch (error) {
      toast.update(id, {
        render: t("RepositorySynchronizationFiled") + error,
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }
  return (
    <>
      <RepoSideBarSettingItem
        href={`/workspace/${username}/o/${reponame}/${first_path}`}
        selected={page_type != "~visi"}
        icon_name={"IoBook"}
        text={t("RepoHome")}
        theme_color={theme_color}
      />
      {authname == username && (
        <>
          <RepoSideBarSettingItem
            onClick={() => {
              manualSyncRepoHandler();
            }}
            selected={false}
            href="#"
            icon_name={"MdCloudSync"}
            text={t("SyncRepo")}
            theme_color={theme_color}
          />

          {visibility_level == "chosen" && (
            <RepoSideBarSettingItem
              href={`/workspace/${username}/o/${reponame}/~visi`}
              text={t("VisibleTo")}
              selected={page_type == "~visi"}
              icon_name={"MdOutlineVisibility"}
              theme_color={theme_color}
            />
          )}
        </>
      )}
      {anchors.map((item, index) => (
        <RepoSideBarSettingItem
          key={index}
          href={item.url}
          selected={false}
          icon_name={item.icon}
          text={item.name}
          theme_color={theme_color}
        />
      ))}
    </>
  );
}
