"use client";
import { useEffect, useState, useRef } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";
import { useRouter } from "next/navigation";
import { ChevronDown, ChevronUp, Check } from "lucide-react";
import { useSession } from "next-auth/react";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { RequireManager } from "@/src/app/components/RequireManager";
import LoadingScreen from "@/src/app/components/LoadingScreen";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Track = {
  id: number;
  name: string;
};

type User = {
  id: number;
  first_name: string;
  last_name: string;
};

type EventForm = {
  title: string;
  description: string;
  date: string;
  location: string;
  track_ids: number[];
  user_ids: number[];
};

type Subgroup = {
  id: number;
  name: string;
  users: number[];
};

export default function AddEvent() {
  const { groupId } = useGroup();
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();

  const [eventForm, setEventForm] = useState<EventForm>({
    title: "",
    description: "",
    date: "",
    location: "",
    track_ids: [],
    user_ids: [],
  });

  const [tracks, setTracks] = useState<Track[]>([]);
  const [availableUsers, setAvailableUsers] = useState<User[]>([]);
  const [subgroups, setSubgroups] = useState<Subgroup[]>([]);
  const [isTracksDropdownOpen, setIsTracksDropdownOpen] = useState(false);
  const [isSubgroupsDropdownOpen, setIsSubgroupsDropdownOpen] = useState(false);
  const [isUsersDropdownOpen, setIsUsersDropdownOpen] = useState(false);
  const tracksDropdownRef = useRef<HTMLDivElement>(null);
  const usersDropdownRef = useRef<HTMLDivElement>(null);
  const subgroupsDropdownRef = useRef<HTMLDivElement>(null);
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });

  useEffect(() => {
    if (sessionStatus === "loading") return;
    const fetchData = async () => {
      try {
        const tracksResponse = await fetch(
          `/api/track/group/${groupId}/${session?.user?.id}`,
        );
        if (!tracksResponse.ok) {
          throw new Error("Failed to fetch tracks");
        }
        const tracksData = await tracksResponse.json();
        setTracks(tracksData.tracks);

        const usersResponse = await fetch(
          `/api/group/members/${groupId}/${session?.user?.id}`,
        );
        if (!usersResponse.ok) {
          throw new Error("Failed to fetch users");
        }
        const usersData = await usersResponse.json();
        setAvailableUsers(usersData.members);

        const subgroupResponse = await fetch(
          `/api/subgroup/group/${groupId}/${session?.user?.id}`,
        );
        if (!subgroupResponse.ok) {
          throw new Error("Failed to fetch subgroups");
        }
        const subgroupData = await subgroupResponse.json();
        setSubgroups(subgroupData.subgroups);
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

  const isSubgroupSelected = (subgroup: Subgroup) => {
    return subgroup.users.every((userId) =>
      eventForm.user_ids.includes(userId),
    );
  };

  const toggleTrack = (trackId: number) => {
    setEventForm((prev) => ({
      ...prev,
      track_ids: prev.track_ids.includes(trackId)
        ? prev.track_ids.filter((id) => id !== trackId)
        : [...prev.track_ids, trackId],
    }));
  };

  const toggleUser = (userId: number) => {
    setEventForm((prev) => ({
      ...prev,
      user_ids: prev.user_ids.includes(userId)
        ? prev.user_ids.filter((id) => id !== userId)
        : [...prev.user_ids, userId],
    }));
  };

  const toggleSubgroup = (subgroup: Subgroup) => {
    setEventForm((prev) => {
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

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const formattedDate = new Date(eventForm.date).toISOString();
      const response = await fetch("/api/event/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ...eventForm,
          date: formattedDate,
          group_id: groupId,
          user_id: session?.user?.id,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to create event");
      }

      router.push("/events");
    } catch (error) {
      console.error("Error creating event:", error);
    }
  };

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        tracksDropdownRef.current &&
        !tracksDropdownRef.current.contains(event.target as Node)
      ) {
        setIsTracksDropdownOpen(false);
      }

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
          Failed to load events. Please try again later.
        </div>
      </RequireGroup>
    );
  }

  return (
    <RequireGroup>
      <RequireManager>
        <div className="p-6 max-w-4xl mx-auto">
          <h1 className="text-3xl font-bold mb-6">Add new event</h1>

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block font-medium mb-1">
                Event title <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                required
                value={eventForm.title}
                onChange={(e) =>
                  setEventForm((prev) => ({ ...prev, title: e.target.value }))
                }
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block font-medium mb-1">
                Event description <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                required
                value={eventForm.description}
                onChange={(e) =>
                  setEventForm((prev) => ({
                    ...prev,
                    description: e.target.value,
                  }))
                }
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block font-medium mb-1">
                Date <span className="text-red-500">*</span>
              </label>
              <input
                type="datetime-local"
                required
                value={eventForm.date}
                onChange={(e) =>
                  setEventForm((prev) => ({ ...prev, date: e.target.value }))
                }
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block font-medium mb-1">
                Location <span className="text-red-500">*</span>
              </label>
              <input
                type="text"
                required
                value={eventForm.location}
                onChange={(e) =>
                  setEventForm((prev) => ({
                    ...prev,
                    location: e.target.value,
                  }))
                }
                className="w-full px-4 py-2 border bg-background border-customGray rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div className="relative" ref={tracksDropdownRef}>
              <label className="block font-medium mb-1">Select tracks</label>
              <button
                type="button"
                onClick={() => setIsTracksDropdownOpen(!isTracksDropdownOpen)}
                className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
              >
                <span>
                  {eventForm.track_ids.length
                    ? `Selected tracks: ${eventForm.track_ids.length}`
                    : "Select tracks"}
                </span>
                {isTracksDropdownOpen ? (
                  <ChevronUp className="h-5 w-5" />
                ) : (
                  <ChevronDown className="h-5 w-5" />
                )}
              </button>

              {isTracksDropdownOpen && (
                <div className="absolute z-10 w-full mt-1 bg-background border border-customGray rounded shadow-lg max-h-60 overflow-y-auto">
                  {tracks.length > 0 ? (
                    tracks.map((track) => {
                      const isSelected = eventForm.track_ids.includes(track.id);
                      return (
                        <button
                          type="button"
                          key={track.id}
                          onClick={() => toggleTrack(track.id)}
                          className="w-full px-4 py-2 flex items-center justify-between hover:bg-headerHoverGray transition"
                        >
                          <span className="truncate">{track.name}</span>
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
                      No tracks available. Please create a track first.
                    </p>
                  )}
                </div>
              )}
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
                  {subgroups.filter((sg) => isSubgroupSelected(sg)).length
                    ? `Selected subgroups: ${subgroups.filter((sg) => isSubgroupSelected(sg)).length}`
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

            <div className="relative" ref={usersDropdownRef}>
              <label className="block font-medium mb-1">
                Select participants
              </label>
              <button
                type="button"
                onClick={() => setIsUsersDropdownOpen(!isUsersDropdownOpen)}
                className="w-full px-4 py-2 bg-background border border-customGray rounded flex justify-between items-center hover:bg-headerHoverGray"
              >
                <span>
                  {eventForm.user_ids.length
                    ? `Selected participants: ${eventForm.user_ids.length}`
                    : "Select participants"}
                </span>
                {isUsersDropdownOpen ? (
                  <ChevronUp className="h-5 w-5" />
                ) : (
                  <ChevronDown className="h-5 w-5" />
                )}
              </button>

              {isUsersDropdownOpen && (
                <div className="absolute z-10 w-full mt-1 bg-background border border-customGray rounded shadow-lg max-h-60 overflow-y-auto">
                  {availableUsers.length > 0 ? (
                    availableUsers.map((user) => {
                      const isSelected = eventForm.user_ids.includes(user.id);
                      return (
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
                      No available users in this group.
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
                Add event
              </button>
            </div>
          </form>
        </div>
      </RequireManager>
    </RequireGroup>
  );
}
