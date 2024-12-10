'use client';

import { useState, useEffect } from 'react';
import { useParams } from 'next/navigation';
import axios from 'axios';

type Song = {
  ID: number;
  Slug: string;
  Title: string;
  Artist: { Name: string };
  Category: { Name: string };
};

export default function SongDetails() {
  const params = useParams(); // Menggunakan useParams untuk mendapatkan params
  const [song, setSong] = useState<Song | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchSong = async () => {
      if (!params?.slug) return;

      try {
        const response = await axios.get(`http://localhost:8080/songs/${params.slug}`);
        setSong(response.data);
      } catch (err) {
        console.error('Error fetching song details:', err);
        setError('Failed to load song details.');
      }
    };

    fetchSong();
  }, [params]);

  if (!song) {
    return (
      <main className="p-4">
        <h1 className="text-2xl font-bold">Loading...</h1>
        {error && <p className="text-red-500">{error}</p>}
      </main>
    );
  }

  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold">{song.Title}</h1>
      <p className="text-lg">Artist: {song.Artist.Name}</p>
      <p className="text-lg">Category: {song.Category.Name}</p>
    </main>
  );
}
