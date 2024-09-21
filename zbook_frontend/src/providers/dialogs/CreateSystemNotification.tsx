"use client";
import { useRef, useContext } from "react";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import { OperationContext } from "@/providers/OperationProvider";
import DialogComponent from "../../components/DialogComponent";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import FormGroupWrapper from "@/components/wrappers/FormGroupWrapper";
import FormInputWrapper from "@/components/wrappers/FormInputWrapper";
import FormCommitWrapper from "@/components/wrappers/FormCommitWrapper";
import { useTranslations } from "next-intl";
import FormTextAreaWrapper from "@/components/wrappers/FormTextAreaWrapper";
import { FetchError } from "@/fetchs/util";
export default function CreateSystemNotification() {
  const cancelButtonRef = useRef(null);
  const t = useTranslations("SystemNotification");
  const { CreateSystemNotificationOpen, setCreateSystemNotificationOpen } =
    useContext(OperationContext);
  function createSystemNotificationValidate(values: any) {
    const errors: { [key: string]: string } = {};
    if (!values.title) {
      errors.title = t("Required");
    }
    if (!values.contents) {
      errors.contents = t("Required");
    }
    if (!values.username) {
      errors.username = t("Required");
    }
    return errors;
  }
  const formik = useFormik({
    initialValues: {
      username: "",
      title: "",
      contents: "",
      redirect_url: "",
    },
    validate: createSystemNotificationValidate,
    onSubmit: handleSubmit,
  });
  async function handleSubmit(values: any) {
    const id = toast(t("Creatingnotification"), {
      type: "info",
      isLoading: true,
    });
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_SYSTEM_NOTIFICATION,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          username: values.username,
          title: values.title,
          contents: values.contents,
          redirect_url: values.redirect_url,
        },
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      formik.resetForm();
      setCreateSystemNotificationOpen(false);
      toast.update(id, {
        render: t("CreateSystemNotificationsSucc"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
    } catch (error) {
      toast.update(id, {
        render: t("FailedCreateSystemNotifications"),
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }
  return (
    <DialogComponent
      showDialog={CreateSystemNotificationOpen}
      setShowDialog={setCreateSystemNotificationOpen}
    >
      <form className="px-4 py-4 text-slate-700" onSubmit={formik.handleSubmit}>
        <div className="sm:overflow-hidden sm:rounded-md w-full">
          <div className="grid grid-cols-6 gap-4 gap-x-8 p-2">
            <div className="col-span-6 justify-center flex-center">
              <label
                htmlFor="system_notification_title"
                className="block text-center pt-2 pb-6 text-2xl font-bold text-gray-700 dark:text-slate-50"
              >
                {t("CreateSystemNotifications")}
              </label>
            </div>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-6"
              nameKey="title"
              showName={t("Title")}
              formik={formik}
            >
              <div className="flex items-center space-x-2">
                <FormInputWrapper
                  show={true}
                  name="title"
                  placeholder={t("TitleTip")}
                  formik={formik}
                  error={formik.errors.title}
                />
              </div>
            </FormGroupWrapper>

            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="username"
              showName={t("Username")}
              formik={formik}
            >
              <div className="flex items-center space-x-2">
                <FormInputWrapper
                  show={true}
                  name="username"
                  placeholder={t("UsernameTip")}
                  formik={formik}
                  error={formik.errors.username}
                />
              </div>
            </FormGroupWrapper>

            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="redirect_url"
              showName={t("LinkOptional")}
              formik={formik}
            >
              <div className="flex items-center space-x-2">
                <FormInputWrapper
                  show={true}
                  name="redirect_url"
                  placeholder={t("LinkOptionalTip")}
                  formik={formik}
                  error={formik.errors.redirect_url}
                />
              </div>
            </FormGroupWrapper>

            <FormGroupWrapper
              classType="col-span-6 sm:col-span-6"
              nameKey="contents"
              showName={t("Content")}
              formik={formik}
            >
              <FormTextAreaWrapper
                name="contents"
                placeholder={t("ContentTip")}
                row={3}
                formik={formik}
                error={formik.errors.contents}
              />
            </FormGroupWrapper>
          </div>
          <FormCommitWrapper
            setOpen={setCreateSystemNotificationOpen}
            cancelButtonRef={cancelButtonRef}
          />
        </div>
      </form>
    </DialogComponent>
  );
}
