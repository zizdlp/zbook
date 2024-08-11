export default function FormTextAreaWrapper({
  name,
  formik,
  placeholder,
  row,
  error,
}: {
  name: string;
  formik: any;
  placeholder: string;
  row: number;
  error: string | undefined;
}) {
  return (
    <textarea
      rows={row}
      className={`block w-full resize-none p-3 rounded-md text-sm border border-slate-300 dark:border-slate-500  dark:text-slate-400 grow  dark:bg-slate-800
      placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base 
    ${
      error
        ? "focus:outline-pink-400 dark:focus:outline-pink-600 outline-1"
        : "focus:outline-sky-500 dark:focus:outline-sky-600 outline-1"
    }
    `}
      placeholder={placeholder}
      {...formik.getFieldProps(name)}
    />
  );
}
