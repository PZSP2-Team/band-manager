"use client";
import { RequireDashboard } from "../../components/RequireDashboard";

/**
 * Landing page for groups section.
 * Displays guidance message when no group is selected.
 * Requires user to be logged in to access dashboard.
 */
export default function GroupsPage() {
  return (
    <RequireDashboard>
      <div className="h-full flex-1 flex items-center justify-center p-96">
        <div className="text-center text-customGray">
          <h2 className="text-xl mb-2">
            Select a group from the sidebar or create a new one to get started
          </h2>
        </div>
      </div>
    </RequireDashboard>
  );
}
