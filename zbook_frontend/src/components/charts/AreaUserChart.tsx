"use client";
import dynamic from "next/dynamic";
import { useTranslations } from "next-intl";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";
import EnableElement from "./EnableElement";
import { getAreaChartOptions } from "@/utils/const_value";
import { useEffect, useState } from "react";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";

interface WebTrafficProps {
  allow_login: boolean;
  allow_registration: boolean;
  allow_invitation: boolean;
}

export default function AreaUserChart({
  allow_login,
  allow_registration,
  allow_invitation,
}: WebTrafficProps) {
  const { theme } = useTheme();
  const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
  const [newUserCounts, setNewUserCounts] = useState<number[]>([]);
  const [activeUserCounts, setActiveUserCounts] = useState<number[]>([]);
  const [dates, setDates] = useState<string[]>([]);

  useEffect(() => {
    fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.DAILY_CREATE_USER_COUNT,
      xforward: "",
      agent: "",
      tags: [],
      values: {
        time_zone: timezone,
        ndays: 7,
      },
    }).then((data) => {
      setNewUserCounts(data.counts);

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

    fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.DAILY_ACTIVE_USER_COUNT,
      xforward: "",
      agent: "",
      tags: [],
      values: { time_zone: timezone, ndays: 7 },
    }).then((data) => {
      setActiveUserCounts(data.counts);
    });
  }, []);

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
