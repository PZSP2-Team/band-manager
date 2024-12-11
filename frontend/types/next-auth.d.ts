import "next-auth";
import { DefaultSession } from "next-auth";

declare module 'next-auth' {
    interface User {
        id?: number | null;
        role?: string | null;
        groupId?: number | null;
    }

    interface Session {
        user: {
            id?: number | null;
            role?: string | null;
            groupId?: number | null;
        } & DefaultSession["user"];
    }
}
