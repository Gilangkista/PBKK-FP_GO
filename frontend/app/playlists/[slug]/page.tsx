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

  useEffect(() => {
    const fetchPlaylistAndSongs = async () => {
      if (!slug) return;

      try {
        const playlistResponse = await axios.get(`http://localhost:8080/playlists/detail/${slug}`);
        setPlaylist(playlistResponse.data);
        setDescription(playlistResponse.data.Description || '');
        
        const songsResponse = await axios.get('http://localhost:8080/songs/');
        setSongs(songsResponse.data);

      } catch (err) {
        console.error('Error fetching playlist details:', err);
        setError('Failed to load playlist details.');
      }
    };

    fetchPlaylistAndSongs();
  }, [slug]);

  const getAvailableSongs = () => {
    if (!playlist) return [];
    
    return songs.filter(song => !playlist.Songs.some(playlistSong => playlistSong.Slug === song.Slug));
  };

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

  const handleAddSong = async () => {
    if (!slug || !selectedSongSlug) return;

    try {
        const response = await axios.post(`http://localhost:8080/playlists/detail/${slug}/songs`, null, {
            params: { songSlug: selectedSongSlug },
        });

        console.log('Add song response:', response.data);

        setSuccess('Song added successfully!');
        setError(null);
        setSelectedSongSlug(null);

        const updatedPlaylistResponse = await axios.get(`http://localhost:8080/playlists/detail/${slug}`);
        setPlaylist(updatedPlaylistResponse.data);

    } catch (err) {
        console.error('Error adding song:', err);
        setError('Failed to add song.');
        setSuccess(null);
    }
  };

  const handleRemoveSong = async (songSlug: string) => {
    if (!slug) return;
  
    try {
      const response = await axios.delete(`http://localhost:8080/playlists/detail/${slug}/songs/${songSlug}`);
      console.log('Remove song response:', response.data);

      setSuccess('Song removed successfully!');
      setError(null);
      
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
      <main className="bg-gray-100 min-h-screen flex items-center justify-center">
        <div className="text-center">
          <h1 className="text-2xl font-bold">Loading...</h1>
          {error && <p className="text-red-500 mt-2">{error}</p>}
        </div>
      </main>
    );
  }

  return (
    <main className="bg-gray-100 min-h-screen flex flex-col items-center p-6">
      <div className="bg-white shadow-md rounded-lg p-6 w-full max-w-4xl">
        <h1 className="text-3xl font-bold mb-6 text-center text-gray-800">{playlist.Name}</h1>
        {success && <p className="text-green-500 text-center mb-4">{success}</p>}
        {error && <p className="text-red-500 text-center mb-4">{error}</p>}

        <form onSubmit={handleUpdateDescription} className="mb-6 bg-gray-50 p-4 rounded-lg shadow">
          <h2 className="text-xl font-semibold mb-4 text-gray-700">Update Description</h2>
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Update playlist description"
            className="border border-gray-300 rounded p-2 w-full mb-2 focus:outline-none focus:ring-2 focus:ring-blue-400 text-black"
          />
          <button
            type="submit"
            className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600 transition"
          >
            Update Description
          </button>
        </form>

        <div className="mb-6">
          <h2 className="text-xl font-semibold mb-4 text-gray-700">Add Song to Playlist</h2>
          <select
            value={selectedSongSlug || ''}
            onChange={(e) => setSelectedSongSlug(e.target.value)}
            className="border border-gray-300 rounded p-2 w-full mb-2 focus:outline-none focus:ring-2 focus:ring-green-400 text-black"
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
            className="w-full bg-green-500 text-white py-2 rounded hover:bg-green-600 transition"
          >
            Add Song
          </button>
        </div>

        <h2 className="text-xl font-semibold mt-6 mb-4 text-gray-700">Songs in this Playlist:</h2>
        <ul className="divide-y divide-gray-200">
          {playlist.Songs.length > 0 ? (
            playlist.Songs.map((song) => (
              <li
                key={song.Slug}
                className="py-4 flex justify-between items-center hover:bg-gray-50 rounded-md transition duration-200"
              >
                <div className='text-black'>
                  {song.Title}
                </div>
                <button
                  onClick={() => handleRemoveSong(song.Slug)}
                  className="text-red-500 hover:text-red-700 transition"
                >
                  Remove
                </button>
              </li>
            ))
          ) : (
            <p className="text-center text-gray-500">No songs available in this playlist.</p>
          )}
        </ul>
      </div>
    </main>
  );
}
