"use client"; // This is a client component

import { useState } from "react";
import { GreeterClient } from "../GreeterServiceClientPb"; // Adjust path if you moved the files
import { HelloRequest } from "../greeter_pb"; // Adjust path

export default function Home() {
  const [name, setName] = useState("World");
  const [greeting, setGreeting] = useState("");
  const [error, setError] = useState("");

  const getGreeting = () => {
    // The proxy will run on port 8080
    const client = new GreeterClient("http://localhost:8080");
    const request = new HelloRequest();
    request.setName(name);

    setGreeting("");
    setError("");

    client.sayHello(request, {}, (err, response) => {
      if (err) {
        console.error(`Unexpected error for sayHello: code = ${err.code}, message = ${err.message}`);
        setError(`Error: ${err.message}`);
        return;
      }
      if (response) {
        setGreeting(response.getMessage());
      }
    });
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24 bg-gray-50">
      <div className="z-10 w-full max-w-2xl items-center justify-center font-mono text-sm lg:flex flex-col p-8 border rounded-lg bg-white shadow-md">
        <h1 className="text-4xl font-bold mb-6">gRPC-Web Client</h1>
        <div className="flex items-center space-x-4">
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="px-4 py-2 border rounded-md text-lg"
            placeholder="Enter a name"
          />
          <button
            onClick={getGreeting}
            className="px-6 py-2 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700 transition-colors"
          >
            Say Hello
          </button>
        </div>
        
        {greeting && (
          <div className="mt-8 p-4 bg-green-100 border border-green-300 text-green-800 rounded-md w-full text-center">
            <p className="font-semibold">Server Response:</p>
            <p className="text-lg">{greeting}</p>
          </div>
        )}

        {error && (
            <div className="mt-8 p-4 bg-red-100 border border-red-300 text-red-800 rounded-md w-full text-center">
              <p className="font-semibold">Error:</p>
              <p className="text-lg">{error}</p>
            </div>
        )}
      </div>
    </main>
  );
}
