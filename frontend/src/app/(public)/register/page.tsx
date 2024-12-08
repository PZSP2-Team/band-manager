'use client';

import { useState } from 'react';
import Image from 'next/image';
import Link from 'next/link';

export default function Register() {
  const [name, setName] = useState('');
  const [surname, setSurname] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log('Register submit');
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
          <h2 className="text-2xl font-bold mb-4 text-center">Sign up for Band Manager</h2>
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label htmlFor="name" className="ml-1 block text-sm font-medium text-white">Name</label>
              <input
                type="text"
                id="name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <div>
              <label htmlFor="surname" className="ml-1 block text-sm font-medium text-white">Surname</label>
              <input
                type="text"
                id="surname"
                value={surname}
                onChange={(e) => setSurname(e.target.value)}
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <div>
              <label htmlFor="email" className="ml-1 block text-sm font-medium text-white">Email</label>
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
              <label htmlFor="password" className="ml-1 block text-sm font-medium text-white">Password</label>
              <input
                type="password"
                id="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                className="px-2 mt-2 block w-full rounded-md bg-background border border-customGray shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
              />
            </div>
            <button type="submit" className="w-full border-customGray border hover:bg-hoverGray font-bold py-2 px-4 rounded">
              Sign up
            </button>
          </form>
          <p className="mt-4 text-center">
            Already have an account? <Link href="/login" className="text-blue-500 hover:text-blue-700">Sign in</Link>
          </p>
        </div>
      </div>
    </div>
  );
}
