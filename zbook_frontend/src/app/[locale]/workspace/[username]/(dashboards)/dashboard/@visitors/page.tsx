import SomeThingWrong from "@/components/SomeThingWrong";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import { getTranslations } from "next-intl/server";
import DonuChart from "@/components/charts/DonutChart";
import EarthChart from "@/components/charts/EarthChart";
import NDay from "@/components/charts/NDay";
import { promises as fs } from "fs";

import BarChart from "@/components/charts/BarChart";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
interface Visitor {
  ip: string;
  agent?: string;
  count: number;
  city?: string;
  lat?: number;
  long?: number;
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
  const user_agent = headers().get("User-Agent") ?? "";
  try {
    const data = await fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.GetDailyVisitors,
      xforward,
      agent: user_agent,
      tags: [],
      values: {
        lang: params.locale == "zh" ? "zh-CN" : params.locale,
        ndays: ndays,
      },
    });
    if (data.error) {
      throw new FetchError(data.message, data.status);
    }

    const { visitors, agent_count } = data;
    // 提取前5个IP和对应的count
    let ips: string[] = [];
    let counts: number[] = [];
    let cities: string[] = [];
    if (Array.isArray(visitors) && visitors.length > 0) {
      visitors.slice(0, 5).forEach((visitor: Visitor) => {
        ips.push(visitor.ip);
        counts.push(visitor.count ?? 0);
        cities.push(visitor.city ?? "");
      });
    } else {
      // 处理visitors为空或undefined的情况，比如初始化为空数组
      ips = [];
      counts = [];
      cities = [];
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
    console.log("all dona====");
    return (
      <>
        <div className="xl:col-span-2 md:col-span-2 col-span-1 md:row-span-2">
          <div className="w-full flex-col flex items-center justify-center bg-gray-50 dark:bg-slate-800/50 dark:shadow-lg rounded-md p-4 md:p-6  border-[0.01rem] border-slate-300 dark:border-0">
            <div className="flex justify-between items-start w-full py-4">
              <div className="flex justify-between mb-5">
                <div>
                  <h5 className="leading-none text-3xl font-bold text-gray-900 dark:text-white pb-2">
                    {visitors?.length ?? 0}
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
                markers={visitors}
                isSmall={true}
              />
            </div>
            <div className="hidden md:block">
              <EarthChart
                key={ndays}
                landData={landData}
                lakeData={lakeData}
                riverData={riverData}
                markers={visitors}
                isSmall={false}
              />
            </div>
          </div>
        </div>
        <div className="col-span-1">
          <DonuChart agentCounts={agent_count} />
        </div>
        <div className="col-span-1">
          <BarChart
            ips={ips}
            counts={counts}
            cities={cities}
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
