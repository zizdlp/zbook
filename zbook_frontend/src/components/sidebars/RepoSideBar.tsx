"use client";
import SubMenu from "./SubMenu";
import { MenuStruct } from "@/types/interface";
import SideBarWrapper from "@/components/sidebars/SideBarWrapper";
import RepoSideBarSetting from "./RepoSideBarSettting";
import { usePathname } from "@/navigation";
import { useLocale } from "next-intl";
import SideBarSearchButton from "./SideBarSearchButton";
import { SearchType } from "@/utils/const_value";
export default function RepoSideBar({
  repo_id,
  sublayouts,
  username,
  reponame,
  authname,
  visibility_level,
}: {
  repo_id: number;
  sublayouts: MenuStruct[];
  username: string;
  reponame: string;
  authname: string;
  visibility_level: string;
}) {
  const pathname = usePathname();
  const locale = useLocale();
  return (
    <SideBarWrapper>
      <div className="sticky top-0 -ml-0.5 pointer-events-none">
        <div className="h-10 bg-white dark:bg-gray-900"></div>
        <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
          <SideBarSearchButton
            username={username}
            repo_name={reponame}
            searchType={SearchType.REPO_DOCUMENT}
          />
        </div>
        <div className="h-6 bg-gradient-to-b from-white dark:from-slate-900"></div>
      </div>

      <RepoSideBarSetting
        username={username}
        repo_id={repo_id}
        reponame={reponame}
        authname={authname}
        visibility_level={visibility_level}
      />
      <SubMenu
        prefix={`/workspace/${username}/o/${reponame}/`}
        menus={sublayouts}
        layer={1}
        pathname={pathname}
        locale={locale}
      />
    </SideBarWrapper>
  );
}
