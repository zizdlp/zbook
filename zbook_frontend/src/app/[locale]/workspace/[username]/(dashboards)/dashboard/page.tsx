import SomeThingWrong from "@/components/SomeThingWrong";
import PieChart from "@/components/charts/PieChart";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { Metadata } from "next";
import { getTranslations } from "next-intl/server";
import AreaChart from "@/components/charts/AreaChart";
import AreaUserChart from "@/components/charts/AreaUserChart";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
export async function generateMetadata({
  params,
}: {
  params: { username: string };
}): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: params.username + " - " + t("DashBoard") + " - " + t("OverView"),
  };
}
export default async function AdminOverviewPage({
  params,
  searchParams,
}: {
  params: { locale: string };
  searchParams?: { ndays?: string };
}) {
  const t = await getTranslations("AdminOverView");
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";
  try {
    const [
      repoCountRes,
      commentCountRes,
      commentReportCountRes,
      userCountRes,
      allow_registration,
      allow_login,
      allow_invitation,
    ] = await Promise.all([
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_REPO_COUNT,
        xforward,
        agent: agent,
        tags: [],
        values: { query: "" },
      }),
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_COUNT,
        xforward,
        agent: agent,
        tags: [],
        values: { query: "" },
      }),
      fetchServerWithAuthWrapper({
        endpoint:
          FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_REPORT_COUNT,
        xforward,
        agent: agent,
        tags: [],
        values: { query: "" },
      }),
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_USER_COUNT,
        xforward,
        agent: agent,
        tags: [],
        values: { query: "" },
      }),
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GetConfiguration,
        xforward,
        agent: agent,
        tags: [],
        values: { config_name: "allow_registration" },
      }),
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GetConfiguration,
        xforward,
        agent: agent,
        tags: [],
        values: { config_name: "allow_login" },
      }),
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GetConfiguration,
        xforward,
        agent: agent,
        tags: [],
        values: { config_name: "allow_invitation" },
      }),
    ]);
    if (repoCountRes.error) {
      throw new FetchError(repoCountRes.message, repoCountRes.status);
    }

    if (commentCountRes.error) {
      throw new FetchError(commentCountRes.message, commentCountRes.status);
    }
    if (commentReportCountRes.error) {
      throw new FetchError(
        commentReportCountRes.message,
        commentReportCountRes.status
      );
    }
    if (userCountRes.error) {
      throw new FetchError(userCountRes.message, userCountRes.status);
    }

    if (allow_registration.error) {
      throw new FetchError(
        allow_registration.message,
        allow_registration.status
      );
    }
    if (allow_login.error) {
      throw new FetchError(allow_login.message, allow_login.status);
    }

    if (allow_invitation.error) {
      throw new FetchError(allow_invitation.message, allow_invitation.status);
    }
    return (
      <>
        <div className="xl:col-span-3 md:col-span-2 col-span-1">
          <AreaChart title={t("DailyVisitors")} label={t("NewVisitor")} />
        </div>
        <div className="col-span-1 md:col-span-2">
          <AreaUserChart
            allow_login={allow_login.config_value}
            allow_registration={allow_registration.config_value}
            allow_invitation={allow_invitation.config_value}
          />
        </div>
        <div className="col-span-1">
          <PieChart
            repo_count={repoCountRes.count ?? 0}
            comment_count={commentCountRes.count ?? 0}
            comment_report_count={commentReportCountRes.count ?? 0}
            user_count={userCountRes.count ?? 0}
          />
        </div>
      </>
    );
  } catch (error) {
    let currentError = error as FetchError;
    logger.error(
      `AdminOverviewPage Error:${currentError.status} ${currentError.message}`,
      currentError.status
    );
    return <SomeThingWrong />;
  }
}
