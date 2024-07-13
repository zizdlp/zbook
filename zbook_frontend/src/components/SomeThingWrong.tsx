import { useTranslations } from "next-intl";
import { MdError } from "react-icons/md";

export default function SomeThingWrong() {
  const t = useTranslations("SomeThingWrong");
  return (
    <div className="h-full flex w-full items-center justify-center ">
      <div className="flex flex-col items-center space-y-2">
        <MdError className="h-16 w-16 dark:text-slate-700/50 text-slate-400/50" />
        <p className="py-2 font-semibold text-xl dark:text-slate-600/50 text-slate-500/50">
          {t("SomeThingWrong")}
        </p>
      </div>
    </div>
  );
}
