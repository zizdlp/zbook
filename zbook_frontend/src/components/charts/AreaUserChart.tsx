"use client";
import dynamic from "next/dynamic";
import { useTranslations } from "next-intl";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
import EnableElement from "./EnableElement";

interface TrafficData {
  count: number;
  date: string;
}

interface WebTrafficProps {
  newUserCounts: TrafficData[];
  activeUserCounts: TrafficData[];
  allow_login: boolean;
  allow_registration: boolean;
}
function getArrary({
  counts,
  dates,
}: {
  counts: TrafficData[];
  dates: string[];
}) {
  // 如果counts未定义，将其设置为空数组
  counts = counts || [];
  dates = dates || [];

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
  return countsArray;
}
export default function AreaUserChart({
  newUserCounts,
  activeUserCounts,
  allow_login,
  allow_registration,
}: WebTrafficProps) {
  const { theme } = useTheme();
  const t = useTranslations("AdminOverView");
  const dates: string[] = [];
  for (let i = 0; i <= 6; ++i) {
    const date = new Date();
    date.setDate(date.getDate() - i);
    dates.push(date.toISOString().split("T")[0]); // 只获取日期部分
  }
  let newUserArray = getArrary({
    counts: newUserCounts,
    dates: dates,
  }).reverse();
  let activeUserArray = getArrary({
    counts: activeUserCounts,
    dates: dates,
  }).reverse();

  let options = {
    grid: {
      show: true,
      strokeDashArray: 4,
      borderColor: theme == "dark" ? "#1e293b" : "#cbd5e1", // 设置网格颜色
      borderOpacity: 0.1, // 设置网格透明度
      padding: {
        left: 2,
        right: 2,
        top: -26,
      },
    },
    chart: {
      height: "100%",
      maxWidth: "100%",
      type: "area" as "area",
      fontFamily: "Inter, sans-serif",
      dropShadow: {
        enabled: false,
      },
      toolbar: {
        show: false,
      },
    },
    tooltip: {
      enabled: true,
      x: {
        show: false,
      },
    },
    legend: {
      show: true,
      labels: {
        colors: theme == "dark" ? "#CBD5E1" : "#334155",
      },
    },
    fill: {
      type: "gradient",
      gradient: {
        opacityFrom: 0.55,
        opacityTo: 0,
        shade: "#1C64F2",
        gradientToColors: ["#1C64F2"],
      },
    },
    dataLabels: {
      enabled: false,
    },
    stroke: {
      width: 6,
    },
    xaxis: {
      categories: dates,
      labels: {
        show: false,
      },
      axisBorder: {
        show: false,
      },
      axisTicks: {
        show: false,
      },
    },
    yaxis: {
      show: false,
    },
  };
  let series = [
    {
      name: t("newUserLabel"),
      data: newUserArray,
      color: "#7E3BF2",
    },
    {
      name: t("activeUserLabel"),
      data: activeUserArray,
      color: "#1A56DB",
    },
  ];
  return (
    <div className="w-full h-[400px]  bg-gray-50 dark:bg-slate-800/50 dark:shadow-lg rounded-md p-4 md:p-6  border-[0.01rem] border-slate-300 dark:border-0">
      <div className="flex justify-between items-start w-full py-4">
        <h5 className="leading-none text-3xl font-bold text-gray-900 dark:text-white pb-2">
          {t("DailyUsers")}
        </h5>
        <div className="flex space-x-2">
          <EnableElement
            config_name="allow_login"
            label={t("allow_login")}
            initEnabled={allow_login}
          />

          <EnableElement
            config_name="allow_registration"
            label={t("allow_registration")}
            initEnabled={allow_registration}
          />
        </div>
      </div>
      <ApexChart
        type="area"
        options={options}
        series={series}
        height="260px"
        width="100%"
      />
    </div>
  );
}
