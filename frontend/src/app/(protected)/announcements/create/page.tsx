"use client";
import { useEffect, useState } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";
import { useRouter } from "next/navigation";
import { ChevronDown, ChevronUp, Check } from "lucide-react";
import { useSession } from "next-auth/react";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { RequireManager } from "@/src/app/components/RequireManager";

type Subgroup = {
  id: number;
  name: string;
};

type AnnouncementForm = {
  title: string;
  description: string;
  priority: number;
  subgroup_ids: number[];
};

export default function CreateAnnouncement() {
  const { groupId, userRole } = useGroup();
  const router = useRouter();
  const { data: session } = useSession();

  const [announcementForm, setAnnouncementForm] = useState<AnnouncementForm>({
    title: "",
    description: "",
    priority: 0,
    subgroup_ids: [],
  });

  const [subgroups, setSubgroups] = useState<Subgroup[]>([]);
  const [isSubgroupsDropdownOpen, setIsSubgroupsDropdownOpen] = useState(false);

  useEffect(() => {
    const fetchSubgroups = async () => {
      try {
        const response = await fetch(
          `/api/subgroup/group/${groupId}/${session?.user?.id}`,
        );
        if (!response.ok) {
          throw new Error("Failed to fetch subgroups");
        }
        const data = await response.json();
        setSubgroups(data.subgroups);
      } catch (error) {
        console.error("Error fetching subgroups:", error);
      }
    };

    if (groupId && session?.user?.id) {
      fetchSubgroups();
    }
  }, [groupId, session?.user?.id]);

  const toggleSubgroup = (subgroupId: number) => {
    setAnnouncementForm((prev) => ({
      ...prev,
      subgroup_ids: prev.subgroup_ids.includes(subgroupId)
        ? prev.subgroup_ids.filter((id) => id !== subgroupId)
        : [...prev.subgroup_ids, subgroupId],
    }));
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

  if (userRole !== "manager" && userRole !== "moderator") {
    router.push("/announcements");
    return null;
  }

  return (
    <RequireGroup>
      <RequireManager>
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

            <div className="relative">
              <label className="block font-medium mb-1">Select subgroups</label>
              <button
                type="button"
                onClick={() =>
                  setIsSubgroupsDropdownOpen(!isSubgroupsDropdownOpen)
                }
                className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
              >
                <span>
                  {announcementForm.subgroup_ids.length
                    ? `Selected subgroups: ${announcementForm.subgroup_ids.length}`
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
                    subgroups.map((subgroup) => {
                      const isSelected = announcementForm.subgroup_ids.includes(
                        subgroup.id,
                      );
                      return (
                        <button
                          type="button"
                          key={subgroup.id}
                          onClick={() => toggleSubgroup(subgroup.id)}
                          className="w-full px-4 py-2 flex items-center justify-between hover:bg-headerHoverGray transition"
                        >
                          <span className="truncate">{subgroup.name}</span>
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
                    })
                  ) : (
                    <p className="p-4 text-center text-gray-500">
                      No subgroups available in this group.
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
      </RequireManager>
    </RequireGroup>
  );
}
