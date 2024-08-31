"use client";
import { MdComputer } from "react-icons/md";
import { MdDarkMode, MdLightMode } from "react-icons/md";
import { Fragment, useEffect, useState } from "react";
import { Listbox, Transition } from "@headlessui/react";
import { CheckIcon } from "@heroicons/react/20/solid";
import { useTheme } from "next-themes";
import NavBarIcon from "./NavBarIcon";
const themeClass = ["dark", "light", "system"];
export default function NavLang() {
  const [mounted, setMounted] = useState(false);
  const { theme, setTheme } = useTheme();

  // useEffect only runs on the client, so now we can safely show the UI
  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return (
      <div className="block w-6 h-6 bg-slate-200 dark:bg-slate-500  animate-pulse rounded-full" />
    );
  }

  return (
    <Listbox value={theme} onChange={setTheme}>
      <div className="relative">
        <Listbox.Button className="relative flex w-full cursor-pointer rounded-lg text-left focus:outline-none focus-visible:border-indigo-500 focus-visible:ring-2 focus-visible:ring-white/75 focus-visible:ring-offset-2 focus-visible:ring-offset-orange-300 sm:text-sm">
          <NavBarIcon
            Icon={
              theme == "dark"
                ? MdDarkMode
                : theme == "light"
                  ? MdLightMode
                  : MdComputer
            }
            onClick={() => {}}
            mounted={mounted}
          />
        </Listbox.Button>
        <Transition
          as={Fragment}
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <Listbox.Options className="absolute mt-2 max-h-60 w-36 overflow-auto rounded-md bg-white dark:bg-slate-800 py-1 text-base border-[0.01rem] dark:border-slate-600 ring-1 ring-black/5 focus:outline-none sm:text-sm right-0">
            {themeClass.map((lang, langIdx) => (
              <div key={langIdx}>
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
                        {lang}
                      </span>
                      {selected && (
                        <span className="absolute inset-y-0 left-0 flex items-center pl-3 text-sky-600 dark:text-white">
                          <CheckIcon className="h-5 w-5" aria-hidden="true" />
                        </span>
                      )}
                    </>
                  )}
                </Listbox.Option>
              </div>
            ))}
          </Listbox.Options>
        </Transition>
      </div>
    </Listbox>
  );
}
