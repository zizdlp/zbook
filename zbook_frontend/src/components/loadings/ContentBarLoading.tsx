import RightSideBarWrapper from "../sidebars/RightSideBarWrapper";
import LoadingElement from "./LoadingElement";

export default function ContentBarLoading() {
  return (
    <RightSideBarWrapper>
      <div className="animate-pulse flex flex-col w-full px-4 border-l-[0.01rem] border-slate-300 dark:border-slate-700">
        <div className="flex flex-col items-begin justify-center mb-4">
          <LoadingElement className="rounded-md w-1/2 h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-40" />
        </div>
        <div className="flex flex-col items-begin justify-center mb-4">
          <LoadingElement className="rounded-md w-1/2 h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-40" />
        </div>
        <div className="flex flex-col items-begin justify-center mb-4">
          <LoadingElement className="rounded-md w-1/2 h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-2 mb-2.5" />
          <LoadingElement className="rounded-md w-full h-40" />
        </div>
      </div>
    </RightSideBarWrapper>
  );
}
