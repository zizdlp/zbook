import { auth } from "@/auth";
import CallSignIn from "@/components/CallSignIn";
import { getTranslations } from "next-intl/server";
import { redirect } from "@/navigation";
export default async function OAuth({
  searchParams,
}: {
  searchParams?: { oauth_type?: string };
}) {
  const t = await getTranslations("NotFoundPage");

  const oauthType = searchParams?.oauth_type || "";
  const session = await auth();
  if (!session) {
    return <CallSignIn oauthType={oauthType} />;
  } else if (session.access_token) {
    redirect(`/workspace/${session.username}`); // Navigate to the workspace
  } else {
    return (
      <div className="text-slate-600 fixed h-full w-full inset-0 dark:text-slate-300 flex items-center justify-center px-4">
        <div className="text-center">
          <h1 className="text-6xl font-bold mb-4">{t("404")}</h1>
          <p className="text-2xl mb-8">{t("UserNotLinked")}</p>
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
