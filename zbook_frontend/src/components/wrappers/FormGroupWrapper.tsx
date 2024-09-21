export default function FormGroupWrapper({
  classType,
  children,
  formik,
  nameKey,
  showName,
}: {
  classType: string;
  children: React.ReactNode;
  formik: any;
  nameKey: string;
  showName: string;
}) {
  return (
    <div className={classType}>
      <label
        htmlFor={nameKey}
        className="block text-sm font-medium text-gray-700  dark:text-slate-200 pb-1"
      >
        {showName}
      </label>
      {children}
      <div
        className="flex items-center justify-start h-5 pt-1 text-xs text-pink-600 whitespace-nowrap overflow-x-auto  scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-h-[6px]
              scrollbar-thumb-slate-200 scrollbar-track-slate-100
              dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]"
      >
        {formik.errors[nameKey] && (formik.errors[nameKey] as string)}
      </div>
    </div>
  );
}
