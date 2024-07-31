import { Link } from "@/navigation";
import Image from "next/image";
import RegisterForm from "@/app/[locale]/auth/register/RegisterForm";
import { Metadata } from "next";
import { getTranslations } from "next-intl/server";
export async function generateMetadata(): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: t("Register"),
  };
}
export default async function Register({
  searchParams,
}: {
  searchParams?: { invitation_url?: string };
}) {
  const invitation_url = searchParams?.invitation_url || "";
  const t = await getTranslations("RegisterPage");
  return (
    <div className="overflow-hidden">
      <div className="text-slate-700 dark:text-slate-200 dark:backdrop-blur mx-auto flex items-center max-w-4xl md:p-5">
        <div className="md:w-1/2 min-w-80 md:px-12 mx-auto">
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

          <h2 className="font-bold mt-5 text-2xl">{t("Register")}</h2>
          <RegisterForm invitation_url={invitation_url} />
          <div className="mt-6 text-xs border-b-[0.01rem] border-slate-500 dark:border-slate-300 py-4 ">
            <Link href="/auth/forget">
              <p className="font-medium hover:text-sky-500">
                {t("ForgetPassword")}
              </p>
            </Link>
          </div>
          <div className="mt-3 text-xs flex justify-between items-center">
            <p>{t("OwnAnAccount")}</p>
            <Link href="/auth/login">
              <button className="py-2 px-5 text-white rounded-md bg-sky-500 hover:bg-sky-700 text-sm leading-5 font-semibold">
                {t("Login")}
              </button>
            </Link>
          </div>
        </div>
        <div className="md:block hidden w-1/2">
          <Image src="/login.png" alt="Login" width={500} height={50} />
        </div>
      </div>
    </div>
  );
}
