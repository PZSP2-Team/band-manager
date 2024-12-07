import NextAuth, { User, Session } from "next-auth"
import { JWT } from "next-auth/jwt"
import CredentialsProvider from "next-auth/providers/credentials"

interface CustomUser extends User {
    role?: string
}

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
            ): Promise<CustomUser | null> {
                if (!credentials) return null;
                if (
                    credentials?.email === process.env.TEST_EMAIL && 
                    credentials?.password === process.env.TEST_PASSWORD
                ) {
                    return {
                        id: "1",
                        email: credentials.email!,
                        name: "Test User",
                        role: "manager"
                    }
                }
                return null
            }
        })
    ],
    callbacks: {
        async jwt({ token, user }: { token: JWT, user?: CustomUser }): Promise<JWT> {
            if (user) {
                token.role = user.role
            }
            return token
        },
        async session({ session, token }: { session: Session, token: JWT & { role?: string } }): Promise<Session> {
            if (session.user) {
                (session.user as CustomUser).role = token.role
            }
            return session
        }
    },
    pages: {
        signIn: "/login",
    },
    secret: process.env.NEXTAUTH_SECRET
}

const handler = NextAuth(authOptions)
export const GET = handler
export const POST = handler
