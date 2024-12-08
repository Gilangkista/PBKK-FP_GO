'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import axios from 'axios';

type Playlist = {
  ID: number;
  Name: string;
  Slug: string;
};

export default function Playlists() {
  const [playlists, setPlaylists] = useState<Playlist[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPlaylists = async () => {
      try {
        const response = await axios.get('http://localhost:8080/playlists');  // Pastikan API endpoint ini benar
        console.log(response.data);
        setPlaylists(response.data || []);  // Pastikan data yang diterima adalah array
      } catch (err) {
        console.error('Error fetching playlists:', err);
        setError('Failed to fetch playlists');
      }
    };

    fetchPlaylists();
  }, []);

  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-4">All Playlists</h1>

      {error && <p className="text-red-500">{error}</p>}

      <ul className="space-y-2">
        {playlists.length > 0 ? (
          playlists.map((playlist) => (
            <li key={playlist.ID} className="border-b py-2">
              <Link href={`/playlists/${playlist.Slug}`} className="text-blue-500">
                {playlist.Name}
              </Link>
            </li>
          ))
        ) : (
          <li>No playlists available</li>
        )}
      </ul>
    </main>
  );
}
