import { MenuStruct } from "@/types/interface";
import { Link } from "@/navigation";
import { isSameUrl } from "@/utils/util";
import { ThemeColor } from "../TableOfContent";
function getSideBarColorClasses(color: ThemeColor) {
  return {
    activeClass: `bg-${color}-400/10 text-${color}-900 font-semibold dark:text-${color}-400 dark:bg-${color}-500/10`,
  };
}
export default function SubMenuFileItem({
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
  let { activeClass } = getSideBarColorClasses(theme_color);
  return (
    <Link
      href={prefix + menu.relative_path.toLocaleLowerCase()}
      className={`group mt-2 lg:mt-0 ${
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
          ${
            isSameUrl(
              pathname,
              locale,
              prefix,
              menu.relative_path.toLocaleLowerCase()
            )
              ? activeClass
              : "hover:bg-gray-600/5 dark:hover:bg-gray-200/5 text-gray-700 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
          }
          `}
    >
      <div className="flex-1 flex items-center space-x-2.5">
        <div> {menu.title}</div>
      </div>
    </Link>
  );
}
