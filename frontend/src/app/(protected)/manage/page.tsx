"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";

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

export default function ManageGroupPage() {
  const { groupId } = useGroup(); // Fetch group ID from context
  const router = useRouter();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [groupInfo, setGroupInfo] = useState<GroupInfo | null>(null);
  const [members, setMembers] = useState<GroupMember[]>([]);

  useEffect(() => {
    const fetchGroupData = async () => {
      try {
        // Mock data
        const mockGroupInfo: GroupInfo = {
          name: "Jazz Band",
          description: "A group of passionate jazz musicians.",
          access_token: "12345-abcdef-67890",
        };

        const mockMembers: GroupMember[] = [
          {
            id: 1,
            firstName: "Alice",
            lastName: "Smith",
            email: "alice@example.com",
            role: "manager",
          },
          {
            id: 2,
            firstName: "Bob",
            lastName: "Johnson",
            email: "bob@example.com",
            role: "member",
          },
          {
            id: 3,
            firstName: "Charlie",
            lastName: "Brown",
            email: "charlie@example.com",
            role: "moderator",
          },
        ];

        // Simulate network delay
        setTimeout(() => {
          setGroupInfo(mockGroupInfo);
          setMembers(mockMembers);
          setRenderState({ status: "loaded" });
        }, 1000);
      } catch (error) {
        console.error("Error fetching group data:", error);
        setRenderState({ status: "error" });
      }
    };

    fetchGroupData();
  }, [groupId]);

  const handleRoleChange = (userId: number, newRole: GroupMember["role"]) => {
    setMembers((prev) =>
      prev.map((member) =>
        member.id === userId ? { ...member, role: newRole } : member
      )
    );
  };

  const handleRemoveMember = (userId: number) => {
    setMembers((prev) => prev.filter((member) => member.id !== userId));
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
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">{groupInfo?.name}</h1>
      <p className="text-gray-600 mb-4">{groupInfo?.description}</p>
      {groupInfo?.access_token && (
        <div className="mb-6">
            <p className="text-gray-400">
            <span
                className="hover:cursor-pointer hover:text-gray-300"
                onClick={handleCopyToken}
            >
                {groupInfo.access_token}
            </span>
            </p>
        </div>
        )}


      <h2 className="text-2xl font-semibold mb-4">Members</h2>
      <ul className="space-y-4">
        {members.map((member) => (
          <li
            key={member.id}
            className="flex items-center justify-between p-4 border border-gray-300 rounded shadow"
          >
            <div>
              <p className="font-semibold">
                {member.firstName} {member.lastName}
              </p>
              <p className="text-sm text-gray-500">{member.email}</p>
            </div>
            <div className="flex items-center space-x-4">
              <select
                value={member.role}
                onChange={(e) =>
                  handleRoleChange(member.id, e.target.value as GroupMember["role"])
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
          </li>
        ))}
      </ul>
    </div>
  );
}
