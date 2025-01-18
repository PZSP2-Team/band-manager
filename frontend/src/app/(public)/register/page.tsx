"use client";
import { useState } from "react";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";

/**
 * Registration page component for new user sign up.
 * Collects user information and handles account creation through API.
 */
export default function Register() {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const router = useRouter();

  /**
   * Handles form submission for user registration
   * Side effects:
   * - Sends registration data to API
   * - Updates success/error message states based on response
   * - Redirects to login page after successful registration
   */
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrorMessage("");
    setSuccessMessage("");

    try {
      const response = await fetch("/api/verify/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          email: email,
          password: password,
        }),
      });

      if (!response.ok) {
        const errorText = await response.text();
        setErrorMessage(errorText);
        return;
      }

      const data = await response.json();
      if (data.message) {
        setSuccessMessage(data.message);
        setTimeout(() => {
          router.push("/login");
        }, 2000);
      }
    } catch (err) {
      console.error(`Error: ${err}`);
      setErrorMessage("Something went wrong with the registration");
    }
  };
  return (
    <div className="min-h-screen flex">
      <div className="relative w-1/2">
        <Image
          src="/register.jpg"
          alt="Register image"
          fill
          className="object-cover"
          priority
        />
      </div>
      <div className="w-1/2 flex items-center justify-center">
        <div className="w-96 border border-customGray p-8 rounded-md shadow-md">
          <h2 className="text-2xl font-bold mb-4 text-center">
            Sign up for Band Manager
          </h2>
          {errorMessage && (
            <div className="mb-4 p-2 bg-red-100 border border-red-400 text-red-700 rounded">
              {errorMessage}
            </div>
          )}
          {successMessage && (
            <div className="mb-4 p-2 bg-green-100 border border-green-400 text-green-700 rounded">
              {successMessage}
            </div>
          )}
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label
                htmlFor="name"
                className="ml-1 block text-sm font-medium text-white"
              >
                Name
              </label>
              <input
                type="text"
                id="name"
                value={firstName}
                onChange={(e) => setFirstName(e.target.value)}
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <div>
              <label
                htmlFor="surname"
                className="ml-1 block text-sm font-medium text-white"
              >
                Surname
              </label>
              <input
                type="text"
                id="surname"
                value={lastName}
                onChange={(e) => setLastName(e.target.value)}
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
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
                value={email}
                onChange={(e) => setEmail(e.target.value)}
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
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <button
              type="submit"
              className="w-full border-customGray border hover:bg-hoverGray font-bold py-2 px-4 rounded"
            >
              Sign up
            </button>
          </form>
          <p className="mt-4 text-center">
            Already have an account?{" "}
            <Link href="/login" className="text-blue-500 hover:text-blue-700">
              Sign in
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
