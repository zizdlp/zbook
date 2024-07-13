"use client";
import { ListCommentInfo } from "@/types/model";
import React, { useContext, useState, useEffect, useRef } from "react";
interface ListLevelOneCommentProps {
  markdown_id: number;
  authname: string;
}
import { getCurrentDateTime } from "@/utils/util";
import { OperationContext } from "@/providers/OperationProvider";
import { useTranslations } from "next-intl";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import InfCard from "../notifications/InfCard";
import CommentLevelOne from "./CommentLevelOne";
import CommentLoadingList from "./CommentLoadingList";
export default function ListLevelOneComment(props: ListLevelOneCommentProps) {
  const {
    operationMarkdownID,
    mutationCreateComment,
    createCommentContent,
    operationCommentID,
    operationRootID,
    operationParentID,
  } = useContext(OperationContext);
  const t = useTranslations("Notifications");
  const [currentPage, setCurrentPage] = useState(1);
  const isFetchingData = useRef(false);
  const [hasMore, setHasMore] = useState(true);
  const [listModelInfo, setListModelInfo] = useState<Array<ListCommentInfo>>(
    []
  );
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);
  useEffect(() => {
    async function fetchMoreCommentData() {
      if (isFetchingData.current) {
        return;
      }
      isFetchingData.current = true;
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_LEVEL_ONE,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          page_id: currentPage,
          page_size: 5,
          markdown_id: props.markdown_id,
        },
      }).then((data: any) => {
        if (data?.comments && data?.comments.length > 0) {
          setListModelInfo((prevState) => [...prevState, ...data?.comments]);
        } else {
          setHasMore(false);
        }
        isFetchingData.current = false;
      });
    }
    if (!isFetchingData.current) {
      fetchMoreCommentData();
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentPage]);

  useEffect(() => {
    if (
      mounted &&
      operationMarkdownID == props.markdown_id &&
      !operationRootID
    ) {
      const newComment: ListCommentInfo = {
        comment_id: operationCommentID,
        markdown_id: operationMarkdownID,
        parent_id: operationParentID,
        username: props.authname,
        pusername: "",
        comment_content: createCommentContent,
        created_at: getCurrentDateTime(),
        like_count: 0,
        reply_count: 0,
        is_liked: false,
        is_disliked: false,
        is_shared: false,
        is_reported: false,
      };
      setListModelInfo((prevList) => [newComment, ...prevList]);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [mutationCreateComment]);

  return (
    <>
      {listModelInfo.map((model: ListCommentInfo, index) => (
        <InfCard
          key={model.comment_id}
          isLast={index === listModelInfo.length - 1}
          newLimit={() => {
            if (isFetchingData.current || !hasMore) {
              return;
            } else {
              setCurrentPage(currentPage + 1);
            }
          }}
        >
          <CommentLevelOne
            ListCommentInfo={model as ListCommentInfo}
            markdown_id={props.markdown_id}
            authname={props.authname}
          />
        </InfCard>
      ))}
      {hasMore && <CommentLoadingList itemCount={3} />}
    </>
  );
}
