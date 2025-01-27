"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { Megaphone, X } from "lucide-react";

/**
 * Represents the component's render state
 */
type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

/**
 * Represents an announcement data structure
 */
type Announcement = {
  id: number;
  title: string;
  description: string;
  created_at: string;
  group_id: number;
  priority: number;
  sender: Sender;
};

/**
 * Represents the sender of an announcement
 */
type Sender = {
  first_name: string;
  last_name: string;
};

/**
 * Page component displaying list of announcements for a group.
 * Allows managers to delete announcements and moderators/managers to create new ones.
 * Requires group membership to access.
 */
export default function AnnouncementsPage() {
  const { groupId, userRole } = useGroup();
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [announcements, setAnnouncements] = useState<Announcement[]>([]);

  /**
   * Fetches and filters announcements for the current group.
   * Dependencies: groupId, sessionStatus, session?.user?.id
   *
   * Side effects:
   * - Updates announcements state with fetched and filtered data
   * - Updates renderState based on fetch result
   * - Sorts announcements by creation date (newest first)
   */
  useEffect(() => {
    if (sessionStatus === "loading") return;
    const fetchAnnouncements = async () => {
      setRenderState({ status: "loading" });
      try {
        const response = await fetch(
          `/api/announcement/user/${session?.user?.id}`,
        );
        if (!response.ok) {
          throw new Error("Failed to fetch announcements");
        }
        const data = await response.json();
        const filteredAnnouncements = data.announcements
          .filter(
            (announcement: Announcement) => announcement.group_id === groupId,
          )
          .sort(
            (a: Announcement, b: Announcement) =>
              new Date(b.created_at).getTime() -
              new Date(a.created_at).getTime(),
          );
        setAnnouncements(filteredAnnouncements);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching announcements:", error);
        setRenderState({ status: "error" });
      }
    };
    fetchAnnouncements();
  }, [groupId, sessionStatus, session?.user?.id]);

  /**
   * Deletes an announcement and updates the UI.
   * Only available to managers.
   *
   * Side effect: Removes announcement from the announcements state
   */
  const removeAnnouncement = async (announcementId: number) => {
    try {
      const response = await fetch(
        `/api/announcement/delete/${announcementId}/${session?.user?.id}`,
        {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      if (!response.ok) {
        throw new Error("Failed to delete announcement");
      }
      setAnnouncements(
        announcements.filter(
          (announcement) => announcement.id !== announcementId,
        ),
      );
    } catch (error) {
      console.error("Error deleting announcement:", error);
    }
  };

  /**
   * Returns CSS class for priority-based text color
   * Green for low (0), Yellow for medium (1), Red for high (2)
   */
  const getPriorityColor = (priority: number) => {
    switch (priority) {
      case 0:
        return "text-green-500";
      case 1:
        return "text-yellow-500";
      case 2:
        return "text-red-500";
      default:
        return "text-gray-500";
    }
  };

  if (renderState.status === "loading") {
    return <LoadingScreen />;
  }

  if (renderState.status === "error") {
    return (
      <RequireGroup>
        <div className="text-center mt-10">
          Failed to load announcements. Please try again later.
        </div>
      </RequireGroup>
    );
  }

  const canCreateAnnouncement =
    userRole === "manager" || userRole === "moderator";

  return (
    <RequireGroup>
      <div className="flex flex-col py-10 px-10">
        <div className="flex flex-row items-center justify-between mb-6">
          <h1 className="text-3xl font-bold text-left">Announcements</h1>
          {canCreateAnnouncement && (
            <button
              className="px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
              onClick={() => router.push("/announcements/create")}
            >
              Create new announcement
            </button>
          )}
        </div>
        <ul className="space-y-4 text-left">
          {announcements.length > 0 ? (
            announcements.map((announcement) => (
              <li
                key={announcement.id}
                className="flex flex-row p-4 border border-customGray items-center justify-between rounded shadow transition"
              >
                <div
                  className="text-white flex justify-between items-center cursor-pointer space-x-4"
                  onClick={() =>
                    router.push(`/announcements/${announcement.id}`)
                  }
                >
                  <Megaphone
                    className={`${getPriorityColor(announcement.priority)}`}
                  />
                  <div>
                    <h2 className="font-semibold">{announcement.title}</h2>
                    <p className="text-sm text-gray-600">
                      {new Date(announcement.created_at).toLocaleDateString()} -
                      by {announcement.sender.first_name}{" "}
                      {announcement.sender.last_name}
                    </p>
                  </div>
                </div>
                {userRole === "manager" && (
                  <button
                    onClick={() => removeAnnouncement(announcement.id)}
                    className="p-2 text-red-500 hover:bg-red-100 rounded-full transition"
                  >
                    <X className="h-5 w-5" />
                  </button>
                )}
              </li>
            ))
          ) : (
            <p className="text-customGray text-xl text-center">
              You have no announcements.
            </p>
          )}
        </ul>
      </div>
    </RequireGroup>
  );
}
