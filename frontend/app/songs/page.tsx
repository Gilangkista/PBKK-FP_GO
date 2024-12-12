'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import axios from 'axios';

type Song = {
  ID: number;
  Slug: string;
  Title: string;
  Artist: { Name: string };
  Category: { Name: string };
};

export default function Songs() {
  const [songs, setSongs] = useState<Song[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchSongs = async () => {
      try {
        const response = await axios.get('http://localhost:8080/songs/');
        setSongs(response.data || []);
      } catch (err) {
        console.error('Error fetching songs:', err);
        setError('Failed to fetch songs');
      }
    };

    fetchSongs();
  }, []);

  return (
    <main className="bg-gray-100 min-h-screen flex items-center justify-center">
      <div className="container max-w-4xl mx-auto p-4">
        <h1 className="text-3xl font-bold mb-6 text-center text-gray-800">All Songs</h1>

        {error && <p className="text-red-500 text-center mb-4">{error}</p>}

        <div className="bg-white shadow-md rounded-lg p-6">
          {songs && songs.length > 0 ? (
            <ul className="divide-y divide-gray-200">
              {songs.map((song) => (
                <li key={song.ID} className="py-4 flex items-center justify-between hover:bg-gray-50">
                  <div>
                    <Link href={`/songs/${song.Slug}`} className="text-lg font-medium text-blue-600 hover:underline">
                      {song.Title}
                    </Link>
                    <p className="text-sm text-gray-500">
                      {song.Artist.Name} - {song.Category.Name}
                    </p>
                  </div>
                </li>
              ))}
            </ul>
          ) : (
            <p className="text-center text-gray-500">No songs available</p>
          )}
        </div>
      </div>
    </main>
  );
}
