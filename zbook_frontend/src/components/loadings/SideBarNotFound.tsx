import LeftSideBarWrapper from "@/components/sidebars/LeftSideBarWrapper";
import { useTranslations } from "next-intl";
export default function SideBarNotFound() {
  const t = useTranslations("NotFoundPage");
  return (
    <LeftSideBarWrapper small={false}>
      <div className="text-slate-600 min-h-[70vh] dark:text-slate-300 flex items-center justify-center px-4">
        <div className="text-center">
          <h1 className="text-6xl font-bold mb-4">{t("404")}</h1>
          <p className="text-2xl mb-8">{t("PageNotFound")}</p>
          <a
            href="/"
            className="bg-sky-500 hover:bg-sky-700 text-white font-bold py-2 px-4 rounded"
          >
            {t("BackToHome")}
          </a>
        </div>
      </div>
    </LeftSideBarWrapper>
  );
}
