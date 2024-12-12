import Link from 'next/link';

export default function Home() {
  return (
    <main className="p-4">
      <h1 className="text-3xl font-bold mb-6 text-center text-black">Welcome to Music Library</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <Link href="/songs" legacyBehavior>
          <a className="block p-6 bg-green-500 text-white rounded-lg shadow-md hover:bg-green-600">
            <h2 className="text-xl font-semibold">Songs</h2>
            <p className="mt-2">Explore and listen to your favorite songs.</p>
          </a>
        </Link>
        <Link href="/artists" legacyBehavior>
          <a className="block p-6 bg-red-500 text-white rounded-lg shadow-md hover:bg-red-600">
            <h2 className="text-xl font-semibold">Artists</h2>
            <p className="mt-2">Discover artists and their music.</p>
          </a>
        </Link>
        <Link href="/playlists" legacyBehavior>
          <a className="block p-6 bg-blue-500 text-white rounded-lg shadow-md hover:bg-blue-600">
            <h2 className="text-xl font-semibold">Playlists</h2>
            <p className="mt-2">Browse and manage your playlists.</p>
          </a>
        </Link>
      </div>
    </main>
  );
}