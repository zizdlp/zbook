import { MenuStruct } from "@/types/interface";
import { Link } from "@/navigation";
import { isSameUrl } from "@/utils/util";
export default function SubMenuFileItem({
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
  return (
    <Link
      href={prefix + menu.relative_path}
      className={`group mt-2 lg:mt-0 ${
        layer == 1
          ? " pl-4 "
          : layer == 2
          ? "pl-7"
          : layer == 3
          ? "pl-10"
          : "pl-14"
      } flex items-center  pr-3 py-1.5 rounded-lg my-0.5
          ${
            isSameUrl(pathname, locale, prefix, menu.relative_path)
              ? "bg-sky-400/10 text-sky-900 font-semibold dark:text-sky-100 dark:bg-sky-500/10"
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
