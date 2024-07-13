import { Suspense } from "react";
import SearchList from "@/components/SearchList";
import LoadingList from "@/components/loadings/LoadingList";
import ListElementWrapper from "./ListElementWrapper";
import { ListDataType } from "@/fetchs/model";
export default async function ListPageWrapper({
  params,
  searchParams,
  listType,
}: {
  params: { username: string };
  searchParams?: { query?: string; page?: string };
  listType: ListDataType;
}) {
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;
  return (
    <>
      <SearchList listType={listType} repo_id={0} />
      <Suspense
        key={query + currentPage}
        fallback={<LoadingList itemCount={10} />}
      >
        <ListElementWrapper
          username={params.username}
          query={query}
          currentPage={currentPage}
          listType={listType}
          repo_id={0}
        />
      </Suspense>
    </>
  );
}
