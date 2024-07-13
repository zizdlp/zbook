interface MarkdownWrapperProps {
  children: React.ReactNode;
}
export default function MarkdownWrapper(props: MarkdownWrapperProps) {
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
