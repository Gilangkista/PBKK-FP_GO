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
    <main className="bg-gray-100 min-h-screen flex items-center justify-center p-6">
      <div className="bg-white shadow-md rounded-lg p-6 w-full max-w-3xl">
        <h1 className="text-3xl font-bold mb-6 text-center text-gray-800">All Artists</h1>

        {error && <p className="text-red-500 text-center mb-4">{error}</p>}

        {artists && artists.length > 0 ? (
          <ul className="divide-y divide-gray-200">
            {artists.map((artist) => (
              <li key={artist.ID} className="py-4 hover:bg-gray-50">
                <Link href={`/artists/${artist.Slug}`} className="text-lg font-medium text-blue-600 hover:underline">
                  {artist.Name}
                </Link>
              </li>
            ))}
          </ul>
        ) : (
          <p className="text-center text-gray-500">No artists available</p>
        )}
      </div>
    </main>
  );
}
