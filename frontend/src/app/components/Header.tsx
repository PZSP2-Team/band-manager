"use client";

import Image from "next/image";
import { LogOut } from "lucide-react";
import { signOut } from "next-auth/react";
import { useRouter } from "next/navigation";

/**
 * Custom sign out function that clears local storage before signing out.
 * Extends NextAuth's signOut functionality.
 *
 * Side effects:
 * - Clears all items from localStorage
 * - Triggers NextAuth sign out process
 */

/**
 * Header component for the application.
 * Displays logo, app name and logout button.
 *
 * Features:
 * - Brand identity (logo and name)
 * - Logout functionality with localStorage cleanup
 */
export default function Header() {
  const router = useRouter();
  const customSignOut = async () => {
    localStorage.clear();
    await signOut({ redirect: false });
    router.push("/");
  };
  return (
    <header className="bg-headerGray p-4 shadow-xl">
      <nav className="flex justify-between items-center">
        <div className="flex items-center gap-4">
          <Image
            src="/live-music.png"
            alt="Band Manager Logo"
            width={32}
            height={32}
          />
          <span className="font-bold text-2xl">Band Manager</span>
        </div>
        <button
          onClick={customSignOut}
          className="bg-headerGray flex items-center gap-2 border border-customGray text-customGray hover:bg-headerHoverGray px-4 py-2 rounded transition-colors"
        >
          <LogOut className="h-4 w-4" />
          Log out
        </button>
      </nav>
    </header>
  );
}
