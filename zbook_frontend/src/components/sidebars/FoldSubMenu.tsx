import type { SubMenuProps } from "@/types/interface";
import SubMenuDirItem from "./SubMenuDirItem";
import SubMenuFileItem from "./SubMenuFileItem";

export default function FoldSubMenu({
  prefix,
  menus,
  layer,
  pathname,
  locale,
  collapse,
  theme_color,
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
              theme_color={theme_color}
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
              theme_color={theme_color}
            />
          </li>
        )
      )}
    </ul>
  );
}
