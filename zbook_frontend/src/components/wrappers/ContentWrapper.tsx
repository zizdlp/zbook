interface ContentWrapperProps {
  children: React.ReactNode;
}
export default function ContentWrapper(props: ContentWrapperProps) {
  return (
    <div
      className="z-30 px-4 lg:px-12 mr-auto lg:max-w-6xl lg:ml-[max(19rem,calc(50%-26rem))] 2xl:ml-[max(22rem,calc(50%-28rem))]
      "
    >
      <div className="h-10"></div>
      {props.children}
    </div>
  );
}
