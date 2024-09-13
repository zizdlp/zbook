"use client";
import dynamic from "next/dynamic";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
import { getAreaChartOptions } from "@/utils/const_value";
import { useEffect, useState } from "react";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";

export default function AreaVisitorChart({
  title,
  label,
}: {
  title: string;
  label: string;
}) {
  const { resolvedTheme } = useTheme();
  const [counts, setCounts] = useState<number[]>([]);
  const [dates, setDates] = useState<string[]>([]);
  const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;

  useEffect(() => {
    fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.DAILY_VISITOR_COUNT,
      xforward: "",
      agent: "",
      tags: [],
      values: { ndays: 31, time_zone: timezone },
    }).then((data) => {
      setCounts(data.counts);
      // 生成与 counts 等长的日期序列
      const today = new Date();
      const generatedDates = Array.from(
        { length: data.counts.length },
        (_, i) => {
          const date = new Date();
          date.setDate(today.getDate() - i);
          return date.toISOString().split("T")[0]; // 获取 YYYY-MM-DD 格式的日期
        }
      ).reverse();
      setDates(generatedDates);
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const totalCount = counts.reduce((sum, count) => sum + count, 0);
  let options = getAreaChartOptions(resolvedTheme, dates);
  let series = [
    {
      name: label,
      data: counts,
      color: "#7E3BF2",
    },
  ];
  return (
    <div className="w-full h-[500px]  bg-gray-50 dark:bg-slate-800/50 dark:shadow-lg rounded-md p-4 md:p-6  border-[0.01rem] border-slate-300 dark:border-0">
      <div className="flex justify-between items-start w-full py-4">
        <div className="flex justify-between mb-5">
          <div>
            <h5 className="leading-none text-3xl font-bold text-gray-900 dark:text-white pb-2">
              {totalCount}
            </h5>
            <p className="text-base font-normal text-gray-500 dark:text-gray-400">
              {title}
            </p>
          </div>
        </div>
      </div>

      <ApexChart
        type="area"
        options={options}
        series={series}
        height="320px"
        width="100%"
      />
    </div>
  );
}
