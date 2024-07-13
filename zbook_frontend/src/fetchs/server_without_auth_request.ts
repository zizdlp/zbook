interface createUserRequest {
  username: string;
  password: string;
  email: string;
}
interface SendEmailToResetPasswordRequest {
  email: string;
}
interface SendEmailToVerifyEmailRequest {
  email: string;
}
interface ResetPasswordRequest {
  verification_id: string;
  password: string;
  email: string;
}
interface createOAuthLinkRequest {
  oauth_type: string;
  app_id: string;
}
interface refreshTokenRequest {
  refresh_token: string;
}
interface getRepoIDRequest {
  username: string;
  repo_name: string;
}
interface LoginUserRequest {
  email: string;
  password: string;
}
interface LoginByOAuthRequest {
  oauth_type: string;
  app_id: string;
  access_token: string;
}
export type {
  createUserRequest,
  SendEmailToResetPasswordRequest,
  SendEmailToVerifyEmailRequest,
  ResetPasswordRequest,
  createOAuthLinkRequest,
  refreshTokenRequest,
  getRepoIDRequest,
  LoginUserRequest,
  LoginByOAuthRequest,
};
