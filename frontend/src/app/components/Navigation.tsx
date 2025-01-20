"use client";
import { useGroup } from "@/src/app/contexts/GroupContext";
import Link from "next/link";
import { usePathname } from "next/navigation";

/**
 * Navigation bar component providing main application routing.
 * Shows different navigation items based on user role and group context.
 * Highlights current active route.
 *
 * Features:
 * - Conditional rendering based on group selection
 * - Role-based navigation (additional items for managers)
 * - Visual indication of current route
 * - Responsive hover states
 *
 * Navigation Items:
 * Default (all users):
 * - Events
 * - Announcements
 *
 * Manager-only:
 * - Tracks
 * - Subgroups
 * - Manage
 */
export default function NavigationBar() {
  const { groupId, userRole } = useGroup();
  const pathname = usePathname();

  if (!groupId) return null;

  const defaultUserItems = [
    { href: "/events", label: "Events" },
    { href: "/announcements", label: "Announcements" },
  ];

  const managerItems = [
    { href: "/tracks", label: "Tracks" },
    { href: "/subgroups", label: "Subgroups" },
    { href: "/manage", label: "Manage" },
  ];

  return (
    <nav className="border-b border-customGray px-8 py-4">
      <div className="flex gap-4">
        {defaultUserItems.map((item) => (
          <Link
            key={item.href}
            href={item.href}
            className={`px-4 py-2 rounded-md ${
              pathname === item.href
                ? "bg-sidebarButtonYellow text-white"
                : "hover:bg-headerHoverGray"
            }`}
          >
            {item.label}
          </Link>
        ))}
        {userRole === "manager" &&
          managerItems.map((item) => (
            <Link
              key={item.href}
              href={item.href}
              className={`px-4 py-2 rounded-md ${
                pathname === item.href
                  ? "bg-sidebarButtonYellow text-white"
                  : "hover:bg-headerHoverGray"
              }`}
            >
              {item.label}
            </Link>
          ))}
      </div>
    </nav>
  );
}
