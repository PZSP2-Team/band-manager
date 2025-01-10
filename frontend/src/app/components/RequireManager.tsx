import { ReactNode } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";

interface RequireManagerProps {
  children: ReactNode;
  fallback?: ReactNode;
}

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
