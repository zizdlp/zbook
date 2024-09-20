import SideBarSearchButton from "./SideBarSearchButton";
import LeftSideBarWrapper from "@/components/sidebars/LeftSideBarWrapper";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { headers } from "next/headers";
import UserSideBarSetting from "./UserSideBarSetting";
import ShowComponent from "../ShowComponent";
import UserSideBarSocial from "./UserSideBarSocial";
import UserSideBarAdmin from "./UserSideBarAdmin";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { getTranslations } from "next-intl/server";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
import OAuthWrapper from "./OAuthWrapper";
import { auth } from "@/auth";
import { SearchType } from "@/utils/const_value";
import UserSideBarProfile from "./UserSideBarProfile";
import { BiError } from "react-icons/bi";
export default async function UserSideBar({
  username,
  authname,
  authrole,
}: {
  username: string;
  authname: string;
  authrole: string;
}) {
  const t = await getTranslations("SideBar");
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";
  const session = await auth();
  try {
    // const delay = Math.floor(Math.random() * 4000) + 1400;
    // await new Promise((resolve) => setTimeout(resolve, delay));
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_USER_INFO,
      xforward: xforward,
      agent: agent,
      tags: [],
      values: {
        username: username,
        user_basic: true,
        user_count: true,
        user_image: true,
      },
    });
    if (data.error) {
      throw new FetchError(data.message, data.status);
    }
    let follow_status = false;
    if (username != authname) {
      const follow = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_FOLLOWE_STATUS,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: {
          username: username,
        },
      });
      if (follow.error) {
        throw new FetchError(follow.message, follow.status);
      }
      follow_status = follow.is_following;
    }
    let github = false;
    let google = false;
    if (authname == username) {
      const authData = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CHECK_OAUTH_STATUS,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: {},
      });
      if (authData.error) {
        throw new FetchError(authData.message, authData.status);
      }
      github = authData?.github;
      google = authData?.google;
    }

    return (
      <LeftSideBarWrapper small={true}>
        <SideBarSearchButton
          username={username}
          repo_name=""
          searchType={SearchType.USER_DOCUMENT}
        />

        <div
          className="absolute inset-0 z-10 pb-10 pt-32 lg:pt-24 px-4
        scrollbar scrollbar-thumb-sky-700 scrollbar-track-sky-300 overflow-y-scroll"
        >
          <UserSideBarProfile
            avatar={data.user_image_info?.avatar}
            email={data.user_basic_info?.email}
            username={data.user_basic_info?.username}
            motto={data.user_basic_info?.motto}
            bio={t("Bio")}
          />

          <div className="w-full pt-4 rounded-md sm:pt-6">
            <p className="font-semibold text-slate-800 dark:text-slate-400 border-slate-300 dark:border-slate-700 border-b-[0.01rem] pb-0.5">
              {t("Social")}
            </p>
            <UserSideBarSocial
              username={username}
              count_following={data.user_count_info?.count_following ?? 0}
              count_follower={data.user_count_info?.count_follower ?? 0}
              count_likes={data.user_count_info?.count_likes ?? 0}
              count_repos={data.user_count_info?.count_repos ?? 0}
            />
          </div>

          <div className="w-full pt-4 rounded-md sm:pt-6">
            <p className="font-semibold text-slate-800 dark:text-slate-400 border-slate-300 dark:border-slate-700 border-b-[0.01rem] pb-0.5">
              {t("Settings")}
            </p>
            <UserSideBarSetting
              username={username}
              authname={authname}
              authrole={authrole}
              follow_status={follow_status}
            />
          </div>

          <ShowComponent show={authname == username && authrole == "admin"}>
            <div className="w-full pt-4 rounded-md sm:pt-6">
              <p className="font-semibold text-slate-800 dark:text-slate-400 border-slate-300 dark:border-slate-700 border-b-[0.01rem] pb-0.5">
                {t("Admin")}
              </p>
              <UserSideBarAdmin
                username={username}
                authname={authname}
                authrole={authrole}
              />
            </div>
          </ShowComponent>
          <ShowComponent show={authname == username}>
            <div className="w-full pt-4 rounded-md sm:pt-6">
              <p className="font-semibold text-slate-800 dark:text-slate-400 border-slate-300 dark:border-slate-700 border-b-[0.01rem] pb-0.5">
                {t("LinkedAccount")}
              </p>
              <ul className="my-4 space-y-3">
                <li>
                  <OAuthWrapper
                    oauth_type={"github"}
                    status={github}
                    access_token={session?.access_token ?? ""}
                  />
                </li>
                <li>
                  <OAuthWrapper
                    oauth_type={"google"}
                    status={google}
                    access_token={session?.access_token ?? ""}
                  />
                </li>
              </ul>
            </div>
          </ShowComponent>
        </div>
      </LeftSideBarWrapper>
    );
  } catch (error) {
    let currentError = error as FetchError;
    logger.error(
      `UserSideBar Error:${currentError.status} ${currentError.message}`
    );
    return (
      <LeftSideBarWrapper small={true}>
        <div className="absolute inset-0 flex items-center justify-center z-50">
          <div className="text-center">
            <BiError className="text-red-600 dark:text-red-400 w-12 h-12 mx-auto" />
            <p className="text-lg font-semibold text-gray-700 dark:text-gray-200">
              {t("ErrorLoadingUserSideBar")}
            </p>
          </div>
        </div>
      </LeftSideBarWrapper>
    );
  }
}
