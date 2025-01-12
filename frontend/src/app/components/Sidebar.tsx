"use client";
import { useSession } from "next-auth/react";
import { LucideUsers, UserRoundPlus, UsersRound } from "lucide-react";
import { useEffect, useState } from "react";
import { useGroup } from "../contexts/GroupContext";
import { useRouter } from "next/navigation";
import LoadingScreen from "./LoadingScreen";

type Group = {
  id: number;
  name: string;
  role: string;
};

type RenderState =
  | { status: "loading" }
  | { status: "loaded" }
  | { status: "error" };

export default function Sidebar() {
  const { groupId, setGroupId, setUserRole } = useGroup();
  const router = useRouter();
  const [renderState, setRenderState] = useState<RenderState>({
    status: "loading",
  });
  const [groupList, setGroupList] = useState<Group[]>([]);
  const { data: session, status: sessionStatus } = useSession();
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
    if (sessionStatus === "loading") return;
    const fetchGroups = async () => {
      try {
        const response = await fetch(`/api/group/user/${session?.user?.id}`);

        if (!response.ok) {
          throw new Error("Failed to fetch groups");
        }

        const data = await response.json();
        setGroupList(data.groups);
        setRenderState({ status: "loaded" });
      } catch (error) {
        console.error("Error fetching groups:", error);
      }
    };

    if (session?.user?.id) {
      fetchGroups();
    }
  }, [sessionStatus, session?.user?.id]);

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
          {renderState.status === "loaded" ? (
            groupList?.length > 0 ? (
              <div className="space-y-2">
                {groupList.map((group) => (
                  <button
                    key={group.id}
                    onClick={() => handleGroupSelect(group.id, group.role)}
                    className={`flex w-full items-center gap-3 p-2 rounded-lg transition-colors text-customGray
            ${
              groupId === group.id
                ? "bg-sidebarButtonYellow text-white"
                : "hover:bg-headerHoverGray"
            }`}
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
            )
          ) : (
            <LoadingScreen />
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
        <div className="fixed inset-0 bg-black bg-opacity-80 flex items-center justify-center z-50">
          <div className="bg-background text-white p-8 rounded-lg border border-customGray shadow-lg w-[28rem] relative">
            <button
              onClick={() => {
                setShowCreateModal(false);
                setErrorMessage("");
                setGroupName("");
                setGroupDescription("");
              }}
              className="absolute top-4 right-4 text-customGray hover:text-white transition-colors"
            >
              ✕
            </button>
            {!success ? (
              <>
                <h2 className="text-2xl font-bold mb-2 text-center">
                  Create New Group
                </h2>
                <p className="text-customGray mb-6 text-center text-sm">
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
                    className="px-4 py-2 mb-4 block w-full rounded bg-background border border-customGray text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                  <textarea
                    value={groupDescription}
                    onChange={(e) => setGroupDescription(e.target.value)}
                    placeholder="Group description"
                    required
                    className="px-4 py-2 mb-6 block w-full rounded bg-background border border-customGray text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                  <button
                    type="submit"
                    className="w-full px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
                  >
                    Create
                  </button>
                  {errorMessage && (
                    <p className="mt-4 text-red-500 text-center text-sm">
                      {errorMessage}
                    </p>
                  )}
                </form>
              </>
            ) : (
              <div className="text-center py-4">
                <h2 className="text-2xl font-bold mb-2">
                  Group Created Successfully!
                </h2>
                <button
                  onClick={() => {
                    setShowCreateModal(false);
                    setSuccess(false);
                  }}
                  className="mt-6 px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
                >
                  Close
                </button>
              </div>
            )}
          </div>
        </div>
      )}

      {showJoinModal && (
        <div className="fixed inset-0 bg-black bg-opacity-80 flex items-center justify-center z-50">
          <div className="bg-background text-white p-8 rounded-lg border border-customGray shadow-lg w-[28rem] relative">
            <button
              onClick={() => {
                setShowJoinModal(false);
                setErrorMessage("");
                setJoinToken("");
              }}
              className="absolute top-4 right-4 text-customGray hover:text-white transition-colors"
            >
              ✕
            </button>
            {!success ? (
              <>
                <h2 className="text-2xl font-bold mb-2 text-center">
                  Join group with code
                </h2>
                <p className="text-customGray mb-6 text-center text-sm">
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
                    className="px-4 py-2 mb-6 block w-full rounded bg-background border border-customGray text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                  <button
                    type="submit"
                    className="w-full px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
                  >
                    Join
                  </button>
                  {errorMessage && (
                    <p className="mt-4 text-red-500 text-center text-sm">
                      Try again or check your code
                    </p>
                  )}
                </form>
              </>
            ) : (
              <div className="text-center py-4">
                <h2 className="text-2xl font-bold mb-2">
                  Successfully joined!
                </h2>
                <button
                  onClick={() => {
                    setShowJoinModal(false);
                    setSuccess(false);
                  }}
                  className="mt-6 px-6 py-2 bg-blue-600 text-white rounded shadow hover:bg-blue-500 transition"
                >
                  Close
                </button>
              </div>
            )}
          </div>
        </div>
      )}
    </aside>
  );
}
