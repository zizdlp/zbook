import { fetchServerWithoutAuthWrapper } from "@/fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";

export const dynamic = "force-dynamic"; // defaults to auto
export async function GET(request: Request) {
  try {
    const urlParams = new URLSearchParams(request.url.split("?")[1]);
    const username = urlParams.get("username");
    const repo_name = urlParams.get("repo_name");
    const sync_token = urlParams.get("sync_token");

    const data = await fetchServerWithoutAuthWrapper({
      endpoint: FetchServerWithoutAuthWrapperEndPoint.AUTO_SYNC_REPO,
      values: {
        username: username ?? "",
        repo_name: repo_name ?? "",
        sync_token: sync_token ?? "",
      },
      xforward: "",
      agent: "",
    });

    return Response.json(data);
  } catch (error) {
    console.error(error);
    return Response.error();
  }
}
