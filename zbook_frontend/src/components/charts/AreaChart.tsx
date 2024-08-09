"use client";
import dynamic from "next/dynamic";
import { useTranslations } from "next-intl";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
import { getAreaChartOptions } from "@/utils/const_value";

export default function AreaChart({
  dates,
  counts,
  title,
  label,
}: {
  dates: string[];
  counts: number[];
  title: string;
  label: string;
}) {
  const { theme } = useTheme();
  const t = useTranslations("AdminOverView");

  counts = counts || [];
  dates = dates || [];
  const totalCount = counts.reduce((sum, count) => sum + count, 0);
  let options = getAreaChartOptions(theme, dates);
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
              {t("DailyVisitors")}
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
