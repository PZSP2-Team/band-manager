import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";

interface RequireGroupProps {
  children: React.ReactNode;
}

/**
 * Higher-order component that restricts access to routes requiring group context.
 * Redirects to dashboard if no group is selected.
 *
 * @component
 * Properties:
 * @param {React.ReactNode} children - Child components to render when group is selected
 *
 * Features:
 * - Checks for active group selection
 * - Automatic redirection to dashboard if no group
 * - Conditional rendering of children
 */
export function RequireGroup({ children }: RequireGroupProps) {
  const router = useRouter();
  const { groupId } = useGroup();

  useEffect(() => {
    if (!groupId) {
      router.replace("/dashboard");
    }
  }, [groupId, router]);

  if (!groupId) {
    return null;
  }

  return <>{children}</>;
}
