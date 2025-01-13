"use client";
import { useEffect, useState, useRef } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";
import { useRouter } from "next/navigation";
import { ChevronDown, ChevronUp, Check } from "lucide-react";
import { useSession } from "next-auth/react";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { RequireModerator } from "@/src/app/components/RequireModerator";
import LoadingScreen from "@/src/app/components/LoadingScreen";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Subgroup = {
  id: number;
  name: string;
  users: number[];
};

type User = {
  id: number;
  first_name: string;
  last_name: string;
};

type AnnouncementForm = {
  title: string;
  description: string;
  priority: number;
  user_ids: number[];
};

export default function CreateAnnouncement() {
  const { groupId } = useGroup();
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [announcementForm, setAnnouncementForm] = useState<AnnouncementForm>({
    title: "",
    description: "",
    priority: 0,
    user_ids: [],
  });

  const [subgroups, setSubgroups] = useState<Subgroup[]>([]);
  const [users, setUsers] = useState<User[]>([]);
  const [isSubgroupsDropdownOpen, setIsSubgroupsDropdownOpen] = useState(false);
  const [isUsersDropdownOpen, setIsUsersDropdownOpen] = useState(false);
  const usersDropdownRef = useRef<HTMLDivElement>(null);
  const subgroupsDropdownRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (sessionStatus === "loading") return;
    const fetchSubgroups = async () => {
      try {
        const subgroupsResponse = await fetch(
          `/api/subgroup/group/${groupId}/${session?.user?.id}`,
        );
        if (!subgroupsResponse.ok) {
          throw new Error("Failed to fetch subgroups");
        }
        const subgroupsData = await subgroupsResponse.json();
        setSubgroups(subgroupsData.subgroups);

        const usersResponse = await fetch(
          `/api/group/members/${groupId}/${session?.user?.id}`,
        );
        if (!usersResponse.ok) {
          throw new Error("Failed to fetch users");
        }
        const usersData = await usersResponse.json();
        setUsers(usersData.members);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching subgroups:", error);
        setRenderState({ status: "error" });
      }
    };

    if (groupId) {
      fetchSubgroups();
    }
  }, [groupId, sessionStatus, session?.user?.id]);

  const isSubgroupSelected = (subgroup: Subgroup) => {
    return subgroup.users.every((userId) =>
      announcementForm.user_ids.includes(userId),
    );
  };

  const toggleSubgroup = (subgroup: Subgroup) => {
    setAnnouncementForm((prev) => {
      const newUserIds = new Set(prev.user_ids);

      if (isSubgroupSelected(subgroup)) {
        subgroup.users.forEach((userId) => newUserIds.delete(userId));
      } else {
        subgroup.users.forEach((userId) => newUserIds.add(userId));
      }

      return {
        ...prev,
        user_ids: Array.from(newUserIds),
      };
    });
  };

  const toggleUser = (userId: number) => {
    setAnnouncementForm((prev) => {
      const newUserIds = prev.user_ids.includes(userId)
        ? prev.user_ids.filter((id) => id !== userId)
        : [...prev.user_ids, userId];

      return {
        ...prev,
        user_ids: newUserIds,
      };
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await fetch("/api/announcement/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ...announcementForm,
          group_id: groupId,
          sender_id: session?.user?.id,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to create announcement");
      }

      router.push("/announcements");
    } catch (error) {
      console.error("Error creating announcement:", error);
    }
  };

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        usersDropdownRef.current &&
        !usersDropdownRef.current.contains(event.target as Node)
      ) {
        setIsUsersDropdownOpen(false);
      }

      if (
        subgroupsDropdownRef.current &&
        !subgroupsDropdownRef.current.contains(event.target as Node)
      ) {
        setIsSubgroupsDropdownOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  if (renderState.status === "loading") {
    return <LoadingScreen />;
  }

  if (renderState.status === "error") {
    return (
      <RequireGroup>
        <div className="text-center mt-10">
          Failed to data. Please try again later.
        </div>
      </RequireGroup>
    );
  }

  return (
    <RequireGroup>
      <RequireModerator>
        <div className="p-6 max-w-4xl mx-auto">
          <h1 className="text-3xl font-bold mb-6">Create new announcement</h1>

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block font-medium mb-1">
                Title <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                required
                value={announcementForm.title}
                onChange={(e) =>
                  setAnnouncementForm((prev) => ({
                    ...prev,
                    title: e.target.value,
                  }))
                }
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block font-medium mb-1">
                Description <span className="text-red-500">*</span>
              </label>
              <textarea
                required
                value={announcementForm.description}
                onChange={(e) =>
                  setAnnouncementForm((prev) => ({
                    ...prev,
                    description: e.target.value,
                  }))
                }
                rows={4}
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block font-medium mb-1">
                Priority <span className="text-red-500">*</span>
              </label>
              <select
                required
                value={announcementForm.priority}
                onChange={(e) =>
                  setAnnouncementForm((prev) => ({
                    ...prev,
                    priority: parseInt(e.target.value),
                  }))
                }
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value={0}>Low</option>
                <option value={1}>Medium</option>
                <option value={2}>High</option>
              </select>
            </div>

            <div className="relative" ref={subgroupsDropdownRef}>
              <label className="block font-medium mb-1">Select subgroups</label>
              <button
                type="button"
                onClick={() =>
                  setIsSubgroupsDropdownOpen(!isSubgroupsDropdownOpen)
                }
                className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
              >
                <span>
                  {subgroups.filter((subgroup) => isSubgroupSelected(subgroup))
                    .length
                    ? `Selected subgroups: ${subgroups.filter((subgroup) => isSubgroupSelected(subgroup)).length}`
                    : "Select subgroups"}
                </span>
                {isSubgroupsDropdownOpen ? (
                  <ChevronUp className="h-5 w-5" />
                ) : (
                  <ChevronDown className="h-5 w-5" />
                )}
              </button>

              {isSubgroupsDropdownOpen && (
                <div className="absolute z-10 w-full mt-1 bg-background border border-customGray rounded shadow-lg max-h-60 overflow-y-auto">
                  {subgroups.length > 0 ? (
                    subgroups.map((subgroup) => (
                      <button
                        type="button"
                        key={subgroup.id}
                        onClick={() => toggleSubgroup(subgroup)}
                        className="w-full px-4 py-2 flex items-center justify-between hover:bg-headerHoverGray transition"
                      >
                        <span className="truncate">{subgroup.name}</span>
                        <div
                          className={`w-5 h-5 border rounded flex items-center justify-center 
                        ${isSubgroupSelected(subgroup) ? "bg-blue-500 border-blue-500" : "border-customGray"}`}
                        >
                          {isSubgroupSelected(subgroup) && (
                            <Check className="h-4 w-4 text-white" />
                          )}
                        </div>
                      </button>
                    ))
                  ) : (
                    <p className="p-4 text-center text-gray-500">
                      No subgroups available in this group.
                    </p>
                  )}
                </div>
              )}
            </div>
            <div className="relative mt-4" ref={usersDropdownRef}>
              <label className="block font-medium mb-1">Select users</label>
              <button
                type="button"
                onClick={() => setIsUsersDropdownOpen(!isUsersDropdownOpen)}
                className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
              >
                <span>
                  {announcementForm.user_ids.length
                    ? `Selected users: ${announcementForm.user_ids.length}`
                    : "Select users"}
                </span>
                {isUsersDropdownOpen ? (
                  <ChevronUp className="h-5 w-5" />
                ) : (
                  <ChevronDown className="h-5 w-5" />
                )}
              </button>

              {isUsersDropdownOpen && (
                <div className="absolute z-10 w-full mt-1 bg-background border border-customGray rounded shadow-lg max-h-60 overflow-y-auto">
                  {users.length > 0 ? (
                    users.map((user) => (
                      <button
                        type="button"
                        key={user.id}
                        onClick={() => toggleUser(user.id)}
                        className="w-full px-4 py-2 flex items-center justify-between hover:bg-headerHoverGray transition"
                      >
                        <span className="truncate">
                          {user.first_name} {user.last_name}
                        </span>
                        <div
                          className={`w-5 h-5 border rounded flex items-center justify-center 
                        ${announcementForm.user_ids.includes(user.id) ? "bg-blue-500 border-blue-500" : "border-customGray"}`}
                        >
                          {announcementForm.user_ids.includes(user.id) && (
                            <Check className="h-4 w-4 text-white" />
                          )}
                        </div>
                      </button>
                    ))
                  ) : (
                    <p className="p-4 text-center text-gray-500">
                      No users available in this group.
                    </p>
                  )}
                </div>
              )}
            </div>

            <div className="mt-8">
              <button
                type="submit"
                className="w-full px-6 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition"
              >
                Create announcement
              </button>
            </div>
          </form>
        </div>
      </RequireModerator>
    </RequireGroup>
  );
}
