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
          className={`z-40 hidden lg:block fixed bottom-0 right-auto w-[19rem]  top-[46px] lg:top-[54px]`}
        >
          {children}
        </div>
      </div>
    </>
  );
}
