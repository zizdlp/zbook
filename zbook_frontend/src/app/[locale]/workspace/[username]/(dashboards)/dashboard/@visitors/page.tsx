import SomeThingWrong from "@/components/SomeThingWrong";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { Metadata } from "next";
import { getTranslations } from "next-intl/server";
import DonuChart from "@/components/charts/DonutChart";
import EarthChart from "@/components/charts/EarthChart";
import NDay from "@/components/charts/NDay";
import { promises as fs } from "fs";

import BarChart from "@/components/charts/BarChart";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
import { parseUserAgent } from "@/utils/util";
interface Visitor {
  IP: string;
  Agent?: string;
  Count: number;
  city?: string;
  lat?: number;
  long?: number;
}
function hasValidCoordinates(visitor: Visitor): visitor is Visitor {
  return typeof visitor.lat === "number" && typeof visitor.long === "number";
}

export default async function AdminOverviewPage({
  params,
  searchParams,
}: {
  params: { locale: string };
  searchParams?: { ndays?: string };
}) {
  const ndays = Number(searchParams?.ndays) || 1;
  const t = await getTranslations("AdminOverView");
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";
  try {
    const dailyVisitors = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GetDailyVisitors,
      xforward,
      agent: agent,
      tags: [],
      values: {
        lang: params.locale == "zh" ? "zh-CN" : params.locale,
        ndays: ndays,
      },
    });
    if (dailyVisitors.error) {
      throw new FetchError(dailyVisitors.message, dailyVisitors.status);
    }
    const landFile = await fs.readFile(
      process.cwd() + "/public/ne_110m_land.geojson",
      "utf8"
    );
    const lakeFile = await fs.readFile(
      process.cwd() + "/public/ne_110m_lakes.geojson",
      "utf8"
    );
    const riverFile = await fs.readFile(
      process.cwd() + "/public/ne_110m_rivers_lake_centerlines.geojson",
      "utf8"
    );
    const landData = JSON.parse(landFile);
    const lakeData = JSON.parse(lakeFile);
    const riverData = JSON.parse(riverFile);

    // 聚合 IP 计数
    const aggregated = dailyVisitors.visitors.reduce(
      (acc: any, visitor: Visitor) => {
        if (!acc[visitor.IP]) {
          acc[visitor.IP] = 0;
        }
        acc[visitor.IP] += visitor.Count;
        return acc;
      },
      {} as Record<string, number>
    );

    // 转换为数组并排序
    const sorted = Object.entries(aggregated)
      .sort((a: any, b: any) => b[1] - a[1])
      .slice(0, 5);

    // 分别提取 IP 和计数
    const ips = sorted.map((entry) => entry[0]);
    const counts = sorted.map((entry) => entry[1] as number);

    // 分类统计
    const agentCounts = {
      computer: 0,
      phone: 0,
      tablet: 0,
      bot: 0,
      unknown: 0,
    };

    dailyVisitors.visitors.forEach((visitor: any) => {
      const agentString = parseUserAgent(visitor.Agent).platform.toLowerCase();
      const osString = parseUserAgent(visitor.Agent).os.toLowerCase();
      const agent = visitor.Agent;
      let visited = false;
      if (agent && typeof agent === "string") {
        const agentLowerCase = agent.toLowerCase();
        if (
          agentLowerCase.includes("bot") ||
          agentLowerCase.includes("spider")
        ) {
          visited = true;
          agentCounts.bot++;
        }
      }
      if (visited) {
      } else if (
        agentString.includes("windows") ||
        agentString.includes("macintosh") ||
        (agentString.includes("linux") && !osString.includes("android"))
      ) {
        agentCounts.computer++;
      } else if (
        agentString.includes("iphone") ||
        agentString.includes("android") ||
        osString.includes("android") ||
        agentString.includes("windows phone") ||
        agentString.includes("blackberry")
      ) {
        agentCounts.phone++;
      } else if (
        agentString.includes("ipad") ||
        agentString.includes("android tablet") ||
        agentString.includes("kindle")
      ) {
        agentCounts.tablet++;
      } else {
        agentCounts.unknown++;
      }
    });
    let filteredVisitors = dailyVisitors.visitors.filter(hasValidCoordinates);
    return (
      <>
        <div className="xl:col-span-2 md:col-span-2 col-span-1 md:row-span-2">
          <div className="w-full flex-col flex items-center justify-center bg-gray-50 dark:bg-slate-800/50 dark:shadow-lg rounded-md p-4 md:p-6  border-[0.01rem] border-slate-300 dark:border-0">
            <div className="flex justify-between items-start w-full py-4">
              <div className="flex justify-between mb-5">
                <div>
                  <h5 className="leading-none text-3xl font-bold text-gray-900 dark:text-white pb-2">
                    {dailyVisitors.visitors.length}
                  </h5>
                  <p className="text-base font-normal text-gray-500 dark:text-gray-400">
                    {t("VisitorRegion")}
                  </p>
                </div>
              </div>
              <NDay oldNdays={ndays} />
            </div>
            <div className="md:hidden block">
              <EarthChart
                key={ndays}
                landData={landData}
                lakeData={lakeData}
                riverData={riverData}
                markers={filteredVisitors}
                isSmall={true}
              />
            </div>
            <div className="hidden md:block">
              <EarthChart
                key={ndays}
                landData={landData}
                lakeData={lakeData}
                riverData={riverData}
                markers={filteredVisitors}
                isSmall={false}
              />
            </div>
          </div>
        </div>
        <div className="col-span-1">
          <DonuChart agentCounts={agentCounts} />
        </div>
        <div className="col-span-1">
          <BarChart
            ips={ips}
            counts={counts}
            title={t("TopVisitors")}
            label={t("VisitedCount")}
          />
        </div>
      </>
    );
  } catch (error) {
    let currentError = error as FetchError;
    logger.error(
      `Dashboard visitors Error:${currentError.status} ${currentError.message}`
    );
    return <SomeThingWrong />;
  }
}
