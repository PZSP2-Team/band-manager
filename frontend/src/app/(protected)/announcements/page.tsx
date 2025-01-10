"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
// import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Announcement = {
  id: number;
  title: string;
  description: string;
  date: string;
  priority: number;
};

export default function AnnouncementsPage() {
  // const { userRole } = useGroup();
  const userRole = "manager";
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [announcements, setAnnouncements] = useState<Announcement[]>([]);
  const [showDeleteConfirmation, setShowDeleteConfirmation] = useState<number | null>(null); // ID of announcement to delete

  useEffect(() => {
    if (sessionStatus === "loading") return;

    const fetchAnnouncements = async () => {
      try {
        console.log("User role:", userRole);

        const mockAnnouncements: Announcement[] = [
          {
            id: 1,
            title: "Welcome to the Group!",
            description: "We are thrilled to have you here. Let’s work together!",
            date: "2025-01-10",
            priority: 1,
          },
          {
            id: 2,
            title: "Meeting Schedule Update",
            description: "Please check the updated schedule for the monthly meeting.",
            date: "2025-01-12",
            priority: 2,
          },
          {
            id: 3,
            title: "Project Deadline Reminder",
            description: "Don’t forget to submit your project proposals by the end of the month!",
            date: "2025-01-15",
            priority: 3,
          },
        ];

        // Uncomment the following lines to use the actual API:
        // const response = await fetch(`/api/announcement/group/${groupId}/${userId}`);
        // const data = await response.json();
        // const announcementsFromAPI = data.announcements;

        const filteredAnnouncements =
          userRole === "member"
            ? mockAnnouncements.slice(0, 2) // Show limited announcements for members
            : mockAnnouncements; // Show all announcements for manager or moderator

        setTimeout(() => {
          setAnnouncements(filteredAnnouncements);
          setRenderState({ status: "loaded" });
        }, 1000);
      } catch (error) {
        console.error("Error fetching announcements:", error);
        setRenderState({ status: "error" });
      }
    };

    fetchAnnouncements();
  }, [sessionStatus, userRole]);

  const handleDeleteAnnouncement = async (announcementId: number) => {
    try {
      console.log(`Deleting announcement with ID: ${announcementId}`);
      setAnnouncements((prev) =>
        prev.filter((announcement) => announcement.id !== announcementId)
      );

      // Uncomment to use actual API:
      // const userId = session?.user.id; // Replace with actual user ID
      // const response = await fetch(`/api/announcement/delete/${announcementId}/${userId}`, {
      //   method: "DELETE",
      // });
      // if (!response.ok) {
      //   throw new Error("Failed to delete announcement");
      // }
      // const data = await response.json();
      // console.log(data.message);

      setShowDeleteConfirmation(null); // Close the dialog
    } catch (error) {
      console.error("Error deleting announcement:", error);
      alert("Failed to delete the announcement. Please try again.");
    }
  };

  if (sessionStatus === "loading" || renderState.status === "loading") {
    return <LoadingScreen />;
  }

  if (renderState.status === "error") {
    return (
      <div className="text-center mt-10">
        Failed to load announcements. Please try again later.
      </div>
    );
  }

  return (
    <div className="flex flex-col mt-10 p-6">
      <div className="w-1/4">
        <h1 className="text-3xl font-bold mb-6 text-left">Announcements</h1>
        <ul className="space-y-4 text-left">
          {announcements.map((announcement, index) => (
            <li
              key={announcement.id}
              className="p-4 border border-gray-300 rounded shadow hover:bg-white hover:text-black transition"
              style={{ opacity: 0.8 }}
            >
              <div className="flex justify-between items-center">
                <div>
                  <h2 className="text-lg font-semibold">
                    {index + 1}. {announcement.title}
                  </h2>
                  <p className="text-gray-600 text-sm">
                    {new Date(announcement.date).toLocaleDateString()}
                  </p>
                  <p className="text-gray-400 mt-2">{announcement.description}</p>
                </div>
                {userRole === "manager" && (
                  <button
                    onClick={() => setShowDeleteConfirmation(announcement.id)}
                    className="text-gray-500 hover:text-red-600 text-xl"
                  >
                    🗑️
                  </button>
                )}
              </div>
            </li>
          ))}
        </ul>

        {(userRole === "manager" || userRole === "moderator") && (
          <button
            className="mt-6 px-6 py-2 bg-green-600 text-white rounded-lg shadow hover:bg-green-500 transition"
            onClick={() => router.push("/announcements/create")}
          >
            Create New Announcement
          </button>
        )}
      </div>

      {showDeleteConfirmation !== null && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
          <div className="bg-white p-6 rounded shadow">
            <p className="text-lg font-bold text-gray-800 mb-4">
              Are you sure you want to delete this announcement?
            </p>
            <div className="flex justify-end space-x-4">
              <button
                onClick={() => setShowDeleteConfirmation(null)}
                className="px-4 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-200"
              >
                No
              </button>
              <button
                onClick={() => handleDeleteAnnouncement(showDeleteConfirmation)}
                className="px-4 py-2 bg-red-600 text-white rounded hover:bg-red-500"
              >
                Yes
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
