"use client";
import React, { useContext } from "react";
import { OperationContext } from "@/providers/OperationProvider";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import WarngingDialog from "../WarningDialog";
import { useTranslations } from "next-intl";
import { toast } from "react-toastify";
import { FetchError } from "@/fetchs/util";

export default function GlobalDeleteCommentDialog() {
  const t = useTranslations("Dialog");
  const {
    mutationDeleteComment,
    setMutationDeleteComment,
    setOperationMarkdownID,
    deleteCommentOpen,
    setDeleteCommentOpen,
    operationCommentID,
  } = useContext(OperationContext);
  return (
    <WarngingDialog
      title={t("DeleteComment")}
      cancelTitle={t("Cancel")}
      submitTitle={t("Delete")}
      showDialog={deleteCommentOpen}
      setShowDialog={setDeleteCommentOpen}
      cancelFunc={() => setDeleteCommentOpen(false)}
      submitFunc={async () => {
        try {
          const data = await fetchServerWithAuthWrapper({
            endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_COMMENT,
            xforward: "",
            agent: "",
            tags: [],
            values: {
              comment_id: operationCommentID,
            },
          });
          if (data.error) {
            throw new FetchError(data.message, data.status);
          }
          setMutationDeleteComment(!mutationDeleteComment);
          setDeleteCommentOpen(false);
          refreshPage("/", true, false);
        } catch (error) {
          toast(t("FailedDeleteComment"), {
            type: "error",
            isLoading: false,
            autoClose: 1500,
          });
        }
      }}
    />
  );
}
