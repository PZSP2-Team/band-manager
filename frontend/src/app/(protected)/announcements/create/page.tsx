"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState, useEffect } from "react";

type AnnouncementForm = {
  title: string;
  description: string;
  date: string;
  priority: number;
  subgroup_ids: number[];
};

type Subgroup = {
  id: number;
  group_id: number;
  name: string;
  description: string;
};

export default function CreateAnnouncementPage() {
  const router = useRouter();
  const { data: session } = useSession();
  const [formData, setFormData] = useState<AnnouncementForm>({
    title: "",
    description: "",
    date: "",
    priority: 1,
    subgroup_ids: [],
  });
  const [subgroups, setSubgroups] = useState<Subgroup[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isDropdownOpen, setIsDropdownOpen] = useState(false); // State to control dropdown visibility

  useEffect(() => {
    const fetchSubgroups = async () => {
      try {
        console.log("Fetching subgroups...");

        // Mock API call
        const mockSubgroups: Subgroup[] = [
          { id: 1, group_id: 101, name: "Development Team", description: "Handles coding tasks" },
          { id: 2, group_id: 101, name: "Marketing Team", description: "Handles promotions" },
          { id: 3, group_id: 101, name: "Design Team", description: "Handles UI/UX design" },
        ];

        setTimeout(() => {
          setSubgroups(mockSubgroups);
        }, 1000);

        // Uncomment the following lines to use the actual API:
        // const userId = session?.user.id; // Replace with actual user ID
        // const groupId = 101; // Replace with actual group ID
        // const response = await fetch(`/api/subgroup/group/${groupId}/${userId}`);
        // if (!response.ok) {
        //   throw new Error("Failed to fetch subgroups");
        // }
        // const data = await response.json();
        // setSubgroups(data.subgroups);
      } catch (error) {
        console.error("Error fetching subgroups:", error);
      }
    };

    fetchSubgroups();
  }, [session]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name === "priority" ? parseInt(value) : value,
    }));
  };

  const handleCheckboxChange = (subgroupId: number) => {
    setFormData((prev) => {
      const isSelected = prev.subgroup_ids.includes(subgroupId);
      return {
        ...prev,
        subgroup_ids: isSelected
          ? prev.subgroup_ids.filter((id) => id !== subgroupId)
          : [...prev.subgroup_ids, subgroupId],
      };
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      console.log("Submitting announcement:", formData);

      // Mock API call
      setTimeout(() => {
        alert("Announcement created successfully!");
        setIsSubmitting(false);
        router.push("/announcements"); // Redirect to announcements page
      }, 1000);

      // Uncomment the following lines to use the actual API:
      // const response = await fetch("/api/announcement/create", {
      //   method: "POST",
      //   headers: {
      //     "Content-Type": "application/json",
      //   },
      //   body: JSON.stringify({
      //     ...formData,
      //     group_id: 101, // Replace with actual group_id
      //     sender_id: session?.user.id, // Replace with actual sender_id
      //   }),
      // });
      // if (response.ok) {
      //   alert("Announcement created successfully!");
      //   router.push("/announcements");
      // } else {
      //   throw new Error("Failed to create announcement");
      // }
    } catch (error) {
      console.error("Error creating announcement:", error);
      alert("Failed to create announcement. Please try again.");
      setIsSubmitting(false);
    }
  };

  return (
    <div className="flex flex-col items-center justify-center mt-10 p-6">
      <h1 className="text-3xl font-bold mb-6">Create Announcement</h1>
      <form
        onSubmit={handleSubmit}
        className="w-full max-w-lg p-4 border border-gray-300 rounded shadow"
      >
        <div className="mb-4">
          <label htmlFor="title" className="block text-gray-700 font-semibold mb-2">
            Title
          </label>
          <input
            type="text"
            id="title"
            name="title"
            value={formData.title}
            onChange={handleInputChange}
            className="w-full p-2 border rounded"
            placeholder="Enter announcement title"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="description" className="block text-gray-700 font-semibold mb-2">
            Description
          </label>
          <textarea
            id="description"
            name="description"
            value={formData.description}
            onChange={handleInputChange}
            className="w-full p-2 border rounded"
            placeholder="Enter announcement description"
            rows={4}
            required
          ></textarea>
        </div>
        <div className="mb-4">
          <label htmlFor="date" className="block text-gray-700 font-semibold mb-2">
            Date
          </label>
          <input
            type="date"
            id="date"
            name="date"
            value={formData.date}
            onChange={handleInputChange}
            className="w-full p-2 border rounded"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="priority" className="block text-gray-700 font-semibold mb-2">
            Priority
          </label>
          <select
            id="priority"
            name="priority"
            value={formData.priority}
            onChange={handleInputChange}
            className="w-full p-2 border rounded"
            required
          >
            <option value={1}>Low</option>
            <option value={2}>Medium</option>
            <option value={3}>High</option>
          </select>
        </div>
        <div className="mb-4">
          <label className="block text-gray-700 font-semibold mb-2">Select Subgroups</label>
          <div className="relative">
            <button
              type="button"
              className="w-full p-2 bg-gray-800 text-white rounded"
              onClick={() => setIsDropdownOpen((prev) => !prev)}
            >
              {isDropdownOpen ? "Close Subgroup Selection" : "Select Subgroups"}
            </button>
            {isDropdownOpen && (
              <div className="absolute mt-2 w-full bg-gray-800 text-white border border-gray-700 rounded shadow p-2">
                {subgroups.map((subgroup) => (
                  <div key={subgroup.id} className="flex items-center mb-2">
                    <input
                      type="checkbox"
                      id={`subgroup-${subgroup.id}`}
                      checked={formData.subgroup_ids.includes(subgroup.id)}
                      onChange={() => handleCheckboxChange(subgroup.id)}
                      className="mr-2"
                    />
                    <label htmlFor={`subgroup-${subgroup.id}`} className="text-white">
                      {subgroup.name}
                    </label>
                  </div>
                ))}
              </div>
            )}
          </div>
        </div>
        {formData.subgroup_ids.length > 0 && (
          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2">Selected Subgroups</label>
            <ul className="list-disc pl-5">
              {formData.subgroup_ids.map((id) => {
                const subgroup = subgroups.find((sg) => sg.id === id);
                return (
                  <li key={id} className="text-gray-700">
                    {subgroup?.name}
                  </li>
                );
              })}
            </ul>
          </div>
        )}
        <div className="flex justify-end">
          <button
            type="button"
            className="px-4 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-200 mr-2"
            onClick={() => router.push("/announcements")}
          >
            Cancel
          </button>
          <button
            type="submit"
            className={`px-4 py-2 bg-green-600 text-white rounded hover:bg-green-500 ${
              isSubmitting ? "opacity-50 cursor-not-allowed" : ""
            }`}
            disabled={isSubmitting}
          >
            {isSubmitting ? "Submitting..." : "Create Announcement"}
          </button>
        </div>
      </form>
    </div>
  );
}
