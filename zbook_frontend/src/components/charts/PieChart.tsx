"use client";
import dynamic from "next/dynamic";
import { useTranslations } from "next-intl";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
export default function PieChart({
  repo_count,
  comment_count,
  comment_report_count,
  user_count,
}: {
  repo_count: string;
  comment_count: string;
  comment_report_count: string;
  user_count: string;
}) {
  const { theme } = useTheme();
  const t = useTranslations("AdminOverView");
  let repo_count_n = parseInt(repo_count);
  let comment_count_n = parseInt(comment_count);
  let comment_report_count_n = parseInt(comment_report_count);
  let user_count_n = parseInt(user_count);

  let options = {
    colors: ["#1C64F2", "#16BDCA", "#9061F9", "#8FAAF9"],
    labels: [t("Repositories"), t("Comments"), t("CommentReports"), t("Users")],
    legend: {
      fontSize: "14px",
      fontFamily: "Inter, sans-serif",
      position: "bottom" as "bottom",
      labels: {
        colors: theme == "dark" ? "#CBD5E1" : "#334155",
      },
    },
    total: {
      color: "#373d3f",
    },
    stroke: {
      colors: ["white"],
      width: theme == "dark" ? 0 : 1,
      lineCap: "square" as "square",
    },
  };
  let series = [
    repo_count_n,
    comment_count_n,
    comment_report_count_n,
    user_count_n,
  ];

  return (
    <div className="w-full h-[400px]  bg-gray-50 dark:bg-slate-800/50 dark:shadow-lg rounded-md p-4 md:p-6  border-[0.01rem] border-slate-300 dark:border-0">
      <div className="flex justify-between items-start w-full py-4">
        <div className="flex-col items-center">
          <div className="flex items-center mb-1">
            <h5 className="text-xl font-bold leading-none text-gray-900 dark:text-white me-1">
              {t("ResourceDistribution")}
            </h5>
          </div>
        </div>
      </div>
      <ApexChart
        type="pie"
        options={options}
        series={series}
        height="320px"
        width="100%"
      />
    </div>
  );
}
