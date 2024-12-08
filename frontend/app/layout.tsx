import type { Metadata } from "next";
import localFont from "next/font/local";
import Link from "next/link"; // Import Link untuk navigasi antar halaman
import "./globals.css";

const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export const metadata: Metadata = {
  title: "My Next.js App",
  description: "An awesome Next.js app with a custom navbar.",
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        {/* Navbar */}
        <nav className="bg-gray-800 text-white p-4">
          <div className="max-w-7xl mx-auto flex justify-between items-center">
            <div>
              <Link href="/" className="text-2xl font-semibold">
                My App
              </Link>
            </div>
            <div className="space-x-4">
              <Link href="/" className="hover:text-gray-300">Home</Link>
              <Link href="/songs" className="hover:text-gray-300">Songs</Link>
              <Link href="/artists" className="hover:text-gray-300">Artists</Link>
              <Link href="/playlists" className="hover:text-gray-300">Playlists</Link>
            </div>
          </div>
        </nav>

        {/* Content */}
        <main>{children}</main>
      </body>
    </html>
  );
}
