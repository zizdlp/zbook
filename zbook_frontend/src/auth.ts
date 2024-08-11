import NextAuth, { Session, User } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import GitHubProvider from "next-auth/providers/github";
import GoogleProvider from "next-auth/providers/google";
import { FetchError } from "@/fetchs/util";
import { JWT } from "next-auth/jwt";
import { fetchServerWithoutAuthWrapper } from "./fetchs/server_without_auth";
import { FetchServerWithoutAuthWrapperEndPoint } from "./fetchs/server_without_auth_util";
import { logger } from "./utils/logger";
export const { handlers, signIn, signOut, auth } = NextAuth({
  providers: [
    CredentialsProvider({
      credentials: {
        email: {},
        password: {},
      },
      name: "Credentials",
      async authorize(credentials, request: Request) {
        const headers = request.headers;
        if (credentials != undefined) {
          try {
            const data = await fetchServerWithoutAuthWrapper({
              endpoint: FetchServerWithoutAuthWrapperEndPoint.LOGIN_USER,
              xforward: headers.get("x-forwarded-for") ?? "",
              agent: headers.get("User-Agent") ?? "",
              values: {
                email: (credentials.email as string) ?? "",
                password: (credentials.password as string) ?? "",
              },
            });
            if (data.error) {
              throw new FetchError(data.message, data.status);
            }
            const user: User = {
              id: data.username,
              username: data.username,
              role: data.role,
              access_token: data.access_token,
              access_token_expires_at: new Date(
                data.access_token_expires_at
              ).getTime(),
              refresh_token: data.refresh_token,
              refresh_token_expires_at: new Date(
                data.refresh_token_expires_at
              ).getTime(),
            };
            return user;
          } catch (error) {
            throw error;
          }
        }
        return Promise.reject(new Error("cred is invalid"));
      },
    }),
    GitHubProvider({
      clientId: process.env.GITHUB_ID as string,
      clientSecret: process.env.GITHUB_SECRET as string,
    }),
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID as string,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET as string,
    }),
  ],
  secret: process.env.AUTH_SECRET,
  session: {
    strategy: "jwt", // trigger jwt callback
  },
  callbacks: {
    async jwt({ token, user }: { token: JWT; user?: User }) {
      // first call of jwt function just user object is provided
      if (user?.username) {
        return { ...token, ...user };
      }
      // on subsequent calls, token is provided and we need to check if it's expired
      if (
        token?.access_token_expires_at &&
        Date.now() + 3000 > token?.access_token_expires_at
      ) {
        return refreshAccessToken(token);
      }
      return { ...token, ...user };
    },
    async session({
      session,
      token,
    }: {
      session: Session;
      token: JWT;
    }): Promise<Session> {
      if (token.app_id) {
        // oauth
        session.app_id = token.app_id;
        return Promise.resolve(session);
      }
      if (
        !token?.access_token_expires_at ||
        Date.now() > token?.access_token_expires_at
      ) {
        return Promise.reject({
          error: new Error(
            "token has expired. Please log in again to get a new token."
          ),
        });
      }
      session.username = token.username;
      (session.role = token.role), (session.access_token = token?.access_token);
      return Promise.resolve(session);
    },

    async signIn({ user, account }) {
      if (account?.provider === "github" || account?.provider === "google") {
        let check_token = {
          oauth_type: account.provider,
          app_id: account.providerAccountId,
          access_token: account.access_token ?? "",
        };
        try {
          const data = await fetchServerWithoutAuthWrapper({
            endpoint: FetchServerWithoutAuthWrapperEndPoint.LOGIN_BY_OAUTH,
            xforward: "",
            agent: "",
            values: check_token,
          });
          if (data.error) {
            throw new FetchError(data.message, data.status);
          } else {
            user.username = data.username;
            user.role = data.role;
            user.access_token = data.access_token;
            user.access_token_expires_at = new Date(
              data.access_token_expires_at
            ).getTime();
            user.refresh_token = data.refresh_token;
            user.refresh_token_expires_at = new Date(
              data.refresh_token_expires_at
            ).getTime();
            return true;
          }
        } catch (error) {
          let e = error as FetchError;
          if (e.status == 404) {
            user.error = "oauth user not link error";
            user.app_id = account.providerAccountId;
            return true;
          } else {
            logger.error(
              `error founded in oauth process:${e.message}`,
              e.status
            );
            return false;
          }
        }
      } else {
        return true;
      }
    },
  },
});

async function refreshAccessToken(token: JWT) {
  const data = await fetchServerWithoutAuthWrapper({
    endpoint: FetchServerWithoutAuthWrapperEndPoint.REFRESH_TOKEN,
    xforward: "",
    agent: "",
    values: { refresh_token: token.refresh_token },
  });
  if (data.error == true) {
    logger.error(`refresh token failed:${data.message}`, data.status);
  } else {
    token.access_token = data.access_token;
    token.access_token_expires_at = new Date(
      data.access_token_expires_at
    ).getTime();
  }
  return {
    ...token,
  };
}
