import NextAuth, { User, Session, CallbacksOptions } from "next-auth"
import { JWT } from "next-auth/jwt"
import CredentialsProvider from "next-auth/providers/credentials"

export const authOptions = {
    providers: [
        CredentialsProvider({
            name: "Credentials",
            credentials: {
                email: { 
                    label: "Email", 
                    type: "email" 
                },
                password: { 
                    label: "Password", 
                    type: "password" 
                }
            },
            async authorize(
                credentials: {
                    email: string,
                    password: string
                } | undefined
            ): Promise<User | null> {
                if (!credentials) return null;

                try {
                    const response = await fetch(`${process.env.API_URL}/api/auth/login`, {
                        method: "POST",
                        headers: {
                            "Content-Type": "application/json"
                        },
                        body: JSON.stringify({
                            email: credentials.email,
                            password: credentials.password
                        })
                    });
                    
                    if (!response.ok) {
                        const errorText = await response.text();
                        throw new Error(errorText);
                    }

                    const data = await response.json();

                    return {
                        id: data.id,
                        name: data.first_name + data.last_name,
                        email: data.email,
                        role: data.role,
                        groupId: data.group_id
                    };
                } catch (err) {
                    console.error("Login error:", err);
                    return null;
                }
            }
        })
    ],
    callbacks: {
        async jwt({ token, user, trigger, session }: Parameters<CallbacksOptions["jwt"]>[0]): Promise<JWT> {
            if (user) {
                token.role = user.role;
                token.groupId = user.groupId;
                console.log("JWT token updated:", token);
            }
            if (trigger == "update" && session?.user) {
                token.groupId = session.user.groupId;
            }
            return token;
        },
        async session({ session, token }: { session: Session, token: JWT & { role?: string, groupId?: number | null } }): Promise<Session> {
            if (session.user) {
                session.user.role = token.role;
                session.user.groupId = token.groupId;
                console.log("Session updated:", session);
            }
            return session;
        }
    },

    pages: {
        signIn: "/login",
    },
    secret: process.env.NEXTAUTH_SECRET
}

const handler = NextAuth(authOptions);
export const GET = handler;
export const POST = handler;
