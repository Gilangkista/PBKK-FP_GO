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
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-4">All Songs</h1>

      {error && <p className="text-red-500">{error}</p>}

      <ul className="space-y-2">
        {songs && songs.length > 0 ? (
          songs.map((song) => (
            <li key={song.ID} className="border-b py-2">
              <Link href={`/songs/${song.Slug}`} className="text-blue-500">
                {song.Title} - {song.Artist.Name} ({song.Category.Name})
              </Link>
            </li>
          ))
        ) : (
          <li>No songs available</li>
        )}
      </ul>
    </main>
  );
}
