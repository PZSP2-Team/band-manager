"use client";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";
import { UserRoundPlus, UsersRound } from "lucide-react";

export default function GroupPage() {
  const { data: session } = useSession();
  const [group, setGroup] = useState<{ id: number; name: string; description: string } | null>(null);
  const [showJoinModal, setShowJoinModal] = useState(false);
  const [joinCode, setJoinCode] = useState("");
  const [error, setError] = useState(false);
  const [success, setSuccess] = useState(false);

  const TEST_JOIN_CODE = process.env.NEXT_PUBLIC_TEST_JOIN_CODE; // Accessing the code from .env

  useEffect(() => {
    if (session?.user?.groupId) {
      setGroup({
        id: session.user.groupId,
        name: "testgroup",
        description: "test group description",
      });
    }
  }, [session]);

  const handleJoinGroup = () => {
    if (joinCode === TEST_JOIN_CODE) {
      setSuccess(true);
      setError(false);
    } else {
      setError(true);
      setJoinCode("");
    }
  };

  if (group) {
    return (
      <div className="max-w-md mx-auto mt-10 p-6">
        <h1 className="text-2xl font-bold mb-6">You belong to: {group.name}</h1>
        <p className="text-gray-600">{group.description}</p>
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
          onClick={() => setShowJoinModal(true)}
          className="flex flex-col items-center justify-center w-full p-4 border border-customGray text-customGray rounded hover:bg-hoverGray group"
        >
          <UserRoundPlus size={48} className="mb-4 transform transition-transform group-hover:-translate-y-4" />
          <span className="text-lg font-semibold mb-2">Join group with code</span>
          <p className="text-sm text-center">Use an invite code to join an existing group</p>
        </button>
      </div>

      {showJoinModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white p-6 rounded-md shadow-lg w-96 relative">
            <button
              onClick={() => setShowJoinModal(false)}
              className="absolute top-2 right-2 text-gray-500 hover:text-gray-800"
            >
              ✕
            </button>
            {!success ? (
              <>
                <h2 className="text-xl font-bold mb-4 text-center">Join group with code</h2>
                <p className="text-gray-600 mb-4 text-center">Paste your code</p>
                <input
                  type="text"
                  value={joinCode}
                  onChange={(e) => setJoinCode(e.target.value)}
                  placeholder="Enter group code"
                  className="w-full px-4 py-2 mb-4 border border-gray-300 rounded-md"
                />
                <button
                  onClick={handleJoinGroup}
                  className="w-full py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
                >
                  Join
                </button>
                {error && <p className="mt-2 text-red-500 text-center">Try again or check your code</p>}
              </>
            ) : (
              <div className="text-center">
                <h2 className="text-xl font-bold mb-4">Successfully joined!</h2>
                <span className="text-4xl">🎉</span>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
