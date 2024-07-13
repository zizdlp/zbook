export default function InputRow({
  show,
  name,
  formik,
  error,
  place_holder,
}: {
  show: boolean;
  name: string;
  formik: any;
  error: string | undefined;
  place_holder: string;
}) {
  return (
    <input
      type={`${show ? "text" : "password"}`}
      placeholder={place_holder}
      autoComplete="on"
      className={`p-3 rounded-md border border-slate-300 dark:border-slate-500  dark:text-slate-300 grow  dark:bg-slate-800 bg-slate-100/50
        placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base 
           ${
             error
               ? "focus:outline-pink-400 dark:focus:outline-pink-600"
               : "focus:outline-sky-500 dark:focus:outline-sky-600"
           }  `}
      {...formik.getFieldProps(name)}
    />
  );
}
