"use client";

import { useRef, useState } from "react";
import Image from "next/image";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { signIn } from "next-auth/react";

export default function Login() {
  const [errorMessage, setErrorMessage] = useState("");
  const router = useRouter();
  const formRef = useRef<HTMLFormElement>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMessage("");

    const form = formRef.current;
    const email = form?.email?.value;
    const password = form?.password?.value;

    try {
      const result = await signIn("credentials", {
        email,
        password,
        redirect: false,
      });
      if (result?.error) {
        setErrorMessage("Invalid email or password");
        return;
      }
      if (result?.ok) {
        router.push("/dashboard");
        router.refresh();
      }
    } catch (err) {
      console.error(`Error: ${err}`);
      setErrorMessage(
        "An unexpected error occured. Please try again in few minutes.",
      );
    }
  };

  return (
    <div className="min-h-screen flex ">
      <div className="relative w-1/2">
        <Image
          src="/login.jpg"
          alt="Login image"
          fill
          className="object-cover"
          priority
        />
      </div>
      <div className="w-1/2 flex items-center justify-center">
        <div className="w-96 border border-customGray p-8 rounded-md shadow-md">
          <h2 className="text-2xl font-bold mb-4 text-center">
            Sign in to Band Manager
          </h2>
          {errorMessage && (
            <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
              {errorMessage}
            </div>
          )}
          <form ref={formRef} onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label
                htmlFor="email"
                className="ml-1 block text-sm font-medium text-white"
              >
                Email
              </label>
              <input
                type="email"
                id="email"
                name="email"
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <div>
              <label
                htmlFor="password"
                className="ml-1 block text-sm font-medium text-white"
              >
                Password
              </label>
              <input
                type="password"
                id="password"
                name="password"
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <button
              type="submit"
              className="w-full border-customGray border hover:bg-hoverGray font-bold py-2 px-4 rounded"
            >
              Sign in
            </button>
          </form>
          <p className="mt-4 text-center">
            New to Band Manager?{" "}
            <Link
              href="/register"
              className="text-blue-500 hover:text-blue-700"
            >
              Sign up
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
