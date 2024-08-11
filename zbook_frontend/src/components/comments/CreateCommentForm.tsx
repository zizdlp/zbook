"use client";
import { OperationContext } from "@/providers/OperationProvider";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import React, { useContext } from "react";
import { BsFillSendFill } from "react-icons/bs";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import AvatarImageClient from "../AvatarImageClient";
import { useTranslations } from "next-intl";
import { FetchError } from "@/fetchs/util";
import { ThemeColor } from "../TableOfContent";

function getCommentColorClasses(color: ThemeColor) {
  return {
    textAreaClass: `border-${color}-400 dark:border-${color}-800 focus:border-${color}-500 dark:focus:border-${color}-600`,
    buttonClass: `bg-${color}-600 hover:bg-${color}-700 dark:bg-${color}-700/50 hover:dark:bg-${color}-800/50`,
  };
}
export default function CreateCommentForm({
  markdownID,
  parentID,
  username,
  theme_color,
}: {
  markdownID: number;
  parentID: number;
  username: string;
  theme_color: ThemeColor;
}) {
  let { textAreaClass, buttonClass } = getCommentColorClasses(theme_color);
  const t = useTranslations("Dialog");
  const {
    mutationCreateComment,
    setMutationCreateComment,
    setCreateCommentContent,
    setOperationCommentID,
    setOperationRootID,
    setOperationMarkdownID,
    setOperationParentID,
    operationParentID,
    setCreateCommentOpen,
  } = useContext(OperationContext);
  const formik = useFormik({
    initialValues: {
      markdown_id: 0,
      parent_id: 0,
      comment_content: "",
    },
    onSubmit: handleSubmit,
  });

  async function handleSubmit(values: any) {
    values.markdown_id = markdownID;
    values.parent_id = parentID;
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_COMMENT,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          markdown_id: values.markdown_id,
          comment_content: values.comment_content,
          parent_id: values.parent_id,
        },
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      setOperationCommentID(data.comment.comment_id);
      setOperationMarkdownID(data.comment.markdown_id);
      setOperationRootID(data.comment.root_id ?? 0);
      setOperationParentID(data.comment.parent_id ?? 0);
      setCreateCommentContent(data.comment.comment_content);
      setMutationCreateComment(!mutationCreateComment);
      setCreateCommentOpen(false);
      formik.resetForm();
    } catch {
      toast(t("FailedCreateComment"), {
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }

  return (
    <form onSubmit={formik.handleSubmit} className="w-full">
      <div className="flex flex-row space-x-4 my-2 py-2 items-center">
        <AvatarImageClient
          username={username}
          className="h-10 w-10 rounded-full"
        />
        <textarea
          id="comment_content"
          autoComplete="comment_content"
          placeholder={t("RespectComments")}
          className={`grow rounded-md border-2 p-2 resize-none dark:bg-slate-800/25 bg-slate-100/75
          placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base 
          ${textAreaClass}  h-16
          outline-0 ring-0`}
          {...formik.getFieldProps("comment_content")}
        />
        <button
          type="submit"
          className={`inline-flex items-center p-2 md:p-4 font-medium
          leading-6 text-sm md:text-base h-16
          rounded-md
          text-slate-100 dark:text-slate-300 
          ${buttonClass}
          `}
        >
          <BsFillSendFill className="w-4 h-4 mr-2" />
          {t("Publish")}
        </button>
      </div>
    </form>
  );
}
