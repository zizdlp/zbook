"use client";
import { SideBarContext } from "@/providers/SideBarProvider";
import { Link } from "@/navigation";
import { useContext } from "react";
import { BsFillBookmarkCheckFill } from "react-icons/bs";
import { FaUserFriends } from "react-icons/fa";
import { FaUsersLine } from "react-icons/fa6";
import { MdBook } from "react-icons/md";
import { usePathname } from "@/navigation";
import SideBarLiContent from "./SideBarLiContent";
import { useTranslations } from "next-intl";
import { Lexend_Tera } from "next/font/google";
export default function UserSideBarSocial({
  username,
  count_following,
  count_follower,
  count_repos,
  count_likes,
}: {
  username: string;
  count_following: number;
  count_follower: number;
  count_repos: number;
  count_likes: number;
}) {
  const t = useTranslations("SideBar");
  const { sideBarOpen, setSideBarOpen } = useContext(SideBarContext);
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
  const handleOnClick = () => {
    setSideBarOpen(false);
    localStorage.setItem("sidebarValue", JSON.stringify(false));
  };

  return (
    <ul className="my-4 space-y-3">
      <li>
        <Link onClick={handleOnClick} href={`/workspace/${username}`}>
          <SideBarLiContent isSelected={page_type == ""}>
            <MdBook className="text-indigo-700 dark:text-indigo-400 w-6 h-6" />
            <span className={`flex-1 ms-3 whitespace-nowrap`}>
              {t("Repositories")}
            </span>
            <span className="inline-flex items-center justify-center px-2 py-0.5 ms-3 text-xs font-medium text-gray-500 bg-gray-200 rounded dark:bg-gray-700 dark:text-gray-400">
              {count_repos ?? 0}
            </span>
          </SideBarLiContent>
        </Link>
      </li>
      <li>
        <Link onClick={handleOnClick} href={`/workspace/${username}/favorite`}>
          <SideBarLiContent isSelected={page_type == "favorite"}>
            <BsFillBookmarkCheckFill className="text-cyan-700 dark:text-cyan-400 w-6 h-6" />
            <span className={`flex-1 ms-3 whitespace-nowrap`}>
              {t("Favorite")}
            </span>
            <span className="inline-flex items-center justify-center px-2 py-0.5 ms-3 text-xs font-medium text-gray-500 bg-gray-200 rounded dark:bg-gray-700 dark:text-gray-400">
              {count_likes ?? 0}
            </span>
          </SideBarLiContent>
        </Link>
      </li>
      <li>
        <Link onClick={handleOnClick} href={`/workspace/${username}/following`}>
          <SideBarLiContent isSelected={page_type == "following"}>
            <FaUserFriends className="text-sky-700 dark:text-sky-400 w-6 h-6" />
            <span className={`flex-1 ms-3 whitespace-nowrap`}>
              {t("Following")}
            </span>
            <span className="inline-flex items-center justify-center px-2 py-0.5 ms-3 text-xs font-medium text-gray-500 bg-gray-200 rounded dark:bg-gray-700 dark:text-gray-400">
              {count_following ?? 0}
            </span>
          </SideBarLiContent>
        </Link>
      </li>
      <li>
        <Link onClick={handleOnClick} href={`/workspace/${username}/follower`}>
          <SideBarLiContent isSelected={page_type == "follower"}>
            <FaUsersLine className="text-green-700 dark:text-green-400 w-6 h-6" />
            <span className={`flex-1 ms-3 whitespace-nowrap`}>
              {t("Follower")}
            </span>
            <span className="inline-flex items-center justify-center px-2 py-0.5 ms-3 text-xs font-medium text-gray-500 bg-gray-200 rounded dark:bg-gray-700 dark:text-gray-400">
              {count_follower ?? 0}
            </span>
          </SideBarLiContent>
        </Link>
      </li>
    </ul>
  );
}
