"use client";
import { useState, useEffect } from "react";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { useSession } from "next-auth/react";
import { User } from "lucide-react";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type GroupMember = {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
  role: "manager" | "moderator" | "member";
};

type GroupInfo = {
  name: string;
  description: string;
  access_token: string;
};

interface GroupInfoResponse {
  name: string;
  description: string;
  access_token: string;
}

interface GroupMemberResponse {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
  role: "manager" | "moderator" | "member";
}

interface GroupMembersResponse {
  members: GroupMemberResponse[];
}

export default function ManageGroupPage() {
  const { groupId } = useGroup();
  const { data: session } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [groupInfo, setGroupInfo] = useState<GroupInfo | null>(null);
  const [members, setMembers] = useState<GroupMember[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      setRenderState({ status: "loading" });

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

    if (groupId && session?.user?.id) {
      fetchData();
    }
  }, [groupId, session?.user?.id]);

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

  const handleCopyToken = () => {
    if (groupInfo?.access_token) {
      navigator.clipboard.writeText(groupInfo.access_token);
      alert("Access token copied to clipboard!");
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
        <div className="mb-6">
          <p className="text-customGray">
            <span className="hover:cursor-pointer" onClick={handleCopyToken}>
              {groupInfo.access_token}
            </span>
          </p>
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
