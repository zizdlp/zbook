"use client";
import { MenuStruct } from "@/types/interface";
import { useContext, useEffect, useState } from "react";
import { BiChevronDown, BiChevronLeft } from "react-icons/bi";
import FoldSubMenu from "./FoldSubMenu";
import { isPrefix } from "@/utils/util";
import { SideBarContext } from "@/providers/SideBarProvider";
import { ThemeColor } from "../TableOfContent";
export default function SubMenuDirItem({
  layer,
  menu,
  pathname,
  locale,
  prefix,
  collapse,
  theme_color,
}: {
  layer: number;
  menu: MenuStruct;
  pathname: string;
  locale: string;
  prefix: string;
  collapse: boolean;
  theme_color: ThemeColor;
}) {
  const { sideBarReload, setSideBarReload } = useContext(SideBarContext);
  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (
      isPrefix(pathname, locale, prefix, menu.relative_path.toLocaleLowerCase())
    ) {
      setOpen(true);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [sideBarReload]);
  return (
    <>
      <div
        onClick={() => {
          setOpen(!open);
        }}
        className={`group mt-2 lg:mt-0 cursor-pointer  ${
          layer == 1
            ? collapse
              ? "pl-0"
              : "pl-4"
            : layer == 2
              ? collapse
                ? "pl-4"
                : "pl-7"
              : layer == 3
                ? collapse
                  ? "pl-7"
                  : "pl-10"
                : collapse
                  ? "pl-10"
                  : "pl-14"
        } flex items-center  pr-3 py-1.5 rounded-lg my-0.5
        ${"hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"}
          `}
      >
        <div className="flex-1 flex items-center space-x-2.5">
          <div> {menu.title}</div>
        </div>

        <div
          className={`${(open || !menu.sublayouts) && "hidden"}`}
          onClick={() => {
            setOpen(true);
          }}
        >
          <BiChevronLeft className="hover:text-sky-500 duration-300 cursor-pointer md:h-4 md:w-4 h-6 w-6" />
        </div>
        <div
          className={`${(!open || !menu.sublayouts) && "hidden"}`}
          onClick={() => {
            setOpen(false);
          }}
        >
          <BiChevronDown className="hover:text-sky-500 duration-300 cursor-pointer md:h-4 md:w-4 h-6 w-6" />
        </div>
      </div>
      {open && (
        <FoldSubMenu
          prefix={prefix}
          menus={menu.sublayouts}
          layer={layer + 1}
          pathname={pathname}
          locale={locale}
          collapse={collapse}
          theme_color={theme_color}
        />
      )}
    </>
  );
}
