"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import LoadingScreen from "@/src/app/components/LoadingScreen";

// Type Definitions
type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Group = {
  id: number;
  name: string;
  description: string;
  access_token: string;
  members: User[];
};

type User = {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
  role: "manager" | "moderator" | "member";
};

export default function ManagePage() {
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [groups, setGroups] = useState<Group[]>([]);
  const [expandedGroup, setExpandedGroup] = useState<number | null>(null);
  const [showUserDropdown, setShowUserDropdown] = useState<number | null>(null);
  const mockUsers: User[] = [
    { id: 111, first_name: "John", last_name: "Doe", email: "john@example.com", role: "member" },
    { id: 112, first_name: "Jane", last_name: "Smith", email: "jane@example.com", role: "member" },
    { id: 113, first_name: "Jim", last_name: "Beam", email: "jim@example.com", role: "member" },
  ];

  useEffect(() => {
    if (sessionStatus === "loading") return;

    // Mock data for testing
    const mockGroups: Group[] = [
      {
        id: 1,
        name: "Orchestra",
        description: "The main orchestra group",
        access_token: "abc123xyz789",
        members: [
          {
            id: 101,
            first_name: "Alice",
            last_name: "Smith",
            email: "alice@example.com",
            role: "manager",
          },
          {
            id: 102,
            first_name: "Bob",
            last_name: "Jones",
            email: "bob@example.com",
            role: "member",
          },
        ],
      },
      {
        id: 2,
        name: "Choir",
        description: "Choir group",
        access_token: "def456uvw123",
        members: [
          {
            id: 103,
            first_name: "Carol",
            last_name: "Taylor",
            email: "carol@example.com",
            role: "moderator",
          },
          {
            id: 104,
            first_name: "Dave",
            last_name: "Wilson",
            email: "dave@example.com",
            role: "member",
          },
        ],
      },
    ];

    setTimeout(() => {
      setGroups(mockGroups);
      setRenderState({ status: "loaded" });
    }, 1000);
  }, [sessionStatus]);

  const handleRemoveUser = async (groupId: number, userId: number) => {
    try {
      // Commented out for mock implementation
      // await fetch(/api/group/remove/${groupId}/${session?.user.id}/${userId}, {
      //   method: "DELETE",
      // });

      setGroups((prevGroups) =>
        prevGroups.map((group) =>
          group.id === groupId
            ? {
                ...group,
                members: group.members.filter((member) => member.id !== userId),
              }
            : group
        )
      );
    } catch (error) {
      console.error("Error removing user:", error);
    }
  };

  const handleRoleChange = (groupId: number, userId: number, newRole: "manager" | "moderator" | "member") => {
    setGroups((prevGroups) =>
      prevGroups.map((group) =>
        group.id === groupId
          ? {
              ...group,
              members: group.members.map((member) =>
                member.id === userId ? { ...member, role: newRole } : member
              ),
            }
          : group
      )
    );
  };

  const handleAddUser = async (groupId: number, userId: number) => {
    try {
      const accessToken = groups.find((group) => group.id === groupId)?.access_token;
      if (!accessToken) throw new Error("Access token not found");

      const newUser: User = {
        id: userId,
        first_name: "New",
        last_name: "User",
        email: `newuser${Date.now()}@example.com`,
        role: "member",
      };

      setGroups((prevGroups) =>
        prevGroups.map((group) =>
          group.id === groupId
            ? {
                ...group,
                members: [...group.members, newUser],
              }
            : group
        )
      );

      // Close the dropdown after adding the user
      setShowUserDropdown(null);
    } catch (error) {
      console.error("Error adding user:", error);
    }
  };

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
    alert("Access token copied to clipboard");
  };

  if (sessionStatus === "loading" || renderState.status === "loading") {
    return <LoadingScreen />;
  }

  if (renderState.status === "error") {
    return (
      <div className="text-center mt-10">
        Failed to load groups. Please try again later.
      </div>
    );
  }

  return (
    <div className="flex flex-col mt-10 p-6">
      <h1 className="text-3xl font-bold mb-6 text-left">Manage Groups</h1>
      <ul className="space-y-4">
        {groups.map((group) => (
          <li key={group.id} className="p-4 border border-gray-300 rounded shadow">
            <div
              className="flex justify-between items-center hover:cursor-pointer"
              onClick={() =>
                setExpandedGroup((prev) => (prev === group.id ? null : group.id))
              }
            >
              <div>
                <h2 className="text-lg font-semibold">{group.name}</h2>
                <p className="text-gray-600 text-sm">{group.description}</p>
              </div>
              <div className="flex items-center space-x-4">
                <span
                  className="text-sm text-gray-400 hover:cursor-pointer"
                  onClick={(e) => {
                    e.stopPropagation();
                    copyToClipboard(group.access_token);
                  }}
                >
                  {group.access_token.slice(0, 4)}...{group.access_token.slice(-4)}
                </span>
                
                <button className="text-blue-600 hover:underline">
                  {expandedGroup === group.id ? "▲" : "▼"}
                </button>
              </div>
            </div>
            {expandedGroup === group.id && (
              <div className="mt-4 border-t border-gray-200 pt-4">
                <h3 className="text-md font-bold mb-2">Participants</h3>
                <ul className="space-y-2">
                {group.members.map((member, index) => (
                    <li
                    key={member.id}
                    className="flex justify-between items-center space-x-4" // добавлен space-x для отступов между элементами
                    >
                    <div className="flex flex-col space-y-1"> {/* flex-col для вертикальной ориентации внутри div */}
                        <p className="font-medium">
                        {index + 1}. {member.first_name} {member.last_name}
                        </p>
                        <p className="text-sm text-gray-500">{member.email}</p>
                    </div>

                    <div className="flex items-center space-x-2"> {/* добавлен flex для горизонтального расположения */}
                        <select
                        value={member.role}
                        onChange={(e) =>
                            handleRoleChange(
                            group.id,
                            member.id,
                            e.target.value as "manager" | "moderator" | "member"
                            )
                        }
                        className="p-1 border-gray-600 rounded bg-gray-800 text-white"
                        >
                        <option value="manager">Manager</option>
                        <option value="moderator">Moderator</option>
                        <option value="member">Member</option>
                        </select>
                        <button
                        onClick={() => handleRemoveUser(group.id, member.id)}
                        className="px-2 py-1 bg-red-600 text-white rounded hover:bg-red-500"
                        >
                        Remove
                        </button>
                    </div>
                    </li>
                ))}
                </ul>


                
                
                <button
                  className="mt-4 px-4 py-2 bg-green-600 text-white rounded hover:bg-green-500"
                  onClick={() =>
                    setShowUserDropdown((prev) =>
                      prev === group.id ? null : group.id
                    )
                  }
                >
                  Add User
                </button>
                {showUserDropdown === group.id && (
                  <div className="mt-2 p-2 bg-gray-100 rounded shadow">
                    <ul>
                      {mockUsers.map((user) => (
                        <li
                          key={user.id}
                          className="p-2 hover:bg-gray-200 cursor-pointer"
                          onClick={() => handleAddUser(group.id, user.id)}
                        >
                          {user.first_name} {user.last_name}
                        </li>
                      ))}
                    </ul>
                    <button
                      className="mt-2 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-500"
                      onClick={() => setShowUserDropdown(null)}
                    >
                      Submit
                    </button>
                  </div>
                )}
              </div>
            )}
          </li>
        ))}
      </ul>
    </div>
  );
}
