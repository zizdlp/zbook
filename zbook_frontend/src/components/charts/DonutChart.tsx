"use client";
import dynamic from "next/dynamic";
import { useTranslations } from "next-intl";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
interface AgentCount {
  computer: number;
  phone: number;
  tablet: number;
  bot: number;
  unknown: number;
}
export default function DonuChart({
  agentCounts,
}: {
  agentCounts: AgentCount;
}) {
  const { theme } = useTheme();
  const t = useTranslations("AdminOverView");
  let options = {
    colors: ["#14b8a6", "#0ea5e9", "#f28e4b", "#e11d48", "#6366f1"],
    chart: {
      height: "100%",
      width: "100%",
      type: "donut" as "donut",
    },
    stroke: {
      colors: ["transparent"],
    },
    grid: {
      padding: {
        top: -2,
      },
    },
    labels: [t("Computer"), t("Phone"), t("Tablet"), t("Bot"), t("Unknown")],
    dataLabels: {
      enabled: false,
    },
    legend: {
      position: "bottom" as "bottom",
      fontFamily: "Inter, sans-serif",
      labels: {
        colors: theme == "dark" ? "#CBD5E1" : "#334155",
      },
    },
    yaxis: {
      labels: {
        formatter: function (value: any) {
          return value;
        },
      },
    },
    xaxis: {
      labels: {
        formatter: function (value: any) {
          return value;
        },
      },
      axisTicks: {
        show: false,
      },
      axisBorder: {
        show: false,
      },
    },
  };
  let series = [
    agentCounts.computer ?? 0,
    agentCounts.phone ?? 0,
    agentCounts.tablet ?? 0,
    agentCounts.bot ?? 0,
    agentCounts.unknown ?? 0,
  ];
  return (
    <div className="w-full h-[400px]  bg-gray-50 dark:bg-slate-800/50 dark:shadow-lg rounded-md p-4 md:p-6  border-[0.01rem] border-slate-300 dark:border-0">
      <div className="flex justify-between items-start w-full">
        <div className="flex justify-between mb-5">
          <div>
            <h5 className="leading-none text-3xl font-bold text-gray-900 dark:text-white pb-2">
              {t("VisitorAnalysis")}
            </h5>
          </div>
        </div>
      </div>

      <ApexChart
        type="donut"
        options={options}
        series={series}
        height="320px"
        width="100%"
      />
    </div>
  );
}
