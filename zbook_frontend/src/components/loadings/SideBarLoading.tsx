import SideBarWrapper from "../sidebars/SideBarWrapper";

export default function SideBarLoading() {
  return (
    <SideBarWrapper>
      <div className="h-10"></div>
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
    </SideBarWrapper>
  );
}
