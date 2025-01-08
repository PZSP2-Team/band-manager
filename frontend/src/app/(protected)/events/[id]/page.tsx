"use client";
import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useSession } from "next-auth/react"; // Для получения данных о сессии пользователя
import LoadingScreen from "@/src/app/components/LoadingScreen";

type Event = {
  id: number;
  name: string;
  date: string;
  type: "concert" | "rehearsal";
  time: string;
  materials: string[];
};

export default function EventDetailPage() {
  const { id } = useParams(); // Получаем ID мероприятия из маршрута
  const router = useRouter();
  const { data: session } = useSession(); // Получаем сессию пользователя, включая роль
  const [event, setEvent] = useState<Event | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const mockEvents: Event[] = [
      {
        id: 1,
        name: "Rock Festival",
        date: "2025-01-15",
        type: "concert",
        time: "18:00",
        materials: ["К Элизе", "Ляляля", "4 времени года зима"],
      },
      {
        id: 2,
        name: "Jazz Night",
        date: "2025-01-20",
        type: "concert",
        time: "20:00",
        materials: ["Saxophone", "Piano", "Bass"],
      },
      {
        id: 3,
        name: "Classical Evening",
        date: "2025-02-01",
        type: "rehearsal",
        time: "15:00",
        materials: ["Sheet Music", "Violin", "Conductor's Baton"],
      },
    ];

    const fetchedEvent = mockEvents.find((e) => e.id === Number(id));
    if (fetchedEvent) {
      // Фильтруем материалы в зависимости от роли пользователя
      const filteredMaterials = session?.user?.role === "manager"
        ? fetchedEvent.materials // Менеджер видит все материалы
        : fetchedEvent.materials.slice(0, 2); // Обычный пользователь и модератор видят только первые два материала

      setEvent({
        ...fetchedEvent,
        materials: filteredMaterials, // Применяем фильтрацию материалов
      });
    } else {
      setEvent(null); // Если событие не найдено, устанавливаем event в null
    }

    setIsLoading(false);
  }, [id, session?.user?.role]); // Следим за id события и ролью пользователя

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
        {new Date(event.date).toLocaleDateString()} • {event.time} • {event.type === "concert" ? "Concert" : "Rehearsal"}
      </p>

      {/* Список материалов */}
      <div className="mb-8">
        <h2 className="text-2xl font-semibold uppercase mb-4">Materials</h2>
        <div className="space-y-4">
          {event.materials.map((material, idx) => (
            <div
              key={idx}
              className="p-4 bg-gray-800 border border-gray-600 rounded-lg hover:bg-gray-700 transition"
            >
              <p className="text-lg font-medium text-white">{material}</p>
              {/* Добавить подгруппы и ноты */}
              {session?.user?.role === "manager" && (
                <button
                  className="mt-2 px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-500 transition"
                  onClick={() => alert(`Upload materials for ${material}`)}
                >
                  Upload Notes for Subgroups
                </button>
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Кнопка создания нового события для менеджера */}
      {session?.user?.role === "manager" && (
        <button
          className="mt-6 px-6 py-2 bg-green-600 text-white rounded-lg shadow hover:bg-green-500 transition"
          onClick={() => router.push("/events/create")}
        >
          Create New Event
        </button>
      )}

      {/* Кнопка редактирования события для менеджера */}
      {session?.user?.role === "manager" && (
        <button
          className="mt-6 ml-4 px-6 py-2 bg-blue-600 text-white rounded-lg shadow hover:bg-blue-500 transition"
          onClick={() => router.push(`/events/${id}/edit`)}
        >
          Edit Event
        </button>
      )}
    </div>
  );
}
