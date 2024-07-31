"use client";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import { useState } from "react";
import { useRouter } from "@/navigation";
import { HiAtSymbol, HiFingerPrint, HiOutlineUser } from "react-icons/hi";
import { useTranslations } from "next-intl";
import FormRow from "@/components/forms/FormRow";
import InputRow from "@/components/forms/InputRow";
import { emailRegex } from "@/utils/const_value";
import { fetchServerWithoutAuthWrapper } from "@/fetchs/server_without_auth";
import { createUserRequest } from "@/fetchs/server_without_auth_request";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";
import { toastError, toastInfo } from "@/utils/util";
export default function RegisterForm({
  invitation_url,
}: {
  invitation_url: string;
}) {
  const t = useTranslations("RegisterForm");

  const [show, setShow] = useState({ password: false, cpassword: false });
  const router = useRouter();

  function registerValidate(values: { [key: string]: string }) {
    const errors: { [key: string]: string } = {};
    const usernameRegex = /^[a-z0-9_]+$/; // Regular expression to match lowercase letters, digits, and underscores.
    if (!values.username) {
      errors.username = t("Required");
    } else if (values.username.includes(" ")) {
      errors.username = t("InvalidUsername");
    } else if (!usernameRegex.test(values.username)) {
      errors.username = t("OnlyLower");
    } else if (values.username.length < 4 || values.username.length > 20) {
      errors.username = t("UsernameCount");
    }

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
  const formik = useFormik({
    initialValues: {
      username: "",
      email: "",
      password: "",
      cpassword: "",
      invitation_url: "",
    },
    validate: registerValidate,
    onSubmit: onSubmit,
  });
  async function onSubmit(values: createUserRequest) {
    const id = toast(t("Registering"), {
      type: "info",
      isLoading: true,
    });
    try {
      values.invitation_url = invitation_url;
      const data = await fetchServerWithoutAuthWrapper({
        endpoint: FetchServerWithoutAuthWrapperEndPoint.CREATE_USER,
        xforward: "",
        agent: "",
        values: values,
      });
      if (data.error) {
        toastError(id, data.status + " " + data.message);
      } else {
        toastInfo(id, t("RegistrationSuccessful"));
        router.push(`/auth/login`);
      }
    } catch (error) {}
  }

  return (
    <form className="flex flex-col my-6" onSubmit={formik.handleSubmit}>
      <FormRow
        error={formik.errors.username}
        show_error={formik.touched.username}
      >
        <InputRow
          show={true}
          name={"username"}
          formik={formik}
          place_holder={t("Username")}
          error={formik.errors.username}
        />
        <div className="absolute inset-y-0 right-0 pr-2 flex items-center">
          <HiOutlineUser size={25} />
        </div>
      </FormRow>

      <FormRow error={formik.errors.email} show_error={formik.touched.email}>
        <InputRow
          show={true}
          name={"email"}
          formik={formik}
          place_holder={t("Email")}
          error={formik.errors.email}
        />
        <div className="absolute inset-y-0 right-0 pr-2 flex items-center">
          <HiAtSymbol size={25} />
        </div>
      </FormRow>

      <FormRow
        error={formik.errors.password}
        show_error={formik.touched.password}
      >
        <InputRow
          show={show.password}
          name={"password"}
          formik={formik}
          place_holder={t("Password")}
          error={formik.errors.password}
        />
        <div className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer">
          <HiFingerPrint
            size={25}
            onClick={() => setShow({ ...show, password: !show.password })}
          />
        </div>
      </FormRow>

      <FormRow
        error={formik.errors.cpassword}
        show_error={formik.touched.cpassword}
      >
        <InputRow
          show={show.cpassword}
          name={"cpassword"}
          formik={formik}
          place_holder={t("CPassword")}
          error={formik.errors.cpassword}
        />
        <div className="absolute inset-y-0 right-0 pr-2 flex items-center cursor-pointer">
          <HiFingerPrint
            size={25}
            onClick={() => setShow({ ...show, cpassword: !show.cpassword })}
          />
        </div>
      </FormRow>

      <button
        className="text-white rounded-md bg-sky-500 hover:bg-sky-700 text-sm leading-5 font-semibold p-3"
        type="submit"
      >
        {t("Register")}
      </button>
    </form>
  );
}
