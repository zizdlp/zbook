export default function FormInputWrapper({
  show,
  name,
  formik,
  placeholder,
  error,
}: {
  show: boolean;
  name: string;
  formik: any;
  placeholder: string;
  error: string | undefined;
}) {
  return (
    <input
      type={`${show ? "text" : "password"}`}
      placeholder={placeholder}
      autoComplete="on"
      className={`p-3 rounded-md border border-slate-300 dark:border-slate-500  dark:text-slate-400 grow  dark:bg-slate-800 w-full text-sm
      placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base
         ${
           error
             ? "focus:outline-pink-400 dark:focus:outline-pink-600 outline-1"
             : "focus:outline-sky-500 dark:focus:outline-sky-600 outline-1"
         }  `}
      {...formik.getFieldProps(name)}
    />
  );
}
