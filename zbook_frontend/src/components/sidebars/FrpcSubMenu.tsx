/* eslint-disable react/jsx-no-literals */
"use client";
import type { MenuStruct, SubMenuProps } from "@/types/interface";
import SubMenuFileItem from "./SubMenuFileItem";
import SubMenu from "./SubMenu";

export default function FrpcSubMenu({
  prefix,
  menus,
  layer,
  pathname,
  locale,
}: any) {
  return (
    <>
      {menus.map((menu: MenuStruct, index: any) =>
        menu.isdir ? (
          <li key={index}>
            <div className="mt-12 lg:mt-8">
              <h5 className="pl-4 mb-3.5 lg:mb-2.5 font-semibold text-gray-900 dark:text-gray-200">
                {menu.title}
              </h5>

              <SubMenu
                prefix={prefix}
                menus={menu.sublayouts}
                layer={layer + 1}
                pathname={pathname}
                locale={locale}
                collapse={true}
              />
            </div>
          </li>
        ) : (
          <li key={index}>
            <SubMenuFileItem
              layer={layer}
              menu={menu}
              pathname={pathname}
              locale={locale}
              prefix={prefix}
              collapse={false}
            />
          </li>
        )
      )}
    </>
  );
}
