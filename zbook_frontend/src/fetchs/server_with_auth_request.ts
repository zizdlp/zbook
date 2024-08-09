interface ListUserLikeRepoRequest {
  username: string;
  page_id: number;
  page_size: number;
  query: string;
}
interface ListUserOwnRepoRequest {
  username: string;
  page_id: number;
  page_size: number;
  query: string;
}

interface ListFollowerRequest {
  username: string;
  page_id: number;
  page_size: number;
  query: string;
}
interface ListFollowingRequest {
  username: string;
  page_id: number;
  page_size: number;
  query: string;
}
interface UpdateUserBlockRequest {
  username: string;
  blocked: boolean;
}
interface GetFollowerCountRequest {
  username: string;
  query: string;
}
interface GetFollowingCountRequest {
  username: string;
  query: string;
}

interface GetListUserOwnRepoCountRequest {
  username: string;
  query: string;
}
interface GetListUserLikeRepoCountRequest {
  username: string;
  query: string;
}

interface GetListUserCountRequest {
  query: string;
}
interface GetListSessionCountRequest {
  query: string;
}
interface GetListCommentReportCountRequest {
  query: string;
}
interface GetListCommentCountRequest {
  query: string;
}
interface GetListRepoCountRequest {
  query: string;
}
interface GetDailyVisitorCountRequest {
  time_zone: string;
  ndays: number;
}
interface GetDailyActiveUserCountRequest {
  time_zone: string;
  ndays: number;
}
interface GetDailyCreateUserCountRequest {
  time_zone: string;
  ndays: number;
}

interface ListUserRequest {
  page_id: number;
  page_size: number;
  query: string;
}
interface ListSessionRequest {
  page_id: number;
  page_size: number;
  query: string;
}
interface ListCommentReportRequest {
  page_id: number;
  page_size: number;
  query: string;
}
interface ListCommentRequest {
  page_id: number;
  page_size: number;
  query: string;
}
interface ListRepoRequest {
  page_id: number;
  page_size: number;
  query: string;
}
interface ListRepoVisibilityRequest {
  page_id: number;
  page_size: number;
  query: string;
  username: string;
  repo_name: string;
}
interface GetRepoVisibilityCountRequest {
  username: string;
  repo_name: string;
}

interface GetUserInfoRequest {
  username: string;
  user_count: boolean;
  user_basic: boolean;
  user_image: boolean;
}
interface GetRepoBasicInfoRequest {
  username: string;
  repo_name: string;
}

interface GetRepoConfigRequest {
  username: string;
  repo_name: string;
}
interface GetMarkdownContentRequest {
  username: string;
  repo_name: string;
  relative_path: string;
}

interface GetMarkdownImageRequest {
  username: string;
  repo_name: string;
  file_path: string;
}
interface CreateRepoRequest {
  repo_name: string;
  repo_description: string;
  git_addr: string;
  git_access_token: string;
  sync_token: string;
  visibility_level: string;
  home_page: string;
  theme_sidebar: string;
  theme_color: string;
}
interface UpdateUserRequest {
  motto: string;
  password: string;
  avatar: string;
}

interface createRepoRelationRequest {
  username: string;
  repo_name: string;
  relation_type: string;
}
interface deleteRepoRelationRequest {
  username: string;
  repo_name: string;
  relation_type: string;
}
interface CreateRepoVisibilityRequest {
  repo_username: string;
  repo_name: string;
  username: string;
}
interface DeleteRepoVisibilityRequest {
  repo_username: string;
  repo_name: string;
  username: string;
}
interface ListNotificationRequest {
  page_id: number;
  page_size: number;
}

interface ListCommentLevelOneRequest {
  markdown_id: number;
  page_id: number;
  page_size: number;
}

interface ListCommentLevelTwoRequest {
  root_id: number;
  page_id: number;
  page_size: number;
}

interface MarkFollowerNotificationReadedRequest {
  noti_id: number;
}
interface MarkSystemNotificationReadedRequest {
  noti_id: number;
}
interface MarkCommentNotificationReadedRequest {
  noti_id: number;
}
interface MarkRepoNotificationReadedRequest {
  noti_id: number;
}
interface UpdateCommentReportStatusRequest {
  report_id: number;
  processed: boolean;
}

interface CreateSystemNotificationRequest {
  username: string;
  title: string;
  contents: string;
  redirect_url: string;
}
interface CreateInvitationRequest {
  email: string;
}
interface DeleteRepoRequest {
  username: string;
  repo_name: string;
}
interface DeleteUserRequest {
  username: string;
}

interface DeleteCommentRequest {
  comment_id: number;
}

