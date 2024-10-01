import { getTranslations } from "next-intl/server";
import SendForgetForm from "./SendForgetForm";
import { Metadata } from "next";
export async function generateMetadata(): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: t("ForgetPassword"),
    description: t("ForgetPassword"),
  };
}
export default async function forget() {
  return <SendForgetForm />;
}
