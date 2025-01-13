"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { useGroup } from "../../contexts/GroupContext";
import LoadingScreen from "@/src/app/components/LoadingScreen";
import { RequireGroup } from "@/src/app/components/RequireGroup";
import { RequireManager } from "../../components/RequireManager";
import { Music4, X } from "lucide-react";

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

type Track = {
  id: number;
  name: string;
  description: string;
};

export default function TracksPage() {
  const { groupId } = useGroup();
  const router = useRouter();
  const { data: session, status: sessionStatus } = useSession();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [tracks, setTracks] = useState<Track[]>([]);

  useEffect(() => {
    if (sessionStatus === "loading") return;
    const fetchTracks = async () => {
      try {
        const response = await fetch(
          `/api/track/group/${groupId}/${session?.user?.id}`,
        );

        if (!response.ok) {
          throw new Error("Failed to fetch events");
        }

        const data = await response.json();
        setTracks(data.tracks);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching events:", error);
        setRenderState({ status: "error" });
      }
    };

    if (groupId) {
      fetchTracks();
    }
  }, [sessionStatus, groupId, session?.user?.id]);

  const removeTrack = async (trackId: number) => {
    try {
      const response = await fetch(
        `/api/track/delete/${trackId}/${session?.user?.id}`,
        {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
        },
      );

      if (!response.ok) {
        throw new Error("Failed to delete track");
      }

      setTracks((prev) => prev.filter((track) => track.id !== trackId));
    } catch (error) {
      console.error("Error deleting track:", error);
    }
  };
  if (renderState.status === "loading") {
    return <LoadingScreen />;
  }

  if (renderState.status === "error") {
    return (
      <RequireGroup>
        <div className="text-center mt-10">
          Failed to load tracks. Please try again later.
        </div>
      </RequireGroup>
    );
  }

  return (
    <RequireGroup>
      <RequireManager>
        <div className="flex flex-col py-10 px-10">
          <div className="flex flex-row items-center justify-between mb-6">
            <h1 className="text-3xl font-bold text-left">Available tracks</h1>
            <button
              className="px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
              onClick={() => router.push("/tracks/add")}
            >
              Add new track
            </button>
          </div>
          <ul className="space-y-4 text-left">
            {tracks.length > 0 ? (
              tracks.map((track) => (
                <li
                  key={track.id}
                  className="flex flex-row p-4 border border-customGray items-center justify-between rounded shadow transition"
                >
                  <div className="text-white flex flex-row space-x-4">
                    <Music4></Music4>
                    <h2 className="font-semibold">{track.name}</h2>
                  </div>
                  <button
                    onClick={() => removeTrack(track.id)}
                    className="p-2 text-red-500 hover:bg-red-100 rounded-full transition"
                  >
                    <X className="h-5 w-5" />
                  </button>
                </li>
              ))
            ) : (
              <p className="text-customGray text-xl text-center ">
                This group does not contain any tracks.
              </p>
            )}
          </ul>
        </div>
      </RequireManager>
    </RequireGroup>
  );
}
