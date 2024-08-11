"use client";

import DialogComponent from "../../components/DialogComponent";
import { OperationContext } from "@/providers/OperationProvider";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import React, { useEffect, useContext, useState } from "react";
import { BsFillSendFill } from "react-icons/bs";

import AvatarImageClient from "@/components/AvatarImageClient";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { useTranslations } from "next-intl";
import { FetchError } from "@/fetchs/util";
import { useSession } from "next-auth/react";
export default function CreateCommentReportDialog() {
  const t = useTranslations("Toast");
  const [username, setUsername] = useState("");
  const { data, status } = useSession();
  useEffect(() => {
    if (data?.username) {
      setUsername(data.username);
    }
  }, [data]);
  const {
    mutationUpdateComment,
    setMutationUpdateComment,
    setOperationCommentID,
    createCommentReportOpen,
    setCreateCommentReportOpen,
    operationCommentID,
  } = useContext(OperationContext);

  const formik = useFormik({
    initialValues: {
      comment_id: 0,
      report_content: "",
    },
    onSubmit: handleSubmit,
  });

  async function handleSubmit(values: any) {
    const id = toast(t("CreatingCommentReport"), {
      type: "info",
      isLoading: true,
    });
    values.comment_id = operationCommentID;

    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_COMMENT_REPORT,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          comment_id: values.comment_id,
          report_content: values.report_content,
        },
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      setCreateCommentReportOpen(false);
      toast.update(id, {
        render: t("ReportSuccessfully"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
      setOperationCommentID(values.comment_id);
      setMutationUpdateComment(!mutationUpdateComment);
      formik.resetForm();
    } catch (error) {
      toast.update(id, {
        render: t("FailedCreateReport"),
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }

  return (
    <DialogComponent
      showDialog={createCommentReportOpen}
      setShowDialog={setCreateCommentReportOpen}
    >
      <header className="">
        <div className="my-4 items-center justify-center mx-4">
          <form onSubmit={formik.handleSubmit}>
            <div className="flex flex-row space-x-4 my-2 py-4  items-center">
              <AvatarImageClient
                username={username}
                className="h-10 w-10 md:h-14 md:w-14 rounded-full"
              />
              <textarea
                id="report_content"
                autoComplete="report_content"
                placeholder={t("WhyReport")}
                className="grow rounded-md border  p-2 resize-none dark:bg-slate-800/25 bg-slate-100/75 
                placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base 
                 border-sky-400 dark:border-sky-800 focus:border-sky-400 focus:ring-sky-400 dark:focus:ring-sky-800 focus:ring-1"
                {...formik.getFieldProps("report_content")}
              />
              <button
                type="submit"
                className="inline-flex items-center px-2 py-2 md:p-4 font-semibold leading-6  text-sm md:text-lg shadow rounded-md text-slate-100 dark:text-slate-300 border-2 border-sky-400 bg-sky-500 dark:bg-transparent dark:border-sky-800 scale-95 hover:scale-100"
              >
                <BsFillSendFill className="w-6 h-6 mr-2 dark:text-sky-500" />
                {t("Submit")}
              </button>
            </div>
          </form>
        </div>
      </header>
    </DialogComponent>
  );
}
