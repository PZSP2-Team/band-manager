import type { NextConfig } from "next";

/**
 * Next.js configuration with proxy routing setup
 * Handles routing between frontend and backend services
 */
const nextConfig: NextConfig = {
  ...(process.env.NODE_ENV === "production" ? { output: "standalone" } : {}),
  trailingSlash: false,
  /**
   * Rewrite rules for API request proxying
   *
   * Rules:
   * 1. Notesheet creation endpoint
   *    - Source: /api/track/notesheet/create
   *    - Requires user-id header
   *
   * 2. Next.js Authentication endpoints
   *    - Source: /api/auth/*
   *    - Handles all auth-related routes
   *
   * 3. Backend Authentication endpoints
   *    - Source: /api/verify/*
   *    - Proxies to backend authentication service
   *    - No auth header required
   *
   * 4. General API endpoints
   *    - Source: /api/*
   *    - Proxies to backend API
   *    - Requires user-id header
   */
  async rewrites() {
    return [
      {
        source: "/api/track/notesheet/create",
        destination: "/api/track/notesheet/create",
        has: [
          {
            type: "header",
            key: "user-id",
          },
        ],
      },
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
