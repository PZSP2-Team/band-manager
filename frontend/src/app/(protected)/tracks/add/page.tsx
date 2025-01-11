"use client";
import { useEffect, useState } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";
import { useRouter } from "next/navigation";
import {
  Paperclip,
  PlusCircle,
  ChevronDown,
  ChevronUp,
  X,
  Check,
} from "lucide-react";
import { useSession } from "next-auth/react";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { RequireManager } from "@/src/app/components/RequireManager";

type Notesheet = {
  file?: File;
  subgroup_ids: number[];
  isDropdownOpen?: boolean;
};

type Subgroup = {
  id: number;
  name: string;
};

export default function AddTrack() {
  const { groupId } = useGroup();
  const [subgroups, setSubgroups] = useState<Subgroup[]>([]);
  const [trackTitle, setTrackTitle] = useState("");
  const { data: session } = useSession();
  const [trackDescription, setTrackDescription] = useState("");
  const [notesheets, setNotesheets] = useState<Notesheet[]>([]);
  const router = useRouter();

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

  const handleFileSelect = (index: number, file: File) => {
    setNotesheets(
      notesheets.map((sheet, i) => {
        if (i === index) {
          return {
            ...sheet,
            file: file,
          };
        }
        return sheet;
      }),
    );
  };

  const addNotesheet = () => {
    setNotesheets([
      ...notesheets,
      {
        subgroup_ids: [],
        isDropdownOpen: false,
      },
    ]);
  };

  const toggleDropdown = (index: number) => {
    setNotesheets(
      notesheets.map((sheet, i) =>
        i === index
          ? { ...sheet, isDropdownOpen: !sheet.isDropdownOpen }
          : sheet,
      ),
    );
  };

  const toggleSubgroup = (notesheetIndex: number, subgroupId: number) => {
    setNotesheets(
      notesheets.map((sheet, i) => {
        if (i === notesheetIndex) {
          const newSubgroupIds = sheet.subgroup_ids.includes(subgroupId)
            ? sheet.subgroup_ids.filter((id) => id !== subgroupId)
            : [...sheet.subgroup_ids, subgroupId];
          return {
            ...sheet,
            subgroup_ids: newSubgroupIds,
          };
        }
        return sheet;
      }),
    );
  };

  const removeNotesheet = (index: number) => {
    setNotesheets(notesheets.filter((_, i) => i !== index));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const trackResponse = await fetch("/api/track/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title: trackTitle,
          description: trackDescription,
          group_id: groupId,
          user_id: session?.user?.id,
        }),
      });

      if (!trackResponse.ok) {
        throw new Error("Failed to create track");
      }

      const track = await trackResponse.json();
      const trackId = track.id;

      for (const notesheet of notesheets) {
        if (!notesheet.file) {
          console.log("No file in notesheet", notesheet);
          continue;
        }

        const notesheetFormData = new FormData();
        notesheetFormData.append("file", notesheet.file);
        notesheetFormData.append("track_id", trackId.toString());
        notesheetFormData.append(
          "user_id",
          session?.user?.id?.toString() || "",
        );
        const subgroupIdsString = JSON.stringify(notesheet.subgroup_ids);
        notesheetFormData.append("subgroup_ids", subgroupIdsString);

        const notesheetResponse = await fetch(
          `http://${process.env.NEXT_PUBLIC_BACKEND_HOST}:${process.env.NEXT_PUBLIC_BACKEND_PORT}/api/track/notesheet/create/`,
          {
            method: "POST",
            body: notesheetFormData,
          },
        );
        if (!notesheetResponse.ok) {
          const errorText = await notesheetResponse.text();
          console.error("Error response:", errorText);
          throw new Error("Failed to create notesheet");
        }
      }
      router.push("/tracks");
    } catch (error) {
      console.error("Error creating track:", error);
    }
  };

  return (
    <RequireGroup>
      <RequireManager>
        <div className="p-6 max-w-4xl mx-auto">
          <h1 className="text-3xl font-bold mb-6 text-white">Add new track</h1>

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block font-semibold mb-1">
                Track title <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                required
                value={trackTitle}
                onChange={(e) => setTrackTitle(e.target.value)}
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block font-semibold mb-1">
                Track description <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                required
                value={trackDescription}
                onChange={(e) => setTrackDescription(e.target.value)}
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            {notesheets.map((notesheet, index) => (
              <div
                key={index}
                className="p-4 mt-4 border bg-background border-customGray rounded"
              >
                <div className="flex justify-between items-center mb-2">
                  <div className="flex items-center gap-4">
                    <div className="font-medium">Notesheet {index + 1}</div>
                    <div>
                      <input
                        type="file"
                        id={`file-${index}`}
                        className="hidden"
                        onChange={(e) => {
                          if (e.target.files?.[0]) {
                            handleFileSelect(index, e.target.files[0]);
                          }
                        }}
                      />
                      <label
                        htmlFor={`file-${index}`}
                        className="flex items-center gap-2 px-3 py-1.5 bg-background border border-customGray rounded cursor-pointer hover:bg-headerHoverGray"
                      >
                        <Paperclip className="h-4 w-4" />
                      </label>
                    </div>
                    {notesheet.file?.name && (
                      <span className="text-sm text-customGray">
                        {notesheet.file.name}
                      </span>
                    )}
                  </div>
                  <button
                    type="button"
                    onClick={() => removeNotesheet(index)}
                    className="p-2 text-red-500 hover:bg-red-100 rounded-full transition"
                  >
                    <X className="h-5 w-5" />
                  </button>
                </div>

                <div className="relative">
                  <button
                    type="button"
                    onClick={() => toggleDropdown(index)}
                    className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
                  >
                    <span>
                      {notesheet.subgroup_ids.length
                        ? `Selected subgroups: ${notesheet.subgroup_ids.length}`
                        : "Select subgroups"}
                    </span>
                    {notesheet.isDropdownOpen ? (
                      <ChevronUp className="h-5 w-5" />
                    ) : (
                      <ChevronDown className="h-5 w-5" />
                    )}
                  </button>

                  {notesheet.isDropdownOpen && (
                    <div className="absolute top-full left-0 right-0 mt-1 bg-background border border-customGray rounded shadow-lg z-10 max-h-60 overflow-y-auto">
                      {subgroups.length > 0 ? (
                        subgroups.map((subgroup) => {
                          const isSelected = notesheet.subgroup_ids.includes(
                            subgroup.id,
                          );
                          return (
                            <button
                              type="button"
                              key={subgroup.id}
                              onClick={() => toggleSubgroup(index, subgroup.id)}
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
                          This group does not contain any subgroup. Please
                          create a subgroup first.
                        </p>
                      )}
                    </div>
                  )}
                </div>
              </div>
            ))}

            <button
              type="button"
              onClick={addNotesheet}
              className="flex items-center gap-2 mt-4 px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 transition"
            >
              <PlusCircle className="h-5 w-5" />
              Add Notesheet
            </button>

            <div className="mt-8">
              <button
                type="submit"
                className="w-full px-6 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition"
              >
                Add track
              </button>
            </div>
          </form>
        </div>
      </RequireManager>
    </RequireGroup>
  );
}
