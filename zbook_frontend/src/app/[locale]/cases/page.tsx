import ListPageWrapper from "@/components/elements/ListPageWrapper";
import { ListDataType } from "@/fetchs/model";
import { Metadata } from "next";
import { getTranslations } from "next-intl/server";
export async function generateMetadata(): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: t("Cases"),
    description: t("Cases"),
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
    overflow-y-auto scrollbar scrollbar-thumb-rounded-full scrollbar-track-rounded-full scrollbar-w-[6px]
        scrollbar-thumb-slate-200 scrollbar-track-slate-100
        dark:scrollbar-thumb-slate-500/50 dark:scrollbar-track-slate-500/[0.16]  bg-white dark:bg-gray-900"
    >
      <ListPageWrapper
        params={params}
        searchParams={searchParams}
        listType={ListDataType.LIST_PUBLIC_REPO}
      />
    </div>
  );
}
