"use client";
import { Anchor, MenuStruct } from "@/types/interface";
import LeftSideBarWrapper from "@/components/sidebars/LeftSideBarWrapper";
import RepoSideBarSetting from "./RepoSideBarSettting";
import { usePathname } from "@/navigation";
import { useLocale } from "next-intl";
import SideBarSearchButton from "./SideBarSearchButton";
import { SearchType } from "@/utils/const_value";
import UnfoldSubMenu from "./UnfoldSubMenu";
import FoldSubMenu from "./FoldSubMenu";
import { ThemeColor } from "../TableOfContent";
export default function RepoSideBar({
  sublayouts,
  anchors,
  username,
  reponame,
  authname,
  theme_sidebar,
  theme_color,
  visibility_level,
  first_path,
}: {
  sublayouts: MenuStruct[];
  anchors: Anchor[];
  username: string;
  reponame: string;
  authname: string;
  theme_color: ThemeColor;
  theme_sidebar: string;
  visibility_level: string;
  first_path: string;
}) {
  const pathname = usePathname();
  const locale = useLocale();
  return (
    <LeftSideBarWrapper small={true}>
      <SideBarSearchButton
        username={username}
        repo_name={reponame}
        searchType={SearchType.REPO_DOCUMENT}
      />

      <div className="absolute inset-0 z-10 overflow-auto pb-10 pt-32 lg:pt-24 px-4">
        <div className="pt-4"></div>
        <ul id="navigation-items">
          <RepoSideBarSetting
            username={username}
            reponame={decodeURIComponent(reponame)}
            authname={authname}
            anchors={anchors ?? []}
            visibility_level={visibility_level}
            theme_color={theme_color}
            first_path={first_path}
          />
          {theme_sidebar == "theme_sidebar_fold" ? (
            <FoldSubMenu
              prefix={`/workspace/${username}/o/${decodeURIComponent(
                reponame
              )}/`}
              menus={sublayouts}
              layer={1}
              pathname={pathname}
              locale={locale}
              collapse={false}
              theme_color={theme_color}
            />
          ) : (
            <UnfoldSubMenu
              prefix={`/workspace/${username}/o/${decodeURIComponent(
                reponame
              )}/`}
              menus={sublayouts}
              layer={1}
              pathname={pathname}
              locale={locale}
              theme_color={theme_color}
            />
          )}
        </ul>
      </div>
    </LeftSideBarWrapper>
  );
}
