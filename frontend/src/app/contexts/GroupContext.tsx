"use client";

import { createContext, useContext, useState } from "react";

type GroupContextType = {
  groupId: number | null;
  userRole: string | null;
  setGroupId: (id: number | null) => void;
  setUserRole: (role: string | null) => void;
};

const GroupContext = createContext<GroupContextType | undefined>(undefined);

export function GroupProvider({ children }: { children: React.ReactNode }) {
  const [groupId, setGroupId] = useState<number | null>(() => {
    if (typeof window !== "undefined") {
      const saved = localStorage.getItem("groupId");
      return saved ? Number(saved) : null;
    }
    return null;
  });

  const [userRole, setUserRole] = useState<string | null>(() => {
    if (typeof window !== "undefined") {
      return localStorage.getItem("userRole");
    }
    return null;
  });

  const handleSetGroupId = (id: number | null) => {
    setGroupId(id);
    if (id) {
      localStorage.setItem("groupId", String(id));
    } else {
      localStorage.removeItem("groupId");
    }
  };

  const handleSetUserRole = (role: string | null) => {
    setUserRole(role);
    if (role) {
      localStorage.setItem("userRole", role);
    } else {
      localStorage.removeItem("userRole");
    }
  };

  return (
    <GroupContext.Provider
      value={{
        groupId,
        userRole,
        setGroupId: handleSetGroupId,
        setUserRole: handleSetUserRole,
      }}
    >
      {children}
    </GroupContext.Provider>
  );
}

export function useGroup() {
  const context = useContext(GroupContext);
  if (context === undefined) {
    throw new Error("useGroup must be used within a GroupProvider");
  }
  return context;
}
