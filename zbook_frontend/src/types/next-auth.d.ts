import NextAuth from "next-auth";

declare module "next-auth" {
  interface User {
    username: string;
    role: string;
    access_token: string;
    access_token_expires_at: number;
    refresh_token: string;
    refresh_token_expires_at: number;
    error?: string;
    app_id?: string;
  }

  interface Session {
    access_token?: string;
    username: string;
    role: string;
    error?: string;
    app_id?: string;
  }
}

declare module "next-auth/jwt" {
  /** Returned by the `jw` callback and `getToken`, when using JWT sessions */
  interface JWT {
    username: string;
    role: string;
    access_token: string;
    access_token_expires_at: number;
    refresh_token: string;
    refresh_token_expires_at: number;
    exp?: number;
    iat?: number;
    jti?: string;
    error?: string;
    app_id?: string;
  }
}
