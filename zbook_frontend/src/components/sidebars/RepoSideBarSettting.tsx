"use client";
import { AiOutlineSync } from "react-icons/ai";
import { toast } from "react-toastify";
import { IoBookOutline } from "react-icons/io5";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { useTranslations } from "next-intl";
import RepoSideBarButton from "./RepoSideBarButton";
import { FetchError } from "@/fetchs/util";
import { MdOutlineVisibility } from "react-icons/md";
export default function RepoSideBarSetting({
  username,
  reponame,
  repo_id,
  authname,
  visibility_level,
}: {
  username: string;
  reponame: string;
  repo_id: number;
  authname: string;
  visibility_level: string;
}) {
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
          repo_name: reponame,
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
        render: t("RepositorySynchronizationFiled"),
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }

  return (
    <div className="mb-4 md:mt-0">
      <RepoSideBarButton
        url={`/workspace/${username}/o/${reponame}`}
        onclick={() => {}}
        title={t("RepoHome")}
        className="group-hover:bg-sky-500"
      >
        <IoBookOutline
          className={`h-4 w-4 group-hover:text-white text-gray-700 dark:text-gray-400`}
        />
      </RepoSideBarButton>

      {authname == username && (
        <>
          <RepoSideBarButton
            onclick={() => {
              manualSyncRepoHandler();
            }}
            url="#"
            title={t("SyncRepo")}
            className="group-hover:bg-teal-500"
          >
            <AiOutlineSync
              className={`h-4 w-4 group-hover:text-white text-gray-700 dark:text-gray-400`}
            />
          </RepoSideBarButton>
          {visibility_level == "chosen" && (
            <RepoSideBarButton
              onclick={() => {
                {
                }
              }}
              url={`/workspace/${username}/o/${reponame}/~visi`}
              title={t("VisibleTo")}
              className="group-hover:bg-indigo-500"
            >
              <MdOutlineVisibility
                className={`h-4 w-4 group-hover:text-white text-gray-700 dark:text-gray-400`}
              />
            </RepoSideBarButton>
          )}
        </>
      )}
    </div>
  );
}
