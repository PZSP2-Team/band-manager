import NextAuth from "next-auth";

import { authOptions } from "@/auth.config";

/**
 * NextAuth authentication handler for both GET and POST requests.
 * Implements authentication endpoints using configuration from auth.config.
 *
 * Uses configuration defined in auth.config.ts including:
 * - Authentication providers
 * - Callback functions
 * - Session strategy
 * - JWT settings
 */
const handler = NextAuth(authOptions);
export const GET = handler;
export const POST = handler;
