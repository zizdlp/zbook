"use client";
import { useContext, useRef } from "react";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import { FormikErrors } from "formik";
import { useState } from "react";
import { HiFingerPrint } from "react-icons/hi";
import DialogComponent from "../../components/DialogComponent";
import { OperationContext } from "../OperationProvider";
import FormInputWrapper from "../../components/wrappers/FormInputWrapper";
import FormGroupWrapper from "../../components/wrappers/FormGroupWrapper";
import FormCommitWrapper from "../../components/wrappers/FormCommitWrapper";
import {
  fetchServerWithAuthWrapper,
  refreshPage,
} from "@/fetchs/server_with_auth";
import { useTranslations } from "next-intl";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { CreateRepoRequest } from "@/fetchs/server_with_auth_request";
import FormTextAreaWrapper from "@/components/wrappers/FormTextAreaWrapper";
import { isValidateRepoName, isValidGitURL } from "@/utils/validate";
import FormListBox from "@/components/wrappers/FormListBox";
import { FetchError } from "@/fetchs/util";
import { ThemeColor } from "@/components/TableOfContent";

export default function CreateRepoDialog() {
  const { createRepoOpen, setCreateRepoOpen } = useContext(OperationContext);
  const t = useTranslations("Repo");
  const cancelButtonRef = useRef(null);
  const [show, setShow] = useState({ sync_token: false, access_token: false });
  // 将枚举值转换为 options 数组
  const themeColorOptions = Object.values(ThemeColor).map((color) => ({
    value: color,
    label: t(`ThemeColor${color}`), // 使用国际化，如果需要
  }));

  interface FormValues {
    repo_name: string;
    repo_description: string;
    git_addr: string;
    git_access_token: string;
    sync_token: string;
    visibility_level: string;
    theme_color: string;
    theme_sidebar: string;
  }
  function validateCreateRepo(values: any) {
    let errors: FormikErrors<FormValues> = {};
    if (!values.repo_name) {
      errors.repo_name = t("Required");
    } else if (!isValidateRepoName(values.repo_name)) {
      errors.repo_name = t("InvalidRepoName");
    }
    if (!values.git_addr) {
      errors.git_addr = t("Required");
    }
    if (!values.theme_color) {
      errors.theme_color = t("Required");
    }
    if (!values.theme_sidebar) {
      errors.theme_sidebar = t("Required");
    }
    if (!isValidGitURL(values.git_addr)) {
      errors.git_addr = t("InValidGitUrl");
    }
    if (!values.repo_description) {
      errors.repo_description = t("Required");
    }
    if (
      values.visibility_level != "public" &&
      values.visibility_level != "private" &&
      values.visibility_level != "chosen" &&
      values.visibility_level != "signed"
    ) {
      errors.visibility_level = t("Required");
    }
    return errors;
  }
  const formik = useFormik({
    initialValues: {
      repo_name: "",
      repo_description: "",
      git_addr: "",
      git_access_token: "",
      visibility_level: "",
      sync_token: "",
      theme_sidebar: "",
      theme_color: "",
    },
    validate: validateCreateRepo,
    onSubmit: onSubmit,
  });
  async function onSubmit(values: CreateRepoRequest) {
    const id = toast(t("CreatingRepository"), {
      type: "info",
      isLoading: true,
    });
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_REPO,
        xforward: "",
        agent: "",
        tags: [],
        values: values,
        timeout: 600000, //600s
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      await refreshPage("/workspace/[username]", true, false);
      formik.resetForm();
      setCreateRepoOpen(false);
      toast.update(id, {
        render: t("RepositoryCreatedSucc"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
    } catch (error) {
      toast.update(id, {
        render: t("FailedCreateRepository") + error,
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }
  return (
    <DialogComponent
      showDialog={createRepoOpen}
      setShowDialog={setCreateRepoOpen}
    >
      <form
        className="w-fll flex lg:block md:overflow-hidden grow  justify-center overflow-auto  items-center  px-4 py-4  text-slate-700 "
        onSubmit={formik.handleSubmit}
      >
        <div className="sm:overflow-hidden sm:rounded-md w-full">
          <div className="grid grid-cols-6 gap-4 gap-x-8 p-2">
            <div className="col-span-6 justify-center flex-center">
              <label
                htmlFor="repo_name"
                className="block text-center pt-2 pb-6 text-2xl font-bold text-gray-700 dark:text-slate-50"
              >
                {t("CreateRepository")}
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
              nameKey="git_addr"
              showName={t("GitUrl")}
              formik={formik}
            >
              <div className="flex rounded-md shadow-sm">
                <span className="p-3 rounded-l-md border border-r-0 border-slate-300 dark:border-slate-500  dark:text-slate-400 grow  dark:bg-slate-800  text-sm">
                  {t("GitProtocol")}
                </span>

                <input
                  type={"text"}
                  id={"git_addr"}
                  autoComplete={"git_addr"}
                  placeholder={t("GitUrlTip")}
                  className={`p-3 rounded-r-md border border-slate-300 dark:border-slate-500  dark:text-slate-400 grow  dark:bg-slate-800 w-full text-sm
                  placeholder:text-slate-400/75 dark:placeholder:text-slate-500/75 placeholder:text-sm  placeholder:font-base 
                  ${
                    formik.errors.git_addr
                      ? "focus:outline-pink-400 dark:focus:outline-pink-600 outline-1"
                      : "focus:outline-sky-500 dark:focus:outline-sky-600 outline-1"
                  }  `}
                  {...formik.getFieldProps("git_addr")}
                />
              </div>
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
            setOpen={setCreateRepoOpen}
            cancelButtonRef={cancelButtonRef}
          />
        </div>
      </form>
    </DialogComponent>
  );
}
