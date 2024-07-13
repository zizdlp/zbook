import {
  AiOutlineHeart,
  AiFillHeart,
  AiFillDislike,
  AiOutlineDislike,
} from "react-icons/ai";
import { useContext } from "react";

import { CommentCountInfo } from "@/types/model";
import React, { useState, useEffect } from "react";
import { OperationContext } from "@/providers/OperationProvider";

import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import TimeElement from "../TimeElement";
import { useTranslations } from "next-intl";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
interface CommentOperationLineProps {
  created_at: string | undefined;
  commentCountInfo: CommentCountInfo | undefined;
  setCommentCountInfo: React.Dispatch<
    React.SetStateAction<CommentCountInfo | undefined>
  >;
  comment_id: number;
  markdown_id: number;
}

export default function CommentOperationCount(
  props: CommentOperationLineProps
) {
  const t = useTranslations("Dialog");
  const {
    mutationUpdateComment,
    setMutationUpdateComment,
    setOperationCommentID,
    operationCommentID,
  } = useContext(OperationContext);
  const {
    createCommentOpen,
    setCreateCommentOpen,
    setOperationParentID,
    setOperationMarkdownID,
    setOperationRootID,
  } = useContext(OperationContext);

  const [isMounted, setIsMounted] = useState(false);

  const IconText = ({
    Icon,
    text,
    onClick,
  }: {
    Icon: any;
    text: number;
    onClick: () => void;
  }) => (
    <div onClick={onClick} className="flex items-center mr-4 cursor-pointer">
      <Icon className="mr-1" />
      {text}
    </div>
  );
  const IconFill = ({
    Icon,
    onClick,
  }: {
    Icon: any;

    onClick: () => void;
  }) => (
    <span onClick={onClick} className="flex items-center mr-4 cursor-pointer">
      <Icon className="mr-1" />
    </span>
  );

  useEffect(() => {
    setIsMounted(true);
  }, []);
  useEffect(() => {
    if (isMounted && operationCommentID == props.comment_id) {
      fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_COMMENT_COUNT_INFO,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          comment_id: props.comment_id,
        },
      }).then((data) => {
        if (data?.comment_count_info) {
          props.setCommentCountInfo(data?.comment_count_info);
        }
      });
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [mutationUpdateComment]);

  return (
    <div className="flex">
      <span className="mr-5">
        <TimeElement timeInfo={props.created_at ?? ""} />{" "}
      </span>
      <IconText
        Icon={props.commentCountInfo?.is_liked ? AiFillHeart : AiOutlineHeart}
        text={props.commentCountInfo?.like_count ?? 0}
        onClick={
          props.commentCountInfo?.is_liked
            ? async () => {
                try {
                  const data = await fetchServerWithAuthWrapper({
                    endpoint:
                      FetchServerWithAuthWrapperEndPoint.DELETE_COMMENT_RELATION,
                    xforward: "",
                    agent: "",
                    tags: [],
                    values: {
                      comment_id: props.comment_id,
                      relation_type: "like",
                    },
                  });
                  if (data.error) {
                    throw new FetchError(data.message, data.status);
                  } else {
                    setOperationCommentID(props.comment_id);
                    setMutationUpdateComment(!mutationUpdateComment);
                  }
                } catch (error) {
                  let e = error as FetchError;
                  logger.error(
                    `delete comment relation failed:${e.message}`,
                    e.status
                  );
                }
              }
            : async () => {
                try {
                  const data = await fetchServerWithAuthWrapper({
                    endpoint:
                      FetchServerWithAuthWrapperEndPoint.CREATE_COMMENT_RELATION,
                    xforward: "",
                    agent: "",
                    tags: [],
                    values: {
                      comment_id: props.comment_id,
                      relation_type: "like",
                    },
                  });
                  if (data.error) {
                    throw new FetchError(data.message, data.status);
                  } else {
                    setOperationCommentID(props.comment_id);
                    setMutationUpdateComment(!mutationUpdateComment);
                  }
                } catch (error) {
                  let e = error as FetchError;
                  logger.error(
                    `create comment relation failed:${e.message}`,
                    e.status
                  );
                }
              }
        }
      />
      <IconFill
        Icon={
          props.commentCountInfo?.is_disliked ? AiFillDislike : AiOutlineDislike
        }
        onClick={
          props.commentCountInfo?.is_disliked
            ? () => {
                fetchServerWithAuthWrapper({
                  endpoint:
                    FetchServerWithAuthWrapperEndPoint.DELETE_COMMENT_RELATION,
                  tags: [],
                  xforward: "",
                  agent: "",
                  values: {
                    comment_id: props.comment_id,
                    relation_type: "dislike",
                  },
                }).then((data) => {
                  setOperationCommentID(props.comment_id);
                  setMutationUpdateComment(!mutationUpdateComment);
                });
              }
            : () => {
                fetchServerWithAuthWrapper({
                  endpoint:
                    FetchServerWithAuthWrapperEndPoint.CREATE_COMMENT_RELATION,
                  xforward: "",
                  agent: "",
                  tags: [],
                  values: {
                    comment_id: props.comment_id,
                    relation_type: "dislike",
                  },
                }).then((data) => {
                  setOperationCommentID(props.comment_id);
                  setMutationUpdateComment(!mutationUpdateComment);
                });
              }
        }
      />

      <span
        onClick={() => {
          setOperationParentID(props.comment_id);
          setOperationMarkdownID(props.markdown_id);
          setOperationRootID(props.comment_id);
          setCreateCommentOpen(!createCommentOpen);
        }}
        className="cursor-pointer"
      >
        {t("Reply")}
      </span>
    </div>
  );
}
