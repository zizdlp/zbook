"use client";
import dynamic from "next/dynamic";
import { useTranslations } from "next-intl";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
import EnableElement from "./EnableElement";
import { getAreaChartOptions } from "@/utils/const_value";

interface WebTrafficProps {
  newUserCounts: number[];
  activeUserCounts: number[];
  allow_login: boolean;
  allow_registration: boolean;
  allow_invitation: boolean;
  dates: string[];
}

export default function AreaUserChart({
  newUserCounts,
  dates,
  activeUserCounts,
  allow_login,
  allow_registration,
  allow_invitation,
}: WebTrafficProps) {
  const { theme } = useTheme();
  const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
  console.log("client timezone:", timezone);
  const t = useTranslations("AdminOverView");

  let options = getAreaChartOptions(theme, dates);

  let series = [
    {
      name: t("newUserLabel"),
      data: newUserCounts,
      color: "#7E3BF2",
    },
    {
      name: t("activeUserLabel"),
      data: activeUserCounts,
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
          <EnableElement
            config_name="allow_invitation"
            label={t("allow_invitation")}
            initEnabled={allow_invitation}
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
