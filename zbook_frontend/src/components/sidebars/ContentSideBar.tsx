interface ContentSideBarProps {
  children: React.ReactNode;
}

export default function ContentSideBar(props: ContentSideBarProps) {
  // return (
  //   <div
  //     className="z-10 hidden xl:flex flex-none pl-10 w-[19rem]"
  //     id="table-of-contents"
  //   >
  //     <div className="fixed text-gray-600 text-sm leading-6 w-[16.5rem] overflow-y-auto space-y-2 h-[calc(100%-7rem)]">
  //       {props.children}
  //     </div>
  //   </div>
  // );
  return (
    <div
      className={`hidden xl:block fixed z-40 right-0 2xl:right-[max(0px,calc(50%-50rem))] w-[19rem] 2xl:w-[22rem] inset-y-0 pr-[4rem]
        top-8 lg:top-12
        scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md  overflow-y-scroll 
        `}
    >
      <div className="h-10 bg-white dark:bg-gray-900"></div>
      {props.children}
    </div>
  );
}
