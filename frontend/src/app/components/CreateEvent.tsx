"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { PlusCircle } from "lucide-react";

type Song = {
  name: string;
  subgroups: string[];
};

export default function CreateEvent() {
  const router = useRouter();
  const [eventName, setEventName] = useState("");
  const [eventDate, setEventDate] = useState("");
  const [eventType, setEventType] = useState("concert");
  const [eventTime, setEventTime] = useState("");
  const [songs, setSongs] = useState<Song[]>([]);

  const addSong = () => {
    setSongs([...songs, { name: "", subgroups: [] }]);
  };

  const updateSong = <K extends keyof Song>(
    index: number,
    key: K,
    value: Song[K],
  ) => {
    const updatedSongs = [...songs];
    updatedSongs[index][key] = value;
    setSongs(updatedSongs);
  };

  const addSubgroupToSong = (index: number) => {
    const updatedSongs = [...songs];
    updatedSongs[index].subgroups.push("");
    setSongs(updatedSongs);
  };

  const updateSubgroup = (
    songIndex: number,
    subgroupIndex: number,
    value: string,
  ) => {
    const updatedSongs = [...songs];
    updatedSongs[songIndex].subgroups[subgroupIndex] = value;
    setSongs(updatedSongs);
  };

  const removeSong = (index: number) => {
    const updatedSongs = songs.filter((_, i) => i !== index);
    setSongs(updatedSongs);
  };

  const handleSubmit = () => {
    const newEvent = {
      name: eventName,
      date: eventDate,
      type: eventType,
      time: eventTime,
      songs,
    };

    console.log("Event created:", newEvent);

    // TODO: Add backend integration for saving the event
    router.push("/events"); // Redirect to events page after creation
  };

  return (
    <div className="p-6 max-w-4xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">Create New Event</h1>

      {/* Event Details */}
      <div className="space-y-4">
        <div>
          <label className="block font-medium mb-1">Event Name</label>
          <input
            type="text"
            value={eventName}
            onChange={(e) => setEventName(e.target.value)}
            className="w-full px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label className="block font-medium mb-1">Event Date</label>
          <input
            type="date"
            value={eventDate}
            onChange={(e) => setEventDate(e.target.value)}
            className="w-full px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label className="block font-medium mb-1">Event Time</label>
          <input
            type="time"
            value={eventTime}
            onChange={(e) => setEventTime(e.target.value)}
            className="w-full px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label className="block font-medium mb-1">Event Type</label>
          <select
            value={eventType}
            onChange={(e) => setEventType(e.target.value)}
            className="w-full px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          >
            <option value="concert">Concert</option>
            <option value="rehearsal">Rehearsal</option>
          </select>
        </div>
      </div>

      {/* Songs Section */}
      <div className="mt-8">
        <h2 className="text-2xl font-semibold mb-4">Songs</h2>
        {songs.map((song, index) => (
          <div
            key={index}
            className="p-4 mb-4 border border-gray-300 rounded bg-gray-100"
          >
            <div className="flex justify-between items-center">
              <input
                type="text"
                placeholder="Song Name"
                value={song.name}
                onChange={(e) => updateSong(index, "name", e.target.value)}
                className="w-3/4 px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <button
                onClick={() => removeSong(index)}
                className="text-red-500 hover:underline ml-4"
              >
                Remove
              </button>
            </div>
            <div className="mt-4">
              <h3 className="font-medium mb-2">Subgroups</h3>
              {song.subgroups.map((subgroup, subgroupIndex) => (
                <div
                  key={subgroupIndex}
                  className="flex items-center gap-2 mb-2"
                >
                  <input
                    type="text"
                    placeholder="Subgroup Name"
                    value={subgroup}
                    onChange={(e) =>
                      updateSubgroup(index, subgroupIndex, e.target.value)
                    }
                    className="w-full px-4 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                </div>
              ))}
              <button
                onClick={() => addSubgroupToSong(index)}
                className="mt-2 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition"
              >
                Add Subgroup
              </button>
            </div>
          </div>
        ))}
        <button
          onClick={addSong}
          className="flex items-center gap-2 mt-4 px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600 transition"
        >
          <PlusCircle className="h-5 w-5" />
          Add Song
        </button>
      </div>

      {/* Submit Button */}
      <div className="mt-8">
        <button
          onClick={handleSubmit}
          className="px-6 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 transition"
        >
          Create Event
        </button>
      </div>
    </div>
  );
}
