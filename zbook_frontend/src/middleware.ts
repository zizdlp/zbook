import { NextRequest } from "next/server";
import createIntlMiddleware from "next-intl/middleware";
import { locales } from "./navigation";
// import { auth } from "@/auth";
const privatePages = ["/workspace", "/link"];

const intlMiddleware = createIntlMiddleware({
  locales,
  localePrefix: "as-needed",
  defaultLocale: "en",
});
// const authMiddleware = auth((req) => {
//   if (!req.auth) {
//     const url = req.url.replace(req.nextUrl.pathname, "/auth/login");
//     return Response.redirect(url);
//   }
//   return intlMiddleware(req);
// });

export default function middleware(req: NextRequest) {
  const privatePathnameRegex = RegExp(
    `^(/(${locales.join("|")}))?(${privatePages
      .flatMap((p) => (p === "/" ? ["", "/"] : p))
      .join("|")})/?`,
    "i"
  );
  const isPrivatePage = privatePathnameRegex.test(req.nextUrl.pathname);
  if (isPrivatePage) {
    return intlMiddleware(req);
    // return (authMiddleware as any)(req); // Error: TODO:
  } else {
    return intlMiddleware(req);
  }
}

export const config = {
  // Skip all paths that should not be internationalized
  matcher: ["/((?!api|_next|.*\\..*).*)"],
};
