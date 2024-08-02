/* eslint-disable react/jsx-no-literals */
import React from "react";
import { FaBookOpen, FaGithub } from "react-icons/fa";
import { IconType } from "react-icons/lib";

interface MenuItemProps {
  href: string;
  selected: boolean;
  text: string;
  icon: IconType;
}
const IconText = ({
  Icon,
  selected,
}: {
  Icon: IconType;
  selected: boolean;
}) => (
  <Icon
    className={`h-4 w-4  ${selected ? "fill-white" : "group-hover:fill-white"}`}
  />
);
export default function MenuItem({
  href,
  selected,
  text,
  icon,
}: MenuItemProps) {
  return (
    <li>
      <a
        href={href}
        className={`pl-4 group flex items-center lg:text-sm lg:leading-6 mb-5 sm:mb-4 font-medium 
            ${
              selected
                ? "text-purple-600"
                : "text-gray-600 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300"
            }
            `}
      >
        <div
          className={`mr-4 rounded-md p-1 ring-1 ring-slate-300 dark:ring-slate-700

            ${
              selected
                ? "bg-gradient-to-r from-purple-600 to-pink-400 ring-0"
                : "group-hover:bg-gradient-to-r group-hover:from-purple-600 group-hover:to-pink-400 group-hover:ring-0"
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
