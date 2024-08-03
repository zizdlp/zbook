"use client";
import SubMenu from "./SubMenu";
import { Anchor, MenuStruct } from "@/types/interface";
import SideBarWrapper from "@/components/sidebars/SideBarWrapper";
import RepoSideBarSetting from "./RepoSideBarSettting";
import { usePathname } from "@/navigation";
import { useLocale } from "next-intl";
import SideBarSearchButton from "./SideBarSearchButton";
import { SearchType } from "@/utils/const_value";
import FrpcSubMenu from "./FrpcSubMenu";
export default function RepoSideBar({
  sublayouts,
  anchors,
  username,
  reponame,
  authname,
  visibility_level,
}: {
  sublayouts: MenuStruct[];
  anchors: Anchor[];
  username: string;
  reponame: string;
  authname: string;
  visibility_level: string;
}) {
  const pathname = usePathname();
  const locale = useLocale();
  return (
    <div
      className="z-20 hidden lg:block fixed bottom-0 right-auto w-[18rem] top-[4rem]"
      id="sidebar"
    >
      <div
        className="absolute inset-0 z-10 overflow-auto pr-8 pb-10"
        id="sidebar-content"
      >
        <div className="relative lg:text-sm lg:leading-6">
          <div className="sticky top-0 pointer-events-none">
            <div className="h-10 bg-white dark:bg-gray-900"></div>
            <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
              <SideBarSearchButton
                username={username}
                repo_name={decodeURIComponent(reponame)}
                searchType={SearchType.REPO_DOCUMENT}
              />
            </div>
          </div>
          <div className="sticky top-0 h-8"></div>
          <ul id="navigation-items">
            <RepoSideBarSetting
              username={username}
              reponame={decodeURIComponent(reponame)}
              authname={authname}
              anchors={anchors}
              visibility_level={visibility_level}
            />
            <FrpcSubMenu
              prefix={`/workspace/${username}/o/${decodeURIComponent(
                reponame
              )}/`}
              menus={sublayouts}
              layer={1}
              pathname={pathname}
              locale={locale}
            />
          </ul>
        </div>
      </div>
    </div>
  );
  // return <FrpcSideBar />;
  // return (
  //   <SideBarWrapper>
  //     <div className="sticky top-0 -ml-0.5 pointer-events-none">
  //       <div className="h-10 bg-white dark:bg-gray-900"></div>
  //       <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
  //         <SideBarSearchButton
  //           username={username}
  //           repo_name={decodeURIComponent(reponame)}
  //           searchType={SearchType.REPO_DOCUMENT}
  //         />
  //       </div>
  //       <div className="h-6 bg-gradient-to-b from-white dark:from-slate-900"></div>
  //     </div>

  //     <RepoSideBarSetting
  //       username={username}
  //       reponame={decodeURIComponent(reponame)}
  //       authname={authname}
  //       visibility_level={visibility_level}
  //     />
  //     <SubMenu
  //       prefix={`/workspace/${username}/o/${decodeURIComponent(reponame)}/`}
  //       menus={sublayouts}
  //       layer={1}
  //       pathname={pathname}
  //       locale={locale}
  //     />
  //   </SideBarWrapper>
  // );
}
