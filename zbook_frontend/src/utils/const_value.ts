const emailRegex =
  /^(([^<>()$$$$\\.,;:\s@"]+(\.[^<>()$$$$\\.,;:\s@"]+)*)|(".+"))@(($$[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$$)|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export { emailRegex };

// 定义枚举类型
export enum SearchType {
  DOCUMENT = 0,
  USER = 1,
  USER_DOCUMENT = 2,
  REPO_DOCUMENT = 3,
  VISI_USER = 4,
}

export function getAreaChartOptions(theme: any, dates: string[]): any {
  return {
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
      theme: theme == "dark" ? "dark" : "light",
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
}
