import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
import { headers } from "next/headers";
export default async function LogVisitor() {
  const xforward = headers().get("x-forwarded-for") ?? "";
  const userAgent = headers().get("User-Agent") ?? "";
  try {
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.LOG_VISITOR,
      xforward: xforward,
      agent: userAgent,
      tags: [],
      values: {},
    });
    if (data.error) {
      throw new FetchError(data.message, data.status);
    }
    return <></>;
  } catch (error) {
    let e = error as FetchError;
    logger.error(`LogVisitor failed:${e.message}`, e.status);
    return <></>;
  }
}
