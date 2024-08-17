import { redirect } from "@/navigation";
import NotFound from "@/components/loadings/NotFound";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { FetchError } from "@/fetchs/util";
export default async function RepoDetail({
  params: { reponame, username, locale },
}: {
  params: { reponame: string; username: string; locale: string };
}) {
  let home_page = "readme";
  try {
    const xforward = headers().get("x-forwarded-for") ?? "";
    const agent = headers().get("User-Agent") ?? "";
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_FIRST_DOCUMENT,
      xforward: xforward,
      agent: agent,
      tags: [],
      values: {
        username: username,
        repo_name: decodeURIComponent(reponame),
        lang: locale == "" ? "en" : locale,
      },
    });
    console.log("error:", data);
    if (data.error) {
      throw new FetchError(data.message, data.status);
    }
    home_page = data.relative_path;
  } catch (error) {
    return <NotFound />;
  }
  redirect(`/workspace/${username}/o/${reponame}/${home_page}`); // Navigate to the new post page
}
