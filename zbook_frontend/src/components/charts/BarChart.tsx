"use client";
import dynamic from "next/dynamic";
const ApexChart = dynamic(() => import("react-apexcharts"), { ssr: false });
import { useTheme } from "next-themes";

interface WebTrafficProps {
  ips: string[];
  counts: number[];
  cities: string[];
  title: string;
  label: string;
}

// 定义函数，根据条件生成结果数组
function getResults(ipArray: string[], cityArray: string[]): string[] {
  const ret: string[] = [];
  const length = Math.max(ipArray.length, cityArray.length);

  for (let i = 0; i < length; i++) {
    const ipValue = ipArray[i];
    const cityValue = cityArray[i];

    // 使用条件运算符选择 ip 或 city
    ret.push(cityValue ? cityValue + " " + ipValue : ipValue);
  }

  return ret;
}
export default function BarChart({
  ips,
  counts,
  cities,
  title,
  label,
}: WebTrafficProps) {
  const { resolvedTheme } = useTheme();
  var options = {
    chart: {
      type: "bar" as "bar",
      toolbar: {
        show: false,
      },
    },
    grid: {
      borderColor: resolvedTheme === "dark" ? "#334155" : "#E2E8F0",
      xaxis: {
        lines: {
          show: false,
        },
      },
      yaxis: {
        lines: {
          show: true,
        },
      },
    },
    tooltip: {
      theme: resolvedTheme == "dark" ? "dark" : "light",
      enabled: true,
      x: {
        show: false,
      },
    },
    plotOptions: {
      bar: {
        borderRadius: 4,
        borderRadiusApplication: "end" as "end",
        horizontal: true,
      },
    },
    dataLabels: {
      enabled: false,
    },
    legend: {
      show: true,
      labels: {
        colors: resolvedTheme == "dark" ? "#CBD5E1" : "#334155",
      },
    },

    stroke: {
      width: 2,
    },
    xaxis: {
      categories: getResults(ips, cities),
      labels: {
        style: {
          colors: resolvedTheme === "dark" ? "#CBD5E1" : "#334155",
        },
      },
    },
    yaxis: {
      show: true,
      labels: {
        style: {
          colors: resolvedTheme === "dark" ? "#CBD5E1" : "#334155",
        },
      },
    },
  };
  let series = [
    {
      name: label,
      data: counts,
    },
  ];
  return (
    <div className="w-full h-[400px] bg-gray-50 dark:bg-slate-800/50 dark:shadow-lg rounded-md p-4 md:p-6  border-[0.01rem] border-slate-300 dark:border-0">
      <div className="flex justify-between items-start w-full">
        <div className="flex justify-between mb-5">
          <div>
            <h5 className="leading-none text-3xl font-bold text-gray-900 dark:text-white pb-2">
              {title}
            </h5>
          </div>
        </div>
      </div>
      <ApexChart
        type="bar"
        options={options}
        series={series}
        height="320px"
        width="100%"
      />
    </div>
  );
}
