import LeftSideBarWrapper from "../sidebars/LeftSideBarWrapper";
import GroupLoading from "./GroupLoading";
import LoadingElement from "@/components/loadings/LoadingElement";
export default function UserSideBarLoading({
  username,
  authname,
  authrole,
}: {
  username: string;
  authname: string;
  authrole: string;
}) {
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
      <div
        className="absolute inset-0 z-10 overflow-y-auto scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
        scrollbar-thumb-slate-200 scrollbar-track-slate-100
        dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16] pb-10 pt-32 lg:pt-24 px-4"
      >
        <div className="animate-pulse">
          <div className="flex flex-col items-center justify-center">
            <LoadingElement className="h-24 rounded-full w-24 my-2" />
            <LoadingElement className="h-4 rounded-md w-16 my-2" />
            <LoadingElement className="h-4 rounded-md w-32 my-1" />
          </div>

          <div className="flex items-center w-full">
            <LoadingElement className="h-4 rounded-md w-32 my-1" />
            <LoadingElement className="h-4 ms-2 rounded-md w-full" />
          </div>

          <GroupLoading itemCount={4} showRight={true} />
          <GroupLoading itemCount={3} showRight={false} />
          <GroupLoading itemCount={6} showRight={true} />
          <GroupLoading itemCount={2} showRight={true} />
        </div>
      </div>
    </LeftSideBarWrapper>
  );
}
