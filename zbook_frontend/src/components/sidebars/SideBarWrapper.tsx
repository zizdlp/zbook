import SideBarTransition from "./SideBarTransition";
export default function SideBarWrapper({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      <div className="block lg:hidden">
        <SideBarTransition>{children}</SideBarTransition>
      </div>
      <div className="hidden lg:block">
        <div
          className={`hidden lg:block fixed left-[max(0px,calc(50%-45rem))] 2xl:left-[max(0px,calc(50%-50rem))] z-40 inset-0 w-[19rem] 2xl:w-[22rem] top-0 pt-8 lg:pt-12 px-4 border-r-[0.01rem] border-slate-300 dark:border-slate-700
          overflow-y-scroll
          scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md bg-white dark:bg-gray-900`}
        >
          {children}
        </div>
      </div>
    </>
  );
}
