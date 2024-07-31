"use client";
import { Link } from "@/navigation";
import Image from "next/image";
import { useState } from "react";
import { useFormik } from "formik";

import { useRouter } from "@/navigation";
import { toast } from "react-toastify";
import { HiFingerPrint, HiOutlineUser } from "react-icons/hi";
import { useTranslations } from "next-intl";
import FormRow from "@/components/forms/FormRow";
import InputRow from "@/components/forms/InputRow";
import { emailRegex } from "@/utils/const_value";
import { fetchServerWithoutAuthWrapper } from "@/fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";
import { FetchError } from "@/fetchs/util";
export default function ResetPasswordForm({
  verification_url,
}: {
  verification_url: string;
}) {
  const t = useTranslations("ForgetForm");

  const [show, setShow] = useState({ password: false, cpassword: false });
  const router = useRouter();

  function forget_check_validate(values: { [key: string]: string }) {
    const errors: { [key: string]: string } = {};
    // validation for email
    if (!values.email) {
      errors.email = t("Required");
    } else if (!emailRegex.test(values.email)) {
      errors.email = t("InvalidEmailFormat");
    }
    if (!values.password) {
      errors.password = t("Required");
    } else if (values.password.length < 8 || values.password.length > 20) {
      errors.password = t("CharacterCount");
    } else if (values.password.includes(" ")) {
      errors.password = t("InvaliPassword");
    }

    // validate confirm password
    if (!values.cpassword) {
      errors.cpassword = t("Required");
    } else if (values.cpassword.length < 8 || values.cpassword.length > 20) {
      errors.cpassword = t("CharacterCount");
    } else if (values.cpassword.includes(" ")) {
      errors.cpassword = t("InvaliPassword");
    } else if (values.password !== values.cpassword) {
      errors.cpassword = t("PasswordNotMatch");
    }
    return errors;
  }

  const formikResetPassword = useFormik({
    initialValues: {
      email: "",
      password: "",
      verification_url: "",
      cpassword: "",
    },
    validate: forget_check_validate,
    onSubmit: onSubmitResetPassword,
  });
  async function onSubmitResetPassword(values: {
    email: string;
    verification_url: string;
    password: string;
  }) {
    values.verification_url = verification_url;
    const id = toast(t("ResettingPassword"), {
      type: "info",
      isLoading: false,
    });
    try {
      const data = await fetchServerWithoutAuthWrapper({
        endpoint: FetchServerWithoutAuthWrapperEndPoint.RESET_PASSWORD,
        values: values,
        xforward: "",
        agent: "",
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      toast.update(id, {
        render: t("PasswordResetSuccessful"),
        type: "success",
        isLoading: false,
        autoClose: 500,
      });

      router.push(`/auth/login`);
    } catch (error) {
      toast.update(id, {
        render: t("PasswordResetFailed"),
        type: "error",
        isLoading: false,
        autoClose: 1500,
      });
    }
  }

  return (
    <div className="overflow-hidden">
      <div className="text-slate-700 dark:text-slate-200 dark:backdrop-blur mx-auto flex items-center max-w-4xl md:p-5">
        <div className="md:w-1/2 min-w-80 md:px-12 mx-auto">
          <Image
            src="/logo_256.png"
            alt="LOGO"
            width={256}
            height={256}
            className="mx-auto h-16 w-auto rounded-lg"
          />

          <div className="flex justify-center items-center flex-shrink-0">
            <h1 className="font-bold text-2xl">
              <span className="bg-gradient-to-r from-teal-500 to-blue-500 inline-block text-transparent bg-clip-text">
                {t("AppName")}
              </span>
            </h1>
          </div>

          <form
            className="flex flex-col my-6"
            onSubmit={formikResetPassword.handleSubmit}
          >
            <FormRow
              error={formikResetPassword.errors.email}
              show_error={formikResetPassword.touched.email}
            >
              <InputRow
                show={true}
                name={"email"}
                formik={formikResetPassword}
                place_holder={t("Email")}
                error={formikResetPassword.errors.email}
              />
              <div className="absolute inset-y-0 right-0 pr-2 flex items-center">
                <HiOutlineUser size={25} />
              </div>
            </FormRow>

            <FormRow
              error={formikResetPassword.errors.password}
              show_error={formikResetPassword.touched.password}
            >
              <InputRow
                show={show.password}
                name={"password"}
                formik={formikResetPassword}
                place_holder={t("NewPassword")}
                error={formikResetPassword.errors.password}
              />
              <div className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer">
                <HiFingerPrint
                  size={25}
                  onClick={() => setShow({ ...show, password: !show.password })}
                />
              </div>
            </FormRow>
            <FormRow
              error={formikResetPassword.errors.cpassword}
              show_error={formikResetPassword.touched.cpassword}
            >
              <InputRow
                show={show.cpassword}
                name={"cpassword"}
                formik={formikResetPassword}
                place_holder={t("CPassword")}
                error={formikResetPassword.errors.cpassword}
              />
              <div className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer">
                <HiFingerPrint
                  size={25}
                  onClick={() =>
                    setShow({ ...show, cpassword: !show.cpassword })
                  }
                />
              </div>
            </FormRow>

            <button
              className="text-white rounded-md bg-sky-500 hover:bg-sky-700 text-sm leading-5 font-semibold p-3"
              type="submit"
            >
              {t("Reset")}
            </button>
          </form>

          <div className="mt-6 text-xs border-b-[0.01rem] border-slate-500 dark:border-slate-300 py-4">
            <Link href="/auth/forget">
              <p className="font-medium  hover:text-sky-500">
                {t("NotReceiveEmail")}
              </p>
            </Link>
          </div>

          <div className="mt-3 text-xs flex justify-between items-center ">
            <p>{t("NeedAnAccount")}</p>
            <Link href="/auth/register">
              <button className="py-2 px-5 text-white rounded-md bg-sky-500 hover:bg-sky-700 text-sm leading-5 font-semibold">
                {t("Register")}
              </button>
            </Link>
          </div>
        </div>
        <div className="md:block hidden w-1/2">
          <Image src="/login.png" alt="LOGO" width={500} height={500} />
        </div>
      </div>
    </div>
  );
}
