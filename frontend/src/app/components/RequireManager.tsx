import { ReactNode } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";

interface RequireManagerProps {
  children: ReactNode;
  fallback?: ReactNode;
}

/**
 * Higher-order component that restricts access to manager-only content.
 * Shows fallback content for non-manager users.
 *
 * @component
 * Properties:
 * @param {ReactNode} children - Content to render for managers
 * @param {ReactNode} [fallback] - Optional content to show for non-managers
 *                                 Defaults to permission denied message
 *
 * Features:
 * - Role-based access control
 * - Configurable fallback content
 * - Uses GroupContext for role verification
 */
export const RequireManager = ({
  children,
  fallback = (
    <div className="text-center mt-10">
      You need manager permissions to access this page.
    </div>
  ),
}: RequireManagerProps) => {
  const { userRole } = useGroup();

  if (userRole !== "manager") {
    return <>{fallback}</>;
  }

  return <>{children}</>;
};
