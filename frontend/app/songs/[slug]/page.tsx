import axios from 'axios';

type Song = {
  ID: number;
  Slug: string;
  Title: string;
  Artist: { Name: string };
  Category: { Name: string };
};

export default async function SongDetails({ params }: { params: { slug: string } }) {
  const { slug } = params;

  try {
    const response = await axios.get(`http://localhost:8080/songs/${slug}`);
    const song: Song = response.data;

    return (
      <main className="p-4">
        <h1 className="text-2xl font-bold">{song.Title}</h1>
        <p className="text-lg">Artist: {song.Artist.Name}</p>
        <p className="text-lg">Category: {song.Category.Name}</p>
      </main>
    );
  } catch (error) {
    console.error('Error fetching song details:', error);
    return (
      <main className="p-4">
        <h1 className="text-2xl font-bold">Song Not Found</h1>
        <p className="text-red-500">The song you are looking for does not exist.</p>
      </main>
    );
  }
}
