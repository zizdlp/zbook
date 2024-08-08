"use client";
import dynamic from "next/dynamic";
import { useTranslations } from "next-intl";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
import { getAreaChartOptions } from "@/utils/const_value";
interface TrafficData {
  count: number;
  date: string;
}

interface WebTrafficProps {
  counts: TrafficData[];
  title: string;
  label: string;
}
export default function AreaChart({ counts, title, label }: WebTrafficProps) {
  const { theme } = useTheme();
  const t = useTranslations("AdminOverView");
  const dates: string[] = [];
  for (let i = 30; i >= 0; i--) {
    const date = new Date();
    date.setDate(date.getDate() - i);
    dates.push(date.toISOString().split("T")[0]); // 只获取日期部分
  }

  counts = counts || [];
  // 构建日期到计数的映射
  const countsMap = new Map();
  counts.forEach(({ date, count }) => {
    let datePart = date;
    if (date.includes("T")) {
      datePart = date.split("T")[0]; // 只获取日期部分
    }
    countsMap.set(datePart, count || 0);
  });

  // 生成计数数组
  const countsArray: number[] = [];
  dates.forEach((date) => {
    countsArray.push(
      parseInt(countsMap.has(date) ? countsMap.get(date) : 0) || 0
    );
  });
  const totalCount = countsArray.reduce((sum, count) => sum + count, 0);
  let options = getAreaChartOptions(theme, dates);
  let series = [
    {
      name: label,
      data: countsArray,
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
