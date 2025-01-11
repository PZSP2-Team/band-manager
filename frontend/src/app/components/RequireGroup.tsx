import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";

interface RequireGroupProps {
  children: React.ReactNode;
}

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
