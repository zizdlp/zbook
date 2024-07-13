"use client";
import { useSearchParams } from "next/navigation";
import { usePathname, useRouter } from "@/navigation";
import { useTranslations } from "next-intl";
export default function NDay({ oldNdays }: { oldNdays: number }) {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();
  const t = useTranslations("AdminOverView");
  function handleSearch(ndays: string) {
    if (searchParams) {
      const params = new URLSearchParams(searchParams);
      if (ndays) {
        params.set("ndays", ndays);
      } else {
        params.delete("ndays");
      }
      replace(`${pathname}?${params.toString()}`);
    }
  }
  return (
    <div className="flex items-center space-x-1 bg-white dark:bg-slate-800 border border-slate-300 dark:border-slate-700 rounded-md px-2 py-1 my-1 md:text-sm text-xs">
      <div
        onClick={() => {
          handleSearch("1");
        }}
        className={`cursor-pointer md:px-4 px-2 md:py-2 py-1.5 rounded-md ${
          oldNdays == 1
            ? "bg-sky-500 text-white dark:bg-sky-600"
            : "text-slate-600 dark:text-slate-200"
        }`}
      >
        {t("LastDay")}
      </div>
      <div
        onClick={() => {
          handleSearch("7");
        }}
        className={`cursor-pointer md:px-4 px-2 md:py-2 py-1.5 rounded-md ${
          oldNdays == 7
            ? "bg-sky-500 text-white dark:bg-sky-600"
            : "text-slate-600 dark:text-slate-200"
        }`}
      >
        {t("LastWeek")}
      </div>
      <div
        onClick={() => {
          handleSearch("31");
        }}
        className={`cursor-pointer md:px-4 px-2 md:py-2 py-1.5 rounded-md ${
          oldNdays == 31
            ? "bg-sky-500 text-white dark:bg-sky-600"
            : "text-slate-600 dark:text-slate-200"
        }`}
      >
        {t("LastMonth")}
      </div>
    </div>
  );
}
