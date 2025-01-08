"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import LoadingScreen from "@/src/app/components/LoadingScreen";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Event = {
  id: number;
  name: string;
  date: string;
  type: "concert" | "rehearsal";
  time: string;
  materials: string[];
};

export default function EventsPage() {
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({ status: "loading" });
  const [events, setEvents] = useState<Event[]>([]);

  useEffect(() => {
    if (sessionStatus === "loading") return;

    const fetchEvents = async () => {
      try {
        console.log("User role:", session?.user?.role);

        const mockEvents: Event[] = [
          {
            id: 1,
            name: "Rock Festival",
            date: "2025-01-15",
            type: "concert",
            time: "18:00",
            materials: ["Guitar", "Drums", "Microphone"]
          },
          {
            id: 2,
            name: "Jazz Night",
            date: "2025-01-20",
            type: "concert",
            time: "20:00",
            materials: ["Saxophone", "Piano", "Bass"]
          },
          {
            id: 3,
            name: "Classical Evening",
            date: "2025-02-01",
            type: "rehearsal",
            time: "15:00",
            materials: ["Sheet Music", "Violin", "Conductor's Baton"]
          },
        ];

        const filteredEvents = session?.user?.role === "manager"
          ? mockEvents
          : mockEvents.slice(0, 2);

        setTimeout(() => {
          setEvents(filteredEvents);
          setRenderState({ status: "loaded" });
        }, 1000);
      } catch (error) {
        console.error("Error fetching events:", error);
        setRenderState({ status: "error" });
      }
    };

    fetchEvents();
  }, [sessionStatus, session?.user?.role]);

  if (sessionStatus === "loading" || renderState.status === "loading") {
    return <LoadingScreen />;
  }

  if (renderState.status === "error") {
    return <div className="text-center mt-10">Failed to load events. Please try again later.</div>;
  }

  return (
    <div className="flex flex-col mt-10 p-6">
      <div className="w-1/4">
        <h1 className="text-3xl font-bold mb-6 text-left">Upcoming Events</h1>
        <ul className="space-y-4 text-left">
          {events.map((event, index) => (
            <li
              key={event.id}
              onClick={() => router.push(`/events/${event.id}`)}
              className="p-4 border border-gray-300 rounded shadow hover:cursor-pointer hover:bg-white hover:text-black transition"
              style={{ opacity: 0.8 }}
            >
              <h2 className="text-lg font-semibold">
                {index + 1}. {event.name}
              </h2>
              <p className="text-gray-600">{new Date(event.date).toLocaleDateString()}</p>
            </li>
          ))}
        </ul>

        {session?.user?.role === "manager" && (
          <button
            className="mt-6 px-6 py-2 bg-green-600 text-white rounded-lg shadow hover:bg-green-500 transition"
            onClick={() => router.push("/events/create")}
          >
            Create New Event
          </button>
        )}
      </div>
    </div>
  );
}
