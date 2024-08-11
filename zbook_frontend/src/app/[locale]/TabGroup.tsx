"use client";
import { Tab } from "@headlessui/react";
import { useTranslations } from "next-intl";
import { AiOutlineBook } from "react-icons/ai";
import { MdDashboard } from "react-icons/md";
import { AiOutlineFileSearch } from "react-icons/ai";
import { IoIosNotifications } from "react-icons/io";

import Image from "next/image";
export default function TabGroup() {
  function classNames(...classes: string[]): string {
    return classes.filter(Boolean).join(" ");
  }
  const t = useTranslations("HomePage");
  let categories = [
    t("RepoHome"),
    t("DashBoard"),
    t("SearchDoc"),
    t("Notification"),
  ];
  return (
    <div className="px-2 md:px-4 pb-2 md:pb-4 max-w-5xl mx-auto py-24">
      <Tab.Group>
        <Tab.List className="flex space-x-1 p-1 max-w-xl mx-auto border border-[#65b1e8]/50 rounded-full">
          {categories.map((category, index) => (
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
              <div className="z-10 flex items-center justify-center py-1.5 px-2 md:px-4 md:space-x-2.5">
                {index == 0 ? (
                  <AiOutlineBook className="w-6 h-6 hidden md:block" />
                ) : index == 1 ? (
                  <MdDashboard className="w-6 h-6 hidden md:block" />
                ) : index == 2 ? (
                  <AiOutlineFileSearch className="w-6 h-6 hidden md:block" />
                ) : (
                  <IoIosNotifications className="w-6 h-6 hidden md:block" />
                )}
                <span>{category}</span>
              </div>
            </Tab>
          ))}
        </Tab.List>

        <Tab.Panels className="my-6 p-2 rounded-lg bg-[#65b1e8]  bg-opacity-20 border border-[#65b1e8]/20">
          <Tab.Panel key={0}>
            <div className="block dark:hidden">
              <Image
                src="/feature_light.png"
                className="rounded-md"
                alt="Picture of dark mac md head"
                width={1728}
                height={1080}
              />
            </div>
            <div className="hidden dark:block">
              <Image
                src="/feature_dark.png"
                className="rounded-md"
                alt="Picture of dark mac md head"
                width={1728}
                height={1080}
              />
            </div>
          </Tab.Panel>
          <Tab.Panel key={1}>
            <div className="block dark:hidden">
              <Image
                src="/dashboard_light.png"
                className="rounded-md"
                alt="Picture of dark mac md head"
                width={1728}
                height={1026}
              />
            </div>
            <div className="hidden dark:block">
              <Image
                src="/dashboard_dark.png"
                className="rounded-md"
                alt="Picture of dark mac md head"
                width={1728}
                height={1026}
              />
            </div>
          </Tab.Panel>
          <Tab.Panel key={2}>
            <div className="block dark:hidden">
              <Image
                src="/search_light.png"
                className="rounded-md"
                alt="Picture of light search"
                width={1728}
                height={1080}
              />
            </div>
            <div className="hidden dark:block">
              <Image
                src="/search_dark.png"
                className="rounded-md"
                alt="Picture of dark search"
                width={1728}
                height={1080}
              />
            </div>
          </Tab.Panel>
          <Tab.Panel key={3}>
            <div className="block dark:hidden">
              <Image
                src="/notification_light.png"
                className="rounded-md"
                alt="Picture of light search"
                width={1728}
                height={1080}
              />
            </div>
            <div className="hidden dark:block">
              <Image
                src="/notification_dark.png"
                className="rounded-md"
                alt="Picture of dark search"
                width={1728}
                height={1080}
              />
            </div>
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
}
