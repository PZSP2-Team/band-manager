"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { RequireManager } from "../../components/RequireManager";
import { Calendar, List, X } from "lucide-react";
import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";

// import "@fullcalendar/core/main.css";
// import "@fullcalendar/daygrid/main.css";

const RenderState = {
  LOADING: "loading",
  LOADED: "loaded",
  ERROR: "error",
};

type Event = {
  id: string;
  title: string;
  location: string;
  description: string;
  date: string;
};

export default function EventsPage() {
  const { groupId, userRole } = useGroup();
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState(RenderState.LOADING);
  const [events, setEvents] = useState<Event[]>([]);
  const [viewMode, setViewMode] = useState("list");

  useEffect(() => {
    if (sessionStatus === "loading") return;

    const fetchEvents = async () => {
      setRenderState(RenderState.LOADING);
      try {
        const response = await fetch(`/api/event/user/${session?.user?.id}`);
        if (!response.ok) {
          throw new Error("Failed to fetch events");
        }
        const data = await response.json();
        const filteredEvents = data.events
        .filter((event: Event) => event.group_id === groupId)
        .sort(
          (a: Event, b: Event) =>
            new Date(a.date).getTime() - new Date(b.date).getTime()
        );

        setEvents(filteredEvents);
        setRenderState(RenderState.LOADED);
      } catch (error) {
        console.error("Error fetching events:", error);
        setRenderState(RenderState.ERROR);
      }
    };

    if (groupId && session?.user?.id) {
      fetchEvents();
    }
  }, [sessionStatus, groupId, session?.user?.id]);

  const removeEvent = async (eventId: string) => {
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

  const handleDateClick = (info: { dateStr: string }) => {
    router.push(`/events/create?date=${info.dateStr}`);
  };

  const handleEventClick = (info: { event: { id: string } }) => {
    router.push(`/events/${info.event.id}`);
  };

  if (sessionStatus === "loading" || renderState === RenderState.LOADING) {
    return <LoadingScreen />;
  }

  if (renderState === RenderState.ERROR) {
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
        <div className="flex justify-end space-x-4 mb-6">
          <button
            className={`p-2 rounded-full ${
              viewMode === "list" ? "bg-gray-700" : "hover:bg-gray-600"
            } transition`}
            onClick={() => setViewMode("list")}
          >
            <List className="text-white" />
          </button>
          <button
            className={`p-2 rounded-full ${
              viewMode === "calendar" ? "bg-gray-700" : "hover:bg-gray-600"
            } transition`}
            onClick={() => setViewMode("calendar")}
          >
            <Calendar className="text-white" />
          </button>
        </div>
        {viewMode === "list" ? (
          <ul className="space-y-4 text-left">
            {events.length > 0 ? (
              events.map((event) => (
                <li
                  key={event.id}
                  className="flex flex-row p-4 border border-gray-700 items-center justify-between rounded shadow transition"
                >
                  <div
                    className="text-white flex justify-between items-center cursor-pointer space-x-4"
                    onClick={() => router.push(`/events/${event.id}`)}
                  >
                    <Calendar />
                    <div>
                      <h2 className="font-semibold">{event.title}</h2>
                      <p className="text-sm text-gray-400">
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
              <p className="text-gray-500 text-xl text-center">
                You have no upcoming events.
              </p>
            )}
          </ul>
        ) : (
          <FullCalendar
            plugins={[dayGridPlugin]}
            initialView="dayGridMonth"
            events={events.map((event) => ({
              id: event.id,
              title: event.title,
              date: event.date,
            }))}
            dateClick={handleDateClick}
            eventClick={handleEventClick}
            themeSystem="dark"
            height="auto"
            contentHeight="auto"
            headerToolbar={{
              left: "prev,next today",
              center: "title",
              right: "dayGridMonth,dayGridWeek",
            }}
            buttonText={{
              today: "Today",
              month: "Month",
              week: "Week",
            }}
          />
        )}
      </div>
    </RequireGroup>
  );
}
