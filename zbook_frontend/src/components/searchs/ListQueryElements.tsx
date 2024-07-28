"use client";

import React, { useState, useEffect, useRef, useContext } from "react";
import { useTranslations } from "next-intl";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import NoItemFound from "../NoItemFound";
import { OperationContext } from "@/providers/OperationProvider";
import InfCard from "../notifications/InfCard";
import SearchMarkdownComponent from "./SearchMarkdownComponent";
import ListUserElementForVisibility from "./ListUserElementForVisibility";
import SearchProfileComponent from "./SearchProfileComponent";
import { FetchError } from "@/fetchs/util";
import { SearchType } from "@/utils/const_value";
import LoadingSearchList from "./LoadingSearchList";

export default function ListQueryElements({
  searchType,
  query,
  setShowDialog,
}: {
  searchType: SearchType;
  query: string;
  setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const { operationUsername, operationRepoName } = useContext(OperationContext);
  const t = useTranslations("Searchs");
  const [currentPage, setCurrentPage] = useState(1);
  const isFetchingData = useRef(false);
  const [hasMore, setHasMore] = useState(true);
  const [listModelInfo, setListModelInfo] = useState<any[]>([]);

  useEffect(() => {
    async function fetchMoreData() {
      if (isFetchingData.current) {
        return;
      }
      isFetchingData.current = true;
      try {
        let data = [];
        if (query != "") {
          switch (searchType) {
            case SearchType.DOCUMENT:
              data = await fetchServerWithAuthWrapper({
                endpoint: FetchServerWithAuthWrapperEndPoint.QUERY_MARKDOWN,
                xforward: "",
                agent: "",
                tags: [],
                values: {
                  plain_to_tsquery: query,
                  page_id: currentPage,
                  page_size: 10,
                },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              break;
            case SearchType.USER_DOCUMENT:
              data = await fetchServerWithAuthWrapper({
                endpoint:
                  FetchServerWithAuthWrapperEndPoint.QUERY_USER_MARKDOWN,
                xforward: "",
                agent: "",
                tags: [],
                values: {
                  plain_to_tsquery: query,
                  username: operationUsername,
                  page_id: currentPage,
                  page_size: 10,
                },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              break;
            case SearchType.REPO_DOCUMENT:
              data = await fetchServerWithAuthWrapper({
                endpoint:
                  FetchServerWithAuthWrapperEndPoint.QUERY_REPO_MARKDOWN,
                xforward: "",
                agent: "",
                tags: [],
                values: {
                  username: operationUsername,
                  repo_name: operationRepoName,
                  plain_to_tsquery: query,
                  page_id: currentPage,
                  page_size: 10,
                },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              break;
            case SearchType.VISI_USER:
              data = await fetchServerWithAuthWrapper({
                endpoint:
                  FetchServerWithAuthWrapperEndPoint.LIST_REPO_VISIBILITY,
                xforward: "",
                agent: "",
                tags: [],
                values: {
                  username: operationUsername,
                  repo_name: operationUsername,
                  query: query,
                  page_id: currentPage,
                  page_size: 10,
                },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              break;
            case SearchType.USER:
              data = await fetchServerWithAuthWrapper({
                endpoint: FetchServerWithAuthWrapperEndPoint.QUERY_USER,
                xforward: "",
                agent: "",
                tags: [],
                values: {
                  query: query,
                  page_id: currentPage,
                  page_size: 10,
                },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              break;
            default:
              break;
          }
        } else {
          setListModelInfo([]);
        }

        if (data.elements != undefined && data.elements.length > 0) {
          setListModelInfo((prevState) => [...prevState, ...data.elements]);
        } else {
          setHasMore(false);
        }
        isFetchingData.current = false;
      } catch (error) {
        setHasMore(false);
        isFetchingData.current = false;
      }
    }
    if (!isFetchingData.current) {
      fetchMoreData();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentPage]);
  return (
    <>
      {listModelInfo.map((doc: any, index) => (
        <InfCard
          key={index}
          isLast={index === listModelInfo.length - 1}
          newLimit={() => {
            if (isFetchingData.current || !hasMore) {
              return;
            } else {
              setCurrentPage(currentPage + 1);
            }
          }}
        >
          {searchType === SearchType.VISI_USER && (
            <ListUserElementForVisibility
              repo_username={operationUsername}
              repo_name={operationRepoName}
              is_visible={doc.is_repo_visible}
              username={doc.username}
              email={doc.email}
              is_following={doc.is_following ?? false}
              repo_count={doc.repo_count ?? 0}
              updated_at={doc.updated_at}
            />
          )}

          {searchType != SearchType.VISI_USER && (
            <div onClick={() => setShowDialog(false)}>
              {(searchType === SearchType.REPO_DOCUMENT ||
                searchType === SearchType.DOCUMENT ||
                searchType === SearchType.USER_DOCUMENT) && (
                <SearchMarkdownComponent MarkdownBasicInfo={doc} />
              )}
              {searchType === SearchType.USER && (
                <SearchProfileComponent ListUserInfo={doc} />
              )}
            </div>
          )}
        </InfCard>
      ))}
      {hasMore && <LoadingSearchList itemCount={3} />}
      {!hasMore &&
        currentPage == 1 &&
        listModelInfo.length == 0 &&
        (searchType === SearchType.USER ||
        searchType === SearchType.VISI_USER ? (
          <NoItemFound title={t("NoUser")} />
        ) : (
          <NoItemFound title={t("NoMarkdown")} />
        ))}
    </>
  );
}
