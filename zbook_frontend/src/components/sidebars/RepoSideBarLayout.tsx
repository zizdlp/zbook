import RepoSideBar from "@/components/sidebars/RepoSideBar";
import NotFound from "@/components/loadings/NotFound";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { auth } from "@/auth";
import { headers } from "next/headers";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
import RepoSideBarLoading from "../loadings/RepoSideBarLoading";
import LeftSideBarWrapper from "./LeftSideBarWrapper";
import { BiError } from "react-icons/bi";
export default async function RepoSideBarLayout({
  reponame,
  username,
}: {
  reponame: string;
  username: string;
}) {
  try {
    let authname = "";
    const session = await auth();
    if (session && session.username && session.role) {
      authname = session.username;
    }
    // const delay = Math.floor(Math.random() * 4000) + 1400;
    // await new Promise((resolve) => setTimeout(resolve, delay));
    const xforward = headers().get("x-forwarded-for") ?? "";
    const agent = headers().get("User-Agent") ?? "";
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_REPO_CONFIG,
      xforward: xforward,
      agent: agent,
      tags: [],
      values: {
        repo_name: decodeURIComponent(reponame),
        username: username,
      },
    });

    if (data.error) {
      throw new FetchError(data.message, data.status);
    }
    if (data && data.config) {
      const stringConfig = data.config;
      const jsonConfig = JSON.parse(stringConfig);
      const layout = jsonConfig.layout;
      return (
        <RepoSideBar
          sublayouts={layout}
          anchors={jsonConfig.anchors}
          theme_sidebar={data.theme_sidebar}
          theme_color={data.theme_color}
          reponame={reponame}
          username={data.username}
          authname={authname}
          visibility_level={data.visibility_level}
        />
      );
    } else {
      throw new Error(`Error fetching Repo Layout:${data.message}`);
    }
  } catch (error) {
    let e = error as FetchError;
    logger.error(`Fetch SideBarlayout:${e.message}`, e.status);
    return (
      <LeftSideBarWrapper small={true}>
        <div className="absolute inset-0 flex items-center justify-center   z-50">
          <div className="text-center">
            <BiError className="text-red-600 dark:text-red-400 w-12 h-12 mx-auto" />
            <p className="text-lg font-semibold text-gray-700 dark:text-gray-200">
              Error loading SideBar data
            </p>
          </div>
        </div>
      </LeftSideBarWrapper>
    );
  }
}
