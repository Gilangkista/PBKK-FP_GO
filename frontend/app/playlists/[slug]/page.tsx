'use client';

import { useState, useEffect } from 'react';
import { useParams } from 'next/navigation';
import axios from 'axios';

type Song = {
  Slug: string;
  Title: string;
  Artist: { Name: string };
  Category: { Name: string };
};

type PlaylistDetail = {
  ID: number;
  Name: string;
  Description?: string;
  Songs: Song[];
};

export default function PlaylistDetail() {
  const params = useParams();
  const [slug, setSlug] = useState<string | null>(null);
  const [playlist, setPlaylist] = useState<PlaylistDetail | null>(null);
  const [description, setDescription] = useState<string>('');
  const [songs, setSongs] = useState<Song[]>([]);
  const [selectedSongSlug, setSelectedSongSlug] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

  useEffect(() => {
    if (typeof params?.slug === 'string') {
      setSlug(params.slug);
    } else {
      setError('Invalid playlist slug.');
    }
  }, [params]);

  // Ambil data playlist dan daftar lagu yang bisa dipilih
  useEffect(() => {
    const fetchPlaylistAndSongs = async () => {
      if (!slug) return;

      try {
        // Ambil data playlist
        const playlistResponse = await axios.get(`http://localhost:8080/playlists/detail/${slug}`);
        setPlaylist(playlistResponse.data);
        setDescription(playlistResponse.data.Description || '');
        
        // Ambil daftar semua lagu
        const songsResponse = await axios.get('http://localhost:8080/songs/');
        setSongs(songsResponse.data); // Semua lagu yang tersedia

      } catch (err) {
        console.error('Error fetching playlist details:', err);
        setError('Failed to load playlist details.');
      }
    };

    fetchPlaylistAndSongs();
  }, [slug]);

  // Fungsi untuk mendapatkan lagu yang belum ada di playlist
  const getAvailableSongs = () => {
    if (!playlist) return []; // Jika playlist belum ada, return array kosong
    
    // Filter lagu yang belum ada di playlist
    return songs.filter(song => !playlist.Songs.some(playlistSong => playlistSong.Slug === song.Slug));
  };

  // Fungsi untuk memperbarui deskripsi playlist
  const handleUpdateDescription = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!slug || !playlist) return;

    try {
      await axios.put(`http://localhost:8080/playlists/detail/${slug}`, {
        Name: playlist.Name,
        Description: description,
      });
      setSuccess('Description updated successfully!');
      setError(null);
    } catch (err) {
      console.error('Error updating description:', err);
      setError('Failed to update description.');
      setSuccess(null);
    }
  };

  // Fungsi untuk menambahkan lagu ke playlist
  const handleAddSong = async () => {
    if (!slug || !selectedSongSlug) return;

    try {
        // Mengirim data songSlug ke backend
        const response = await axios.post(`http://localhost:8080/playlists/detail/${slug}/songs`, null, {
            params: { songSlug: selectedSongSlug },
        });

        // Menampilkan response dari backend (jika perlu debugging)
        console.log('Add song response:', response.data);

        // Menambahkan lagu baru ke playlist yang ada di state
        setSuccess('Song added successfully!');
        setError(null);
        setSelectedSongSlug(null); // Reset pilihan lagu

        // Memperbarui data playlist
        const updatedPlaylistResponse = await axios.get(`http://localhost:8080/playlists/detail/${slug}`);
        setPlaylist(updatedPlaylistResponse.data);

    } catch (err) {
        console.error('Error adding song:', err);
        setError('Failed to add song.');
        setSuccess(null);
    }
};

  // Fungsi untuk menghapus lagu dari playlist
  const handleRemoveSong = async (songSlug: string) => {
    if (!slug) return;
  
    try {
      // Mengirim request DELETE untuk menghapus lagu dari playlist
      const response = await axios.delete(`http://localhost:8080/playlists/detail/${slug}/songs`, { data: { songSlug } });
  
      // Menampilkan response dari backend (jika perlu debugging)
      console.log('Remove song response:', response.data);
  
      // Menambahkan pesan sukses
      setSuccess('Song removed successfully!');
      setError(null);
  
      // Memperbarui data playlist setelah menghapus lagu
      const updatedPlaylistResponse = await axios.get(`http://localhost:8080/playlists/detail/${slug}`);
      setPlaylist(updatedPlaylistResponse.data);
  
    } catch (err) {
      console.error('Error removing song:', err);
      setError('Failed to remove song.');
      setSuccess(null);
    }
  };  

  if (!playlist) {
    return (
      <main className="p-4">
        <h1 className="text-2xl font-bold">Loading...</h1>
        {error && <p className="text-red-500">{error}</p>}
      </main>
    );
  }

  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold mb-2">{playlist.Name}</h1>
      {success && <p className="text-green-500 mb-2">{success}</p>}
      {error && <p className="text-red-500 mb-2">{error}</p>}

      <form onSubmit={handleUpdateDescription} className="mb-4">
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="Update playlist description"
          className="border p-2 w-full mb-2"
        />
        <button type="submit" className="bg-blue-500 text-white p-2">
          Update Description
        </button>
      </form>

      {/* Menambahkan lagu baru */}
      <div className="mb-4">
        <h2 className="text-lg font-semibold">Add Song to Playlist</h2>
        <select
          value={selectedSongSlug || ''}
          onChange={(e) => setSelectedSongSlug(e.target.value)}
          className="border p-2 mb-2 w-full"
        >
          <option value="">Select a Song</option>
          {getAvailableSongs().length > 0 ? (
            getAvailableSongs().map((song) => (
              <option key={song.Slug} value={song.Slug}>
                {song.Title} - {song.Artist.Name}
              </option>
            ))
          ) : (
            <option>No songs available</option>
          )}
        </select>
        <button
          type="button"
          onClick={handleAddSong}
          className="bg-green-500 text-white p-2"
        >
          Add Song
        </button>
      </div>

      {/* Menampilkan daftar lagu di playlist */}
      <h2 className="text-lg font-semibold mt-4">Songs in this Playlist:</h2>
      <ul className="space-y-1">
        {playlist.Songs.length > 0 ? (
          playlist.Songs.map((song) => (
            <li key={song.Slug}>
              {song.Title} - {song.Artist.Name} ({song.Category.Name})
              <button
                onClick={() => handleRemoveSong(song.Slug)}
                className="ml-2 text-red-500"
              >
                Remove
              </button>
            </li>
          ))
        ) : (
          <li>No songs available in this playlist.</li>
        )}
      </ul>
    </main>
  );
}
