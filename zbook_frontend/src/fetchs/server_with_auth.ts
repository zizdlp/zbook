"use server";
import { auth } from "@/auth";
import { unstable_noStore as noStore } from "next/cache";
import { backend_url, api_version } from "@/utils/env_variable";
import { joinUrlParts } from "@/utils/util";
import { FetchError, RequestOptions } from "./util";
import { redirect } from "@/navigation";
import {
  GetDailyActiveUserCountRequest,
  GetDailyCreateUserCountRequest,
  GetDailyVisitorCountRequest,
  GetFollowingCountRequest,
  GetFollowerCountRequest,
  GetListSessionCountRequest,
  GetListCommentCountRequest,
  GetListCommentReportCountRequest,
  GetListRepoCountRequest,
  GetListUserCountRequest,
  GetListUserLikeRepoCountRequest,
  GetListUserOwnRepoCountRequest,
  GetMarkdownContentRequest,
  GetMarkdownImageRequest,
  GetRepoConfigRequest,
  GetRepoVisibilityCountRequest,
  GetUserInfoRequest,
  ListSessionRequest,
  ListCommentReportRequest,
  ListCommentRequest,
  ListFollowingRequest,
  ListFollowerRequest,
  ListRepoRequest,
  ListRepoVisibilityRequest,
  ListUserLikeRepoRequest,
  ListUserOwnRepoRequest,
  ListUserRequest,
  GetRepoBasicInfoRequest,
  UpdateUserBlockRequest,
  CreateRepoRequest,
  UpdateUserRequest,
  GetListFollowerNotificationUnreadedCountRequest,
  GetListRepoNotificationUnreadedCountRequest,
  GetListCommentNotificationUnreadedCountRequest,
  GetListSystemNotificationUnreadedCountRequest,
  GetListCommentLevelOneCountRequest,
  GetListCommentLevelTwoCountRequest,
  GetDailyVisitorsRequest,
  GetConfigurationRequest,
  UpdateConfigurationRequest,
  QueryMarkdownRequest,
  CreateInvitationRequest,
} from "./server_with_auth_request";
import {
  CreateRepoVisibilityRequest,
  CreateSystemNotificationRequest,
  DeleteRepoVisibilityRequest,
  ListCommentLevelOneRequest,
  ListCommentLevelTwoRequest,
  ListNotificationRequest,
  MarkCommentNotificationReadedRequest,
  MarkFollowerNotificationReadedRequest,
  DeleteRepoRequest,
  MarkRepoNotificationReadedRequest,
  MarkSystemNotificationReadedRequest,
  UpdateCommentReportStatusRequest,
  DeleteUserRequest,
  createRepoRelationRequest,
  deleteRepoRelationRequest,
  DeleteCommentRequest,
  CreateCommentRequest,
  CreateCommentReportRequest,
  CreateCommentRelationRequest,
  DeleteCommentRelationRequest,
  GetCommentCountInfoRequest,
  QueryUserMarkdownRequest,
  QueryRepoMarkdownRequest,
  ResetUnreadCountRequest,
  GetUnReadCountRequest,
  ManualSyncRepoRequest,
  UpdateRepoInfoRequest,
  DeleteOAuthLinkRequest,
  CheckOAuthStatusRequest,
  QueryUserRequest,
  CreateFollowRequest,
  DeleteFollowRequest,
  GetFollowStatusRequest,
  UpdateUserOnBoardingRequest,
} from "./server_with_auth_request";
import { revalidatePath } from "next/cache";
import { FetchServerWithAuthWrapperEndPoint } from "./server_with_auth_util";
import { revalidateTag } from "next/cache";

export async function refreshPage(url: string, layout: boolean, page: boolean) {
  if (layout && page) {
    revalidatePath(url);
  } else if (layout) {
    revalidatePath(url, "layout");
  } else if (page) {
    revalidatePath(url, "page");
  }
}
export async function redirectPage(url: string) {
  redirect(`${url}`); // Navigate to the new post page
}
//TODO nextjs bug: revalidateTag refresh all tag in same path
export async function refreshTag(tag: string) {
  revalidateTag(tag);
}

