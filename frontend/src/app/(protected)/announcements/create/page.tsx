"use client";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useState } from "react";

type AnnouncementForm = {
  title: string;
  description: string;
  date: string;
  priority: number;
};

export default function CreateAnnouncementPage() {
  const router = useRouter();
  const { data: session } = useSession();
  const [formData, setFormData] = useState<AnnouncementForm>({
    title: "",
    description: "",
    date: "",
    priority: 1,
  });
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: name === "priority" ? parseInt(value) : value,
    }));
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
      //     group_id: 1, // Replace with actual group_id
      //     sender_id: session?.user.id, // Replace with actual sender_id
      //     subgroup_ids: [], // Add subgroup IDs if applicable
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
            className="w-full p-2 border rounded text-gray-700"
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
            className="w-full p-2 border rounded text-gray-700"
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
            className="w-full p-2 border rounded text-gray-700"
            required
          >
            <option value={1}>Low</option>
            <option value={2}>Medium</option>
            <option value={3}>High</option>
          </select>
        </div>
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
