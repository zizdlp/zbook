export default function ListElementCard({
  header,
  content,
  footer,
}: {
  header: React.ReactNode;
  content: string;
  footer: React.ReactNode;
}) {
  return (
    <div
      className="border border-slate-300/75 dark:border-0 h-44 rounded-lg p-5 flex flex-col dark:shadow
       dark:bg-slate-800/50 dark:hover:bg-slate-800 hover:bg-slate-50"
    >
      <div className="flex-none flex justify-between items-center">
        {header}
      </div>
      <div className="flex-1 text-slate-700 dark:text-slate-400 text-xs my-2 overflow-scroll">
        {content}
      </div>
      <div className="flex-none flex justify-between items-center">
        {footer}
      </div>
    </div>
  );
}
