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
    <html lang="en" className="h-full bg-gray-100">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased h-full`}>
        {/* Navbar */}
        <div className="min-h-full">
          <nav className="bg-gray-800">
            <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
              <div className="flex h-16 items-center justify-between">
                <div className="flex items-center">
                  <div className="shrink-0">
                    <img
                      className="h-8 w-8"
                      src="https://tailwindui.com/plus/img/logos/mark.svg?color=indigo&shade=500"
                      alt="Your Company"
                    />
                  </div>
                  <div className="hidden md:block">
                    <div className="ml-10 flex items-baseline space-x-4">
                      <Link href="/" className="hover:text-gray-300">Home</Link>
                      <Link href="/songs" className="hover:text-gray-300">Songs</Link>
                      <Link href="/artists" className="hover:text-gray-300">Artists</Link>
                      <Link href="/playlists" className="hover:text-gray-300">Playlists</Link>
                    </div>
                  </div>
                </div>
                <div className="-mr-2 flex md:hidden">
                  <button
                    type="button"
                    className="relative inline-flex items-center justify-center rounded-md bg-gray-800 p-2 text-gray-400 hover:bg-gray-700 hover:text-white focus:outline-none focus:ring-2 focus:ring-white focus:ring-offset-2 focus:ring-offset-gray-800"
                    aria-controls="mobile-menu"
                    aria-expanded="false"
                  >
                    <span className="absolute -inset-0.5"></span>
                    <span className="sr-only">Open main menu</span>
                    <svg
                      className="block h-6 w-6"
                      fill="none"
                      viewBox="0 0 24 24"
                      strokeWidth="1.5"
                      stroke="currentColor"
                      aria-hidden="true"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
            <div className="md:hidden" id="mobile-menu">
              <div className="space-y-1 px-2 pb-3 pt-2 sm:px-3">
                <Link href="/" className="hover:text-gray-300">Home</Link>
                <Link href="/songs" className="hover:text-gray-300">Songs</Link>
                <Link href="/artists" className="hover:text-gray-300">Artists</Link>
                <Link href="/playlists" className="hover:text-gray-300">Playlists</Link>
              </div>
            </div>
          </nav>

          <main>
            <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
              {children}
            </div>
          </main>
        </div>
      </body>
    </html>
  );
}
