export default function SideBarSettingLoading() {
  return (
    <div
      className={`hidden 2xl:block fixed pt-[22.5rem] z-40 left-[max(0px,calc(50%-48rem))] w-[20rem] h-full
          top-8 lg:top-12
          scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md  overflow-y-scroll 
          `}
    >
      <div className="hidden xl:block flex-shrink-0 py-3 mx-4  rounded-lg animate-pulse bg-slate-100 dark:bg-slate-800 dark:highlight-white/5">
        <div className="px-4">
          <div className="space-y-3 py-2">
            <div className="grid grid-cols-3 gap-4">
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-2"></div>
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-1"></div>
            </div>
            <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded"></div>
          </div>
          <div className="space-y-3 py-2">
            <div className="grid grid-cols-3 gap-4">
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-2"></div>
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-1"></div>
            </div>
            <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded"></div>
          </div>
          <div className="space-y-3 py-2">
            <div className="grid grid-cols-3 gap-4">
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-2"></div>
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-1"></div>
            </div>
            <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded"></div>
          </div>
          <div className="space-y-3 py-2">
            <div className="grid grid-cols-3 gap-4">
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-2"></div>
              <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded col-span-1"></div>
            </div>
            <div className="h-2 bg-slate-200 dark:bg-slate-700 rounded"></div>
          </div>
        </div>
      </div>
    </div>
  );
}
