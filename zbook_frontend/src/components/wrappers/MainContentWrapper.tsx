interface MainContentWrapperProps {
  children: React.ReactNode;
  right_sidebar?: React.ReactNode;
}
export default function MainContentWrapper(props: MainContentWrapperProps) {
  return (
    <div className="flex flex-row pt-[42px]">
      <div
        className="relative grow overflow-hidden lg:pl-[21rem] lg:pr-[2.5rem]"
        id="content-area"
      >
        {props.children}
      </div>
      {props.right_sidebar}
    </div>
  );
}
