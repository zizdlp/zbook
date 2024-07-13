import React, { useContext } from "react";
import { OperationContext } from "@/providers/OperationProvider";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import WarngingDialog from "../WarningDialog";
import { useTranslations } from "next-intl";
import { toast } from "react-toastify";
import { FetchError } from "@/fetchs/util";
interface DeleteCommentDialog {
  showDialog: boolean;
  setShowDialog: React.Dispatch<React.SetStateAction<boolean>>;
  comment_id: number;
  markdown_id: number;
  setIsDeleted: React.Dispatch<React.SetStateAction<boolean>>;
}
export default function DeleteCommentDialog(props: DeleteCommentDialog) {
  const t = useTranslations("Dialog");
  const {
    mutationDeleteComment,
    setMutationDeleteComment,
    setOperationMarkdownID,
  } = useContext(OperationContext);
  return (
    <WarngingDialog
      title={t("DeleteComment")}
      cancelTitle={t("Cancel")}
      submitTitle={t("Delete")}
      showDialog={props.showDialog}
      setShowDialog={props.setShowDialog}
      cancelFunc={() => props.setShowDialog(false)}
      submitFunc={async () => {
        try {
          const data = await fetchServerWithAuthWrapper({
            endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_COMMENT,
            xforward: "",
            agent: "",
            tags: [],
            values: {
              comment_id: props.comment_id,
            },
          });
          if (data.error) {
            throw new FetchError(data.message, data.status);
          }
          props.setIsDeleted(true);
          setOperationMarkdownID(props.markdown_id);
          setMutationDeleteComment(!mutationDeleteComment);
          props.setShowDialog(false);
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
