import { getVerify } from "@/fetchs/server_without_auth";
import { getTranslations } from "next-intl/server";
export default async function VerifyEmail({
  searchParams,
}: {
  searchParams?: {
    verification_id?: string;
  };
}) {
  const t = await getTranslations("VerifyEmail");
  try {
    const verification_id = searchParams?.verification_id || "";
    const verify_result = await getVerify(verification_id);
    if (verify_result && verify_result.code == undefined) {
      return (
        <div className="text-slate-600 fixed inset-0 dark:text-slate-300 flex items-center justify-center px-4">
          <div className="text-center">
            <p className="text-2xl mb-8">{t("VerifiedSuccess")}</p>
            <a
              href="/"
              className="bg-sky-500 hover:bg-sky-700 text-white font-bold py-2 px-4 rounded"
            >
              {t("BackToHome")}
            </a>
          </div>
        </div>
      );
    } else {
      return (
        <div className="text-slate-600 fixed inset-0 dark:text-slate-300 flex items-center justify-center px-4">
          <div className="text-center">
            <p className="text-2xl mb-8">
              {t("VerifiedFailed")}
              {verify_result.message}
            </p>
            <a
              href="/"
              className="bg-sky-500 hover:bg-sky-700 text-white font-bold py-2 px-4 rounded"
            >
              {t("BackToHome")}
            </a>
          </div>
        </div>
      );
    }
  } catch (error) {
    return (
      <div className="text-slate-600 fixed inset-0 dark:text-slate-300 flex items-center justify-center px-4">
        <div className="text-center">
          <p className="text-2xl mb-8">{t("VerifiedFailed")}</p>
          <a
            href="/"
            className="bg-sky-500 hover:bg-sky-700 text-white font-bold py-2 px-4 rounded"
          >
            {t("BackToHome")}
          </a>
        </div>
      </div>
    );
  }
}
