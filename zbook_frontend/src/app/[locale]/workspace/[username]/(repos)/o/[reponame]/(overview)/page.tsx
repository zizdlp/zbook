import { redirect } from "@/navigation";
import NotFound from "@/components/loadings/NotFound";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { FetchError } from "@/fetchs/util";
export default async function RepoDetail({
  params: { reponame, username },
}: {
  params: { reponame: string; username: string };
}) {
  let home_page = "readme";
  try {
    const xforward = headers().get("x-forwarded-for") ?? "";
    const agent = headers().get("User-Agent") ?? "";
    const repo_data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_REPO_BASIC_INFO,
      xforward: xforward,
      agent: agent,
      tags: [],
      values: {
        username: username,
        repo_name: decodeURIComponent(reponame),
      },
    });
    if (repo_data.error) {
      throw new FetchError(repo_data.message, repo_data.status);
    }
    home_page = repo_data.home_page;
  } catch (error) {
    return <NotFound />;
  }
  redirect(`/workspace/${username}/o/${reponame}/${home_page}`); // Navigate to the new post page
}
