"use client";
import React, { useContext } from "react";
import { toast } from "react-toastify";

import { OperationContext } from "../OperationProvider";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import WarngingDialog from "@/components/WarningDialog";
import { useTranslations } from "next-intl";
import { FetchError } from "@/fetchs/util";
export default function DeleteRepoDialog() {
  const t = useTranslations("Dialog");
  const {
    deleteRepoOpen,
    setDeleteRepoOpen,
    operationUsername,
    operationRepoName,
  } = useContext(OperationContext);
  return (
    <WarngingDialog
      title={t("DeleteRepo")}
      cancelTitle={t("Cancel")}
      submitTitle={t("Delete")}
      showDialog={deleteRepoOpen}
      setShowDialog={setDeleteRepoOpen}
      cancelFunc={() => setDeleteRepoOpen(false)}
      submitFunc={async () => {
        try {
          const data = await fetchServerWithAuthWrapper({
            endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_REPO,
            xforward: "",
            agent: "",
            tags: [],
            values: {
              username: operationUsername,
              repo_name: operationRepoName,
            },
          });
          if (data.error) {
            throw new FetchError(data.message, data.status);
          }
          refreshPage("/workspace/[username]", true, false);
          setDeleteRepoOpen(false);
        } catch (error) {
          toast(t("FailedDeleteRepo"), {
            type: "error",
            isLoading: false,
            autoClose: 1500,
          });
        }
      }}
    />
  );
}
