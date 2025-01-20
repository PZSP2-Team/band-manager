"use client";

import { createContext, useContext, useState } from "react";

/**
 * Type definition for the group context value
 * @property {number | null} groupId - Currently selected group ID
 * @property {string | null} userRole - User's role in current group
 * @property {Function} setGroupId - Updates group ID in context and localStorage
 * @property {Function} setUserRole - Updates user role in context and localStorage
 */
type GroupContextType = {
  groupId: number | null;
  userRole: string | null;
  setGroupId: (id: number | null) => void;
  setUserRole: (role: string | null) => void;
};

const GroupContext = createContext<GroupContextType | undefined>(undefined);

/**
 * Context provider for group-related state management.
 * Manages group selection and user role persistence across the application.
 *
 * Features:
 * - Persists group selection in localStorage
 * - Persists user role in localStorage
 * - Provides group context to child components
 */
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

  /**
   * Updates group ID in state and localStorage
   * Side effect: Syncs group ID between state and localStorage
   */
  const handleSetGroupId = (id: number | null) => {
    setGroupId(id);
    if (id) {
      localStorage.setItem("groupId", String(id));
    } else {
      localStorage.removeItem("groupId");
    }
  };

  /**
   * Updates user role in state and localStorage
   * Side effect: Syncs user role between state and localStorage
   */
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

/**
 * Custom hook to access group context
 * Ensures component is wrapped in GroupProvider
 *
 * @returns {GroupContextType} Group context value
 */
export function useGroup() {
  const context = useContext(GroupContext);
  if (context === undefined) {
    throw new Error("useGroup must be used within a GroupProvider");
  }
  return context;
}
