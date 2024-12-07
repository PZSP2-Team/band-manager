"use client";

import { signOut } from "next-auth/react";
import Image from "next/image";
import { LogOut } from "lucide-react"

export default function Header() {
    return (
        <header className="bg-headerGray p-4 shadow-xl">
            <nav className="flex justify-between items-center">
                <div className="flex items-center gap-2">
                    <Image
                        src="/live-music.png"
                        alt="Band Manager Logo"
                        width={24}
                        height={24}
                    />
                    <span className="font-bold text-xl">Band Manager</span>
                </div>
                <button
                    onClick={() => signOut({ callbackUrl: "/" })}
                    className="bg-headerGray flex items-center gap-2 border border-customGray text-customGray hover:bg-headerHoverGray px-4 py-2 rounded transition-colors"
                >
                    <LogOut className="h-4 w-4" />
                    Log out
                </button>
            </nav>
        </header>
    );
}
