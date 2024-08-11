"use client";
import { Tab } from "@headlessui/react";
import { MdImage } from "react-icons/md";
import { MdOutlineFeaturedVideo } from "react-icons/md";

import Image from "next/image";
export default function FeatureTabGroup({
  categories,
  image_urls,
  video_urls,
}: {
  categories: string[];
  image_urls: string[];
  video_urls: string[];
}) {
  function classNames(...classes: string[]): string {
    return classes.filter(Boolean).join(" ");
  }
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
                    ? "bg-transparent shadow bg-white dark:bg-slate-700/50"
                    : "text-slate-500 hover:bg-gray-100/25 dark:hover:bg-gray-900/25 hover:text-slate-700 dark:hover:text-slate-200"
                )
              }
            >
              <div className="z-10 flex items-center justify-center py-1.5 px-2 md:px-4 md:space-x-2.5">
                {index == 0 ? (
                  <MdImage className="w-6 h-6 hidden md:block" />
                ) : (
                  <MdOutlineFeaturedVideo className="w-6 h-6 hidden md:block" />
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
                src={image_urls[0]}
                className="rounded-md dark:hidden"
                alt="Picturel of light image"
                width={1728}
                height={1080}
              />
            </div>
            <div className="hidden dark:block">
              <Image
                src={image_urls[1]}
                className="rounded-md hidden dark:block"
                alt="Picture of dark image"
                width={1728}
                height={1080}
              />
            </div>
          </Tab.Panel>
          <Tab.Panel key={1}>
            <div className="block dark:hidden">
              <iframe
                className="overflow-hidden rounded-lg my-auto mx-auto shadow-lg block dark:hidden w-full aspect-mac"
                loading="lazy"
                src={video_urls[0]}
                title="This is light video"
                allowFullScreen
              ></iframe>
            </div>
            <div className="hidden dark:block">
              <iframe
                className="overflow-hidden rounded-lg my-auto mx-auto shadow-lg dark:block hidden w-full aspect-mac"
                loading="lazy"
                src={video_urls[1]}
                title="This is dark video"
                allowFullScreen
              ></iframe>
            </div>
          </Tab.Panel>
        </Tab.Panels>
      </Tab.Group>
    </div>
  );
}
