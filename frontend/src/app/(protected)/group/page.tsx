"use client"
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { UserRoundPlus, UsersRound, Info } from "lucide-react";

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
        <div className="min-h-[calc(100vh-74px)] flex flex-col mx-auto py-16 gap-8">
            <div className="flex-grow flex flex-row gap-10 px-40 ">
                <button 
                    onClick={() => {/* create group logic */}}
                    className="flex flex-col items-center justify-center w-full p-4 border border-customGray text-customGray rounded hover:bg-hoverGray group"
                >
                    <UsersRound size={48} className="mb-4 transform transition-transform group-hover:-translate-y-4" />
                    <span className="text-lg font-semibold mb-2">Create new group</span>
                    <p className="text-sm text-center">Start your own group and invite others to join</p>
                </button>
                <button
                    onClick={() => {/* join group logic */}}
                    className="flex flex-col items-center justify-center w-full p-4 border border-customGray text-customGray rounded hover:bg-hoverGray group"
                >
                    <UserRoundPlus size={48} className="mb-4 transform transition-transform group-hover:-translate-y-4" />
                    <span className="text-lg font-semibold mb-2">Join group with code</span>
                    <p className="text-sm text-center">Use an invite code to join an existing group</p>
                </button>
            </div>
        </div>
    );
}
