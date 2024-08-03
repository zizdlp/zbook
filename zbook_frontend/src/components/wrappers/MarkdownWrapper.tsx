interface MarkdownWrapperProps {
  children: React.ReactNode;
}
export default function MarkdownWrapper(props: MarkdownWrapperProps) {
  return (
    <div className="max-w-6xl" id="content-container">
      <div className="flex flex-row items-stretch gap-12 pt-[3.5rem] lg:pt-[3.2rem]">
        <div
          className="relative grow overflow-hidden mx-auto px-1 lg:-ml-12 lg:pl-[23.7rem]"
          id="content-area"
        >
          {props.children}
        </div>
      </div>
    </div>
  );
}
