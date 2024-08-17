import { BiError } from "react-icons/bi";
import LeftSideBarWrapper from "../sidebars/LeftSideBarWrapper";

export default function RepoSideBarLoading() {
  return (
    <LeftSideBarWrapper small={false}>
      <div className="sticky top-0 -ml-0.5 pointer-events-none px-4">
        <div className="h-10 bg-white dark:bg-gray-900"></div>
        <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
          <div
            className="flex w-full items-center text-sm leading-6 text-slate-400 rounded-md ring-none border-[0.1rem] border-slate-200 dark:border-0
                py-3 px-3  dark:highlight-white/5 h-12 bg-gray-200 dark:bg-gray-700/75 animate-pulse"
          ></div>
        </div>
        <div className="h-4 bg-gradient-to-b from-white dark:from-slate-900"></div>
      </div>
      <div className="hidden xl:block flex-shrink-0 mx-auto rounded-lg animate-pulse">
        <div className="px-4">
          <div className="pb-4 flex">
            <div className="w-80 h-6 bg-gray-200 rounded-md dark:bg-gray-700/75" />
          </div>

          <div className="flex items-center justify-between mb-2">
            <div>
              <div className="h-2.5 bg-gray-200 rounded-md dark:bg-gray-700/75 w-32 mb-2.5"></div>
              <div className="w-64 h-2 bg-gray-200 rounded-md dark:bg-gray-700/75"></div>
            </div>
          </div>

          <div className="pb-4 flex">
            <div className="w-80 h-40 bg-gray-200 rounded-md dark:bg-gray-700/75" />
          </div>
          <div className="flex items-center justify-between mb-2">
            <div>
              <div className="h-2.5 bg-gray-200 rounded-md dark:bg-gray-700/75 w-32 mb-2.5"></div>
              <div className="w-64 h-2 bg-gray-200 rounded-md dark:bg-gray-700/75"></div>
            </div>
          </div>
          <div className="pb-4 flex">
            <div className="w-80 h-40 bg-gray-200 rounded-md dark:bg-gray-700/75" />
          </div>
          <div className="flex items-center justify-between mb-2">
            <div>
              <div className="h-2.5 bg-gray-200 rounded-md dark:bg-gray-700/75 w-32 mb-2.5"></div>
              <div className="w-64 h-2 bg-gray-200 rounded-md dark:bg-gray-700/75"></div>
            </div>
          </div>
          <div className="pb-4 flex">
            <div className="w-80 h-40 bg-gray-200 rounded-md dark:bg-gray-700/75" />
          </div>
          <div className="flex items-center justify-between mb-2">
            <div>
              <div className="h-2.5 bg-gray-200 rounded-md dark:bg-gray-700/75 w-32 mb-2.5"></div>
              <div className="w-64 h-2 bg-gray-200 rounded-md dark:bg-gray-700/75"></div>
            </div>
          </div>
          <div className="pb-4 flex">
            <div className="w-80 h-40 bg-gray-200 rounded-md dark:bg-gray-700/75" />
          </div>
          <div className="flex items-center justify-between mb-2">
            <div>
              <div className="h-2.5 bg-gray-200 rounded-md dark:bg-gray-700/75 w-32 mb-2.5"></div>
              <div className="w-64 h-2 bg-gray-200 rounded-md dark:bg-gray-700/75"></div>
            </div>
          </div>
          <div className="pb-4 flex">
            <div className="w-80 h-40 bg-gray-200 rounded-md dark:bg-gray-700/75" />
          </div>
        </div>
      </div>
    </LeftSideBarWrapper>
  );
}
