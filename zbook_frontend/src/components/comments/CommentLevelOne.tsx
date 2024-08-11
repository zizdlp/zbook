"use client";
import { useState, useEffect } from "react";
import { ListCommentInfo, CommentCountInfo } from "@/types/model";
interface CommentLevelOneProps {
  ListCommentInfo: ListCommentInfo;
  markdown_id: number;
  authname: string;
}

import ListLevelTwoComment from "./ListLevelTwoComment";
import CommentOperationMore from "./CommentOperationMore";
import CommentOperationCount from "./CommentOperationCount";
import AvatarImageClient from "../AvatarImageClient";
function getPageNumber(n: number) {
  let nn = parseInt(String(n));
  return Math.floor((nn + parseInt("4")) / 5);
}
export default function CommentLevelOne(props: CommentLevelOneProps) {
  const [commentCountInfo, setCommentCountInfo] = useState<CommentCountInfo>();
  const [isDeleted, setIsDeleted] = useState(false);

  const [isHovered, setIsHovered] = useState(false);

  const handleMouseEnter = () => {
    setIsHovered(true);
  };
  const handleMouseLeave = () => {
    setIsHovered(false);
  };
  useEffect(() => {
    let commentCountInfo: CommentCountInfo = {
      like_count: props.ListCommentInfo.like_count,
      reply_count: props.ListCommentInfo.reply_count,
      is_liked: props.ListCommentInfo.is_liked,
      is_disliked: props.ListCommentInfo.is_disliked,
      is_shared: props.ListCommentInfo.is_shared,
      is_reported: props.ListCommentInfo.is_reported,
    };
    setCommentCountInfo(commentCountInfo);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  if (isDeleted) {
    return <></>;
  }

  return (
    <div className="flex flex-col border-b dark:border-b-slate-700/75 border-b-slate-200/75 pt-4 text-slate-700 dark:text-slate-400 pb-2">
      <div className="flex flex-row space-x-2 md:space-x-4">
        <AvatarImageClient
          username={props.ListCommentInfo.username}
          className="h-10 w-10 rounded-full"
        />
        <div className="grow flex flex-col">
          <div className="leading-7 font-semibold text-sm mb-1">
            {props.ListCommentInfo.username}
          </div>
          <div
            className="pt-0.5"
            onMouseEnter={handleMouseEnter}
            onMouseLeave={handleMouseLeave}
          >
            <div
              className={`break-all text-sm font-normal ${
                isDeleted && "line-through"
              }`}
            >
              {props.ListCommentInfo.comment_content}
            </div>
            <div className="text-sm leading-6 mt-0.5 flex items-center justify-between text-gray-500 dark:text-slate-500">
              <CommentOperationCount
                created_at={props.ListCommentInfo.created_at}
                commentCountInfo={commentCountInfo}
                setCommentCountInfo={setCommentCountInfo}
                comment_id={props.ListCommentInfo.comment_id}
                markdown_id={props.markdown_id}
              />
              {isHovered && (
                <CommentOperationMore
                  comment_id={props.ListCommentInfo.comment_id}
                  commentCountInfo={commentCountInfo}
                  markdown_id={props.markdown_id}
                  setIsDeleted={setIsDeleted}
                  owned={props.ListCommentInfo.username === props.authname}
                />
              )}
            </div>
          </div>
          <ListLevelTwoComment
            pusername={props.ListCommentInfo.username}
            root_comment_id={props.ListCommentInfo.comment_id}
            root_username={props.ListCommentInfo.username}
            markdown_id={props.markdown_id}
            pageNumber={getPageNumber(commentCountInfo?.reply_count ?? 0)}
            authname={props.authname}
          />
        </div>
      </div>
    </div>
  );
}
