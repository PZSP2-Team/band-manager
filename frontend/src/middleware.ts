import { withAuth, NextRequestWithAuth } from "next-auth/middleware";
import { NextResponse } from "next/server";

type PublicPath = "/login" | "/register" | "/";
const PUBLIC_PATHS: PublicPath[] = ["/login", "/register", "/"];

const PUBLIC_PREFIXES = ["/api/verify"] as const;

function isPublicPath(path: string) {
  if (PUBLIC_PATHS.includes(path as PublicPath)) {
    return true;
  }

  return PUBLIC_PREFIXES.some((prefix) => path.startsWith(prefix));
}

export default withAuth(
  function middleware(req: NextRequestWithAuth) {
    const path = req.nextUrl.pathname;

    if (isPublicPath(path) && req.nextauth.token) {
      return NextResponse.redirect(new URL("/dashboard", req.url));
    }

    if (path.startsWith("/api/") && !isPublicPath(path) && req.nextauth.token) {
      const requestHeaders = new Headers(req.headers);
      requestHeaders.set("user-id", req.nextauth.token.id as string);
      requestHeaders.set("x-auth-timestamp", Date.now().toString());

      return NextResponse.next({
        request: {
          headers: requestHeaders,
        },
      });
    }

    return NextResponse.next();
  },
  {
    callbacks: {
      authorized: ({ token, req }) => {
        return isPublicPath(req.nextUrl.pathname) || !!token;
      },
    },
  },
);

export const config = {
  matcher: ["/:path*"],
};
