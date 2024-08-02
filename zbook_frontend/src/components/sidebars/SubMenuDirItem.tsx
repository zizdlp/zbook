"use client";
import { MenuStruct } from "@/types/interface";
import { useEffect, useState } from "react";
import { BiChevronDown, BiChevronLeft } from "react-icons/bi";
import SubMenu from "./SubMenu";
import { isPrefix } from "@/utils/util";
export default function SubMenuDirItem({
  layer,
  menu,
  pathname,
  locale,
  prefix,
}: {
  layer: number;
  menu: MenuStruct;
  pathname: string;
  locale: string;
  prefix: string;
}) {
  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (isPrefix(pathname, locale, prefix, menu.relative_path)) {
      setOpen(true);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return (
    <>
      <div
        onClick={() => {
          setOpen(!open);
        }}
        className={`group mt-2 lg:mt-0 cursor-pointer  ${
          layer == 1
            ? "pl-4"
            : layer == 2
            ? "pl-7"
            : layer == 3
            ? "pl-10"
            : "pl-14"
        } flex items-center -ml-4 pr-3 py-1.5 rounded-lg my-0.5
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
        <SubMenu
          prefix={prefix}
          menus={menu.sublayouts}
          layer={layer + 1}
          pathname={pathname}
          locale={locale}
        />
      )}
    </>
  );
}
