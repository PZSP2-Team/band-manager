"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { ChevronUp, ChevronDown, Check } from "lucide-react";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { useGroup } from "../../contexts/GroupContext";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Subgroup = {
  id: number;
  name: string;
  members: number[];
};

type User = {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
};

export default function SubgroupsPage() {
  const { groupId } = useGroup();
  const router = useRouter();
  const { data: session } = useSession();
  const [availableUsers, setAvailableUsers] = useState<User[]>([]);
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [selectedUserIds, setSelectedUserIds] = useState<number[]>([]);
  const [subgroups, setSubgroups] = useState<Subgroup[]>([]);
  const [expandedGroup, setExpandedGroup] = useState<number | null>(null);
  const [showUserDropdown, setShowUserDropdown] = useState<number | null>(null);
  const [showDeleteConfirmation, setShowDeleteConfirmation] = useState<
    number | null
  >(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const membersResponse = await fetch(
          `/api/group/members/${groupId}/${session?.user?.id}`,
        );

        if (!membersResponse.ok) {
          throw new Error("Failed to fetch members");
        }

        const membersData = await membersResponse.json();
        setAvailableUsers(membersData.members);

        const subgroupsResponse = await fetch(
          `/api/group/subgroups/${groupId}/${session?.user?.id}`,
        );

        if (!subgroupsResponse.ok) {
          throw new Error("Failed to fetch subgroups");
        }

        const subgroupsData = await subgroupsResponse.json();

        setSubgroups(subgroupsData.subgroups);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching data:", error);
        setRenderState({ status: "error" });
      }
    };

    fetchData();
  }, [groupId, session?.user?.id]);

  const handleGroupExpand = (groupId: number) => {
    if (expandedGroup !== groupId) {
      setSelectedUserIds([]);
      setShowUserDropdown(null);
    }
    setExpandedGroup((prev) => (prev === groupId ? null : groupId));
  };

  const handleRemoveUser = async (subgroupId: number, userId: number) => {
    try {
      const response = await fetch(
        `/api/group/subgroups/${groupId}/members/${session?.user?.id}`,
        {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            subgroup_id: subgroupId,
            user_id: userId,
          }),
        },
      );

      if (!response.ok) {
        throw new Error("Failed to remove user");
      }

      setSubgroups((prevSubgroups) =>
        prevSubgroups.map((subgroup) =>
          subgroup.id === subgroupId
            ? {
                ...subgroup,
                members: subgroup.members.filter(
                  (memberId) => memberId !== userId,
                ),
              }
            : subgroup,
        ),
      );
    } catch (error) {
      console.error("Error removing user:", error);
    }
  };

  const handleAddSelectedUsers = async (subgroupId: number) => {
    try {
      const response = await fetch(
        `/api/group/subgroups/${groupId}/members/${session?.user?.id}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            subgroup_id: subgroupId,
            userIds: selectedUserIds,
          }),
        },
      );

      if (!response.ok) {
        throw new Error("Failed to add users to subgroup");
      }

      setSubgroups((prevSubgroups) =>
        prevSubgroups.map((subgroup) =>
          subgroup.id === subgroupId
            ? {
                ...subgroup,
                members: [...subgroup.members, ...selectedUserIds],
              }
            : subgroup,
        ),
      );

      setSelectedUserIds([]);
      setShowUserDropdown(null);
    } catch (error) {
      console.error("Error adding users:", error);
    }
  };

  const handleUserSelect = (userId: number) => {
    setSelectedUserIds((prev) =>
      prev.includes(userId)
        ? prev.filter((id) => id !== userId)
        : [...prev, userId],
    );
  };

  const handleDeleteSubgroup = async (subgroupId: number) => {
    try {
      const response = await fetch(
        `/api/group/subgroups/${groupId}/${session?.user?.id}`,
        {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            subgroup_id: subgroupId,
          }),
        },
      );

      if (!response.ok) {
        throw new Error("Failed to delete subgroup");
      }

      setSubgroups((prevSubgroups) =>
        prevSubgroups.filter((subgroup) => subgroup.id !== subgroupId),
      );
      setShowDeleteConfirmation(null);
    } catch (error) {
      console.error("Error deleting subgroup:", error);
    }
  };

  if (renderState.status === "loading") {
    return <LoadingScreen />;
  }

  if (renderState.status === "error") {
    return (
      <div className="text-center mt-10">
        Failed to load subgroups. Please try again later.
      </div>
    );
  }

  return (
    <div className="flex flex-col p-10">
      <h1 className="text-3xl font-bold mb-6 text-left">Manage Subgroups</h1>
      <ul className="space-y-4">
        {subgroups.map((subgroup) => (
          <li
            key={subgroup.id}
            className="p-4 border border-gray-300 rounded shadow"
          >
            <div
              className="flex justify-between items-center cursor-pointer"
              onClick={() => handleGroupExpand(subgroup.id)}
            >
              <div className="flex flex-col">
                <h2 className="text-lg font-semibold">{subgroup.name}</h2>
              </div>
              <button
                onClick={(e) => {
                  e.stopPropagation();
                  setShowDeleteConfirmation(subgroup.id);
                }}
                className="text-gray-500 hover:text-red-600 text-xl ml-2"
              >
                🗑️
              </button>
            </div>
            {expandedGroup === subgroup.id && (
              <div className="mt-4 border-t border-gray-200 pt-4">
                <h3 className="text-md font-bold mb-2">Participants</h3>
                <div className="max-h-96 overflow-y-auto">
                  <ul className="space-y-2">
                    {subgroup.members
                      .map((memberId) =>
                        availableUsers.find((user) => user.id === memberId),
                      )
                      .filter((user): user is User => user !== undefined)
                      .map((member) => (
                        <li
                          key={member.id}
                          className="flex justify-between items-center p-2"
                        >
                          <div>
                            <p className="font-medium">
                              {member.first_name} {member.last_name}
                            </p>
                            <p className="text-sm text-gray-500">
                              {member.email}
                            </p>
                          </div>
                          <button
                            onClick={() =>
                              handleRemoveUser(subgroup.id, member.id)
                            }
                            className="px-2 py-1 bg-red-600 text-white rounded hover:bg-red-500"
                          >
                            Remove
                          </button>
                        </li>
                      ))}
                  </ul>
                </div>
                <div className="relative mt-4">
                  <button
                    onClick={() =>
                      setShowUserDropdown((prev) =>
                        prev === subgroup.id ? null : subgroup.id,
                      )
                    }
                    className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
                  >
                    <span>
                      {selectedUserIds.length > 0
                        ? `Selected users: ${selectedUserIds.length}`
                        : "Select users"}
                    </span>
                    {showUserDropdown === subgroup.id ? (
                      <ChevronUp className="h-5 w-5" />
                    ) : (
                      <ChevronDown className="h-5 w-5" />
                    )}
                  </button>
                  {showUserDropdown === subgroup.id && (
                    <div className="absolute top-full left-0 right-0 mt-1 bg-background border border-customGray rounded shadow-lg z-10 max-h-60 overflow-y-auto">
                      {(() => {
                        const filteredUsers = availableUsers.filter(
                          (user) => !subgroup.members.includes(user.id),
                        );

                        if (filteredUsers.length > 0) {
                          return filteredUsers.map((user) => {
                            const isSelected = selectedUserIds.includes(
                              user.id,
                            );
                            return (
                              <button
                                key={user.id}
                                onClick={() => handleUserSelect(user.id)}
                                className="w-full px-4 py-2 flex items-center justify-between hover:bg-headerHoverGray transition"
                              >
                                <span className="truncate">
                                  {user.first_name} {user.last_name}
                                </span>
                                <div
                                  className={`w-5 h-5 border rounded flex items-center justify-center 
                  ${isSelected ? "bg-blue-500 border-blue-500" : "border-customGray"}`}
                                >
                                  {isSelected && (
                                    <Check className="h-4 w-4 text-white" />
                                  )}
                                </div>
                              </button>
                            );
                          });
                        }

                        return (
                          <p className="p-4 text-center text-gray-500">
                            No users available to add.
                          </p>
                        );
                      })()}
                    </div>
                  )}
                </div>
                {selectedUserIds.length > 0 && (
                  <div className="mt-2 flex justify-end space-x-2">
                    <button
                      className="px-3 py-1.5 bg-gray-100 text-gray-700 rounded hover:bg-gray-200"
                      onClick={() => {
                        setSelectedUserIds([]);
                        setShowUserDropdown(null);
                      }}
                    >
                      Cancel
                    </button>
                    <button
                      className="px-3 py-1.5 bg-blue-600 text-white rounded hover:bg-blue-700"
                      onClick={() => handleAddSelectedUsers(subgroup.id)}
                    >
                      Add Selected
                    </button>
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
                onClick={() => handleDeleteSubgroup(showDeleteConfirmation!)}
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
