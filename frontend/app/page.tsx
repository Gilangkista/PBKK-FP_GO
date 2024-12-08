import Link from 'next/link';

export default function Home() {
  return (
    <main className="p-4">
      <h1 className="text-3xl font-bold mb-4">Welcome to Music Library</h1>
      <nav className="space-y-2">
        <ul>
          <li>
            <Link href="/playlists" className="text-blue-500">
              View Playlists
            </Link>
          </li>
          <li>
            <Link href="/songs" className="text-blue-500">
              View Songs
            </Link>
          </li>
          <li>
            <Link href="/artists" className="text-blue-500">
              View Artists
            </Link>
          </li>
        </ul>
      </nav>
    </main>
  );
}
