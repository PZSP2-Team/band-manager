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
  const [showDeleteConfirmation, setShowDeleteConfirmation] = useState<number | null>(null);

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
          { id: 101, first_name: "Alice", last_name: "Smith", email: "alice@example.com", role: "manager" },
          { id: 102, first_name: "Bob", last_name: "Jones", email: "bob@example.com", role: "member" },
        ],
      },
      {
        id: 2,
        name: "Choir",
        description: "Choir group",
        access_token: "def456uvw123",
        members: [
          { id: 103, first_name: "Carol", last_name: "Taylor", email: "carol@example.com", role: "moderator" },
          { id: 104, first_name: "Dave", last_name: "Wilson", email: "dave@example.com", role: "member" },
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
      setGroups((prevGroups) =>
        prevGroups.map((group) =>
          group.id === groupId
            ? { ...group, members: group.members.filter((member) => member.id !== userId) }
            : group
        )
      );
    } catch (error) {
      console.error("Error removing user:", error);
    }
  };

  const handleAddUser = async (groupId: number, userId: number) => {
    try {
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
            ? { ...group, members: [...group.members, newUser] }
            : group
        )
      );

      setShowUserDropdown(null); // Close dropdown
    } catch (error) {
      console.error("Error adding user:", error);
    }
  };

  const handleDeleteGroup = (groupId: number) => {
    setGroups((prevGroups) => prevGroups.filter((group) => group.id !== groupId));
    setShowDeleteConfirmation(null);
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
      <h1 className="text-3xl font-bold mb-6 text-left">Manage Subgroups</h1>
      <ul className="space-y-4">
        {groups.map((group) => (
          <li key={group.id} className="p-4 border border-gray-300 rounded shadow">
            <div
              className="flex justify-between items-center cursor-pointer"
              onClick={() => setExpandedGroup((prev) => (prev === group.id ? null : group.id))}
            >
              <div className="flex flex-col">
                <h2 className="text-lg font-semibold">{group.name}</h2>
                <p className="text-gray-600 text-sm">{group.description}</p>
              </div>
              <button
                onClick={(e) => {
                  e.stopPropagation(); // Prevent row click event from triggering
                  setShowDeleteConfirmation(group.id);
                }}
                className="text-gray-500 hover:text-red-600 text-xl ml-2"
              >
                🗑️
              </button>
            </div>
            {expandedGroup === group.id && (
              <div className="mt-4 border-t border-gray-200 pt-4">
                <h3 className="text-md font-bold mb-2">Participants</h3>
                <ul className="space-y-2">
                  {group.members.map((member) => (
                    <li key={member.id} className="flex justify-between items-center">
                      <div>
                        <p className="font-medium">
                          {member.first_name} {member.last_name}
                        </p>
                        <p className="text-sm text-gray-500">{member.email}</p>
                      </div>
                      <button
                        onClick={() => handleRemoveUser(group.id, member.id)}
                        className="px-2 py-1 bg-red-600 text-white rounded hover:bg-red-500"
                      >
                        Remove
                      </button>
                    </li>
                  ))}
                </ul>
                <button
                  className="mt-4 px-4 py-2 bg-green-600 text-white rounded hover:bg-green-500"
                  onClick={() =>
                    setShowUserDropdown((prev) => (prev === group.id ? null : group.id))
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
                  </div>
                )}
              </div>
            )}
          </li>
        ))}
      </ul>
      <button
        className="mt-6 w-48 px-4 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 text-left"
        onClick={() => router.push("/subgroups/createSubgroup")}
      >
        Create New Subgroup
      </button>
      {showDeleteConfirmation !== null && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <div className="bg-white p-6 rounded shadow">
            <p className="text-lg font-bold text-gray-800 mb-4">
              Are you sure you want to delete this subgroup?
            </p>
            <div className="flex justify-end space-x-4">
              <button
                onClick={() => setShowDeleteConfirmation(null)}
                className="px-4 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-200"
              >
                Cancel
              </button>
              <button
                onClick={() => handleDeleteGroup(showDeleteConfirmation!)}
                className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-500"
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
