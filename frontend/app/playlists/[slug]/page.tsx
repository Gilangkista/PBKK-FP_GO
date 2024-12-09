'use client';

import { useEffect, useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';

type PlaylistDetail = {
  ID: number;
  Name: string;
  Description?: string;
};

export default function PlaylistDetail() {
  const [playlist, setPlaylist] = useState<PlaylistDetail | null>(null);
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();
  const { slug } = router.query;  // Ambil slug dari URL

  useEffect(() => {
    if (slug) {
      const fetchPlaylist = async () => {
        try {
          // Menggunakan endpoint yang benar sesuai slug
          const response = await axios.get(`http://localhost:8080/playlists/${slug}`);
          setPlaylist(response.data);
        } catch (err) {
          console.error('Error fetching playlist detail:', err);
          setError('Failed to fetch playlist details');
        }
      };

      fetchPlaylist();
    }
  }, [slug]);  // Jalankan ulang jika slug berubah

  if (error) return <p className="text-red-500">{error}</p>;

  return (
    <main className="p-4">
      {playlist ? (
        <>
          <h1 className="text-2xl font-bold">{playlist.Name}</h1>
          {playlist.Description && <p>{playlist.Description}</p>}
        </>
      ) : (
        <p>Loading playlist...</p>
      )}
    </main>
  );
}
