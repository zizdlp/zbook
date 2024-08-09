interface createUserRequest {
  username: string;
  password: string;
  email: string;
  invitation_url: string;
}
interface SendEmailToResetPasswordRequest {
  email: string;
}
interface SendEmailToVerifyEmailRequest {
  email: string;
}
interface ResetPasswordRequest {
  verification_url: string;
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
interface LoginUserRequest {
  email: string;
  password: string;
}
interface LoginByOAuthRequest {
  oauth_type: string;
  app_id: string;
  access_token: string;
}
interface LogVisitorRequest {}
export type {
  createUserRequest,
  SendEmailToResetPasswordRequest,
  SendEmailToVerifyEmailRequest,
  ResetPasswordRequest,
  createOAuthLinkRequest,
  refreshTokenRequest,
  LoginUserRequest,
  LoginByOAuthRequest,
  LogVisitorRequest,
};
