"use client";
import { useRouter as useNavigationRouter, useParams } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import {
  ChevronLeft,
  MessageSquare,
  Calendar,
  AlertCircle,
} from "lucide-react";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Announcement = {
  id: number;
  title: string;
  description: string;
  created_at: string;
  sender: {
    first_name: string;
    last_name: string;
  };
  priority: number;
};

export default function AnnouncementDetailsPage() {
  const params = useParams();
  const navRouter = useNavigationRouter();
  const id = params.id;
  const { data: session, status: sessionStatus } = useSession();
  const [announcement, setAnnouncement] = useState<Announcement>({
    id: -1,
    title: "",
    description: "",
    created_at: "",
    sender: { first_name: "", last_name: "" },
    priority: -1,
  });
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });

  useEffect(() => {
    if (sessionStatus === "loading") return;
    const fetchAnnouncementDetails = async () => {
      try {
        const response = await fetch(
          `/api/announcement/user/${session?.user?.id}`,
        );
        if (!response.ok) throw new Error("Failed to fetch announcement");
        const data = await response.json();
        const announcementData = data.announcements.filter(
          (announcement: Announcement) =>
            announcement.id === parseInt(id as string),
        )[0];
        setAnnouncement(announcementData);
      } catch (error) {
        console.error("Error fetching announcement details:", error);
        setRenderState({ status: "error" });
      } finally {
        setRenderState({ status: "loaded" });
      }
    };

    fetchAnnouncementDetails();
  }, [id, sessionStatus, session?.user?.id]);

  const getPriorityLabel = (priority: number) => {
    switch (priority) {
      case 0:
        return "Low Priority";
      case 1:
        return "Medium Priority";
      case 2:
        return "High Priority";
      default:
        return "Unknown Priority";
    }
  };

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

  if (renderState.status === "loading") return <LoadingScreen />;
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
      <div className="max-w-4xl mx-auto p-6">
        <button
          onClick={() => navRouter.push("/announcements")}
          className="flex items-center text-gray-400 hover:text-gray-300 mb-6"
        >
          <ChevronLeft className="h-5 w-5 mr-1" />
          Back to Announcements
        </button>

        <div className="text-white bg-background rounded-lg p-6">
          <div className="flex items-start justify-between mb-4">
            <h1 className="text-3xl font-bold">{announcement.title}</h1>
            <div
              className={`flex items-center ${getPriorityColor(announcement.priority)}`}
            >
              <AlertCircle className="h-5 w-5 mr-2" />
              {getPriorityLabel(announcement.priority)}
            </div>
          </div>

          <div className="space-y-4">
            <div className="flex items-center text-gray-300">
              <MessageSquare className="h-5 w-5 mr-2" />
              Posted by {announcement.sender.first_name}{" "}
              {announcement.sender.last_name}
            </div>

            <div className="flex items-center text-gray-300">
              <Calendar className="h-5 w-5 mr-2" />
              {new Date(announcement.created_at).toLocaleDateString()} at{" "}
              {new Date(announcement.created_at).toLocaleTimeString()}
            </div>

            <div className="mt-6 text-gray-300 whitespace-pre-wrap">
              {announcement.description}
            </div>
          </div>
        </div>
      </div>
    </RequireGroup>
  );
}