export async function fetchServer(
  url: string,
  options: RequestOptions,
  xforward: string,
  agent: string,
  useAuth: boolean,
  tags: string[],
  timeout = 15000 //milliseconds,15s
) {
  noStore();
  const controller = new AbortController();
  const { signal } = controller;

  const timeoutId = setTimeout(() => {
    controller.abort();
  }, timeout);

  try {
    let access_token = "";
    if (useAuth) {
      const session = await auth();
      access_token = session?.access_token || "";
    }
    const authHeaders = {
      "Content-Type": "application/json",
      "X-Forwarded-For": xforward,
      "User-Agent": agent,
      Authorization: `Bearer ${access_token}`,
    };
    const response = await fetch(url, {
      ...options,
      next: { tags: tags },
      headers: { ...options.headers, ...authHeaders },
      signal,
    });

    clearTimeout(timeoutId);
    if (!response.ok) {
      const data = await response.json();
      throw new FetchError(data.message, response.status);
    }
    return await response.json();
  } catch (unknownError) {
    clearTimeout(timeoutId);
    // Type guard for AbortError
    const isAbortError = (error: unknown): error is DOMException => {
      return error instanceof DOMException && error.name === "AbortError";
    };

    // Type guard for Connect Timeout Error
    const isConnectTimeoutError = (
      error: unknown
    ): error is { cause: { code: string } } => {
      return (
        typeof error === "object" &&
        error !== null &&
        "cause" in error &&
        (error as any).cause.code === "UND_ERR_CONNECT_TIMEOUT"
      );
    };

    // Type guard for ECONNREFUSED Error
    const isECONNREFUSEDError = (
      error: unknown
    ): error is { cause: { code: string } } => {
      return (
        typeof error === "object" &&
        error !== null &&
        "cause" in error &&
        (error as any).cause.code === "ECONNREFUSED"
      );
    };

    // Type guard for other FetchErrors
    const isFetchError = (error: unknown): error is FetchError => {
      return error instanceof FetchError;
    };

    if (isAbortError(unknownError)) {
      return {
        error: true,
        status: 408,
        message: "Request Timeout",
      };
    } else if (isConnectTimeoutError(unknownError)) {
      return {
        error: true,
        status: 408,
        message: "Connect Timeout",
      };
    } else if (isECONNREFUSEDError(unknownError)) {
      return {
        error: true,
        status: 503,
        message: "Connection Refused",
      };
    } else if (isFetchError(unknownError)) {
      return {
        error: true,
        status: unknownError.status,
        message: unknownError.message,
      };
    }

    // Handle other unknown errors
    return {
      error: true,
      status: 500,
      message: "An unexpected error occurred",
    };
  }
}

