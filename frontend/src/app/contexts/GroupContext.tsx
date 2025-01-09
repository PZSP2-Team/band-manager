"use client";
import { createContext, useContext, useState } from "react";

interface GroupContextType {
  groupId: number | null;
  userRole: string | null;
  setGroupId: (id: number | null) => void;
  setUserRole: (role: string | null) => void;
}

const GroupContext = createContext<GroupContextType | null>(null);

export function GroupProvider({ children }: { children: React.ReactNode }) {
  const [groupId, setGroupId] = useState<number | null>(null);
  const [userRole, setUserRole] = useState<string | null>(null);

  return (
    <GroupContext.Provider
      value={{ groupId, userRole, setGroupId, setUserRole }}
    >
      {children}
    </GroupContext.Provider>
  );
}

export function useGroup() {
  const context = useContext(GroupContext);
  if (!context) {
    throw new Error("useGroup must be used within GroupProvider");
  }
  return context;
}
