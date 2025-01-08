import "next-auth";
import { DefaultSession } from "next-auth";

declare module 'next-auth' {
    interface User {
        id?: number | null;
    }

    interface Session {
        user: {
            id?: number | null;
        } & DefaultSession["user"];
    }
}
