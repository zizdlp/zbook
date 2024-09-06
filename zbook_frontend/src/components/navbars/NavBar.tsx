import React from "react";
import { Link } from "@/navigation";
import Image from "next/image";
import NavTheme from "@/components/navbars/NavTheme";
import NavUserState from "@/components/navbars/NavUserState";
import NavNoti from "@/components/navbars/NavNoti";
import SideBarToggle from "@/components/navbars/SideBarToggle";
import NavLang from "./NavLang";
import { getTranslations } from "next-intl/server";
export default async function NavBar() {
  const t = await getTranslations("HomePage");
  return (
    <div
      className="sticky top-0 flex justify-center z-50 
      border-b border-slate-900/10 dark:border-slate-50/[0.06] print:hidden
      backdrop-blur-xl xl:backdrop-blur supports-backdrop-blur:bg-white/60 bg-transparent"
    >
      <div
        className="w-full flex justify-between items-center 
        py-2 px-4 max-w-[93rem] lg:px-12"
      >
        <div className="flex items-center">
          <div className="lg:hidden">
            <SideBarToggle />
          </div>

          <Link
            href={"/"}
            className="flex items-center justify-center overflow-hidden cursor-pointer"
          >
            <Image
              src="/logo_128.png"
              alt="Picture of logo"
              className="w-8 h-8 md:w-8 md:h-8 my-0.5 hidden lg:block rounded-md"
              width={128}
              height={128}
            />
            <p className="font-bold text-xl md:text-xl mx-2">
              <span className="bg-gradient-to-r from-teal-500 to-blue-500 inline-block text-transparent bg-clip-text">
                {t("AppName")}
              </span>
            </p>
          </Link>
        </div>

        <div className="flex items-center justify-center space-x-2">
          <Link
            className="cursor-pointer hover:text-sky-600 dark:hover:text-sky-400 font-semibold"
            href={"/cases"}
          >
            {t("Cases")}
          </Link>
          <div className="flex gap-x-2 md:gap-x-3 justify-center items-center px-1 md:px-6	mx-3 md:mx-6  md:border-l border-slate-300 dark:border-slate-400">
            <NavUserState />
            <NavTheme />
            <NavNoti />
            <NavLang />
          </div>
        </div>
      </div>
    </div>
  );
}
