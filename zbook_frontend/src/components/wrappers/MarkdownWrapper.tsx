interface MarkdownWrapperProps {
  children: React.ReactNode;
}
export default function MarkdownWrapper(props: MarkdownWrapperProps) {
  return (
    <div className="max-w-6xl" id="content-container">
      <div className="flex flex-row items-stretch gap-12 pt-[3.5rem] lg:pt-[3rem]">
        <div
          className="relative grow overflow-hidden mx-auto px-1 lg:-ml-12 lg:pl-[23.7rem]"
          id="content-area"
        >
          {props.children}
        </div>
      </div>
    </div>
  );
  return (
    <div
      className="z-30 px-4 lg:px-8 lg:ml-[max(19rem,calc(50%-26rem))] 2xl:ml-[max(22rem,calc(50%-28rem))] xl:max-w-[min(56rem,calc(100vw-19rem-max(19rem,calc(50%-26rem))))] 2xl:max-w-[min(56rem,calc(100vw-44rem))]
    overflow-y-scroll scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md  bg-white dark:bg-gray-900"
    >
      <div className="h-10"></div>
      {props.children}
    </div>
  );
}
