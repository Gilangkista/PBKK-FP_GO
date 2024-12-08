'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import axios from 'axios';

type Artist = {
  ID: number;
  Name: string;
  Slug: string;
};

export default function Artists() {
  const [artists, setArtists] = useState<Artist[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchArtists = async () => {
      try {
        const response = await axios.get('http://localhost:8080/artists/');
        setArtists(response.data || []); // Ensure it's an array
      } catch (err) {
        console.error('Error fetching artists:', err);
        setError('Failed to fetch artists');
      }
    };

    fetchArtists();
  }, []);

  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-4">All Artists</h1>

      {error && <p className="text-red-500">{error}</p>}

      <ul className="space-y-2">
        {artists && artists.length > 0 ? (
          artists.map((artist) => (
            <li key={artist.ID} className="border-b py-2">
              <Link href={`/artists/${artist.Slug}`} className="text-blue-500">
                {artist.Name}
              </Link>
            </li>
          ))
        ) : (
          <li>No artists available</li>
        )}
      </ul>
    </main>
  );
}
