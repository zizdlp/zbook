"use client";
import { MenuStruct } from "@/types/interface";
import { useContext, useEffect, useState } from "react";
import { BiChevronDown, BiChevronLeft } from "react-icons/bi";
import SubMenu from "./SubMenu";
import { isPrefix } from "@/utils/util";
import { SideBarContext } from "@/providers/SideBarProvider";
export default function FrpcSubMenuDirItem({
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
  const { sideBarReload, setSideBarReload } = useContext(SideBarContext);
  const [open, setOpen] = useState(false);
  useEffect(() => {
    if (isPrefix(pathname, locale, prefix, menu.relative_path)) {
      setOpen(true);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [sideBarReload]);
  if (layer == 1) {
    return (
      <div className="mt-12 lg:mt-8">
        <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
          {menu.title}
        </h5>
        <li>
          <a
            className="pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg bg-primary/10 text-primary font-semibold dark:text-primary-light dark:bg-primary-light/10"
            href="/introduction"
          >
            <div className="flex-1 flex items-center space-x-2.5">
              <SubMenu
                prefix={prefix}
                menus={menu.sublayouts}
                layer={layer + 1}
                pathname={pathname}
                locale={locale}
              />
            </div>
          </a>
        </li>
      </div>
    );
  }
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
