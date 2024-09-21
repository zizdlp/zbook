"use client";
import { useContext, useRef } from "react";
import { useFormik } from "formik";
import { toast } from "react-toastify";

import { useState } from "react";
import { HiFingerPrint } from "react-icons/hi";
import { fileToBase64 } from "../../utils/util";
import DialogComponent from "../../components/DialogComponent";
import { OperationContext } from "../OperationProvider";
import FormGroupWrapper from "@/components/wrappers/FormGroupWrapper";
import FormInputWrapper from "@/components/wrappers/FormInputWrapper";
import FormCommitWrapper from "@/components/wrappers/FormCommitWrapper";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { useTranslations } from "next-intl";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { FetchError } from "@/fetchs/util";
export default function UpdateUserDialog() {
  const { updateUserOpen, setUpdateUserOpen } = useContext(OperationContext);
  const [fileName, setFileName] = useState("");
  const t = useTranslations("Dialog");
  const cancelButtonRef = useRef(null);
  const [show, setShow] = useState({ password: false, cpassword: false });
  function updateUserValidate(values: any) {
    const errors: { [key: string]: string } = {};
    // validation for password
    if (
      values.password &&
      (values.password.length < 8 || values.password.length > 20)
    ) {
      errors.password = t("CharacterCount");
    }
    if (values.motto && values.password.motto > 128) {
      errors.motto = t("SignatureLength");
    }
    if (values.password && values.password.includes(" ")) {
      errors.password = t("InvalidPassword");
    }
    return errors;
  }

  const formik = useFormik({
    initialValues: {
      motto: "",
      password: "",
      avatar: "",
    },
    validate: updateUserValidate,
    onSubmit: onSubmit,
  });
  async function onSubmit(values: any) {
    // 手动检查表单的有效性
    const errors = updateUserValidate(values);
    if (Object.keys(errors).length > 0) {
      return;
    }
    const id = toast(t("UpdateUserInfo"), {
      type: "info",
      isLoading: true,
    });
    try {
      const base64String = await fileToBase64(values.avatar);
      if (base64String) {
        values.avatar = base64String;
      }
    } catch (error) {}
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_USER,
        xforward: "",
        agent: "",
        tags: [],
        values: values,
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      refreshPage("/workspace/[username]", true, false);
      toast.update(id, {
        render: t("EditUserSucc"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
      setUpdateUserOpen(false);
      formik.resetForm();
    } catch (error) {
      toast.update(id, {
        render: t("FailedEditUser"),
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }
  return (
    <DialogComponent
      showDialog={updateUserOpen}
      setShowDialog={setUpdateUserOpen}
    >
      <form className="px-4 py-4 text-slate-700" onSubmit={formik.handleSubmit}>
        <div className="sm:overflow-hidden sm:rounded-md w-full">
          <div className="grid grid-cols-6 gap-4 gap-x-8 p-2">
            <div className="col-span-6 justify-center flex-center">
              <label
                htmlFor="repo_name"
                className="block text-center pt-2 pb-6 text-2xl font-bold text-gray-700 dark:text-slate-50"
              >
                {t("EditUser")}
              </label>
            </div>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="motto"
              showName={t("Bio")}
              formik={formik}
            >
              <div className="flex items-center space-x-2">
                <FormInputWrapper
                  show={true}
                  name="motto"
                  placeholder={t("BioTip")}
                  formik={formik}
                  error={formik.errors.motto}
                />
              </div>
            </FormGroupWrapper>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="password"
              showName={t("Password")}
              formik={formik}
            >
              <div className="flex items-center space-x-2 relative">
                <FormInputWrapper
                  show={show.password}
                  name="password"
                  placeholder={t("PasswordTip")}
                  formik={formik}
                  error={formik.errors.password}
                />
                <div
                  className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer"
                  onClick={() => setShow({ ...show, password: !show.password })}
                >
                  <HiFingerPrint size={25} />
                </div>
              </div>
            </FormGroupWrapper>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-6"
              nameKey="avatar"
              showName={t("UpdateAvatar")}
              formik={formik}
            >
              <div className="flex items-center space-x-2">
                <label className="block w-full">
                  <div className="w-full cursor-pointer flex">
                    <div
                      className=" flex items-center overflow-scroll justify-center text-sm font-semibold py-2 px-4 bg-sky-400 dark:bg-slate-700 text-white dark:text-slate-200 
                  rounded-full border-0 mr-4 hover:bg-sky-500 dark:hover:bg-sky-600 
                  focus:none focus:outline-sky-200 focus:rounded-full dark:focus:outline-1 dark:focus:outline-slate-700	focus:outline-dashed 
                  "
                    >
                      {t("UploadFile")}
                    </div>
                    {fileName && (
                      <p className="mt-2 text-sm text-slate-500">
                        {t("SelectedFile")} {fileName}
                      </p>
                    )}
                  </div>
                  <input
                    type="file"
                    accept="image/png, image/jpeg"
                    className="hidden"
                    onChange={(event) => {
                      const files = event.currentTarget.files;
                      if (files && files.length > 0) {
                        const selectedFile = files[0];
                        const validTypes = ["image/png", "image/jpeg"];

                        if (
                          selectedFile &&
                          validTypes.includes(selectedFile.type) &&
                          selectedFile.size <= 500 * 1024
                        ) {
                          formik.setFieldValue("avatar", selectedFile);
                          setFileName(selectedFile.name);
                        } else {
                          alert(t("AvatarWarning"));
                          event.currentTarget.value = "";
                          formik.setFieldValue("avatar", null);
                          setFileName("");
                        }
                      }
                    }}
                  />
                </label>
              </div>
            </FormGroupWrapper>
          </div>
          <FormCommitWrapper
            setOpen={setUpdateUserOpen}
            cancelButtonRef={cancelButtonRef}
          />
        </div>
      </form>
    </DialogComponent>
  );
}
