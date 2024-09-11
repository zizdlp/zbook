"use client";

import { ThemeColor } from "../TableOfContent";
import IconItem from "../IconComponent";
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
  icon_name,
  theme_color,
  onClick,
}: {
  href: string;
  selected: boolean;
  text: string;
  icon_name: string;
  theme_color: ThemeColor;
  onClick?: any;
}) {
  let { hoverClass, selectedClass, selectBgClass } =
    getSettingColorClasses(theme_color);

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
          <IconItem
            iconName={icon_name}
            className={`h-4 w-4  ${
              selected ? "fill-white" : "group-hover:fill-white"
            }`}
          />
        </div>
        {text}
      </a>
    </li>
  );
}
