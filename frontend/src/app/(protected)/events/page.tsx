"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { Calendar, X } from "lucide-react";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Event = {
  id: number;
  title: string;
  location: string;
  description: string;
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
        const response = await fetch(`/api/event/user/${session?.user?.id}`);
        if (!response.ok) {
          throw new Error("Failed to fetch events");
        }
        const data = await response.json();
        const filteredEvents = data.events
          .filter((event) => event.group_id === groupId)
          .sort(
            (a: Event, b: Event) =>
              new Date(b.date).getTime() - new Date(a.date).getTime(),
          );
        setEvents(filteredEvents);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching events:", error);
        setRenderState({ status: "error" });
      }
    };
    if (groupId) {
      fetchEvents();
    }
  }, [sessionStatus, groupId, session?.user?.id]);

  const removeEvent = async (eventId: number) => {
    try {
      const response = await fetch(`/api/events/${eventId}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          groupId: groupId,
          userId: session?.user?.id,
        }),
      });
      if (!response.ok) {
        throw new Error("Failed to delete event");
      }
      setEvents(events.filter((event) => event.id !== eventId));
    } catch (error) {
      console.error("Error deleting event:", error);
    }
  };

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
      <div className="flex flex-col py-10 px-10">
        <div className="flex flex-row items-center justify-between mb-6">
          <h1 className="text-3xl font-bold text-left">Upcoming Events</h1>
          {userRole === "manager" && (
            <button
              className="px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
              onClick={() => router.push("/events/create")}
            >
              Add new event
            </button>
          )}
        </div>
        <ul className="space-y-4 text-left">
          {events.length > 0 ? (
            events.map((event) => (
              <li
                key={event.id}
                className="flex flex-row p-4 border border-customGray items-center justify-between rounded shadow transition"
              >
                <div
                  className="text-white flex justify-between items-center cursor-pointer space-x-4"
                  onClick={() => router.push(`/events/${event.id}`)}
                >
                  <Calendar />
                  <div>
                    <h2 className="font-semibold">{event.title}</h2>
                    <p className="text-sm text-gray-600">
                      {new Date(event.date).toLocaleDateString()} -{" "}
                      {event.location}
                    </p>
                  </div>
                </div>
                <button
                  onClick={() => removeEvent(event.id)}
                  className="p-2 text-red-500 hover:bg-red-100 rounded-full transition"
                >
                  <X className="h-5 w-5" />
                </button>
              </li>
            ))
          ) : (
            <p className="text-customGray text-xl text-center">
              You have no upcoming events.
            </p>
          )}
        </ul>
      </div>
    </RequireGroup>
  );
}
