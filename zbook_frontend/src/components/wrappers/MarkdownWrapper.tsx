interface MarkdownWrapperProps {
  children: React.ReactNode;
  contentsidebar: React.ReactNode;
}
export default function MarkdownWrapper(props: MarkdownWrapperProps) {
  return (
    <div id="content-container" className={`border`}>
      <div className="flex flex-row items-stretch gap-12 pt-[42px]">
        <div
          className="relative grow overflow-hidden mx-auto px-1 lg:-ml-12 lg:pl-[23.7rem]"
          id="content-area"
        >
          {props.children}
        </div>
        {props.contentsidebar}
      </div>
    </div>
  );
}
