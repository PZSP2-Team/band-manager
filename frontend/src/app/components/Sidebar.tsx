"use client";
import { useSession } from "next-auth/react";
import { LucideUsers, UserRoundPlus, UsersRound } from "lucide-react";
import Link from "next/link";
import { useEffect, useState } from "react";
import { useGroup } from "../contexts/GroupContext";
import { useRouter } from "next/navigation";

type Group = {
  id: number;
  name: string;
  role: string;
};

export default function Sidebar() {
  const { setGroupId, setUserRole } = useGroup();
  const router = useRouter();
  const [groupList, setGroupList] = useState<Group[]>([]);
  const { data: session } = useSession();
  const [showJoinModal, setShowJoinModal] = useState(false);
  const [showCreateModal, setShowCreateModal] = useState(false);
  const [joinToken, setJoinToken] = useState("");
  const [groupName, setGroupName] = useState("");
  const [groupDescription, setGroupDescription] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [success, setSuccess] = useState(false);

  const handleGroupSelect = (groupId: number, userRole: string) => {
    setGroupId(groupId);
    setUserRole(userRole);
    router.push("/events");
  };

  useEffect(() => {
    const fetchGroups = async () => {
      try {
        const response = await fetch(`/api/group/user/${session?.user?.id}`);

        if (!response.ok) {
          throw new Error("Failed to fetch groups");
        }

        const data = await response.json();
        setGroupList(data.groups);
      } catch (error) {
        console.error("Error fetching groups:", error);
      }
    };

    if (session?.user?.id) {
      fetchGroups();
    }
  }, [session?.user?.id]);

  const handleJoinGroup = async () => {
    setErrorMessage("");

    try {
      const response = await fetch("/api/group/join", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          user_id: session?.user?.id,
          access_token: joinToken,
        }),
      });
      if (!response.ok) {
        const errorText = await response.text();
        setErrorMessage(errorText);
        return;
      }

      const data = await response.json();

      if (groupList) {
        setGroupList((prevGroups) => [...prevGroups, data]);
      } else {
        setGroupList([data]);
      }

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
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          user_id: session?.user?.id,
          name: groupName,
          description: groupDescription,
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        setErrorMessage(errorText);
        return;
      }

      const data = await response.json();
      if (groupList) {
        setGroupList((prevGroups) => [...prevGroups, data]);
      } else {
        setGroupList([data]);
      }
      setSuccess(true);
    } catch (err) {
      console.error("Error creating group:", err);
    }
  };
  return (
    <aside className="h-full w-64 bg-headerGray">
      <div className="h-full flex flex-col p-4">
        <h2 className="text-lg font-semibold text-customGray pb-4 flex-shrink-0">
          My groups
        </h2>
        <div className="flex-1 overflow-y-auto mb-4">
          {groupList?.length > 0 ? (
            <div className="space-y-2">
              {groupList.map((group) => (
                <button
                  key={group.id}
                  onClick={() => handleGroupSelect(group.id, group.role)}
                  className="flex w-full items-center gap-3 p-2 rounded-lg hover:bg-headerHoverGray transition-colors text-customGray"
                >
                  <LucideUsers className="h-5 w-5 flex-shrink-0 min-w-[20px] min-h-[20px]" />
                  <span className="truncate">{group.name}</span>
                </button>
              ))}
            </div>
          ) : (
            <div className="flex flex-col items-center justify-center h-full text-customGray">
              <p className="text-sm">You belong to no group.</p>
            </div>
          )}
        </div>
        <div className="space-y-4 flex-shrink-0">
          <button
            onClick={() => setShowCreateModal(true)}
            className="flex flex-row justify-center bg-sidebarButtonYellow space-x-2 w-full p-4 text-white rounded-lg hover:bg-sidebarButtonHover group"
          >
            <UsersRound
              size={24}
              className="transform transition-transform group-hover:-translate-x-2"
            />
            <span className="flex-1 text-sm font-semibold">
              Create new group
            </span>
          </button>
          <button
            onClick={() => setShowJoinModal(true)}
            className="flex flex-row justify-center bg-sidebarButtonYellow space-x-2 w-full p-4 text-white rounded-lg hover:bg-sidebarButtonHover group"
          >
            <UserRoundPlus
              size={24}
              className="transform transition-transform group-hover:-translate-x-2"
            />
            <span className="flex-1 text-sm font-semibold">
              Join group with code
            </span>
          </button>
        </div>
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
                <h2 className="text-2xl font-bold mb-4 text-center">
                  Create New Group
                </h2>
                <p className="text-gray-400 mb-4 text-center">
                  Enter group details
                </p>
                <form
                  onSubmit={(e) => {
                    e.preventDefault();
                    handleCreateGroup();
                  }}
                >
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
                  {errorMessage && (
                    <p className="mt-2 text-red-500 text-center">
                      {errorMessage}
                    </p>
                  )}
                </form>
              </>
            ) : (
              <div className="text-center">
                <h2 className="text-2xl font-bold mb-4">
                  Group Created Successfully!
                </h2>
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
                <h2 className="text-2xl font-bold mb-4 text-center">
                  Join group with code
                </h2>
                <p className="text-gray-400 mb-4 text-center">
                  Paste your code
                </p>
                <form
                  onSubmit={(e) => {
                    e.preventDefault();
                    handleJoinGroup();
                  }}
                >
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
                  {errorMessage && (
                    <p className="mt-2 text-red-500 text-center">
                      Try again or check your code
                    </p>
                  )}
                </form>
              </>
            ) : (
              <div className="text-center">
                <h2 className="text-2xl font-bold mb-4">
                  Successfully joined!
                </h2>
                <span className="text-4xl">🎉</span>
              </div>
            )}
          </div>
        </div>
      )}
    </aside>
  );
}
