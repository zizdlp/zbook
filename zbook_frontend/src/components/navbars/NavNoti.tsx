import { auth } from "@/auth";
import NavNotification from "./NavNotification";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
export default async function UserState() {
  const session = await auth();
  if (session && session.access_token) {
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GET_UNREAD_COUNT,
      xforward: "",
      agent: "",
      tags: [],
      values: {},
    });
    return (
      <NavNotification
        username={session.username}
        access_token={session.access_token}
        websocket_url={process.env.WEBSOCKET_URL ?? ""}
        unread_count={data.unread_count ?? 0}
      />
    );
  } else {
    return <></>;
  }
}
