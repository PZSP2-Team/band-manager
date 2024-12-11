import type { NextConfig } from "next";

const nextConfig: NextConfig = {
    async rewrites() {
        return [
            {
                source: "/api/auth/:auth*",
                destination: "/api/auth/:auth*"
            },
            {
                source: "/api/:path*",
                destination: `http://${process.env.NEXT_PUBLIC_BACKEND_HOST}:${process.env.NEXT_PUBLIC_BACKEND_PORT}/api/:path*`
            }
        ]
    }
};

export default nextConfig;
