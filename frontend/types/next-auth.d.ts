import "next-auth";
import { DefaultUser } from "next-auth";

declare module "next-auth" {
    interface Session {
        user: DefaultUser & {
            role?: string | null;
            groupId?: number | null;
        };
    }
    
    interface User extends DefaultUser {
        role?: string | null;
        groupId?: number | null;
    }
}
