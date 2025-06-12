export default async function ModView({
  params
}: {
  params: Promise<{ slug: string }>
}) {
  const { slug } = await params

  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24 bg-gray-50">
      <div className="z-10 w-full max-w-2xl items-center justify-center font-mono text-sm lg:flex flex-col p-8 border rounded-lg bg-white shadow-md">
        <h1 className="text-4xl font-bold mb-6">{slug}</h1>
      </div>
    </main>
  );
}
