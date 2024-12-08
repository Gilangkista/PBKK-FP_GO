import axios from 'axios';

type Song = {
  title: string;
  artist: { name: string };
  category: { name: string };
};

export default async function SongDetails({ params }: { params: { id: string } }) {
  const { id } = params;

  const response = await axios.get(`http://localhost:8080/songs/${id}`);
  const song: Song = response.data.song;

  return (
    <main className="p-4">
      <h1 className="text-2xl font-bold">{song.title}</h1>
      <p className="text-lg">Artist: {song.artist.name}</p>
      <p className="text-lg">Category: {song.category.name}</p>
    </main>
  );
}
