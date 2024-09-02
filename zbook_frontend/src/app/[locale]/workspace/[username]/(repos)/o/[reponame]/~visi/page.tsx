import ListElementWrapper from "@/components/elements/ListElementWrapper";
import LoadingList from "@/components/loadings/LoadingList";
import { ListDataType } from "@/fetchs/model";
import { Metadata } from "next";
import { Suspense } from "react";
import SearchList from "@/components/SearchList";
import { getTranslations } from "next-intl/server";
import { auth } from "@/auth";
import MainContentWrapper from "@/components/wrappers/MainContentWrapper";
export async function generateMetadata({
  params,
}: {
  params: { username: string; reponame: string };
}): Promise<Metadata> {
  const t = await getTranslations("GenerateMetaData");
  return {
    title: params.username + " - repo - " + params.reponame + t("WhoCanSee"),
  };
}
export default async function ListRepoVisi({
  params,
  searchParams,
}: {
  params: { username: string; reponame: string; locale: string };
  searchParams?: { query?: string; page?: string };
}) {
  try {
    let authname = "";
    const session = await auth();
    if (session?.access_token) {
      authname = session.username;
    }
    const query = searchParams?.query || "";
    const currentPage = Number(searchParams?.page) || 1;

    return (
      <MainContentWrapper>
        <SearchList
          listType={ListDataType.LIST_REPO_VISI}
          username={params.username}
          repo_name={params.reponame}
        />
        <Suspense
          key={query + currentPage}
          fallback={<LoadingList itemCount={10} />}
        >
          <ListElementWrapper
            username={params.username}
            authname={authname}
            query={query}
            currentPage={currentPage}
            listType={ListDataType.LIST_REPO_VISI}
            repo_name={params.reponame}
            lang={params.locale == "" ? "en" : params.locale}
          />
        </Suspense>
      </MainContentWrapper>
    );
  } catch {}
}
