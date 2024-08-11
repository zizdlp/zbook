"use client";
import React, { useContext } from "react";
import { toast } from "react-toastify";
import { OperationContext } from "../OperationProvider";
import { useTranslations } from "next-intl";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import WarngingDialog from "@/components/WarningDialog";
import { FetchError } from "@/fetchs/util";
export default function DeleteUserDialog() {
  const t = useTranslations("Dialog");
  const { deleteUserOpen, setDeleteUserOpen, operationUsername } =
    useContext(OperationContext);
  return (
    <WarngingDialog
      title={t("DeleteUser")}
      cancelTitle={t("Cancel")}
      submitTitle={t("Delete")}
      showDialog={deleteUserOpen}
      setShowDialog={setDeleteUserOpen}
      cancelFunc={() => setDeleteUserOpen(false)}
      submitFunc={async () => {
        try {
          const data = await fetchServerWithAuthWrapper({
            endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_USER,
            xforward: "",
            agent: "",
            tags: [],
            values: { username: operationUsername },
          });
          if (data.error) {
            throw new FetchError(data.message, data.status);
          }
          refreshPage("/workspace/[username]", true, false);
          setDeleteUserOpen(false);
        } catch (error) {
          toast(t("FailedDeleteUser"), {
            type: "error",
            isLoading: false,
            autoClose: 1500,
          });
        }
      }}
    />
  );
}
