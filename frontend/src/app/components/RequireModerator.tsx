import { ReactNode } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";

interface RequireModeratorProps {
  children: ReactNode;
  fallback?: ReactNode;
}

/**
 * Higher-order component that restricts access to moderator/manager-only content.
 * Shows fallback content for users without sufficient permissions.
 *
 * @component
 * Properties:
 * @param {ReactNode} children - Content to render for moderators and managers
 * @param {ReactNode} [fallback] - Optional content to show for non-authorized users
 *                                 Defaults to permission denied message
 *
 * Features:
 * - Role-based access control (allows both moderator and manager roles)
 * - Configurable fallback content
 * - Uses GroupContext for role verification
 */
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
