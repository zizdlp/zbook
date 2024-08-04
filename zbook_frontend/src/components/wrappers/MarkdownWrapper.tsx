interface MarkdownWrapperProps {
  children: React.ReactNode;
  contentsidebar?: React.ReactNode;
}
export default function MarkdownWrapper(props: MarkdownWrapperProps) {
  return (
    <div className="flex flex-row pt-[42px]">
      <div
        className="relative grow overflow-hidden lg:pl-[21rem] lg:pr-[2rem]"
        id="content-area"
      >
        {props.children}
      </div>
      {props.contentsidebar}
    </div>
  );
}
