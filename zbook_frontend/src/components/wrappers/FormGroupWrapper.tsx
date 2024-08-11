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
      <div className="flex items-center justify-start h-5 pt-1 text-xs text-pink-600 whitespace-nowrap overflow-x-scroll">
        {formik.errors[nameKey] && (formik.errors[nameKey] as string)}
      </div>
    </div>
  );
}
