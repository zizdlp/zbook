/* eslint-disable react/jsx-no-literals */
import type { MenuStruct, SubMenuProps } from "@/types/interface";
import { isSameUrl } from "@/utils/util";

export default function FrpcSubMenu({
  prefix,
  menus,
  layer,
  pathname,
  locale,
}: any) {
  return (
    <>
      {menus.map((item: MenuStruct, index: any) => (
        <div className="mt-12 lg:mt-8">
          <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
            {item.title}
          </h5>
          {item.sublayouts.map((subitem: MenuStruct, index: any) => (
            <li>
              <a
                className={`${
                  isSameUrl(pathname, locale, prefix, subitem.relative_path)
                    ? "pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg bg-purple-600/10 text-purple-700 font-semibold dark:text-purple-400 dark:bg-purple-400/10"
                    : "pl-[1rem] group mt-2 lg:mt-0 flex items-center pr-3 py-1.5 cursor-pointer focus:outline-primary dark:focus:outline-primary-light space-x-3 rounded-lg hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
                }`}
                href={prefix + subitem.relative_path}
              >
                <div className="flex-1 flex items-center space-x-2.5">
                  <div>{subitem.title}</div>
                </div>
              </a>
            </li>
          ))}
        </div>
      ))}
    </>
  );
}
