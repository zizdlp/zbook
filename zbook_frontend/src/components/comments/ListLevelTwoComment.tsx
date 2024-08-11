"use client";
interface CommentReplyListProps {
  root_comment_id: number;
  root_username: string;
  markdown_id: number;
  pageNumber: number;
  pusername: string;
  authname: string;
}

import CommentLevelTwo from "@/components/comments/CommentLevelTwo";
import { ListCommentInfo } from "@/types/model";
import { getCurrentDateTime } from "@/utils/util";
import { useState, useEffect, useContext } from "react";
import { OperationContext } from "@/providers/OperationProvider";
import PageBar from "../PageBar";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
export default function ListLevelTwoComment(props: CommentReplyListProps) {
  const {
    operationParentID,
    mutationCreateComment,
    createCommentContent,
    operationRootID,
    operationCommentID,
  } = useContext(OperationContext);
  const [currentPage, setCurrentPage] = useState(1);
  const [commentlist, setCommentlist] = useState<Array<ListCommentInfo>>([]);
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);
  useEffect(() => {
    async function fetchMoreCommentData() {
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_LEVEL_TWO,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          page_id: currentPage,
          page_size: 5,
          root_id: props.root_comment_id,
        },
      }).then((data: any) => {
        if (data?.comments && data?.comments.length > 0) {
          setCommentlist([...data?.comments]);
        } else {
          setCommentlist([]);
        }
      });
    }
    fetchMoreCommentData();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentPage]);
  useEffect(() => {
    if (mounted && operationRootID == props.root_comment_id) {
      // TODO change pusername
      const newComment: ListCommentInfo = {
        comment_id: operationCommentID,
        markdown_id: props.markdown_id,
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
      setCommentlist((prevList) => [...prevList, newComment]);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [mutationCreateComment]);
  useEffect(() => {
    setCurrentPage(1);
  }, [props.root_comment_id]);
  return (
    <>
      {commentlist.map((model: ListCommentInfo, index) => (
        <CommentLevelTwo
          key={index}
          pusername={model.pusername}
          ListCommentInfo={model as ListCommentInfo}
          markdown_id={props.markdown_id}
          authname={props.authname}
        />
      ))}
      {(currentPage != 1 || commentlist.length != 0) && (
        <PageBar
          totalPages={props.pageNumber}
          currentPage={currentPage}
          setCurrentPage={setCurrentPage}
        />
      )}
    </>
  );
}
