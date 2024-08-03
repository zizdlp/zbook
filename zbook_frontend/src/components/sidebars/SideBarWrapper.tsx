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
          className={`z-40 hidden lg:block fixed bottom-0 right-auto w-[18rem] top-[4rem]`}
        >
          {children}
        </div>
      </div>
    </>
  );
}
