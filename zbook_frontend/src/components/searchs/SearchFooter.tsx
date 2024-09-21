import React from "react";
import { Link } from "@/navigation";
import Image from "next/image";
import { useTranslations } from "next-intl";
export default function SearchFooter() {
  const t = useTranslations("Search");
  return (
    <footer className="px-4 py-4 relative flex justify-end text-slate-500 flex-row items-center border-t-[0.05rem]  border-slate-300 dark:border-slate-600/25">
      <div className="flex flex-row items-center px-4">
        <Link href="/" className="flex-col items-center">
          <h1 className="font-semibold text-sm  dark:text-slate-200 cursor-pointer pr-2">
            <span className="bg-gradient-to-r from-teal-500 to-blue-500 inline-block text-transparent bg-clip-text">
              {t("AppName")}
            </span>
          </h1>
        </Link>
        <Link href="/" className="flex-none overflow-hidden">
          <picture>
            <Image
              src="/logo_128.png"
              alt=""
              className="w-7 h-7 rounded-md"
              width={128}
              height={128}
            />
          </picture>
        </Link>
      </div>
    </footer>
  );
}
