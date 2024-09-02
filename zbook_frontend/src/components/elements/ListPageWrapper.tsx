import { Suspense } from "react";
import SearchList from "@/components/SearchList";
import LoadingList from "@/components/loadings/LoadingList";
import ListElementWrapper from "./ListElementWrapper";
import { ListDataType } from "@/fetchs/model";
import { auth } from "@/auth";
export default async function ListPageWrapper({
  params,
  searchParams,
  listType,
}: {
  params: { username: string; locale: string };
  searchParams?: { query?: string; page?: string };
  listType: ListDataType;
}) {
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;
  let authname = "";
  const session = await auth();
  if (session?.access_token) {
    authname = session.username;
  }
  return (
    <>
      <SearchList listType={listType} username="" repo_name="" />
      <Suspense
        key={query + currentPage}
        fallback={<LoadingList itemCount={10} />}
      >
        <ListElementWrapper
          authname={authname}
          username={params.username}
          query={query}
          currentPage={currentPage}
          listType={listType}
          repo_name={""}
          lang={params.locale == "" ? "en" : params.locale}
        />
      </Suspense>
    </>
  );
}
