export default function SearchItemWrapper({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div
      className="rounded-md md:rounded-xl my-2 md:my-3 py-2 bg-white dark:bg-[#263142]
        hover:dark:bg-[#39609a] hover:bg-sky-500 hover:text-white border-[0.05rem] border-slate-300/75 dark:border-0 flex items-center justify-between"
    >
      {children}
    </div>
  );
}
