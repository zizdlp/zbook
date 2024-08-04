import type { SubMenuProps } from "@/types/interface";
import SubMenuDirItem from "./SubMenuDirItem";
import SubMenuFileItem from "./SubMenuFileItem";

export default function SubMenu({
  prefix,
  menus,
  layer,
  pathname,
  locale,
  collapse,
}: SubMenuProps) {
  return (
    <ul>
      {menus.map((menu, index) =>
        menu.isdir ? (
          <li key={index}>
            <SubMenuDirItem
              layer={layer}
              menu={menu}
              pathname={pathname}
              locale={locale}
              prefix={prefix}
              collapse={collapse}
            />
          </li>
        ) : (
          <li key={index}>
            <SubMenuFileItem
              layer={layer}
              menu={menu}
              pathname={pathname}
              locale={locale}
              prefix={prefix}
              collapse={collapse}
            />
          </li>
        )
      )}
    </ul>
  );
}
