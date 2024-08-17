"use client";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import { HiFingerPrint, HiOutlineUser } from "react-icons/hi";
import { useState } from "react";
import { FaGithub } from "react-icons/fa";
import { FaGoogle } from "react-icons/fa";
import { useTranslations } from "next-intl";
import FormRow from "@/components/forms/FormRow";
import InputRow from "@/components/forms/InputRow";
import { emailRegex } from "@/utils/const_value";
import {
  fetchServerWithoutAuthWrapper,
  serverSignIn,
} from "@/fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";
import { FetchError } from "@/fetchs/util";
import { toastError, toastInfo } from "@/utils/util";
import { redirectPage } from "@/fetchs/server_with_auth";
import { Link } from "@/navigation";
export default function LoginForm({
  xforward,
  agent,
}: {
  xforward: string;
  agent: string;
}) {
  const t = useTranslations("LoginForm");
  const [showPassword, setShowPassword] = useState(false);
  const [sendVerifyEmail, setSendVerifyEmail] = useState(false);
  const [sendEmail, setSendEmail] = useState("");

  function login_validate(values: { [key: string]: string }) {
    const errors: { [key: string]: string } = {};
    // validation for email
    if (!values.email) {
      errors.email = t("Required");
    } else if (!emailRegex.test(values.email)) {
      errors.email = t("InvalidEmailFormat");
    }

    // validation for password
    if (!values.password) {
      errors.password = t("Required");
    } else if (values.password.length < 8 || values.password.length > 20) {
      errors.password = t("CharacterCount");
    } else if (values.password.includes(" ")) {
      errors.password = t("InvaliPassword");
    }

    return errors;
  }

  const formik = useFormik({
    initialValues: {
      email: "",
      password: "",
    },
    validate: login_validate,
    onSubmit: onSubmit,
  });
  async function sendVerify({ email }: { email: string }) {
    const id = toast(t("ResendEmail"), {
      type: "info",
      isLoading: false,
    });
    try {
      const data = await fetchServerWithoutAuthWrapper({
        endpoint:
          FetchServerWithoutAuthWrapperEndPoint.SEND_EMAIL_TO_VERIFY_EMAIL,
        xforward: xforward,
        agent: agent,
        values: {
          email: email,
        },
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      toastInfo(id, t("ResendSuccessful"));
    } catch (error) {
      toastError(id, t("ResendFailed"));
    }
  }
  async function onSubmit(values: { email: string; password: string }) {
    const id = toast(t("LoggingIn"), { type: "info", isLoading: true });

    const status = await serverSignIn({
      email: values.email,
      password: values.password,
    });
    if (status.error) {
      if (status.status == 404) {
        toastError(id, t("LoginFailedExist"));
      } else if (status.status == 401) {
        toastError(id, t("LoginFailedIncorrect"));
      } else if (
        status.status == 403 &&
        status.message == "email not verified for this account"
      ) {
        setSendEmail(values.email);
        setSendVerifyEmail(true);
        toastError(id, t("LoginFailedEmailVerify"));
      } else {
        toastError(id, status.status + ":" + status.message);
      }
    } else {
      toastInfo(id, t("LoginSuccessful"));
      redirectPage(`/workspace/`);
    }
  }
  return (
    <div className="">
      <form className="flex flex-col my-6" onSubmit={formik.handleSubmit}>
        <FormRow error={formik.errors.email} show_error={formik.touched.email}>
          <InputRow
            show={true}
            name={"email"}
            formik={formik}
            place_holder={t("Email")}
            error={formik.errors.email}
          />
          <div className="absolute inset-y-0 right-0 pr-2 flex items-center">
            <HiOutlineUser size={25} />
          </div>
        </FormRow>

        <FormRow
          error={formik.errors.password}
          show_error={formik.touched.password}
        >
          <InputRow
            show={showPassword}
            name={"password"}
            formik={formik}
            place_holder={t("Password")}
            error={formik.errors.password}
          />

          <div className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer">
            <HiFingerPrint
              size={25}
              onClick={() => setShowPassword(!showPassword)}
            />
          </div>
        </FormRow>

        <button
          className="text-white rounded-md bg-sky-500 hover:bg-sky-700 text-sm leading-5 font-semibold p-3"
          type="submit"
        >
          {t("Login")}
        </button>
      </form>

      {sendVerifyEmail && (
        <div
          className="text-xs cursor-pointer hover:text-sky-400"
          onClick={() => {
            sendVerify({ email: sendEmail });
            setSendVerifyEmail(false);
          }}
        >
          {t("ReSendEmail")}
        </div>
      )}
      <div className="flex items-center justify-center py-6 text-xs">
        <div className="w-full border-b-[0.01rem] border-slate-500 dark:border-slate-300"></div>
        <span className="mx-4">{t("LoginType")}</span>
        <div className="w-full border-b-[0.01rem] border-slate-500 dark:border-slate-300"></div>
      </div>
      <div className="flex justify-around items-center px-8 ">
        <Link
          href={`/oauth/oauth?oauth_type=github`}
          className="flex flex-col items-center p-2"
        >
          <FaGithub className="h-7 w-7 p-1" />
          <p className="text-xs">{t("Github")}</p>
        </Link>
        <Link
          href={`/oauth/oauth?oauth_type=google`}
          className="flex flex-col items-center p-2"
        >
          <FaGoogle className="h-7 w-7 p-1" />
          <p className="text-xs">{t("Google")}</p>
        </Link>
      </div>
    </div>
  );
}
