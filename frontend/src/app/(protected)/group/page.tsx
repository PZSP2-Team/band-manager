"use client";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { UserRoundPlus, UsersRound } from "lucide-react";
import LoadingScreen from "@/src/app/components/LoadingScreen";

type RenderState = 
  | { status: "loading" }
  | { status: "loaded_with_group" }
  | { status: "loaded_without_group" };

export default function GroupPage() {
    const { data: session, status: sessionStatus, update } = useSession();
    const [showJoinModal, setShowJoinModal] = useState(false);
    const [renderState, setRenderState] = useState<RenderState>({ status: "loading" });
    const [group, setGroup] = useState<{ name: string; description: string, access_token: string} | null>(null);
    const [showCreateModal, setShowCreateModal] = useState(false);
    const [joinToken, setJoinToken] = useState("");
    const [groupName, setGroupName] = useState("");
    const [groupDescription, setGroupDescription] = useState("");
    const [errorMessage, setErrorMessage] = useState("");
    const [success, setSuccess] = useState(false);

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

    const handleJoinGroup = async () => {
        setErrorMessage("");

        try {
            const response = await fetch("/api/group/join", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    user_id: session?.user?.id,
                    access_token: joinToken
                })
            });
            if (!response.ok) {
                const errorText = await response.text();
                setErrorMessage(errorText);
                return;
            }

            const data = await response.json();
            await update({
                user: {
                    ...session?.user,
                    groupId: data.user_group_id,
                    role: data.user_role
                }
            });
            
            setSuccess(true);
        } catch (err) {
            console.error("Error joining group:", err);
        }
    };

    const handleCreateGroup = async () => {
        if (!groupName || !groupDescription) {
            setErrorMessage("Please fill in all fields");
            return;
        }
        setErrorMessage("");
        
        try {
            const response = await fetch("/api/group/create", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({
                    user_id: session?.user?.id,
                    name: groupName,
                    description: groupDescription
                })
            });

            if (!response.ok) {
                const errorText = await response.text();
                setErrorMessage(errorText);
                return;
            }

            const data = await response.json();
            await update({
                user: {
                    ...session?.user,
                    groupId: data.user_group_id,
                    role: data.user_role
                }
            });
            setSuccess(true);
        } catch (err) {
            console.error("Error creating group:", err);
        }
    };

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
            return (
                <div className="min-h-[calc(100vh-74px)] flex flex-col mx-auto py-16 gap-8">
                <div className="flex-grow flex flex-row gap-10 px-40">
                <button
                onClick={() => setShowCreateModal(true)}
                className="flex flex-col items-center justify-center w-full p-4 border border-customGray text-customGray rounded hover:bg-hoverGray group"
                >
                <UsersRound size={48} className="mb-4 transform transition-transform group-hover:-translate-y-4" />
                <span className="text-lg font-semibold mb-2">Create new group</span>
                <p className="text-sm text-center">Start your own group and invite others to join</p>
                </button>
                <button
                onClick={() => setShowJoinModal(true)}
                className="flex flex-col items-center justify-center w-full p-4 border border-customGray text-customGray rounded hover:bg-hoverGray group"
                >
                <UserRoundPlus size={48} className="mb-4 transform transition-transform group-hover:-translate-y-4" />
                <span className="text-lg font-semibold mb-2">Join group with code</span>
                <p className="text-sm text-center">Use an invite code to join an existing group</p>
                </button>
                </div>

                {showCreateModal && (
                    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
                    <div className="bg-gray-900 text-white p-8 rounded-md shadow-lg w-[28rem] relative">
                    <button
                    onClick={() => {
                        setShowCreateModal(false);
                        setErrorMessage("");
                        setGroupName("");
                        setGroupDescription("");
                    }}
                    className="absolute top-2 right-2 text-customGray hover:text-white"
                    >
                    ✕
                    </button>
                    {!success ? (
                        <>
                        <h2 className="text-2xl font-bold mb-4 text-center">Create New Group</h2>
                        <p className="text-gray-400 mb-4 text-center">Enter group details</p>
                        <form onSubmit={(e) => {
                            e.preventDefault();
                            handleCreateGroup();
                        }}>
                            <input
                                type="text"
                                value={groupName}
                                onChange={(e) => setGroupName(e.target.value)}
                                placeholder="Group name"
                                required
                                className="px-3 py-2 mb-4 block w-full rounded-md bg-gray-700 border border-customGray text-gray-300 focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                            />
                            <textarea
                                value={groupDescription}
                                onChange={(e) => setGroupDescription(e.target.value)}
                                placeholder="Group description"
                                required
                                className="px-3 py-2 mb-4 block w-full rounded-md bg-gray-700 border border-customGray text-gray-300 focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                            />
                            <button
                                type="submit"
                                className="w-full py-2 px-4 rounded bg-cornflowerblue text-white font-bold hover:bg-blue-600 active:bg-blue-700 transition-colors duration-300"
                            >
                                Create
                            </button>
                            {errorMessage && <p className="mt-2 text-red-500 text-center">{errorMessage}</p>}
                        </form>
                        </>
                    ) : (
                    <div className="text-center">
                    <h2 className="text-2xl font-bold mb-4">Group Created Successfully!</h2>
                    <span className="text-4xl">🎉</span>
                    </div>
                    )}
                    </div>
                    </div>
                )}

                {showJoinModal && (
                    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
                    <div className="bg-gray-900 text-white p-8 rounded-md shadow-lg w-[28rem] relative">
                    <button
                    onClick={() => {
                        setShowJoinModal(false);
                        setErrorMessage("");
                        setJoinToken("");
                    }}
                    className="absolute top-2 right-2 text-customGray hover:text-white"
                    >
                    ✕
                    </button>
                    {!success ? (
                        <>
                        <h2 className="text-2xl font-bold mb-4 text-center">Join group with code</h2>
                        <p className="text-gray-400 mb-4 text-center">Paste your code</p>
                        <form onSubmit={(e) => {
                            e.preventDefault();
                            handleJoinGroup();
                        }}>
                            <input
                                type="text"
                                value={joinToken}
                                onChange={(e) => setJoinToken(e.target.value)}
                                placeholder="Enter group code"
                                required
                                className="px-3 py-2 mb-4 block w-full rounded-md bg-gray-700 border border-customGray text-gray-300 focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                            />
                            <button
                                type="submit"
                                className="w-full py-2 px-4 rounded bg-cornflowerblue text-white font-bold hover:bg-blue-600 active:bg-blue-700 transition-colors duration-300"
                            >
                                Join
                            </button>
                            {errorMessage && <p className="mt-2 text-red-500 text-center">Try again or check your code</p>}
                        </form>
                        </>
                    ) : (
                    <div className="text-center">
                    <h2 className="text-2xl font-bold mb-4">Successfully joined!</h2>
                    <span className="text-4xl">🎉</span>
                    </div>
                    )}
                    </div>
                    </div>
                )}
                </div>
            );
    }
}
