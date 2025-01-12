"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useEffect, useState } from "react";
import { ChevronDown, ChevronUp, Check } from "lucide-react";
import { useGroup } from "@/src/app/contexts/GroupContext";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { RequireManager } from "@/src/app/components/RequireManager";

type User = {
  id: number;
  first_name: string;
  last_name: string;
  email: string;
};

export default function CreateSubgroupPage() {
  const router = useRouter();
  const { groupId } = useGroup();
  const { data: session } = useSession();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [availableUsers, setAvailableUsers] = useState<User[]>([]);
  const [selectedUsers, setSelectedUsers] = useState<number[]>([]);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await fetch(
          `/api/group/members/${groupId}/${session?.user?.id}`,
        );

        if (!response.ok) {
          throw new Error("Failed to fetch subgroups");
        }

        const data = await response.json();
        setAvailableUsers(data.members);
      } catch (error) {
        console.error("Error fetching subgroups:", error);
      }
    };

    if (groupId && session?.user?.id) {
      fetchUsers();
    }
  }, [groupId, session?.user?.id]);

  const handleCreateSubgroup = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const subgroupInfoResponse = await fetch("/api/subgroup/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          group_id: groupId,
          name: name,
          description: description,
          user_id: session?.user.id,
        }),
      });

      if (!subgroupInfoResponse.ok) {
        throw new Error("Failed to create subgroup");
      }

      const subgroupInfoData = await subgroupInfoResponse.json();

      const subgroupUsersResponse = await fetch(
        `/api/subgroup/members/add/${subgroupInfoData.id}/${session?.user?.id}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            user_ids: selectedUsers,
          }),
        },
      );

      if (!subgroupUsersResponse.ok) {
        throw new Error("Failed to send user ids");
      }

      router.push("/subgroups");
    } catch (error) {
      console.error("Error creating subgroup:", error);
      alert("Failed to create subgroup. Please try again.");
    }
  };

  const toggleUserSelection = (userId: number) => {
    setSelectedUsers((prev) =>
      prev.includes(userId)
        ? prev.filter((id) => id !== userId)
        : [...prev, userId],
    );
  };

  return (
    <RequireGroup>
      <RequireManager>
        <div className="flex flex-col p-6 max-w-4xl mx-auto">
          <h1 className="text-3xl font-bold mb-6 text-white">
            Create Subgroup
          </h1>

          <form onSubmit={handleCreateSubgroup} className="space-y-4">
            <div>
              <label
                className="block mb-2 font-semibold text-white"
                htmlFor="name"
              >
                Subgroup Name <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                id="name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                className="w-full p-2 border border-customGray rounded text-white bg-background"
                placeholder="Enter subgroup name"
                required
                minLength={1}
              />
            </div>

            <div>
              <label
                className="block mb-2 font-semibold text-white"
                htmlFor="description"
              >
                Description <span className="text-red-500">*</span>
              </label>
              <textarea
                id="description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                className="w-full p-2 border border-customGray rounded text-white bg-background"
                placeholder="Enter subgroup description"
                required
                minLength={1}
              />
            </div>

            <div className="relative">
              <button
                type="button"
                onClick={() => setIsDropdownOpen(!isDropdownOpen)}
                className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
              >
                <span>
                  {selectedUsers.length
                    ? `Selected users: ${selectedUsers.length}`
                    : "Select users"}
                </span>
                {isDropdownOpen ? (
                  <ChevronUp className="h-5 w-5" />
                ) : (
                  <ChevronDown className="h-5 w-5" />
                )}
              </button>

              {isDropdownOpen && (
                <div className="absolute top-full left-0 right-0 mt-1 bg-background border border-customGray rounded shadow-lg z-10 max-h-60 overflow-y-auto">
                  {availableUsers.map((user) => {
                    const isSelected = selectedUsers.includes(user.id);
                    return (
                      <button
                        type="button"
                        key={user.id}
                        onClick={() => toggleUserSelection(user.id)}
                        className="w-full px-4 py-2 flex items-center justify-between hover:bg-headerHoverGray transition"
                      >
                        <span className="truncate">
                          {user.first_name} {user.last_name} ({user.email})
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
                  })}
                </div>
              )}
            </div>

            <button
              type="submit"
              className="w-full mt-6 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-500"
            >
              Create Subgroup
            </button>
          </form>
        </div>
      </RequireManager>
    </RequireGroup>
  );
}
