"use client";
import { useEffect, useState } from "react";
import React, { useContext } from "react";
import { Link } from "@/navigation";
import { ListCommentInfo, CommentCountInfo } from "@/types/model";
interface CommentLevelTwoProps {
  ListCommentInfo: ListCommentInfo;
  markdown_id: number;
  pusername: string;
  authname: string;
}
import { useTranslations } from "next-intl";
import CommentOperationCount from "./CommentOperationCount";
import CommentOperationMore from "./CommentOperationMore";
import AvatarImageClient from "../AvatarImageClient";

export default function CommentLevelTwo(props: CommentLevelTwoProps) {
  const t = useTranslations("Dialog");
  const [commentCountInfo, setCommentCountInfo] = useState<CommentCountInfo>();
  const [isDeleted, setIsDeleted] = useState(false);
  const [isHovered, setIsHovered] = useState(false);
  const handleMouseEnter = () => {
    setIsHovered(true);
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
  const handleMouseLeave = () => {
    setIsHovered(false);
  };
  if (isDeleted) {
    return <></>;
  }
  return (
    <div className="flex flex-col  py-2">
      <div className="flex flex-row space-x-4 text-xs">
        <AvatarImageClient
          username={props.ListCommentInfo.username}
          className="h-6 w-6 rounded-full flex-none mt-1"
        />
        <div
          className="grow flex flex-col"
          onMouseEnter={handleMouseEnter}
          onMouseLeave={handleMouseLeave}
        >
          <div className="text-sm font-normal mb-1">
            <span className="font-semibold pr-2">
              {props.ListCommentInfo.username}
            </span>
            <span
              className={`text-sky-500 px-1 ${props.pusername ? "" : "hidden"}`}
            >
              {t("Reply")} {t("At")}
              <Link
                href={`/workspace/${props.pusername}`}
                className="cursor-pointer"
              >
                {props.pusername}
              </Link>
            </span>{" "}
            <span
              className={`break-all text-sm font-normal ${
                isDeleted && "line-through"
              }`}
            >
              {props.ListCommentInfo.comment_content}
            </span>
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
      </div>
    </div>
  );
}
