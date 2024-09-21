"use client";

import React, { useContext } from "react";
import { SideBarContext } from "@/providers/SideBarProvider";
import { Transition } from "@headlessui/react";
interface SideBarTransitionProps {
  children: React.ReactNode;
}

export default function SideBarTransition(props: SideBarTransitionProps) {
  const { sideBarOpen } = useContext(SideBarContext);
  return (
    <Transition show={sideBarOpen}>
      {/* Background overlay */}
      <Transition.Child
        enter="transition-opacity ease-linear duration-300"
        enterFrom="opacity-100"
        enterTo="opacity-100"
        leave="transition-opacity ease-linear duration-300"
        leaveFrom="opacity-100"
        leaveTo="opacity-100"
      >
        <div className="fixed xl:hidden z-40 bg-white/50 inset-0 w-full  backdrop-blur-sm dark:bg-gray-900/25"></div>
      </Transition.Child>
      <Transition.Child
        enter="transition ease-in-out duration-300 transform"
        enterFrom="-translate-x-full opacity-0"
        enterTo="translate-x-0 opacity-100"
        leave="transition ease-in-out duration-300 transform"
        leaveFrom="translate-x-0 opacity-100"
        leaveTo="-translate-x-full opacity-0"
        className={`fixed z-40 inset-0 w-[20rem] top-0 pt-8 lg:pt-12
        overflow-y-auto  scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16] bg-white dark:bg-gray-900 border-r-[0.01rem] border-slate-300 dark:border-slate-700`}
      >
        {props.children}
      </Transition.Child>
    </Transition>
  );
}
