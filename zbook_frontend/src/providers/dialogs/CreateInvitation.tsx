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
import { FetchError } from "@/fetchs/util";
export default function CreateInvitation() {
  const cancelButtonRef = useRef(null);
  const t = useTranslations("Invitation");
  const { createInvitationOpen, setCreateInvitationOpen } =
    useContext(OperationContext);
  function createInvitationValidate(values: any) {
    const errors: { [key: string]: string } = {};
    if (!values.email) {
      errors.title = t("Required");
    }
    return errors;
  }
  const formik = useFormik({
    initialValues: {
      email: "",
    },
    validate: createInvitationValidate,
    onSubmit: handleSubmit,
  });
  async function handleSubmit(values: any) {
    const id = toast(t("CreatingInvitation"), {
      type: "info",
      isLoading: true,
    });
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CreateInvitation,
        xforward: "",
        agent: "",
        tags: [],
        values: {
          email: values.email,
        },
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      formik.resetForm();
      setCreateInvitationOpen(false);
      toast.update(id, {
        render: t("CreateInvitationSucc"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
    } catch (error) {
      toast.update(id, {
        render: t("FailedCreateInvitation"),
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }
  return (
    <DialogComponent
      showDialog={createInvitationOpen}
      setShowDialog={setCreateInvitationOpen}
    >
      <form
        className="w-fll flex lg:block md:overflow-hidden grow  justify-center overflow-auto  items-center  px-4 py-4  text-slate-700"
        onSubmit={formik.handleSubmit}
      >
        <div className="sm:overflow-hidden sm:rounded-md w-full">
          <div className="grid grid-cols-6 gap-4 gap-x-8 p-2">
            <div className="col-span-6 justify-center flex-center">
              <label
                htmlFor="invitation_email"
                className="block text-center pt-2 pb-6 text-2xl font-bold text-gray-700 dark:text-slate-50"
              >
                {t("CreateInvitation")}
              </label>
            </div>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-6"
              nameKey="email"
              showName={t("Email")}
              formik={formik}
            >
              <div className="flex items-center space-x-2">
                <FormInputWrapper
                  show={true}
                  name="email"
                  placeholder={t("EmailTip")}
                  formik={formik}
                  error={formik.errors.email}
                />
              </div>
            </FormGroupWrapper>
          </div>
          <FormCommitWrapper
            setOpen={setCreateInvitationOpen}
            cancelButtonRef={cancelButtonRef}
          />
        </div>
      </form>
    </DialogComponent>
  );
}
