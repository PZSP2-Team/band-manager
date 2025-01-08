"use client";

import { useEffect, useState } from "react";
import { useRouter, useParams } from "next/navigation";
import { useSession } from "next-auth/react"; // Для проверки роли пользователя
import LoadingScreen from "@/src/app/components/LoadingScreen";

type Event = {
  id: number;
  name: string;
  date: string;
  type: "concert" | "rehearsal";
  time: string;
  materials: { name: string; subgroups: string[] }[];
};

export default function EditEventPage() {
  const { id } = useParams();
  const router = useRouter();
  const { data: session } = useSession(); // Получаем сессию для проверки роли
  const [event, setEvent] = useState<Event | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  // Состояния для редактируемых полей
  const [name, setName] = useState("");
  const [date, setDate] = useState("");
  const [time, setTime] = useState("");
  const [type, setType] = useState<"concert" | "rehearsal">("concert");
  const [materials, setMaterials] = useState<{ name: string; subgroups: string[] }[]>([]);

  useEffect(() => {
    const fetchEvent = async () => {
      try {
        // Имитация получения данных о мероприятии
        const mockEvents: Event[] = [
          {
            id: 1,
            name: "Rock Festival",
            date: "2025-01-15",
            type: "concert",
            time: "18:00",
            materials: [
              { name: "К Элизе", subgroups: ["Guitar", "Drums"] },
              { name: "Ляляля", subgroups: ["Vocals", "Keyboard"] },
              { name: "4 времени года - Зима", subgroups: ["Violins", "Cello"] },
            ],
          },
          {
            id: 2,
            name: "Jazz Night",
            date: "2025-01-20",
            type: "concert",
            time: "20:00",
            materials: [
              { name: "Saxophone", subgroups: ["Solo", "Bass"] },
              { name: "Piano", subgroups: ["Chords", "Strings"] },
            ],
          },
          {
            id: 3,
            name: "Classical Evening",
            date: "2025-02-01",
            type: "rehearsal",
            time: "15:00",
            materials: [
              { name: "Sheet Music", subgroups: ["Violins", "Conductor"] },
              { name: "Conductor's Baton", subgroups: ["Percussion", "Strings"] },
            ],
          },
        ];

        const fetchedEvent = mockEvents.find((e) => e.id === Number(id));
        if (fetchedEvent) {
          setEvent(fetchedEvent);
          setName(fetchedEvent.name);
          setDate(fetchedEvent.date);
          setTime(fetchedEvent.time);
          setType(fetchedEvent.type);
          setMaterials(fetchedEvent.materials);
        } else {
          setEvent(null);
        }
      } catch (error) {
        console.error("Error fetching event:", error);
      }
      setIsLoading(false);
    };

    fetchEvent();
  }, [id]);

  const handleSave = async () => {
    try {
      console.log("Saving updated event:", {
        name,
        date,
        time,
        type,
        materials,
      });

      await new Promise((resolve) => setTimeout(resolve, 1000));

      alert("Event successfully updated!");
      router.push(`/events/${id}`);
    } catch (error) {
      console.error("Error saving event:", error);
      alert("Failed to save event. Please try again.");
    }
  };

  if (isLoading) {
    return <LoadingScreen />;
  }

  if (!event) {
    return <div className="text-center mt-10">Event not found</div>;
  }

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-6">Edit Event</h1>

      {/* Flex container для полей */}
      <div className="flex flex-wrap gap-6">
        {/* Колонка 1: Название и Тип */}
        <div className="flex-1 min-w-[300px]">
          <label className="block text-lg font-semibold mb-2">Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full p-2 border border-gray-600 rounded bg-gray-800 text-white"
          />

          <label className="block text-lg font-semibold mt-4 mb-2">Type</label>
          <select
            value={type}
            onChange={(e) => setType(e.target.value as "concert" | "rehearsal")}
            className="w-full p-2 border border-gray-600 rounded bg-gray-800 text-white"
          >
            <option value="concert">Concert</option>
            <option value="rehearsal">Rehearsal</option>
          </select>
        </div>

        {/* Колонка 2: Дата и Время */}
        <div className="flex-1 min-w-[300px]">
          <label className="block text-lg font-semibold mb-2">Date</label>
          <input
            type="date"
            value={date}
            onChange={(e) => setDate(e.target.value)}
            className="w-full p-2 border border-gray-600 rounded bg-gray-800 text-white"
          />

          <label className="block text-lg font-semibold mt-4 mb-2">Time</label>
          <input
            type="time"
            value={time}
            onChange={(e) => setTime(e.target.value)}
            className="w-full p-2 border border-gray-600 rounded bg-gray-800 text-white"
          />
        </div>
      </div>

      {/* Редактирование материалов */}
      <div className="mb-4 mt-6">
        <h2 className="text-2xl font-semibold mb-4">Materials</h2>
        {materials.map((material, idx) => (
          <div key={idx} className="mb-4 p-4 border border-gray-600 rounded bg-gray-800">
            <input
              type="text"
              value={material.name}
              onChange={(e) => {
                const updatedMaterials = [...materials];
                updatedMaterials[idx].name = e.target.value;
                setMaterials(updatedMaterials);
              }}
              className="w-full p-2 border border-gray-600 rounded bg-gray-800 text-white mb-2"
            />
            <div>
              <h3 className="font-semibold text-white">Subgroups</h3>
              {material.subgroups.map((subgroup, subIdx) => (
                <input
                  key={subIdx}
                  type="text"
                  value={subgroup}
                  onChange={(e) => {
                    const updatedMaterials = [...materials];
                    updatedMaterials[idx].subgroups[subIdx] = e.target.value;
                    setMaterials(updatedMaterials);
                  }}
                  className="w-full p-2 border border-gray-600 rounded bg-gray-800 text-white mb-2"
                />
              ))}
              <button
                className="mt-2 px-4 py-2 bg-blue-600 text-white rounded"
                onClick={() => {
                  const updatedMaterials = [...materials];
                  updatedMaterials[idx].subgroups.push("");
                  setMaterials(updatedMaterials);
                }}
              >
                Add Subgroup
              </button>
            </div>
          </div>
        ))}
      </div>

      {/* Кнопка сохранения */}
      <button
        className="mt-6 px-6 py-2 bg-green-600 text-white rounded-lg shadow hover:bg-green-500 transition"
        onClick={handleSave}
      >
        Save Changes
      </button>
    </div>
  );
}
