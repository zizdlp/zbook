import SideBarSearchButton from "./SideBarSearchButton";
import SideBarWrapper from "@/components/sidebars/SideBarWrapper";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import Image from "next/image";
import { headers } from "next/headers";
import UserSideBarSetting from "./UserSideBarSetting";
import ShowComponent from "../ShowComponent";
import UserSideBarSocial from "./UserSideBarSocial";
import UserSideBarAdmin from "./UserSideBarAdmin";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { getTranslations } from "next-intl/server";
import SomeThingWrong from "../SomeThingWrong";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
import OAuthWrapper from "./OAuthWrapper";
import { auth } from "@/auth";
import { SearchType } from "@/utils/const_value";
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
      <SideBarWrapper>
        <div className="sticky top-0 -ml-0.5 pointer-events-none">
          <div className="h-10 bg-white dark:bg-gray-900"></div>
          <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
            <SideBarSearchButton
              username={username}
              repo_id={0}
              searchType={SearchType.USER_DOCUMENT}
            />
          </div>
          <div className="h-4 bg-gradient-to-b from-white dark:from-slate-900"></div>
        </div>
        <div className="">
          <div className="flex flex-col items-center justify-center p-4">
            {data.user_image_info?.avatar ? (
              <Image
                src={`data:image/png;base64,${data.user_image_info?.avatar}`}
                width={80}
                height={80}
                alt="Picture of the user"
                className={`flex-none w-24 h-24 rounded-full  object-cover ${
                  data.user_image_info?.avatar
                    ? ""
                    : "animate-pulse bg-slate-300 dark:bg-slate-700"
                }`}
              />
            ) : (
              <div className="flex-none w-24 h-24 rounded-full  object-cover  bg-slate-300 dark:bg-slate-700" />
            )}

            <p className="text-lg font-semibold">
              {data.user_basic_info?.username}{" "}
            </p>
            <p className="mt-0.5">{data.user_basic_info?.email}</p>
          </div>

          <blockquote className="text-slate-700 dark:text-slate-300 text-sm">
            <p className="overflow-x-auto">
              <span className="font-bold pr-2 text-sm">{t("Bio")}</span>
              {data.user_basic_info?.motto}
            </p>
          </blockquote>

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
      </SideBarWrapper>
    );
  } catch (error) {
    let currentError = error as FetchError;
    logger.error(
      `UserSideBar Error:${currentError.status} ${currentError.message}`
    );
    return (
      <SideBarWrapper>
        <SomeThingWrong />
      </SideBarWrapper>
    );
  }
}
