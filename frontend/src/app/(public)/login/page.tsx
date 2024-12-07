"use client";

import { useRef, useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { signIn } from "next-auth/react";

export default function Login() {
    const [error, setError] = useState("");
    const router = useRouter();
    const formRef = useRef<HTMLFormElement>(null)

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        setError("");

        const form = formRef.current;
        const email = form?.email?.value;
        const password = form?.password?.value;

        try {
            const result = await signIn("credentials", {
                email,
                password,
                redirect: false,
            })
            if (result?.error) {
                setError("Invalid email or password")
                return
            }
            if (result?.ok) {
                router.push("/group")
                router.refresh()
            }
        } catch (error) {
            setError("An unexpected error occured. Please try again in few minutes.")
        }
    }

    return (
        <div className="min-h-screen flex items-center justify-center">
            <div className="w-96">
                <h2 className="text-2xl font-bold mb-4">Sign in</h2>
                {error && (
                    <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
                        {error}
                    </div>
                )}
                <form ref={formRef} onSubmit={handleSubmit} className="space-y-4">
                    <div>
                        <label htmlFor="email" className="block text-sm font-medium text-white">Email</label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            required
                            className="px-2 bg-gray-900 mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                        />
                    </div>
                    <div>
                        <label htmlFor="password" className="block text-sm font-medium text-white">Password</label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            required
                            className="px-2 bg-gray-900 mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                        />
                    </div>
                    <button type="submit" className="w-full bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                        Sign in
                    </button>
                </form>
                <p className="mt-4 text-center">
                    No account? <Link href="/register" className="text-blue-500 hover:text-blue-700">Sign up</Link>
                </p>
            </div>
    </div>
    )
}
