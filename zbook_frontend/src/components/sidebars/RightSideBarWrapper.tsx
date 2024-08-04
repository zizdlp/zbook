interface RightSideBarWrapperProps {
  children: React.ReactNode;
}

export default function RightSideBarWrapper(props: RightSideBarWrapperProps) {
  return (
    <div
      className="z-10 hidden xl:flex flex-none w-[19rem]"
      id="table-of-contents"
    >
      {props.children}
    </div>
  );
}
