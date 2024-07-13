interface FormCommitWrapperProps {
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  cancelButtonRef: any;
}
import { useTranslations } from "next-intl";
export default function FormCommitWrapper(props: FormCommitWrapperProps) {
  const t = useTranslations("Dialog");
  return (
    <div className="grid grid-cols-2 gap-x-4 sm:gap-x-6 lg:gap-x-4 xl:gap-x-6 p-4 sm:px-6 sm:py-5 lg:p-4 xl:px-6 xl:py-5">
      <button
        type="reset"
        ref={props.cancelButtonRef}
        className="text-base font-medium rounded-lg bg-white border-[0.05rem] border-slate-300 dark:border-0 text-slate-900 py-3 text-center cursor-pointer dark:bg-gray-700/75 dark:text-slate-200 dark:highlight-white/10"
        onClick={() => props.setOpen(false)}
      >
        {t("Cancel")}
      </button>
      <button
        type="submit"
        className="text-base font-medium rounded-lg bg-sky-500 dark:bg-sky-800 text-white py-3 text-center cursor-pointer dark:highlight-white/20"
      >
        {t("Confirm")}
      </button>
    </div>
  );
}
