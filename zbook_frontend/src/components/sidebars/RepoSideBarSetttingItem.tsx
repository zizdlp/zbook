"use client";

import { useTranslations } from "next-intl";
import { IconType } from "react-icons/lib";
import { ThemeColor } from "../TableOfContent";
function getSettingColorClasses(color: ThemeColor) {
  return {
    hoverClass: `group-hover:bg-${color}-500`,
    selectedClass: `text-${color}-600 dark:text-${color}-400`,
    selectBgClass: `bg-${color}-500`,
  };
}
export default function RepoSideBarSettingItem({
  href,
  selected,
  text,
  icon,
  theme_color,
  onClick,
}: {
  href: string;
  selected: boolean;
  text: string;
  icon: IconType;
  theme_color: ThemeColor;
  onClick?: any;
}) {
  let { hoverClass, selectedClass, selectBgClass } =
    getSettingColorClasses(theme_color);
  const t = useTranslations("SideBar");
  const IconText = ({
    Icon,
    selected,
  }: {
    Icon: IconType;
    selected: boolean;
  }) => (
    <Icon
      className={`h-4 w-4  ${
        selected ? "fill-white" : "group-hover:fill-white"
      }`}
    />
  );
  return (
    <li onClick={onClick}>
      <a
        href={href}
        className={`pl-4 group flex items-center lg:text-sm lg:leading-6 mb-5 sm:mb-4 font-medium 
          
          ${
            selected
              ? selectedClass
              : "text-gray-600  dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-300"
          }
          `}
      >
        <div
          className={`mr-4 rounded-md p-1 ${hoverClass} group-hover:ring-0
            ${
              selected
                ? selectBgClass
                : "ring-1 ring-slate-300 dark:ring-slate-700"
            }
          `}
        >
          <IconText Icon={icon} selected={selected} />
        </div>
        {text}
      </a>
    </li>
  );
}
