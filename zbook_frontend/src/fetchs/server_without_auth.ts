"use server";
import { server_api_base_url, server_api_version } from "@/utils/env_variable";
import { joinUrlParts } from "@/utils/util";
import { RequestOptions } from "@/fetchs/util";
import { signIn } from "../auth";
import {
  LoginByOAuthRequest,
  LoginUserRequest,
  ResetPasswordRequest,
  SendEmailToResetPasswordRequest,
  SendEmailToVerifyEmailRequest,
  createOAuthLinkRequest,
  createUserRequest,
  refreshTokenRequest,
  LogVisitorRequest,
} from "./server_without_auth_request";
import { FetchServerWithoutAuthWrapperEndPoint } from "./server_without_auth_util";
import { fetchServer } from "./server_with_auth";
async function fetchServerWithoutAuth(
  url: string,
  options: RequestOptions,
  xforward: string,
  agent: string,
  timeout = 15000
) {
  return await fetchServer(url, options, xforward, agent, false, [], timeout);
}
export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.LOGIN_BY_OAUTH;
  values: LoginByOAuthRequest;
  xforward: string;
  agent: string;
}): Promise<any>;
export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.LOGIN_USER;
  values: LoginUserRequest;
  xforward: string;
  agent: string;
}): Promise<any>;

export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.REFRESH_TOKEN;
  values: refreshTokenRequest;
  xforward: string;
  agent: string;
}): Promise<any>;
export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.SEND_EMAIL_TO_RESET_PASSWORD;
  values: SendEmailToResetPasswordRequest;
  xforward: string;
  agent: string;
}): Promise<any>;
export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.SEND_EMAIL_TO_VERIFY_EMAIL;
  values: SendEmailToVerifyEmailRequest;
  xforward: string;
  agent: string;
}): Promise<any>;
export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.RESET_PASSWORD;
  values: ResetPasswordRequest;
  xforward: string;
  agent: string;
}): Promise<any>;

export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.CREATE_USER;
  values: createUserRequest;
  xforward: string;
  agent: string;
}): Promise<any>;

export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint.LOG_VISITOR;
  values: LogVisitorRequest;
  xforward: string;
  agent: string;
}): Promise<any>;

export async function fetchServerWithoutAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
}: {
  endpoint: FetchServerWithoutAuthWrapperEndPoint;
  values: any;
  xforward: string;
  agent: string;
}) {
  const url = joinUrlParts(server_api_base_url, server_api_version, endpoint);
  const options: RequestOptions = {
    method: "POST",
    body: JSON.stringify(values), // 使用对象解构简化代码
  };
  return fetchServerWithoutAuth(url, options, xforward, agent);
}

export async function createOAuthLink(
  values: createOAuthLinkRequest,
  access_token: string
) {
  let options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${access_token}`,
    },
    body: JSON.stringify(values),
  };

  const response = await fetch(
    `${server_api_base_url}v1/create_oauth_link`,
    options
  );
  const data = await response.json();
  return data;
}

export async function getVerify(verification_url: string) {
  const url = `${
    process.env.BACKEND_URL
  }v1/verify_email?verification_url=${decodeURIComponent(verification_url)}`;
  const decodedUrl = decodeURIComponent(url);
  const res = await fetch(decodedUrl);
  return await res.json();
}

interface CustomError extends Error {
  cause?: {
    err?: {
      status?: number;
      message?: string;
    };
  };
}

const hasCause = (error: any): error is CustomError => {
  return error && typeof error === "object" && "cause" in error;
};
export async function serverSignIn({
  email,
  password,
}: {
  email: string;
  password: string;
}) {
  try {
    const status = await signIn("credentials", {
      redirect: false,
      email: email,
      password: password,
    });
    return status;
  } catch (error) {
    if (hasCause(error)) {
      let ee = error.cause?.err;
      return {
        error: true,
        status: ee?.status ?? 500,
        message: ee?.message ?? "Unknown error",
      };
    } else {
      return {
        error: true,
        status: 500,
        message: "Unknown error",
      };
    }
  }
}

export async function getUserAvatarServer({ username }: { username: string }) {
  try {
    const backend_url = process.env.BACKEND_URL;
    const url = `${backend_url}v1/get_user_avatar?username=${username}`;
    const response = await fetch(url, { next: { revalidate: 3600 } });
    const data = await response.json();
    return data;
  } catch (error) {
    throw new Error("Failed to fetch user avatar");
  }
}
