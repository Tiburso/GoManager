export default function Company({
  params,
}: {
  params: {
    id: string;
  };
}) {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1 className="text-6xl font-bold">Welcome to the Company Page!</h1>
      <p className="text-xl">This is the company page with id {params.id}</p>
    </main>
  );
}

// Can add a generateStaticParams function to generate the static paths
// Need to check what to do relating to the dynamic applications
