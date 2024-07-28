import ListElementWrapper from "@/components/elements/ListElementWrapper";
import LoadingList from "@/components/loadings/LoadingList";
import ContentWrapper from "@/components/wrappers/ContentWrapper";
import { ListDataType } from "@/fetchs/model";
import { fetchServerWithoutAuthWrapper } from "@/fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";
import { FetchError } from "@/fetchs/util";
import { Metadata } from "next";
import { headers } from "next/headers";
import { Suspense } from "react";
import SearchList from "@/components/SearchList";
import { getTranslations } from "next-intl/server";
import { auth } from "@/auth";
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
  params: { username: string; reponame: string };
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
    const xforward = headers().get("x-forwarded-for") ?? "";
    const agent = headers().get("User-Agent") ?? "";

    const data = await fetchServerWithoutAuthWrapper({
      endpoint: FetchServerWithoutAuthWrapperEndPoint.GET_REPO_ID,
      xforward: xforward,
      agent: agent,
      values: { repo_name: params.reponame, username: params.username },
    });

    if (data.error) {
      throw new FetchError(data.message, data.status);
    }

    return (
      <ContentWrapper>
        <SearchList
          listType={ListDataType.LIST_REPO_VISI}
          repo_id={data.repo_id}
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
            repo_id={data.repo_id}
          />
        </Suspense>
      </ContentWrapper>
    );
  } catch {}
}
