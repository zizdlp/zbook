import { useTranslations } from "next-intl";
export default function TimeElement({ timeInfo }: { timeInfo: string }) {
  const ut = useTranslations("Time");
  function convertTime(givenTime: string): string {
    // 解析给定时间字符串
    const ct = new Date(givenTime);

    // 获取当前时间和时区
    const now = new Date();

    // 计算时间差
    const diff = ct.getTime() - now.getTime();

    // 小于一分钟
    if (Math.abs(diff) < 60 * 1000) {
      return ut("JustNow");
    }
    // 小于一小时
    else if (Math.abs(diff) < 60 * 60 * 1000) {
      const minutes = Math.floor(Math.abs(diff) / (60 * 1000));
      if (diff < 0) {
        return ut("MinuteAgo", { duration: minutes });
      } else {
        return ut("MinuteAfter", { duration: minutes });
      }
    }
    // 小于一天
    else if (Math.abs(diff) < 24 * 60 * 60 * 1000) {
      const hours = Math.floor(Math.abs(diff) / (60 * 60 * 1000));
      if (diff < 0) {
        return ut("HourAgo", { duration: hours });
      } else {
        return ut("HourAfter", { duration: hours });
      }
    }
    // 大于一天
    else {
      const days = Math.floor(Math.abs(diff) / (24 * 60 * 60 * 1000));
      // return `${days} 天${diff > 0 ? "后" : "前"}`;
      if (diff < 0) {
        return ut("DayAgo", { duration: days });
      } else {
        return ut("DayAfter", { duration: days });
      }
    }
  }
  return <span className="whitespace-nowrap">{convertTime(timeInfo)}</span>;
}
