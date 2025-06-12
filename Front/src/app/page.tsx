"use client"; // This is a client component

import { useState } from "react";
// import { GreeterClient } from "@/proto/GreeterServiceClientPb"; // Adjust path if you moved the files
// import { HelloRequest } from "@/proto/greeter_pb"; // Adjust path
import { ModServiceClient } from "@/proto/ModServiceClientPb";
import { ListModRequest } from "@/proto/mod_pb";
import { Feedback } from "@/components";

export default function Home() {
  const [name, setName] = useState("World");
  const [greeting, setGreeting] = useState("");
  const [error, setError] = useState("");

  const getGreeting = () => {
    // The proxy will run on port 8080
    // const client = new GreeterClient("http://localhost:8080");
    // const request = new HelloRequest();
    // request.setName(name);

    const client = new ModServiceClient("http://localhost:8080");
    const request = new ListModRequest();
    // request.setModId(name);

    setGreeting("");
    setError("");

    client.listMods(request, {}, (err, response) => {
      if (err) {
        console.error(`Unexpected error for sayHello: code = ${err.code}, message = ${err.message}`);
        setError(`Error: ${err.message}`);
        return;
      }
      if (response) {
        let greet = "HALO";
        response.getModsList().forEach(mod => {
          greet += ` ${mod.getName()}`
        })
        setGreeting(greet);
      }
    });
  };

  return (
    <main className="px-4 pt-2">
      <h1 className="text-4xl mb-6 text-center italic">Faith Conquest CMS</h1>
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
          <Feedback
            type="success"
            label="Server Response"
            content={greeting}
          />
        )}

        {error && (
            <Feedback
              type="success"
              label="Error"
              content={error}
            />
        )}
      </div>
    </main>
  );
}
