import { ReactNode } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";

interface RequireModeratorProps {
  children: ReactNode;
  fallback?: ReactNode;
}

export const RequireModerator = ({
  children,
  fallback = (
    <div className="text-center mt-10">
      You need manager permissions to access this page.
    </div>
  ),
}: RequireModeratorProps) => {
  const { userRole } = useGroup();

  if (userRole !== "manager" && userRole !== "moderator") {
    return <>{fallback}</>;
  }

  return <>{children}</>;
};
