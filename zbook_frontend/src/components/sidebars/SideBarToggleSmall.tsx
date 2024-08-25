"use client";
import React, { useContext } from "react";
import { SideBarContext } from "@/providers/SideBarProvider";
import { RiMenu4Line, RiMenuLine } from "react-icons/ri";
export default function SideBarToggleSmall() {
  const { sideBarOpen, setSideBarOpen } = useContext(SideBarContext);
  const IconText = ({ Icon, onClick }: { Icon: any; onClick: () => void }) => (
    <Icon
      onClick={onClick}
      className="mr-2 md:mr-3 block w-11 h-11  dark:text-slate-200  hover:text-sky-500  cursor-pointer"
    />
  );
  return (
    <div
      className={`lg:hidden fixed bottom-5 z-50 right-5 h-11 w-11 rounded-md  backdrop-blur-sm print:hidden`}
    >
      <IconText
        Icon={sideBarOpen === false ? RiMenuLine : RiMenu4Line}
        onClick={() => {
          setSideBarOpen(!sideBarOpen);
          localStorage.setItem("sidebarValue", JSON.stringify(!sideBarOpen));
        }}
      />
    </div>
  );
}
