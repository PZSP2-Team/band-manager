import { withAuth } from "next-auth/middleware";
import { NextResponse } from "next/server";

export default withAuth(
  function middleware(req) {
    const path = req.nextUrl.pathname;
    const isPublicPath = path === "/login" || path === "/register" || path === "/";

    if (isPublicPath && req.nextauth.token) {
      return NextResponse.redirect(new URL("/group", req.url));
    }
  },
  {
    callbacks: {
      authorized: ({ token, req }) => {
        const path = req.nextUrl.pathname;
        const isPublicPath = path === "/login" || path === "/register" || path === "/";
        
        if (isPublicPath) {
          return true;
        }
        return !!token;
      },
    },
  }
);

export const config = {
  matcher: ["/", "/login", "/register", "/group/:path*", "/events/:path*"],
};
