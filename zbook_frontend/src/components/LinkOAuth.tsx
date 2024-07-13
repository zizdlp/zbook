import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { redirect } from "@/navigation";
import { auth } from "@/auth";
import { createOAuthLink } from "@/fetchs/server_without_auth";
import { FetchError } from "@/fetchs/util";
import CallSignIn from "./CallSignIn";
import { logger } from "@/utils/logger";
export default async function LinkOAuth({
  searchParams,
  oauthType,
}: {
  searchParams?: { access_token?: string };
  oauthType: string;
}) {
  const access_token = searchParams?.access_token || "";
  const session = await auth();
  let redirect2workspace = false;
  try {
    if (session?.app_id) {
      //should created link
      let values = {
        oauth_type: oauthType,
        app_id: session.app_id ?? "",
      };

      const data = await createOAuthLink(values, access_token ?? "");
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      return <CallSignIn oauthType={oauthType} />;
    } else if (session?.access_token) {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CHECK_OAUTH_STATUS,
        xforward: "",
        agent: "",
        tags: [],
        values: {},
      });
      if (oauthType == "github") {
        if (data?.github) {
          redirect2workspace = true;
        } else {
          return <CallSignIn oauthType={oauthType} />;
        }
      } else if (oauthType == "google") {
        if (data?.google) {
          redirect2workspace = true;
        } else {
          return <CallSignIn oauthType={oauthType} />;
        }
      }
    } else {
      return <></>;
    }
  } catch (error) {
    let e = error as FetchError;
    logger.error(`LinkOAuth failed:${e.message}`, e.status);
    return <></>;
  }
  if (redirect2workspace) {
    redirect(`/workspace`); // Navigate to the new post page
  }
  return <></>;
}
