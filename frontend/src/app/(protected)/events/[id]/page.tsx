"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect, use } from "react";
import { useGroup } from "@/src/app/contexts/GroupContext";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import {
  Calendar,
  MapPin,
  Download,
  ChevronLeft,
  ChevronUp,
  ChevronDown,
  Music,
} from "lucide-react";

type Event = {
  id: number;
  title: string;
  description: string;
  location: string;
  date: string;
  tracks: Track[];
};

type Track = {
  id: number;
  name: string;
  description: string;
  notesheets: Notesheet[];
};

type Notesheet = {
  id: number;
  track_id: number;
  filepath: string;
};

export default function EventDetailsPage({
  params,
}: {
  params: { id: string };
}) {
  const { id } = use(params);
  const router = useRouter();
  const { data: session } = useSession();
  const [event, setEvent] = useState<Event | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [expandedTracks, setExpandedTracks] = useState<Record<number, boolean>>(
    {},
  );

  const toggleTrack = (trackId: number) => {
    setExpandedTracks((prev) => ({
      ...prev,
      [trackId]: !prev[trackId],
    }));
  };

  useEffect(() => {
    const fetchEventDetails = async () => {
      try {
        const eventResponse = await fetch(
          `/api/event/info/${id}/${session?.user?.id}`,
        );
        if (!eventResponse.ok) throw new Error("Failed to fetch event");
        const eventData = await eventResponse.json();
        setEvent(eventData);
      } catch (error) {
        console.error("Error fetching event details:", error);
      } finally {
        setIsLoading(false);
      }
    };

    if (session?.user?.id) {
      fetchEventDetails();
    }
  }, [id, session?.user?.id]);

  const handleDownload = async (notesheetId: number, fileName: string) => {
    try {
      const response = await fetch(
        `/api/track/notesheet/file/${notesheetId}/${session?.user?.id}`,
        { method: "GET" },
      );

      if (!response.ok) throw new Error("Failed to download file");

      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = url;
      a.download = fileName;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (error) {
      console.error("Error downloading file:", error);
    }
  };

  if (isLoading) return <LoadingScreen />;
  if (!event) return <div>Event not found</div>;

  return (
    <RequireGroup>
      <div className="max-w-4xl mx-auto p-6">
        <button
          onClick={() => router.push("/events")}
          className="flex items-center text-gray-400 hover:text-gray-300 mb-6"
        >
          <ChevronLeft className="h-5 w-5 mr-1" />
          Back to Events
        </button>

        <div className="text-white bg-background border border-customGray rounded-lg p-6 mb-8">
          <h1 className="text-3xl font-bold mb-4">{event.title}</h1>

          <div className="space-y-4">
            <div className="flex items-center text-gray-300">
              <Calendar className="h-5 w-5 mr-2" />
              {new Date(event.date).toLocaleDateString()} at{" "}
              {new Date(event.date).toLocaleTimeString()}
            </div>

            <div className="flex items-center text-gray-300">
              <MapPin className="h-5 w-5 mr-2" />
              {event.location}
            </div>

            {event.description && (
              <p className="text-gray-300 mt-4">{event.description}</p>
            )}
          </div>
        </div>

        <h2 className="text-white text-2xl font-semibold mb-4">Tracks</h2>
        <div className="space-y-4">
          {event.tracks.map((track) => {
            const trackNotesheets = track.notesheets;

            return (
              <div
                key={track.id}
                className="bg-background border border-customGray rounded-lg overflow-hidden"
              >
                <button
                  onClick={() => toggleTrack(track.id)}
                  className="w-full p-4 flex items-center justify-between text-left hover:bg-headerHoverGray transition-colors"
                >
                  <div className="flex items-center">
                    <Music className="h-5 w-5 mr-2" />
                    <div>
                      <h3 className="text-lg font-medium">{track.name}</h3>
                      {track.description && (
                        <p className="text-gray-400 text-sm">
                          {track.description}
                        </p>
                      )}
                    </div>
                  </div>
                  {expandedTracks[track.id] ? (
                    <ChevronUp className="h-5 w-5 text-gray-400" />
                  ) : (
                    <ChevronDown className="h-5 w-5 text-gray-400" />
                  )}
                </button>

                {expandedTracks[track.id] && (
                  <div className="border-t border-customGray p-4">
                    {trackNotesheets.length > 0 ? (
                      <div className="space-y-2">
                        {trackNotesheets.map((notesheet) => (
                          <div
                            key={notesheet.id}
                            className="flex items-center justify-between bg-background hover:bg-headerHoverGray p-2 rounded"
                          >
                            <span className="text-sm text-gray-300">
                              {notesheet.filepath}
                            </span>
                            <button
                              onClick={() =>
                                handleDownload(notesheet.id, notesheet.filepath)
                              }
                              className="flex items-center text-blue-500 hover:text-blue-400"
                            >
                              <Download className="h-4 w-4" />
                            </button>
                          </div>
                        ))}
                      </div>
                    ) : (
                      <div className="text-center text-gray-400 py-2">
                        You do not have access to any notesheets for this track.
                      </div>
                    )}
                  </div>
                )}
              </div>
            );
          })}
        </div>
      </div>
    </RequireGroup>
  );
}
