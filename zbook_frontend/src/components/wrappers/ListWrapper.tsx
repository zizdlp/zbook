interface ListWrapperProps {
  children: React.ReactNode;
}
export default function ListWrapper(props: ListWrapperProps) {
  return (
    <div
      className="z-30 xl:px-12 mx-auto xl:max-w-[min(70rem,calc(100vw-40rem))] border border-red-500
    overflow-y-scroll  scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md  bg-white dark:bg-gray-900"
    >
      {props.children}
    </div>
  );
}
