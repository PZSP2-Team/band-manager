"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { RequireGroup } from "@/src/app/components/RequireGroup";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Event = {
  id: number;
  title: string;
  description: string;
  location: string;
  date: string;
};

export default function EventsPage() {
  const { groupId, userRole } = useGroup();
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [events, setEvents] = useState<Event[]>([]);

  useEffect(() => {
    if (sessionStatus === "loading") return;

    const fetchEvents = async () => {
      setRenderState({ status: "loading" });

      try {
        const response = await fetch(
          `/api/event/group/${groupId}/${session?.user?.id}`,
        );

        if (!response.ok) {
          throw new Error("Failed to fetch events");
        }

        const data = await response.json();
        setEvents(data.events);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching events:", error);
        setRenderState({ status: "error" });
      }
    };

    if (session?.user?.id) {
      fetchEvents();
    }
  }, [sessionStatus, groupId, session?.user?.id]);
  if (sessionStatus === "loading" || renderState.status === "loading") {
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
      <div className="flex flex-col p-10">
        <div className="flex flex-row items-center justify-between mb-6">
          <h1 className="text-3xl font-bold text-left">Upcoming Events</h1>
          {userRole === "manager" && (
            <button
              className="px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
              onClick={() => router.push("/events/create")}
            >
              Create New Event
            </button>
          )}
        </div>
        <ul className="space-y-4 text-left">
          {events.map((event, index) => (
            <li
              key={event.id}
              onClick={() => router.push(`/events/${event.id}`)}
              className="p-4 border border-gray-300 rounded shadow hover:cursor-pointer hover:bg-white hover:text-black transition"
              style={{ opacity: 0.8 }}
            >
              <h2 className="text-lg font-semibold">
                {index + 1}. {event.title}
              </h2>
              <p className="text-gray-600">
                {new Date(event.date).toLocaleDateString()}
              </p>
            </li>
          ))}
        </ul>
      </div>
    </RequireGroup>
  );
}
