"use client";
import { Anchor, MenuStruct } from "@/types/interface";
import SideBarWrapper from "@/components/sidebars/SideBarWrapper";
import RepoSideBarSetting from "./RepoSideBarSettting";
import { usePathname } from "@/navigation";
import { useLocale } from "next-intl";
import SideBarSearchButton from "./SideBarSearchButton";
import { SearchType } from "@/utils/const_value";
import FrpcSubMenu from "./FrpcSubMenu";
import SubMenu from "./SubMenu";
export default function RepoSideBar({
  sublayouts,
  anchors,
  username,
  reponame,
  authname,
  sidebar_theme,
  visibility_level,
}: {
  sublayouts: MenuStruct[];
  anchors: Anchor[];
  username: string;
  reponame: string;
  authname: string;
  sidebar_theme: string;
  visibility_level: string;
}) {
  const pathname = usePathname();
  const locale = useLocale();
  return (
    <SideBarWrapper>
      <div className="sticky top-0 pointer-events-none z-50 px-4 lg:px-0">
        <div className="h-10 bg-white dark:bg-gray-900"></div>
        <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
          <SideBarSearchButton
            username={username}
            repo_name={reponame}
            searchType={SearchType.REPO_DOCUMENT}
          />
        </div>
        <div className="h-4 bg-gradient-to-b from-white dark:from-slate-900"></div>
      </div>

      <div className="absolute inset-0 z-10 overflow-auto pb-10 pt-32 lg:pt-24 px-4 lg:px-0">
        <div className="pt-4"></div>
        <ul id="navigation-items">
          <RepoSideBarSetting
            username={username}
            reponame={decodeURIComponent(reponame)}
            authname={authname}
            anchors={anchors ?? []}
            visibility_level={visibility_level}
          />
          {sidebar_theme == "theme_sidebar_fold" ? (
            <SubMenu
              prefix={`/workspace/${username}/o/${decodeURIComponent(
                reponame
              )}/`}
              menus={sublayouts}
              layer={1}
              pathname={pathname}
              locale={locale}
              collapse={false}
            />
          ) : (
            <FrpcSubMenu
              prefix={`/workspace/${username}/o/${decodeURIComponent(
                reponame
              )}/`}
              menus={sublayouts}
              layer={1}
              pathname={pathname}
              locale={locale}
            />
          )}
        </ul>
      </div>
    </SideBarWrapper>
  );
}
