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
    <div className="hidden lg:block">
      <div
        className={`hidden lg:block fixed left-[max(0px,calc(50%-45rem))] 2xl:left-[max(0px,calc(50%-50rem))] z-40 inset-0 w-[19rem] 2xl:w-[22rem] top-0 pt-8 lg:pt-12 px-4 border-r-[0.01rem] border-slate-300 dark:border-slate-700
      overflow-y-scroll
      scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md bg-white dark:bg-gray-900`}
      >
        <div className="sticky top-0 -ml-0.5 pointer-events-none">
          <div className="h-10 bg-white dark:bg-gray-900"></div>
          <div className="bg-white dark:bg-gray-900 relative pointer-events-auto">
            <div
              className="flex w-full items-center text-sm leading-6 text-slate-400 rounded-md ring-none border-[0.1rem] border-slate-200 dark:border-0
                py-1.5 pl-2 pr-3  dark:highlight-white/5 h-12 bg-gray-200 dark:bg-gray-700/75 animate-pulse"
            ></div>
          </div>
          <div className="h-4 bg-gradient-to-b from-white dark:from-slate-900"></div>
        </div>
        <div className="animate-pulse">
          <div className="flex flex-col items-center justify-center p-4">
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
    </div>
  );
}
