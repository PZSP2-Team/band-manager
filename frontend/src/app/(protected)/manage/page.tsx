"use client";
import { useState, useEffect } from "react";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { useSession } from "next-auth/react";
import { User, RefreshCw } from "lucide-react";

/**
 * Represents the component's render state
 */
type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

/**
 * Represents a member of the group with their basic information and role
 */
type GroupMember = {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
  role: "manager" | "moderator" | "member";
};

/**
 * Basic information about a group
 */
type GroupInfo = {
  name: string;
  description: string;
  access_token: string;
};

/**
 * API response structure for group information
 */
interface GroupInfoResponse {
  name: string;
  description: string;
  access_token: string;
}

/**
 * API response structure for individual group member data
 */
interface GroupMemberResponse {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
  role: "manager" | "moderator" | "member";
}

/**
 * API response structure containing array of group members
 */
interface GroupMembersResponse {
  members: GroupMemberResponse[];
}

/**
 * Page component for managing group members and their roles
 * Provides interface for viewing group information, managing member roles,
 * and removing members from the group
 */
export default function ManageGroupPage() {
  const { groupId } = useGroup();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [groupInfo, setGroupInfo] = useState<GroupInfo | null>(null);
  const [members, setMembers] = useState<GroupMember[]>([]);
  const [isRefreshingToken, setIsRefreshingToken] = useState(false);

  /**
   * Fetches group information and member data simultaneously
   * Dependencies: groupId, sessionStatus, session?.user?.id
   *
   * Side effects:
   * - Sets groupInfo state with group details (name, description, access token)
   * - Sets members state with mapped member data from API response
   * - Updates renderState to "loaded" on success or "error" on failure
   */
  useEffect(() => {
    if (sessionStatus === "loading") return;
    const fetchData = async () => {
      try {
        const infoResponse = await fetch(
          `/api/group/${groupId}/${session?.user?.id}`,
        );
        if (!infoResponse.ok) throw new Error("Failed to fetch group info");
        const infoData = (await infoResponse.json()) as GroupInfoResponse;
        setGroupInfo({
          name: infoData.name,
          description: infoData.description,
          access_token: infoData.access_token,
        });

        const membersResponse = await fetch(
          `/api/group/members/${groupId}/${session?.user?.id}`,
        );
        if (!membersResponse.ok) throw new Error("Failed to fetch members");
        const membersData =
          (await membersResponse.json()) as GroupMembersResponse;
        setMembers(
          membersData.members.map((member) => ({
            id: member.id,
            firstName: member.first_name,
            lastName: member.last_name,
            email: member.email,
            role: member.role,
          })),
        );

        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching data:", error);
        setRenderState({ status: "error" });
      }
    };

    if (groupId) {
      fetchData();
    }
  }, [groupId, sessionStatus, session?.user?.id]);

  /**
   * Updates a member's role in the group
   * Side effect: Updates members state with new role
   */
  const handleRoleChange = async (
    userId: number,
    newRole: GroupMember["role"],
  ) => {
    try {
      const response = await fetch(
        `/api/group/role/${groupId}/${userId}/${session?.user?.id}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ new_role: newRole }),
        },
      );

      if (!response.ok) {
        throw new Error("Failed to update role");
      }

      setMembers((prev) =>
        prev.map((member) =>
          member.id === userId ? { ...member, role: newRole } : member,
        ),
      );
    } catch (error) {
      console.error("Error updating role:", error);
      alert("Failed to update role. Please try again.");
    }
  };

  /**
   * Removes a member from the group
   * Side effect: Removes member from members state
   */
  const handleRemoveMember = async (userId: number) => {
    try {
      const response = await fetch(
        `/api/group/remove/${groupId}/${session?.user?.id}/${userId}`,
        {
          method: "DELETE",
        },
      );

      if (!response.ok) {
        throw new Error("Failed to remove member");
      }

      setMembers((prev) => prev.filter((member) => member.id !== userId));
    } catch (error) {
      console.error("Error removing member:", error);
    }
  };

  /**
   * Copies group access token to clipboard
   * Side effect: Triggers system clipboard and shows alert
   */
  const handleCopyToken = () => {
    if (groupInfo?.access_token) {
      navigator.clipboard.writeText(groupInfo.access_token);
      alert("Access token copied to clipboard!");
    }
  };

  const handleRefreshToken = async () => {
    setIsRefreshingToken(true);
    try {
      const response = await fetch(
        `/api/group/refresh-token/${groupId}/${session?.user?.id}`,
        {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
        },
      );

      if (!response.ok) {
        throw new Error("Failed to regenerate token");
      }

      const data = await response.json();
      setGroupInfo((prev) =>
        prev ? { ...prev, access_token: data.access_token } : null,
      );
      alert("Access token has been regenerated!");
    } catch (error) {
      console.error("Error regenerating token:", error);
      alert("Failed to regenerate token. Please try again.");
    } finally {
      setIsRefreshingToken(false);
    }
  };

  if (renderState.status === "loading") return <LoadingScreen />;
  if (renderState.status === "error")
    return (
      <div className="text-center mt-10">
        Failed to load group data. Please try again later.
      </div>
    );

  return (
    <div className="p-10">
      <h1 className="text-3xl font-bold mb-4">{groupInfo?.name}</h1>
      <p className="text-gray-600 mb-4">{groupInfo?.description}</p>
      {groupInfo?.access_token && (
        <div className="mb-6 flex items-center gap-4">
          <p className="text-customGray">
            <span className="hover:cursor-pointer" onClick={handleCopyToken}>
              {groupInfo.access_token}
            </span>
          </p>
          <button
            onClick={handleRefreshToken}
            disabled={isRefreshingToken}
            className="flex items-center gap-2 px-6 py-2 bg-blue-600 transition text-white rounded shadow hover:bg-blue-500 disabled:bg-blue-300"
          >
            <RefreshCw
              className={`h-4 w-4 ${isRefreshingToken ? "animate-spin" : ""}`}
            />
            Refresh token
          </button>
        </div>
      )}

      <h2 className="text-2xl font-semibold mb-4">Members</h2>
      <ul className="space-y-4">
        {members.map((member) => {
          const isCurrentUser = member.id === session?.user?.id;

          return (
            <li
              key={member.id}
              className="flex items-center justify-between p-4 border border-customGray rounded shadow"
            >
              <div className="flex items-center gap-3">
                {isCurrentUser && <User className="h-5 w-5 text-blue-500" />}
                <div>
                  <p className="font-semibold">
                    {member.firstName} {member.lastName}
                  </p>
                  <p className="text-sm text-gray-500">{member.email}</p>
                </div>
              </div>

              {!isCurrentUser && (
                <div className="flex items-center space-x-4">
                  <select
                    value={member.role}
                    onChange={(e) =>
                      handleRoleChange(
                        member.id,
                        e.target.value as GroupMember["role"],
                      )
                    }
                    className="p-2 border rounded bg-gray-800 text-white"
                  >
                    <option value="manager">Manager</option>
                    <option value="moderator">Moderator</option>
                    <option value="member">Member</option>
                  </select>
                  <button
                    className="px-4 py-2 bg-red-500 text-white rounded shadow hover:bg-red-400"
                    onClick={() => handleRemoveMember(member.id)}
                  >
                    Remove
                  </button>
                </div>
              )}
            </li>
          );
        })}
      </ul>
    </div>
  );
}
