"use client";
import { Link } from "@/navigation";
import Image from "next/image";
import { useFormik } from "formik";
import { toast } from "react-toastify";
import { HiOutlineUser } from "react-icons/hi";
import { useTranslations } from "next-intl";
import FormRow from "@/components/forms/FormRow";
import InputRow from "@/components/forms/InputRow";
import { emailRegex } from "@/utils/const_value";
import { fetchServerWithoutAuthWrapper } from "@/fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "@/fetchs/server_without_auth_util";
import { FetchError } from "@/fetchs/util";
export default function SendForgetForm() {
  const t = useTranslations("ForgetForm");

  function forget_validate(values: { [key: string]: string }) {
    const errors: { [key: string]: string } = {};
    // validation for email
    if (!values.email) {
      errors.email = t("Required");
    } else if (!emailRegex.test(values.email)) {
      errors.email = t("InvalidEmailFormat");
    }
    return errors;
  }

  const formikSendEmail = useFormik({
    initialValues: {
      email: "",
    },
    validate: forget_validate,
    onSubmit: onSubmitSendEmail,
  });
  async function onSubmitSendEmail(values: { email: string }) {
    const id = toast(`Send Verify Code Email...`, {
      type: "info",
      isLoading: false,
    });
    try {
      const data = await fetchServerWithoutAuthWrapper({
        endpoint:
          FetchServerWithoutAuthWrapperEndPoint.SEND_EMAIL_TO_RESET_PASSWORD,
        xforward: "",
        agent: "",
        values: values,
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
      toast.update(id, {
        render: `Send Email Success,please write token blow which received`,
        type: "success",
        isLoading: false,
        autoClose: 500,
      });
    } catch (error) {
      toast.update(id, {
        render: `Send Email failed:` + error,
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
            onSubmit={formikSendEmail.handleSubmit}
          >
            <FormRow
              error={formikSendEmail.errors.email}
              show_error={formikSendEmail.touched.email}
            >
              <InputRow
                show={true}
                name={"email"}
                formik={formikSendEmail}
                place_holder={t("Email")}
                error={formikSendEmail.errors.email}
              />
              <div className="absolute inset-y-0 right-0 pr-2 flex items-center">
                <HiOutlineUser size={25} />
              </div>
            </FormRow>
            <button
              className="text-white rounded-md bg-sky-500 hover:bg-sky-700 text-sm leading-5 font-semibold p-3"
              type="submit"
            >
              {t("SendEmail")}
            </button>
          </form>

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
