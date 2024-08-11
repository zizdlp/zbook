import { fetchServerWithoutAuthWrapper } from "@/fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
import { headers } from "next/headers";
export default async function LogVisitor() {
  const xforward = headers().get("x-forwarded-for") ?? "";
  const userAgent = headers().get("User-Agent") ?? "";
  try {
    const data = await fetchServerWithoutAuthWrapper({
      endpoint: FetchServerWithoutAuthWrapperEndPoint.LOG_VISITOR,
      xforward: xforward,
      agent: userAgent,
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
