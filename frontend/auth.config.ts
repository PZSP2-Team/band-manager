import { User, Session, CallbacksOptions, AuthOptions } from "next-auth";
import { JWT } from "next-auth/jwt";
import { BACKEND_URL } from "@/src/config/api";
import Credentials from "next-auth/providers/credentials";

/**
 * NextAuth configuration options for the application.
 * Sets up authentication with credentials provider and handles session/token management.
 */
export const authOptions: AuthOptions = {
  providers: [
    Credentials({
      name: "Credentials",
      credentials: {
        email: {
          label: "Email",
          type: "email",
        },
        password: {
          label: "Password",
          type: "password",
        },
      },
      /**
       * Credentials provider configuration.
       * Handles email/password authentication against backend API.
       *
       * Features:
       * - Email and password form fields
       * - Backend verification via /api/verify/login
       * - Returns user ID on successful auth
       */
      async authorize(
        credentials:
          | {
              email: string;
              password: string;
            }
          | undefined,
      ): Promise<User | null> {
        if (!credentials) return null;

        try {
          const response = await fetch(`${BACKEND_URL}/api/verify/login`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              email: credentials.email,
              password: credentials.password,
            }),
          });

          if (!response.ok) {
            const errorText = await response.text();
            throw new Error(errorText);
          }

          const data = await response.json();

          return {
            id: data.id,
          };
        } catch (err) {
          console.error("Login error:", err);
          return null;
        }
      },
    }),
  ],
  callbacks: {
    /**
     * JWT callback to customize token contents.
     * Adds user ID to the JWT token when created.
     */
    async jwt({
      token,
      user,
    }: Parameters<CallbacksOptions["jwt"]>[0]): Promise<JWT> {
      if (user) {
        token.id = user.id;
      }
      return token;
    },
    /**
     * Session callback to customize session object.
     * Syncs user ID from token to session.
     */
    async session({
      session,
      token,
    }: {
      session: Session;
      token: JWT & { id?: number };
    }): Promise<Session> {
      if (session.user) {
        session.user.id = token.id;
        console.log("Session updated:", session);
      }
      return session;
    },
  },

  pages: {
    signIn: "/login",
  },
  secret: process.env.NEXTAUTH_SECRET,
};
