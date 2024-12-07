"use client"
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";

export default function GroupPage() {
    const { data: session } = useSession();
    const [group, setGroup] = useState<{id: number, name: string, description: string} | null>(null);

    useEffect(() => {
        if (session?.user?.groupId) {
            //fetch here
            setGroup({
                id: session.user.groupId,
                name: "testgroup", 
                description: "test group description"
            });
        }
    }, [session]);

    if (group) {
        return (
            <div className="max-w-md mx-auto mt-10 p-6">
                <h1 className="text-2xl font-bold mb-6">
                    You belong to: {group.name}
                </h1>
                <p className="text-gray-600">
                    {group.description}
                </p>
            </div>
        );
    }

    return (
        <div className="max-w-md mx-auto mt-10 p-6">
            <h1 className="text-2xl font-bold mb-6">You do not belong to any group.</h1>
            <div className="space-y-4">
                <button 
                    onClick={() => {/* create group logic */}}
                    className="w-full p-4 bg-blue-500 text-white rounded hover:bg-blue-600"
                >
                    Create new group
                </button>
                <button
                    onClick={() => {/* join group logic */}}
                    className="w-full p-4 bg-green-500 text-white rounded hover:bg-green-600"
                >
                    Join group with code
                </button>
            </div>
        </div>
    );
}
