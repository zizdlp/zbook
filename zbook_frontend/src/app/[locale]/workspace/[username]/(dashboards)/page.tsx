import ListPageWrapper from "@/components/elements/ListPageWrapper";
import { ListDataType } from "@/fetchs/model";
import { Metadata } from "next";
import { getTranslations } from "next-intl/server";
export async function generateMetadata({
  params,
}: {
  params: { username: string };
}): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: params.username + " - " + t("Repo"),
  };
}
export default async function OwnRepoPage({
  params,
  searchParams,
}: {
  params: { username: string; locale: string };
  searchParams?: { query?: string; page?: string };
}) {
  return (
    <ListPageWrapper
      params={params}
      searchParams={searchParams}
      listType={ListDataType.LIST_USER_REPO}
    />
  );
}