interface CreateCommentRequest {
  markdown_id: number;
  parent_id: number;
  comment_content: string;
}
interface CreateCommentReportRequest {
  comment_id: number;
  report_content: string;
}
interface CreateCommentRelationRequest {
  comment_id: number;
  relation_type: string;
}
interface DeleteCommentRelationRequest {
  comment_id: number;
  relation_type: string;
}
interface GetCommentCountInfoRequest {
  comment_id: number;
}
interface QueryMarkdownRequest {
  plain_to_tsquery: string;
  page_id: number;
  page_size: number;
}
interface QueryUserMarkdownRequest {
  username: string;
  plain_to_tsquery: string;
  page_id: number;
  page_size: number;
}
interface QueryRepoMarkdownRequest {
  username: string;
  repo_name: string;
  plain_to_tsquery: string;
  page_id: number;
  page_size: number;
}
interface GetUnReadCountRequest {}
interface ResetUnreadCountRequest {}

interface ManualSyncRepoRequest {
  username: string;
  repo_name: string;
}
interface UpdateRepoInfoRequest {
  username: string;
  old_repo_name: string;
  repo_name: string;
  repo_description: string;
  git_access_token: string;
  visibility_level: string;
  sync_token: string;
  theme_sidebar: string;
  theme_color: string;
}

interface CheckOAuthStatusRequest {}
interface DeleteOAuthLinkRequest {
  oauth_type: string;
}

interface QueryUserRequest {
  page_id: number;
  page_size: number;
  query: string;
}
interface UpdateUserOnBoardingRequest {
  onboarding: boolean;
}
interface CreateFollowRequest {
  username: string;
}
interface GetFollowStatusRequest {
  username: string;
}
interface DeleteFollowRequest {
  username: string;
}
interface GetListCommentNotificationUnreadedCountRequest {}
interface GetListRepoNotificationUnreadedCountRequest {}
interface GetListFollowerNotificationUnreadedCountRequest {}
interface GetListSystemNotificationUnreadedCountRequest {}

interface GetListCommentLevelOneCountRequest {
  markdown_id: number;
}

interface GetListCommentLevelTwoCountRequest {
  root_id: number;
}
interface GetDailyVisitorsRequest {
  ndays: number;
  lang: string;
}

interface GetConfigurationRequest {
  config_name: string;
}
interface UpdateConfigurationRequest {
  config_name: string;
  config_value: boolean;
}
export type {
  GetListCommentLevelOneCountRequest,
  GetListCommentLevelTwoCountRequest,
  GetListCommentNotificationUnreadedCountRequest,
  GetListRepoNotificationUnreadedCountRequest,
  GetListFollowerNotificationUnreadedCountRequest,
  GetListSystemNotificationUnreadedCountRequest,
  createRepoRelationRequest,
  deleteRepoRelationRequest,
  CreateRepoVisibilityRequest,
  DeleteRepoVisibilityRequest,
  ListNotificationRequest,
  ListCommentLevelOneRequest,
  ListCommentLevelTwoRequest,
  MarkFollowerNotificationReadedRequest,
  MarkSystemNotificationReadedRequest,
  MarkCommentNotificationReadedRequest,
  MarkRepoNotificationReadedRequest,
  UpdateCommentReportStatusRequest,
  CreateSystemNotificationRequest,
  DeleteRepoRequest,
  DeleteUserRequest,
  DeleteCommentRequest,
  CreateCommentRequest,
  CreateCommentReportRequest,
  CreateCommentRelationRequest,
  DeleteCommentRelationRequest,
  GetCommentCountInfoRequest,
  QueryUserMarkdownRequest,
  QueryRepoMarkdownRequest,
  GetUnReadCountRequest,
  ResetUnreadCountRequest,
  ManualSyncRepoRequest,
  UpdateRepoInfoRequest,
  CheckOAuthStatusRequest,
  DeleteOAuthLinkRequest,
  QueryUserRequest,
  UpdateUserOnBoardingRequest,
  CreateFollowRequest,
  GetFollowStatusRequest,
  DeleteFollowRequest,
};

export type {
  ListUserLikeRepoRequest,
  ListUserOwnRepoRequest,
  ListFollowerRequest,
  ListFollowingRequest,
  GetFollowerCountRequest,
  GetFollowingCountRequest,
  GetListUserOwnRepoCountRequest,
  GetListUserLikeRepoCountRequest,
  GetListSessionCountRequest,
  GetListCommentReportCountRequest,
  GetListCommentCountRequest,
  GetListRepoCountRequest,
  GetDailyVisitorCountRequest,
  GetDailyActiveUserCountRequest,
  GetDailyCreateUserCountRequest,
  GetListUserCountRequest,
  ListUserRequest,
  ListSessionRequest,
  ListCommentReportRequest,
  ListCommentRequest,
  ListRepoRequest,
  ListRepoVisibilityRequest,
  GetRepoVisibilityCountRequest,
  GetUserInfoRequest,
  GetRepoBasicInfoRequest,
  GetRepoConfigRequest,
  GetMarkdownContentRequest,
  GetMarkdownImageRequest,
  UpdateUserBlockRequest,
  CreateRepoRequest,
  UpdateUserRequest,
  GetDailyVisitorsRequest,
  GetConfigurationRequest,
  UpdateConfigurationRequest,
  QueryMarkdownRequest,
  CreateInvitationRequest,
};
