import { Link } from "@/navigation";
import Image from "next/image";

import LoginForm from "@/app/[locale]/auth/login/LoginForm";
import { headers } from "next/headers";
import { Metadata } from "next";
import { auth } from "@/auth";
import { redirect } from "@/navigation";
import { getTranslations } from "next-intl/server";
export async function generateMetadata(): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: t("Login"),
    description: t("Login"),
  };
}
export default async function Login() {
  const t = await getTranslations("LoginPage");
  const xforward = headers().get("x-forwarded-for") ?? "";
  const agent = headers().get("User-Agent") ?? "";
  const session = await auth();
  if (session && session.access_token) {
    redirect(`/workspace/${session.username}`); // Navigate to the workspace
  }
  return (
    <div className="overflow-hidden">
      <div className="text-slate-700 dark:text-slate-200 dark:backdrop-blur mx-auto flex items-center max-w-4xl md:p-5">
        <div className="md:w-1/2 min-w-80 md:px-12 mx-auto md:my-4 my-8">
          <Image
            src="/logo_256.png"
            alt="LOGO"
            width={256}
            height={256}
            className="mx-auto h-16 w-auto rounded-lg"
          />

          <div className="flex justify-center items-center flex-shrink-0">
            <h1 className="font-bold text-2xl cursor-pointer">
              <span className="bg-gradient-to-r from-teal-500 to-blue-500 inline-block text-transparent bg-clip-text">
                {t("AppName")}
              </span>
            </h1>
          </div>

          <h2 className="font-bold mt-5 text-2xl">{t("Login")}</h2>
          <p className="text-xs mt-2">{t("SignInMessage")}</p>
          <LoginForm xforward={xforward} agent={agent} />

          <div className="mt-6 text-xs border-b-[0.01rem] border-slate-500 dark:border-slate-300 py-4">
            <Link href="/auth/forget">
              <p className="font-medium  hover:text-sky-500">
                {t("ForgetPassword")}
              </p>
            </Link>
          </div>
          <div className="mt-3 text-xs flex justify-between items-center ">
            <p>{t("NeedAnAccount")}</p>
            <Link href="/auth/register">
              <button className="py-2 px-5 text-white rounded-md bg-sky-500 hover:bg-sky-700 text-sm leading-5 font-semibold">
                {t("Register")}
              </button>
            </Link>
          </div>
        </div>
        <div className="md:block hidden w-1/2">
          <Image src="/login.png" alt="LOGO" width={500} height={500} />
        </div>
      </div>
    </div>
  );
}
