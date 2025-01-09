"use client";
import { useGroup } from "@/src/app/contexts/GroupContext";
import Link from "next/link";
import { usePathname } from "next/navigation";

export default function NavigationBar() {
  const { groupId } = useGroup();
  const pathname = usePathname();

  // if (!groupId) return null;

  const navigationItems = [
    { href: "/events", label: "Events" },
    { href: "/notifications", label: "Notifications" },
    { href: "/tracks", label: "Tracks" },
    { href: "/subgroups", label: "Subgroups" },
    { href: "/manage", label: "Manage" },
  ];

  return (
    <nav className="border-b border-customGray px-8 py-4">
      <div className="flex gap-4">
        {navigationItems.map((item) => (
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
