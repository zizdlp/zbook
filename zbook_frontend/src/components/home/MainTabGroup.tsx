"use client";
import { Tab } from "@headlessui/react";
import { AiOutlineBook, AiOutlineFileSearch } from "react-icons/ai";
import { MdDashboard } from "react-icons/md";
import { IoIosNotifications } from "react-icons/io";
import Image from "next/image";
import { useTranslations } from "next-intl";

export default function MainTabGroup() {
  const t = useTranslations("HomePage");

  const categories = [
    {
      label: t("RepoHome"),
      icon: AiOutlineBook,
      lightImage: "/feature_light.png",
      darkImage: "/feature_dark.png",
    },
    {
      label: t("DashBoard"),
      icon: MdDashboard,
      lightImage: "/dashboard_light.png",
      darkImage: "/dashboard_dark.png",
    },
    {
      label: t("SearchDoc"),
      icon: AiOutlineFileSearch,
      lightImage: "/search_light.png",
      darkImage: "/search_dark.png",
    },
    {
      label: t("Notification"),
      icon: IoIosNotifications,
      lightImage: "/notification_light.png",
      darkImage: "/notification_dark.png",
    },
  ];

  const classNames = (...classes: string[]) =>
    classes.filter(Boolean).join(" ");

  return (
    <div className="max-w-5xl mx-auto py-24 px-2 md:px-4 pb-2 md:pb-4">
      <Tab.Group>
        <Tab.List className="flex space-x-1 p-1 max-w-xl mx-auto border border-[#65b1e8]/50 rounded-full">
          {categories.map(({ label, icon: Icon }, index) => (
            <Tab
              key={index}
              className={({ selected }) =>
                classNames(
                  "w-full text-xs md:text-sm font-medium leading-3 md:leading-5 focus:outline-none rounded-full",
                  selected
                    ? "bg-transparent shadow text-slate-700 dark:text-slate-200 bg-white dark:bg-slate-700/50"
                    : "text-slate-500 hover:bg-gray-100/25 dark:hover:bg-gray-900/25 hover:text-slate-700 dark:hover:text-slate-200"
                )
              }
            >
              <div className="flex items-center justify-center py-1.5 px-2 md:px-4 md:space-x-2.5">
                <Icon className="w-6 h-6 hidden md:block" />
                <span>{label}</span>
              </div>
            </Tab>
          ))}
        </Tab.List>

        <Tab.Panels className="my-6 p-2 rounded-lg bg-[#65b1e8] bg-opacity-20 border border-[#65b1e8]/20">
          {categories.map(({ lightImage, darkImage }, index) => (
            <Tab.Panel key={index}>
              <div className="block dark:hidden">
                <Image
                  src={lightImage}
                  className="rounded-md"
                  alt={`Light mode ${categories[index].label}`}
                  width={1728}
                  height={1080}
                />
              </div>
              <div className="hidden dark:block">
                <Image
                  src={darkImage}
                  className="rounded-md"
                  alt={`Dark mode ${categories[index].label}`}
                  width={1728}
                  height={1080}
                />
              </div>
            </Tab.Panel>
          ))}
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
}
