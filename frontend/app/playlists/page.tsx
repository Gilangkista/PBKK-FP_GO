'use client'; // Tambahkan 'use client' di sini untuk menandakan ini adalah Client Component

import { useEffect, useState } from 'react';
import Link from 'next/link';
import axios from 'axios';
import { useRouter } from 'next/navigation'; // Gunakan useRouter dari 'next/navigation'

type Playlist = {
  ID: number;
  Name: string;
  Slug: string;
};

export default function Playlists() {
  const [playlists, setPlaylists] = useState<Playlist[]>([]);
  const [newPlaylistName, setNewPlaylistName] = useState<string>('');
  const [newPlaylistDescription, setNewPlaylistDescription] = useState<string>('');
  const [error, setError] = useState<string | null>(null);
  const router = useRouter(); // Pastikan useRouter hanya dipanggil di Client Component

  useEffect(() => {
    const fetchPlaylists = async () => {
      try {
        const response = await axios.get('http://localhost:8080/playlists/');
        setPlaylists(response.data || []); // Pastikan data yang diterima adalah array
      } catch (err) {
        console.error('Error fetching playlists:', err);
        setError('Failed to fetch playlists');
      }
    };

    fetchPlaylists();
  }, []);

  const handleCreatePlaylist = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const newPlaylist = { Name: newPlaylistName, Description: newPlaylistDescription };
      const response = await axios.post('http://localhost:8080/playlists/', newPlaylist);
      setPlaylists([...playlists, response.data]); // Menambah playlist baru ke state
      setNewPlaylistName('');
      setNewPlaylistDescription('');
    } catch (err) {
      console.error('Error creating playlist:', err);
      setError('Failed to create playlist');
    }
  };

  const handleDeletePlaylist = async (id: number) => {
    try {
      await axios.delete(`http://localhost:8080/playlists/${id}`);
      setPlaylists(playlists.filter(playlist => playlist.ID !== id)); // Menghapus playlist dari state
    } catch (err) {
      console.error('Error deleting playlist:', err);
      setError('Failed to delete playlist');
    }
  };

  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-4">All Playlists</h1>

      {error && <p className="text-red-500">{error}</p>}

      <form onSubmit={handleCreatePlaylist} className="mb-4">
        <input
          type="text"
          value={newPlaylistName}
          onChange={(e) => setNewPlaylistName(e.target.value)}
          placeholder="Playlist Name"
          className="border p-2 w-full mb-2"
          required
        />
        <textarea
          value={newPlaylistDescription}
          onChange={(e) => setNewPlaylistDescription(e.target.value)}
          placeholder="Description"
          className="border p-2 w-full mb-2"
        />
        <button type="submit" className="bg-blue-500 text-white p-2">Create Playlist</button>
      </form>

      <ul className="space-y-2">
        {playlists.length > 0 ? (
          playlists.map((playlist) => (
            <li key={playlist.ID} className="border-b py-2 flex justify-between">
              <Link href={`/playlists/${playlist.Slug}`} className="text-blue-500">
                {playlist.Name}
              </Link>
              <button
                onClick={() => handleDeletePlaylist(playlist.ID)}
                className="text-red-500"
              >
                Delete
              </button>
            </li>
          ))
        ) : (
          <li>No playlists available</li>
        )}
      </ul>
    </main>
  );
}
