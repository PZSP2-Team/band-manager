import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  ...(process.env.NODE_ENV === "production" ? { output: "standalone" } : {}),
  async rewrites() {
    return [
      {
        source: "/api/auth/:auth*",
        destination: "/api/auth/:auth*",
      },
      {
        source: "/api/verify/:path*",
        destination: `http://${process.env.NEXT_PUBLIC_BACKEND_HOST}:${process.env.NEXT_PUBLIC_BACKEND_PORT}/api/verify/:path*`,
      },
      {
        source: "/api/:path*",
        destination: `http://${process.env.NEXT_PUBLIC_BACKEND_HOST}:${process.env.NEXT_PUBLIC_BACKEND_PORT}/api/:path*`,
        has: [
          {
            type: "header",
            key: "user-id",
          },
        ],
      },
    ];
  },
};

export default nextConfig;
