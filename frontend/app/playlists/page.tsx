'use client';

import { useEffect, useState } from 'react';
import Link from 'next/link';
import axios from 'axios';
import { useRouter } from 'next/navigation';

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
  const router = useRouter();

  useEffect(() => {
    const fetchPlaylists = async () => {
      try {
        const response = await axios.get('http://localhost:8080/playlists/');
        setPlaylists(response.data || []);
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
      setPlaylists([...playlists, response.data]);
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
      setPlaylists(playlists.filter(playlist => playlist.ID !== id));
    } catch (err) {
      console.error('Error deleting playlist:', err);
      setError('Failed to delete playlist');
    }
  };

  return (
    <main className="bg-gray-100 min-h-screen flex flex-col items-center p-6">
      <div className="bg-white shadow-md rounded-lg p-6 w-full max-w-4xl">
        <h1 className="text-3xl font-bold mb-6 text-center text-gray-800">All Playlists</h1>

        {error && <p className="text-red-500 text-center mb-4">{error}</p>}

        <form onSubmit={handleCreatePlaylist} className="mb-6 bg-gray-50 p-4 rounded-lg shadow">
          <h2 className="text-xl font-semibold mb-4 text-gray-700">Create New Playlist</h2>
          <input
            type="text"
            value={newPlaylistName}
            onChange={(e) => setNewPlaylistName(e.target.value)}
            placeholder="Playlist Name"
            className="border border-gray-300 rounded p-2 w-full mb-2 focus:outline-none focus:ring-2 focus:ring-blue-400 text-black"
            required
          />
          <textarea
            value={newPlaylistDescription}
            onChange={(e) => setNewPlaylistDescription(e.target.value)}
            placeholder="Description"
            className="border border-gray-300 rounded p-2 w-full mb-2 focus:outline-none focus:ring-2 focus:ring-blue-400 text-black"
          />
          <button
            type="submit"
            className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 transition"
          >
            Create Playlist
          </button>
        </form>

        <ul className="divide-y divide-gray-200">
          {playlists.length > 0 ? (
            playlists.map((playlist) => (
              <li
                key={playlist.ID}
                className="py-4 flex justify-between items-center hover:bg-gray-50 rounded-md transition duration-200"
              >
                <Link
                  href={`/playlists/${playlist.Slug}`}
                  className="text-lg font-medium text-blue-600 hover:underline"
                >
                  {playlist.Name}
                </Link>
                <button
                  onClick={() => handleDeletePlaylist(playlist.ID)}
                  className="text-red-500 hover:text-red-700 transition"
                >
                  Delete
                </button>
              </li>
            ))
          ) : (
            <p className="text-center text-gray-500">No playlists available</p>
          )}
        </ul>
      </div>
    </main>
  );
}
