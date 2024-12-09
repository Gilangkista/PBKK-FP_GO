import axios from 'axios';

type Song = {
  Title: string;
  Artist: { Name: string };
  Category: { Name: string };
};

type PlaylistDetail = {
  ID: number;
  Name: string;
  Description?: string;
  Songs: Song[];  // List of songs in the playlist
};

export default async function PlaylistDetail({ params }: { params: { slug: string } }) {
  const { slug } = params;

  try {
    // Mengambil data playlist berdasarkan slug
    const response = await axios.get(`http://localhost:8080/playlists/${slug}`);
    const playlist: PlaylistDetail = response.data;

    return (
      <main className="p-4">
        <h1 className="text-2xl font-bold">{playlist.Name}</h1>
        {playlist.Description && <p>{playlist.Description}</p>}

        <h2 className="mt-4 text-lg">Songs in this Playlist:</h2>
        <ul>
          {playlist.Songs.map((song, index) => (
            <li key={index}>
              {song.Title} - {song.Artist.Name} ({song.Category.Name})
            </li>
          ))}
        </ul>
      </main>
    );
  } catch (error) {
    console.error('Error fetching playlist details:', error);
    return (
      <main className="p-4">
        <h1 className="text-2xl font-bold">Playlist Not Found</h1>
        <p className="text-red-500">The playlist you are looking for does not exist.</p>
      </main>
    );
  }
}
