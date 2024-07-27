"use client";

import React, { useState, useEffect, useRef, useContext } from "react";
import { useTranslations } from "next-intl";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import LoadingNotificationList from "../notifications/LoadingNotificationList";
import NoItemFound from "../NoItemFound";
import { OperationContext } from "@/providers/OperationProvider";
import InfCard from "../notifications/InfCard";
import SearchMarkdownComponent from "./SearchMarkdownComponent";
import ListUserElementForVisibility from "./ListUserElementForVisibility";
import SearchProfileComponent from "./SearchProfileComponent";
import { FetchError } from "@/fetchs/util";

export default function ListQueryElements({
  queryType,
  query,
  setShowDialog,
}: {
  queryType: string;
  query: string;
  setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const { operationUsername, operationRepoID } = useContext(OperationContext);
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
          switch (queryType) {
            case "markdown":
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
            case "userMarkdown":
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
            case "repoMarkdown":
              data = await fetchServerWithAuthWrapper({
                endpoint:
                  FetchServerWithAuthWrapperEndPoint.QUERY_REPO_MARKDOWN,
                xforward: "",
                agent: "",
                tags: [],
                values: {
                  plain_to_tsquery: query,
                  repo_id: operationRepoID,
                  page_id: currentPage,
                  page_size: 10,
                },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              break;
            case "repoVisibleUser":
              data = await fetchServerWithAuthWrapper({
                endpoint:
                  FetchServerWithAuthWrapperEndPoint.LIST_REPO_VISIBILITY,
                xforward: "",
                agent: "",
                tags: [],
                values: {
                  repo_id: operationRepoID,
                  query: query,
                  page_id: currentPage,
                  page_size: 10,
                },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              break;
            case "queryUser":
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
          {queryType == "repoVisibleUser" && (
            <ListUserElementForVisibility
              repo_id={operationRepoID}
              is_visible={doc.is_repo_visible}
              username={doc.username}
              email={doc.email}
              is_following={doc.is_following ?? false}
              repo_count={doc.repo_count ?? 0}
              updated_at={doc.updated_at}
            />
          )}
          {queryType != "repoVisibleUser" && (
            <div onClick={() => setShowDialog(false)}>
              {(queryType == "repoMarkdown" ||
                queryType == "userMarkdown" ||
                queryType == "markdown") && (
                <SearchMarkdownComponent MarkdownBasicInfo={doc} />
              )}
              {queryType == "queryUser" && (
                <SearchProfileComponent ListUserInfo={doc} />
              )}
            </div>
          )}
        </InfCard>
      ))}
      {hasMore && <LoadingNotificationList itemCount={3} />}
      {!hasMore &&
        currentPage == 1 &&
        listModelInfo.length == 0 &&
        (queryType == ("repoVisibleUser" || "queryUser") ? (
          <NoItemFound title={t("NoUser")} />
        ) : (
          <NoItemFound title={t("NoMarkdown")} />
        ))}
    </>
  );
}
