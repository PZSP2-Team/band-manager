"use client";
import { useState, useEffect } from "react";
import LoadingScreen from "@/src/app/components/LoadingScreen";

type RenderState = 
  | { status: "loading" }
  | { status: "loaded_with_group" }
  | { status: "loaded_without_group" };

export default function GroupPage() {
    const { data: session } = useSession();
    const [renderState, setRenderState] = useState<RenderState>({ status: "loading" });
    const [group, setGroup] = useState<{ name: string; description: string, access_token: string} | null>(null);

    useEffect(() => {
        if (sessionStatus === "loading") return;

        const fetchGroupData = async () => {
            if (!session?.user?.groupId) {
                setRenderState({ status: "loaded_without_group" });
            }
            try {
                const response = await fetch(`/api/group/${session?.user.id}`, {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json"
                    }
                });
                if (!response.ok) throw new Error('Failed to fetch group data');

                const groupData = await response.json();
                setRenderState({ status: "loaded_with_group" });
                setGroup(groupData);
            } catch (err) {
                console.error('Error fetching group data:', err);
                setRenderState({ status: "loaded_without_group" });
            }
        };

        fetchGroupData();
    }, [session?.user?.groupId, session?.user?.id, sessionStatus]);


    if (sessionStatus === "loading") {
        return (
            <LoadingScreen/>
        );
    }
    
    switch (renderState.status) {
        case "loading":
            return (
                <LoadingScreen/>
            );
        case "loaded_with_group":
            return (
                <div className="max-w-md mx-auto mt-10 p-6">
                <h1 className="text-2xl font-bold mb-6">You belong to: {group?.name}</h1>
                <p className="text-gray-600">{group?.description}</p>
                {session?.user?.role === "manager" && 
                    <p className="text-gray-600">Access token: {group?.access_token}</p>
                }
                </div>
            );
        case "loaded_without_group":
            return ;
    }
}
