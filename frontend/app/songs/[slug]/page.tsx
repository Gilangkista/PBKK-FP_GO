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
  const params = useParams();
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
      <main className="p-4 flex items-center justify-center min-h-screen bg-gray-100">
        <div className="text-center">
          <h1 className="text-2xl font-bold text-gray-800">Loading...</h1>
          {error && <p className="text-red-500 mt-4">{error}</p>}
        </div>
      </main>
    );
  }

  return (
    <main className="p-6 bg-gray-100 min-h-screen flex items-center justify-center">
      <div className="bg-white shadow-md rounded-lg p-6 max-w-lg w-full">
        <h1 className="text-3xl font-bold mb-4 text-gray-800 text-center">{song.Title}</h1>
        <div className="text-lg text-gray-700 mb-2">
          <p><span className="font-semibold">Artist:</span> {song.Artist.Name}</p>
          <p><span className="font-semibold">Category:</span> {song.Category.Name}</p>
        </div>
        <div className="text-center mt-4">
          <button
            className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-opacity-75"
            onClick={() => window.history.back()}
          >
            Go Back
          </button>
        </div>
      </div>
    </main>
  );
}