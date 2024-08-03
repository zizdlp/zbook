"use client";

import { useTranslations } from "next-intl";
import { IconType } from "react-icons/lib";
export default function RepoSideBarSettingItem({
  href,
  selected,
  text,
  icon,
  onClick,
}: {
  href: string;
  selected: boolean;
  text: string;
  icon: IconType;
  onClick?: any;
}) {
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
              ? "text-purple-600 dark:text-purple-400"
              : "text-gray-600  dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-300"
          }
          `}
      >
        <div
          className={`mr-4 rounded-md p-1 group-hover:bg-gradient-to-r group-hover:from-purple-500 group-hover:to-pink-500 group-hover:ring-0
            ${
              selected
                ? "bg-gradient-to-r from-purple-500 to-pink-500"
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
