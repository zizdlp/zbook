interface ContentSideBarProps {
  children: React.ReactNode;
}

export default function ContentSideBar(props: ContentSideBarProps) {
  return (
    <div
      className="z-10 hidden xl:flex flex-none pl-10 w-[19rem]"
      id="table-of-contents"
    >
      {props.children}
    </div>
  );
}
