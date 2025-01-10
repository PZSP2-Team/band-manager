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
  title: string;
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

  const removeTrack = async (trackId: number) => {
    try {
      const response = await fetch(`/api/tracks/${trackId}`, {
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
        throw new Error("Failed to delete track");
      }

      setTracks(tracks.filter((track) => track.id !== trackId));
    } catch (error) {
      console.error("Error deleting track:", error);
    }
  };
  useEffect(() => {
    if (sessionStatus === "loading") return;

    const fetchTracks = async () => {
      setRenderState({ status: "loading" });

      try {
        // const response = await fetch(
        //   `/api//group/${groupId}/${session?.user?.id}`,
        // );
        //
        // if (!response.ok) {
        //   throw new Error("Failed to fetch events");
        // }
        //
        // const data = await response.json();
        // setTracks(data.tracks);
        setTracks([
          { id: 1, title: "piosenka o marcinie", description: "asshole" },
          {
            id: 2,
            title: "piosenka",
            description: "asshole",
          },
        ]);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching events:", error);
        setRenderState({ status: "error" });
      }
    };

    if (session?.user?.id) {
      fetchTracks();
    }
  }, [sessionStatus, groupId, session?.user?.id]);
  if (sessionStatus === "loading" || renderState.status === "loading") {
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
          <div className="w-full">
            <h1 className="text-3xl font-bold mb-6 text-left">
              Available tracks
            </h1>
            <ul className="space-y-4 text-left">
              {tracks.map((track, index) => (
                <li
                  key={track.id}
                  className="flex flex-row p-4 border border-customGray items-center justify-between rounded shadow transition"
                  style={{ opacity: 0.8 }}
                >
                  <div className="flex flex-row space-x-4">
                    <Music4></Music4>
                    <h2 className="font-semibold">{track.title}</h2>
                  </div>
                  <button
                    onClick={() => removeTrack(index)}
                    className="p-2 text-red-500 hover:bg-red-100 rounded-full transition"
                  >
                    <X className="h-5 w-5" />
                  </button>
                </li>
              ))}
            </ul>

            <button
              className="mt-6 px-6 py-2 bg-green-600 text-white rounded-lg shadow hover:bg-green-500 transition"
              onClick={() => router.push("/tracks/add")}
            >
              Add new track
            </button>
          </div>
        </div>
      </RequireManager>
    </RequireGroup>
  );
}
