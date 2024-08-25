import { redirect } from "@/navigation";
import NotFound from "@/components/loadings/NotFound";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { FetchError } from "@/fetchs/util";

interface RepoDetailParams {
  reponame: string;
  username: string;
  locale: string;
}

async function fetchHomePage({ reponame, username, locale }: RepoDetailParams) {
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";

  try {
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_FIRST_DOCUMENT,
      xforward,
      agent,
      tags: [],
      values: {
        username,
        repo_name: decodeURIComponent(reponame),
        lang: locale || "en",
      },
    });

    if (data.error) {
      throw new FetchError(data.message, data.status);
    }

    return data.relative_path;
  } catch (error) {
    console.error("Error fetching home page:", error);
    return null;
  }
}

export default async function RepoDetail({
  params,
}: {
  params: RepoDetailParams;
}) {
  const { reponame, username, locale } = params;
  const homePage = await fetchHomePage({ reponame, username, locale });

  if (!homePage) {
    return <NotFound />;
  }

  redirect(`/workspace/${username}/o/${reponame}/${homePage}`);
}
