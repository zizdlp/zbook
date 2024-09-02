import ListPageWrapper from "@/components/elements/ListPageWrapper";
import { ListDataType } from "@/fetchs/model";
import { Metadata } from "next";
import { getTranslations } from "next-intl/server";
export async function generateMetadata(): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: t("Cases"),
  };
}

export default async function CasesPage({
  params,
  searchParams,
}: {
  params: { username: string; locale: string };
  searchParams?: { query?: string; page?: string };
}) {
  return (
    <div
      className="z-30 px-4 xl:px-12 mx-auto xl:max-w-6xl pt-10
    overflow-y-scroll scrollbar-thin scrollbar-thumb-rounded-md scrollbar-track-rounded-md  bg-white dark:bg-gray-900"
    >
      <ListPageWrapper
        params={params}
        searchParams={searchParams}
        listType={ListDataType.LIST_PUBLIC_REPO}
      />
    </div>
  );
}
