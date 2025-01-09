"use client";
import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import LoadingScreen from "@/src/app/components/LoadingScreen";

type Event = {
  id: number;
  name: string;
  date: string;
  type: "concert" | "rehearsal";
  time: string;
  materials: { name: string; notes: string[] }[];
};

type Track = {
  name: string;
};

export default function EventDetailPage() {
  const { id } = useParams();
  const router = useRouter(); // To navigate to the edit page
  const { data: session } = useSession();
  const [event, setEvent] = useState<Event | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Для добавления новой песни
  const [isAddingSong, setIsAddingSong] = useState(false);
  const [availableTracks, setAvailableTracks] = useState<Track[]>([]);
  const [selectedTrack, setSelectedTrack] = useState("");

  // Управление раскрытием блоков треков
  const [expandedTracks, setExpandedTracks] = useState<Set<string>>(new Set());

  useEffect(() => {
    const mockEvents: Event[] = [
      {
        id: 1,
        name: "Rock Festival",
        date: "2025-01-15",
        type: "concert",
        time: "18:00",
        materials: [
          { name: "Do Elizy", notes: ["Note 1", "Note 2", "Note 3"] },
          { name: "Ляляля", notes: ["Note A", "Note B", "Note C"] },
          { name: "4 времени года зима", notes: ["Sheet 1", "Sheet 2", "Sheet 3"] },
        ],
      },
    ];

    const fetchedEvent = mockEvents.find((e) => e.id === Number(id));
    if (fetchedEvent) {
      setEvent(fetchedEvent);
    } else {
      setEvent(null);
    }

    // Загружаем доступные треки
    const fetchTracks = async () => {
      const mockTracks: Track[] = [
        { name: "New Song 1" },
        { name: "New Song 2" },
        { name: "New Song 3" },
      ];
      setAvailableTracks(mockTracks);
    };

    fetchTracks();
    setIsLoading(false);
  }, [id]);

  const handleAddTrack = () => {
    if (selectedTrack && event) {
      const updatedMaterials = [
        ...event.materials,
        { name: selectedTrack, notes: [] },
      ];
      setEvent({ ...event, materials: updatedMaterials });
      setIsAddingSong(false);
      setSelectedTrack("");
    }
  };

  const handleDownloadNotes = (note: string) => {
    alert(`Downloading notes: ${note}`);
  };

  const toggleTrack = (trackName: string) => {
    setExpandedTracks((prev) => {
      const newSet = new Set(prev);
      if (newSet.has(trackName)) {
        newSet.delete(trackName);
      } else {
        newSet.add(trackName);
      }
      return newSet;
    });
  };

  if (isLoading) {
    return <LoadingScreen />;
  }

  if (!event) {
    return <div className="text-center mt-10">Event not found</div>;
  }

  return (
    <div className="p-6 ml-8 mr-8">
      <h1 className="text-4xl font-bold uppercase mb-4">{event.name}</h1>
      <p className="text-gray-500 text-lg mb-4">
        {new Date(event.date).toLocaleDateString()} • {event.time} •{" "}
        {event.type === "concert" ? "Concert" : "Rehearsal"}
      </p>

      {/* Кнопка Edit Event */}
      {session?.user?.role === "manager" && (
        <button
          className="mb-6 px-6 py-2 bg-yellow-600 text-white rounded-lg shadow hover:bg-yellow-500 transition"
          onClick={() => router.push(`/events/${id}/edit`)}
        >
          Edit Event
        </button>
      )}

      {/* Список материалов */}
      <div className="mb-8">
        <h2 className="text-2xl font-semibold uppercase mb-4">Tracks</h2>
        <div className="space-y-4">
          {event.materials.map((material, idx) => (
            <div
              key={idx}
              className="p-4 bg-gray-800 border border-gray-600 rounded-lg hover:bg-gray-700 transition"
            >
              <div
                className="text-lg font-medium text-white mb-2 cursor-pointer"
                onClick={() => toggleTrack(material.name)}
              >
                {material.name}
              </div>
              {expandedTracks.has(material.name) && (
                <div className="space-y-2">
                  {(session?.user?.role === "manager"
                    ? material.notes
                    : material.notes.slice(0, 1)
                  ).map((note, noteIdx) => (
                    <div
                      key={noteIdx}
                      className="flex justify-between items-center p-2 bg-gray-700 rounded"
                    >
                      <span className="text-white">{note}</span>
                      <button
                        className="text-blue-400 hover:underline"
                        onClick={() => handleDownloadNotes(note)}
                      >
                        Download
                      </button>
                    </div>
                  ))}
                </div>
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Добавление новой песни */}
      {session?.user?.role === "manager" && (
        <div className="mt-6">
          {isAddingSong ? (
            <div className="p-4 bg-gray-800 border border-gray-600 rounded-lg">
              <h3 className="text-lg font-semibold text-white mb-2">
                Add New Song
              </h3>
              <select
                value={selectedTrack}
                onChange={(e) => setSelectedTrack(e.target.value)}
                className="w-full p-2 bg-gray-700 text-white rounded mb-4"
              >
                <option value="">Select a track</option>
                {availableTracks.map((track, idx) => (
                  <option key={idx} value={track.name}>
                    {track.name}
                  </option>
                ))}
              </select>
              <button
                className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-500 transition"
                onClick={handleAddTrack}
              >
                Add Song
              </button>
              <button
                className="ml-4 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-500 transition"
                onClick={() => setIsAddingSong(false)}
              >
                Cancel
              </button>
            </div>
          ) : (
            <button
              className="px-6 py-2 bg-blue-600 text-white rounded-lg shadow hover:bg-blue-500 transition"
              onClick={() => setIsAddingSong(true)}
            >
              Add New Song
            </button>
          )}
        </div>
      )}
    </div>
  );
}
