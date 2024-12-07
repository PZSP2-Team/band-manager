import "next-auth";
import { DefaultSession } from "next-auth";

declare module 'next-auth' {
    interface User {
        role?: string | null;
        groupId?: number | null;
    }

    interface Session {
        user: {
            role?: string | null;
            groupId?: number | null;
        } & DefaultSession["user"];
    }
}
