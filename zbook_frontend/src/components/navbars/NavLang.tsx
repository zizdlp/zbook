"use client";
import { Fragment, useState } from "react";
import { Listbox, Transition } from "@headlessui/react";
import { CheckIcon } from "@heroicons/react/20/solid";
import { useLocale } from "next-intl";

import { usePathname, Link, locales } from "../../navigation";
const localeMap = {
  en: "English",
  zh: "简体中文",
  de: "Deutsch",
  // Add other locales here
};
export default function NavLink() {
  const locale = useLocale();
  const pathname = usePathname();
  const [selected, setSelected] = useState(locale);
  return (
    <Listbox value={selected} onChange={setSelected}>
      <div className="relative">
        <Listbox.Button
          aria-label="Select language"
          className="relative flex w-full cursor-pointer rounded-lg text-left focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white/75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm"
        >
          <svg
            stroke="currentColor"
            fill="currentColor"
            strokeWidth="0"
            viewBox="0 0 1024 1024"
            className={`block w-6 h-6  hover:text-sky-600 dark:hover:text-sky-400 cursor-pointer`}
            height="1em"
            width="1em"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M580 477.7h-9.8l37.6 209.1c24.8-8.8 47.4-22.2 67.5-39.6-20.6-24.8-37.1-52.5-50-81.9l39.6-5.2c10.8 22.2 23.2 42.2 37.6 59.3 29.3-35.5 51.5-82.9 67.5-142.2l-190 0.5z m149.9 169.5c22.7 19.6 48.4 34 77.2 42.7l18.1 5.7-10.8 38.6-18.1-5.7c-34.5-10.8-66.5-28.8-93.7-53.1-25.3 22.7-55.1 40.2-87.5 51l25.3 141.1H489.8l-20.1 92.2h472.4c22.2 0 40.2-18.1 40.2-40.2v-683c0-22.2-18.1-40.2-40.2-40.2H520.3l31.4 175.2-1-0.5 3.6 19.1 0.5-2.6 8.8 50H661v-40.2h75.3v40.2H862v40.2h-52.5c-17.7 70.6-44.6 127.3-79.6 169.5z m-281.3 220H82.3C38 867.2 2 831.1 2 786.8V104.2c0-44.8 36-80.3 80.3-80.3h401.8l24.8 132.4h432.8c44.3 0 80.3 36 80.3 80.3v683.1c0 44.3-36 80.3-80.3 80.3H419.8l28.8-132.8zM259.1 558.1v-42.2h-79.3v-62.4h73.7v-41.7h-73.7v-53.1h79.3V317H133.3v241.1h125.8z m193.6 0V437.5c0-21.7-5.2-38.6-15-50.5s-24.8-17.5-44.3-17.5c-11.4 0-21.7 2.1-30.4 6.7s-16 11.9-20.6 20.6h-2.6l-6.2-23.7h-35V558h45.3v-87c0-21.7 3.1-37.1 8.8-46.9 5.7-9.3 15-13.9 27.8-13.9 9.3 0 16 3.1 20.6 9.8 4.1 6.7 6.7 16.5 6.7 29.8V558l44.9 0.1z"
              p-id="4274"
            />
          </svg>
        </Listbox.Button>
        <Transition
          as={Fragment}
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <Listbox.Options
            className="absolute mt-2 max-h-60 w-36 overflow-y-auto scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16] rounded-md bg-white dark:bg-slate-800 py-1 text-base border-[0.01rem] dark:border-slate-600 ring-1 ring-black/5 focus:outline-none sm:text-sm right-0"
          >
            {locales.map((lang, langIdx) => (
              <Link
                href={pathname}
                key={langIdx}
                locale={lang == ("en" || "") ? "en" : lang}
              >
                <Listbox.Option
                  key={langIdx}
                  className={({ active }) =>
                    `relative cursor-default select-none py-2 pl-10 pr-4 ${
                      active
                        ? "bg-sky-500 dark:bg-slate-700 text-white dark:text-white"
                        : "text-gray-900 dark:text-white"
                    }`
                  }
                  value={lang}
                >
                  {({ selected }) => (
                    <>
                      <span
                        className={`block truncate ${
                          selected ? "font-medium" : "font-normal"
                        }`}
                      >
                        {localeMap[lang]}
                      </span>
                      {selected && (
                        <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-sky-600 dark:text-white">
                          <CheckIcon className="h-5 w-5" aria-hidden="true" />
                        </span>
                      )}
                    </>
                  )}
                </Listbox.Option>
              </Link>
            ))}
          </Listbox.Options>
        </Transition>
      </div>
    </Listbox>
  );
}
