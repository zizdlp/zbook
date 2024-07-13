"use client";
import { Link } from "@/navigation";
import { BiSolidBookContent } from "react-icons/bi";
import { FaBook } from "react-icons/fa6";
import { TbMessageReport } from "react-icons/tb";
import { FaComment } from "react-icons/fa";
import { useTranslations } from "next-intl";
import React, { useContext } from "react";
import { SideBarContext } from "@/providers/SideBarProvider";
import SideBarLiContent from "./SideBarLiContent";
import { MdAdminPanelSettings } from "react-icons/md";
import { AiFillMessage } from "react-icons/ai";
import ShowComponent from "../ShowComponent";
import { usePathname } from "@/navigation";
export default function UserSideBarAdmin({
  username,
  authname,
  authrole,
}: {
  username: string;
  authname: string;
  authrole: string;
}) {
  const t = useTranslations("SideBar");
  const { sideBarOpen, setSideBarOpen } = useContext(SideBarContext);
  let isCurrentUser = username === authname;
  let isAdmin = authrole === "admin";

  // Function to handle the click event
  const handleOnClick = () => {
    setSideBarOpen(false);
    localStorage.setItem("sidebarValue", JSON.stringify(false));
  };

  const pathname = usePathname();
  let page_type = "";
  if (pathname != undefined) {
    const regex = new RegExp(
      `^\/([^\/]+)?\/?workspace\/${username}\/?([^\/]*)?\/?$`
    );
    const matches = pathname.match(regex);
    if (matches) {
      page_type = matches[2] || "";
    }
  }
  return (
    <ul className="my-4 space-y-3">
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li>
          <Link onClick={handleOnClick} href={`/workspace/admin/dashboard`}>
            <SideBarLiContent isSelected={page_type == "dashboard"}>
              <BiSolidBookContent className="text-rose-600 dark:text-rose-400 w-6 h-6" />
              <span className={`flex-1 ms-3 whitespace-nowrap`}>
                {t("OverView")}
              </span>
            </SideBarLiContent>
          </Link>
        </li>
      </ShowComponent>
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li>
          <Link onClick={handleOnClick} href={`/workspace/admin/users`}>
            <SideBarLiContent isSelected={page_type == "users"}>
              <MdAdminPanelSettings className="text-emerald-600 dark:text-emerald-400 w-6 h-6" />
              <span className={`flex-1 ms-3 whitespace-nowrap`}>
                {t("Users")}
              </span>
            </SideBarLiContent>
          </Link>
        </li>
      </ShowComponent>
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li>
          <Link onClick={handleOnClick} href={`/workspace/admin/sessions`}>
            <SideBarLiContent isSelected={page_type == "sessions"}>
              <AiFillMessage className="text-blue-600 dark:text-blue-400 w-6 h-6" />
              <span className={`flex-1 ms-3 whitespace-nowrap`}>
                {t("Sessions")}
              </span>
            </SideBarLiContent>
          </Link>
        </li>
      </ShowComponent>
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li>
          <Link onClick={handleOnClick} href={`/workspace/admin/repos`}>
            <SideBarLiContent isSelected={page_type == "repos"}>
              <FaBook className="text-yellow-600 dark:text-yellow-400 w-6 h-6" />
              <span className={`flex-1 ms-3 whitespace-nowrap`}>
                {t("Repositories")}
              </span>
            </SideBarLiContent>
          </Link>
        </li>
      </ShowComponent>
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li>
          <Link onClick={handleOnClick} href={`/workspace/admin/comments`}>
            <SideBarLiContent isSelected={page_type == "comments"}>
              <FaComment className="text-fuchsia-600 dark:text-fuchsia-400 w-6 h-6" />
              <span className={`flex-1 ms-3 whitespace-nowrap`}>
                {t("Comments")}
              </span>
            </SideBarLiContent>
          </Link>
        </li>
      </ShowComponent>
      <ShowComponent show={isAdmin && isCurrentUser}>
        <li>
          <Link onClick={handleOnClick} href={`/workspace/admin/reports`}>
            <SideBarLiContent isSelected={page_type == "reports"}>
              <TbMessageReport className="text-sky-600 dark:text-sky-400 w-6 h-6" />
              <span className={`flex-1 ms-3 whitespace-nowrap`}>
                {t("CommentReports")}
              </span>
            </SideBarLiContent>
          </Link>
        </li>
      </ShowComponent>
    </ul>
  );
}
