"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState } from "react";

// Mock user data (this should be replaced by an API call when the backend is ready)
const mockUsers = [
  { id: 111, first_name: "John", last_name: "Doe", email: "john@example.com" },
  { id: 112, first_name: "Jane", last_name: "Smith", email: "jane@example.com" },
  { id: 113, first_name: "Jim", last_name: "Beam", email: "jim@example.com" },
];

export default function CreateSubgroupPage() {
  const router = useRouter();
  const { data: session } = useSession();
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [selectedUsers, setSelectedUsers] = useState<number[]>([]);
  const [showDropdown, setShowDropdown] = useState(false);
  const [userSelections, setUserSelections] = useState<number[]>([]); // Tracks selected checkboxes

  const handleCreateSubgroup = async () => {
    try {
      if (!name || !description) {
        alert("Name and description are required!");
        return;
      }

      // Example API request (replace with the actual API endpoint)
      /*
      await fetch('/api/group/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name,
          description,
          user_id: session?.user.id,
        }),
      });
      */

      console.log("Subgroup created:", { name, description, selectedUsers });
      alert("Subgroup created successfully!");
      router.push("/subgroups"); // Redirect back to the manage page
    } catch (error) {
      console.error("Error creating subgroup:", error);
      alert("Failed to create subgroup. Please try again.");
    }
  };

  const toggleUserSelection = (userId: number) => {
    setUserSelections((prev) =>
      prev.includes(userId) ? prev.filter((id) => id !== userId) : [...prev, userId]
    );
  };

  const handleAddUsers = () => {
    setSelectedUsers((prev) => Array.from(new Set([...prev, ...userSelections]))); // Avoid duplicates
    setUserSelections([]); // Clear selections
    setShowDropdown(false); // Close dropdown
  };

  return (
    <div className="flex flex-col p-6 max-w-lg mx-auto">
      <h1 className="text-3xl font-bold mb-6 text-gray-100">Create Subgroup</h1>

      {/* Subgroup Name */}
      <label className="block mb-2 font-semibold text-gray-100" htmlFor="name">
        Subgroup Name
      </label>
      <input
        type="text"
        id="name"
        value={name}
        onChange={(e) => setName(e.target.value)}
        className="w-full p-2 border rounded mb-4 text-gray-800"
        placeholder="Enter subgroup name"
      />

      {/* Subgroup Description */}
      <label className="block mb-2 font-semibold text-gray-100" htmlFor="description">
        Description
      </label>
      <textarea
        id="description"
        value={description}
        onChange={(e) => setDescription(e.target.value)}
        className="w-full p-2 border rounded mb-4 text-gray-800"
        placeholder="Enter subgroup description"
      />

      {/* Add Users */}
      <div className="relative">
        <button
          className="w-full px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-500"
          onClick={() => setShowDropdown((prev) => !prev)}
        >
          Add Users
        </button>
        {showDropdown && (
          <div className="absolute z-10 mt-2 bg-gray-800 border border-gray-600 rounded shadow w-full p-4">
            <ul className="divide-y divide-gray-600 max-h-48 overflow-y-auto">
              {mockUsers.map((user) => (
                <li key={user.id} className="flex items-center p-2 text-white">
                  <input
                    type="checkbox"
                    id={`user-${user.id}`}
                    checked={userSelections.includes(user.id)}
                    onChange={() => toggleUserSelection(user.id)}
                    className="mr-2"
                  />
                  <label htmlFor={`user-${user.id}`} className="cursor-pointer">
                    {user.first_name} {user.last_name} ({user.email})
                  </label>
                </li>
              ))}
            </ul>
            <button
              onClick={handleAddUsers}
              className="mt-4 w-full px-4 py-2 bg-green-600 text-white rounded hover:bg-green-500"
            >
              Add
            </button>
          </div>
        )}
      </div>

      {/* Selected Users */}
      {selectedUsers.length > 0 && (
        <div className="mt-4">
          <h3 className="font-semibold mb-2 text-gray-200">Selected Users:</h3>
          <ul className="list-disc list-inside text-gray-300">
            {selectedUsers.map((userId) => {
              const user = mockUsers.find((u) => u.id === userId);
              return (
                <li key={userId}>
                  {user?.first_name} {user?.last_name} ({user?.email})
                </li>
              );
            })}
          </ul>
        </div>
      )}

      {/* Submit Button */}
      <button
        onClick={handleCreateSubgroup}
        className="mt-6 px-4 py-2 bg-green-600 text-white rounded hover:bg-green-500"
      >
        Create Subgroup
      </button>
    </div>
  );
}