async function fetchServerWithAuth(
  url: string,
  options: RequestOptions,
  xforward: string,
  agent: string,
  tags: string[],
  timeout = 15000
) {
  return await fetchServer(url, options, xforward, agent, true, tags, timeout);
}
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.UpdateConfiguration;
  values: UpdateConfigurationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GetConfiguration;
  values: GetConfigurationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GetDailyVisitors;
  values: GetDailyVisitorsRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_LEVEL_ONE_COUNT;
  values: GetListCommentLevelOneCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_LEVEL_TWO_COUNT;
  values: GetListCommentLevelTwoCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_FOLLOWER_NOTIFICATION_UNREADED_COUNT;
  values: GetListFollowerNotificationUnreadedCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_REPO_NOTIFICATION_UNREADED_COUNT;
  values: GetListRepoNotificationUnreadedCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_NOTIFICATION_UNREADED_COUNT;
  values: GetListCommentNotificationUnreadedCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_SYSTEM_NOTIFICATION_UNREADED_COUNT;
  values: GetListSystemNotificationUnreadedCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_USER;
  values: UpdateUserRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_REPO;
  values: CreateRepoRequest;
  xforward: string;
  agent: string;
  tags: string[];
  timeout: number;
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_USER_BLOCK;
  values: UpdateUserBlockRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_MARKDOWN_IMAGE;
  values: GetMarkdownImageRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_MARKDOWN_CONTENT;
  values: GetMarkdownContentRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_REPO_CONFIG;
  values: GetRepoConfigRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_REPO_BASIC_INFO;
  values: GetRepoBasicInfoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_USER_LIKE_REPO;
  values: ListUserLikeRepoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_USER_OWN_REPO;
  values: ListUserOwnRepoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_FOLLOWING;
  values: ListFollowingRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_FOLLOWER;
  values: ListFollowerRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_FOLLOER_COUNT;
  values: GetFollowerCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_FOLLOWING_COUNT;
  values: GetFollowingCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_USER_OWN_REPO_COUNT;
  values: GetListUserOwnRepoCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_USER_LIKE_REPO_COUNT;
  values: GetListUserLikeRepoCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_USER_COUNT;
  values: GetListUserCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_SESSION_COUNT;
  values: GetListSessionCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_REPORT_COUNT;
  values: GetListCommentReportCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_COUNT;
  values: GetListCommentCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_REPO_COUNT;
  values: GetListRepoCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DAILY_VISITOR_COUNT;
  values: GetDailyVisitorCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DAILY_ACTIVE_USER_COUNT;
  values: GetDailyActiveUserCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DAILY_CREATE_USER_COUNT;
  values: GetDailyCreateUserCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_USER;
  values: ListUserRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_SESSION;
  values: ListSessionRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_REPORT;
  values: ListCommentReportRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT;
  values: ListCommentRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_REPO;
  values: ListRepoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_REPO_VISIBILITY;
  values: ListRepoVisibilityRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_REPO_VISIBILITY_COUNT;
  values: GetRepoVisibilityCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_USER_INFO;
  values: GetUserInfoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.QUERY_USER;
  values: QueryUserRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_FOLLOW;
  values: CreateFollowRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_FOLLOW;
  values: DeleteFollowRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_FOLLOWE_STATUS;
  values: GetFollowStatusRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_USER_ONBOARDING;
  values: UpdateUserOnBoardingRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_OAUTH_LINK;
  values: DeleteOAuthLinkRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CHECK_OAUTH_STATUS;
  values: CheckOAuthStatusRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
  timeout,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.MANUAL_SYNC_REPO;
  values: ManualSyncRepoRequest;
  xforward: string;
  agent: string;
  tags: string[];
  timeout: number;
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_REPO_INFO;
  values: UpdateRepoInfoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.RESET_UNREAD_COUNT;
  values: ResetUnreadCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_UNREAD_COUNT;
  values: GetUnReadCountRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.QUERY_REPO_MARKDOWN;
  values: QueryRepoMarkdownRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.QUERY_MARKDOWN;
  values: QueryMarkdownRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.QUERY_USER_MARKDOWN;
  values: QueryUserMarkdownRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.GET_COMMENT_COUNT_INFO;
  values: GetCommentCountInfoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_COMMENT;
  values: CreateCommentRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_COMMENT_REPORT;
  values: CreateCommentReportRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_COMMENT_RELATION;
  values: CreateCommentRelationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_COMMENT_RELATION;
  values: DeleteCommentRelationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_REPO_RELATION;
  values: createRepoRelationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_REPO_RELATION;
  values: deleteRepoRelationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_REPO_VISIBILITY;
  values: CreateRepoVisibilityRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_REPO_VISIBILITY;
  values: DeleteRepoVisibilityRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_NOTIFICATION;
  values: ListNotificationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_FOLLOWER_NOTIFICATION;
  values: ListNotificationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_REPO_NOTIFICATION;
  values: ListNotificationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_SYSTEM_NOTIFICATION;
  values: ListNotificationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_LEVEL_ONE;
  values: ListCommentLevelOneRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_LEVEL_TWO;
  values: ListCommentLevelTwoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.MARK_COMMENT_NOTI_READED;
  values: MarkCommentNotificationReadedRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.MARK_FOLLOWER_NOTI_READED;
  values: MarkFollowerNotificationReadedRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.MARK_REPO_NOTI_READED;
  values: MarkRepoNotificationReadedRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.MARK_SYSTEM_NOTI_READED;
  values: MarkSystemNotificationReadedRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_COMMENT_REPORT_STATUS;
  values: UpdateCommentReportStatusRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CREATE_SYSTEM_NOTIFICATION;
  values: CreateSystemNotificationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.CreateInvitation;
  values: CreateInvitationRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_REPO;
  values: DeleteRepoRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_USER;
  values: DeleteUserRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;

export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint.DELETE_COMMENT;
  values: DeleteCommentRequest;
  xforward: string;
  agent: string;
  tags: string[];
}): Promise<any>;
export async function fetchServerWithAuthWrapper({
  endpoint,
  values,
  xforward,
  agent,
  tags,
  timeout,
}: {
  endpoint: FetchServerWithAuthWrapperEndPoint;
  values: any;
  xforward: string;
  agent: string;
  tags: string[];
  timeout?: number;
}) {
  const url = joinUrlParts(backend_url, api_version, endpoint);
  const options: RequestOptions = {
    method: "POST",
    body: JSON.stringify(values), // 使用对象解构简化代码
  };
  return fetchServerWithAuth(url, options, xforward, agent, tags, timeout);
}
