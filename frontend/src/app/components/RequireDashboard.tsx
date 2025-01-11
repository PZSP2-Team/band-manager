import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";

interface RequireDashboardProps {
  children: React.ReactNode;
}

export function RequireDashboard({ children }: RequireDashboardProps) {
  const router = useRouter();
  const { groupId } = useGroup();

  useEffect(() => {
    if (groupId) {
      router.replace("/events");
    }
  }, [groupId, router]);

  if (groupId) {
    return null;
  }

  return <>{children}</>;
}
