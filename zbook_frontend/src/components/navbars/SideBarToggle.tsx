"use client";
import { usePathname } from "@/navigation";
import { useLocale } from "next-intl";
import React, { useContext } from "react";
import { SideBarContext } from "@/providers/SideBarProvider";
import { RiMenu4Line, RiMenuLine } from "react-icons/ri";
import NavBarIcon from "./NavBarIcon";
export default function SideBarToggle() {
  const locale = useLocale();
  const pathname = usePathname();
  const { sideBarOpen, setSideBarOpen } = useContext(SideBarContext);

  const sideBarPages = ["/workspace"];
  const needSideBarPathnameRegex = RegExp(
    `^(/(${locale}))?(${sideBarPages
      .flatMap((p) => (p === "/" ? ["", "/"] : p))
      .join("|")})/?`,
    "i"
  );
  const isNeedSideBar = needSideBarPathnameRegex.test(pathname ?? "");
  if (!isNeedSideBar) {
    return <></>;
  }
  return (
    <NavBarIcon
      Icon={sideBarOpen === false ? RiMenuLine : RiMenu4Line}
      onClick={() => {
        setSideBarOpen(!sideBarOpen);
        localStorage.setItem("sidebarValue", JSON.stringify(!sideBarOpen));
      }}
      aria-label="sidebar toggle"
      mounted={true}
    />
  );
}
