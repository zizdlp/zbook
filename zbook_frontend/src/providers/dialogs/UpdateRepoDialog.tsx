"use client";
import { useContext, useRef, useState } from "react";
import { FormikErrors, FormikValues, useFormik } from "formik";
import { toast } from "react-toastify";

import DialogComponent from "../../components/DialogComponent";
import { OperationContext } from "../OperationProvider";

import FormCommitWrapper from "../../components/wrappers/FormCommitWrapper";
import FormGroupWrapper from "../../components/wrappers/FormGroupWrapper";
import FormInputWrapper from "../../components/wrappers/FormInputWrapper";

import { HiFingerPrint } from "react-icons/hi";
import { useTranslations } from "next-intl";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { UpdateRepoInfoRequest } from "@/fetchs/server_with_auth_request";
import FormTextAreaWrapper from "@/components/wrappers/FormTextAreaWrapper";
import FormListBox from "@/components/wrappers/FormListBox";
import { FetchError } from "@/fetchs/util";
import { isValidateRepoName } from "@/utils/validate";
import { ThemeColor } from "@/components/TableOfContent";
export default function UpdateRepoDialog() {
  const t = useTranslations("Repo");
  const {
    updateRepoOpen,
    setUpdateRepoOpen,
    operationRepoName,
    operationUsername,
  } = useContext(OperationContext);
  const [show, setShow] = useState({ sync_token: false, access_token: false });
  const cancelButtonRef = useRef(null);
  function updateRepoValidate(values: any) {
    let errors: FormikErrors<FormikValues> = {};
    if (values.repo_name != "") {
      if (!isValidateRepoName(values.repo_name)) {
        errors.repo_name = t("InvalidRepoName");
      }
    }
    return errors;
  }
  // 将枚举值转换为 options 数组
  const themeColorOptions = Object.values(ThemeColor).map((color) => ({
    value: color,
    label: t(`ThemeColor${color}`), // 使用国际化，如果需要
  }));

  const formik = useFormik({
    initialValues: {
      username: "",
      old_repo_name: "",
      repo_name: "",
      repo_description: "",
      git_access_token: "",
      visibility_level: "",
      sync_token: "",
      theme_sidebar: "",
      theme_color: "",
    },
    validate: updateRepoValidate,
    onSubmit: onSubmit,
  });
  async function onSubmit(values: UpdateRepoInfoRequest) {
    const id = toast(t("EditingRepository"), {
      type: "info",
      isLoading: true,
    });
    values.old_repo_name = operationRepoName;
    values.username = operationUsername;
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_REPO_INFO,
        xforward: "",
        agent: "",
        tags: [],
        values: values,
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      toast.update(id, {
        render: t("RepositoryEditSucc"),
        type: "success",
        isLoading: false,
        autoClose: 1000,
      });
      setUpdateRepoOpen(false);
      refreshPage("/", true, false);
      formik.resetForm();
    } catch (error) {
      toast.update(id, {
        render: t("FailedEditRepository"),
        type: "error",
        isLoading: false,
        autoClose: 1000,
      });
    }
  }
  return (
    <DialogComponent
      showDialog={updateRepoOpen}
      setShowDialog={setUpdateRepoOpen}
    >
      <form className="px-4 py-4 text-slate-700" onSubmit={formik.handleSubmit}>
        <div className="sm:rounded-md w-full">
          <div className="grid grid-cols-6 gap-4 gap-x-8 p-2">
            <div className="col-span-6 justify-center flex-center">
              <label
                htmlFor="repo_name"
                className="block text-center pt-2 pb-6 text-2xl font-bold text-gray-700 dark:text-slate-50"
              >
                {t("EditRepository")}
              </label>
            </div>

            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="repo_name"
              showName={t("RepositoryName")}
              formik={formik}
            >
              <FormInputWrapper
                show={true}
                name="repo_name"
                placeholder={t("RepositoryNameTip")}
                formik={formik}
                error={formik.errors.repo_name}
              />
            </FormGroupWrapper>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="visibility_level"
              showName={t("VisibleToWho")}
              formik={formik}
            >
              <FormListBox
                options={[
                  { value: "", label: t("ChooseVisibleToWho") },
                  { value: "private", label: t("VisibleOnlyCreator") },
                  { value: "chosen", label: t("VisibleOnlySelected") },
                  { value: "signed", label: t("VisibleOnlyLogin") },
                  { value: "public", label: t("VisibleEveryone") },
                ]}
                nameKey="visibility_level"
                formik={formik}
              />
            </FormGroupWrapper>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="theme_sidebar"
              showName={t("SideBarTheme")}
              formik={formik}
            >
              <FormListBox
                options={[
                  { value: "", label: t("ChooseSideBarTheme") },
                  {
                    value: "theme_sidebar_fold",
                    label: t("ThemeSideBarFold"),
                  },
                  {
                    value: "theme_sidebar_unfold",
                    label: t("ThemeSideBarUnFold"),
                  },
                ]}
                nameKey="theme_sidebar"
                formik={formik}
              />
            </FormGroupWrapper>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="theme_color"
              showName={t("ThemeColor")}
              formik={formik}
            >
              <FormListBox
                options={[
                  { value: "", label: t("ChooseThemeColor") },
                  ...themeColorOptions,
                ]}
                nameKey="theme_color"
                formik={formik}
              />
            </FormGroupWrapper>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-6"
              nameKey="repo_description"
              showName={t("DescribeRepo")}
              formik={formik}
            >
              <FormTextAreaWrapper
                name="repo_description"
                placeholder={t("DescribeRepoTip")}
                row={3}
                formik={formik}
                error={formik.errors.repo_description}
              />
            </FormGroupWrapper>

            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="sync_token"
              showName={t("GitHubSyncToken")}
              formik={formik}
            >
              <div className="flex items-center space-x-2 relative">
                <FormInputWrapper
                  show={show.sync_token}
                  name="sync_token"
                  placeholder={t("GitHubSyncTokenTip")}
                  formik={formik}
                  error={formik.errors.sync_token}
                />
                <div
                  className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer"
                  onClick={() =>
                    setShow({ ...show, sync_token: !show.sync_token })
                  }
                >
                  <HiFingerPrint size={25} />
                </div>
              </div>
            </FormGroupWrapper>
            <FormGroupWrapper
              classType="col-span-6 sm:col-span-3"
              nameKey="git_access_token"
              showName={t("TokenPassword")}
              formik={formik}
            >
              <div className="flex items-center space-x-2 relative">
                <FormInputWrapper
                  show={show.access_token}
                  name="git_access_token"
                  placeholder={t("TokenPasswordTip")}
                  formik={formik}
                  error={formik.errors.git_access_token}
                />
                <div
                  className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer"
                  onClick={() =>
                    setShow({ ...show, access_token: !show.access_token })
                  }
                >
                  <HiFingerPrint size={25} />
                </div>
              </div>
            </FormGroupWrapper>
          </div>
          <FormCommitWrapper
            setOpen={setUpdateRepoOpen}
            cancelButtonRef={cancelButtonRef}
          />
        </div>
      </form>
    </DialogComponent>
  );
}
