import SideBarTransition from "./SideBarTransition";
export default function LeftSideBarWrapper({
  children,
  small,
}: {
  children: React.ReactNode;
  small: boolean;
}) {
  return (
    <>
      {small && (
        <div className="block lg:hidden">
          <SideBarTransition>{children}</SideBarTransition>
        </div>
      )}

      <div className="hidden lg:block">
        <div
          className={`z-40 hidden lg:block fixed bottom-0 right-auto w-[19.5rem]  top-[46px] lg:top-[54px]`}
        >
          {children}
        </div>
      </div>
    </>
  );
}
