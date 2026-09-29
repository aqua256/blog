import type { Metadata } from "next";
import { Geist, Geist_Mono, Recursive } from "next/font/google";
import { site } from "@/lib/site";
import "./globals.css";
import { Providers } from "./providers";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const recursive = Recursive({
  variable: "--font-recursive",
  subsets: ["latin"],
  axes: ["CASL", "MONO", "slnt"],
});

export const metadata: Metadata = {
  // "aqua256@blog" on the home page, "Hello, world · aqua256@blog" elsewhere
  title: {
    default: `${site.user}@${site.host}`,
    template: `%s · ${site.user}@${site.host}`,
  },
  description: site.intro,
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html
      lang="en"
      className={`${geistSans.variable} ${geistMono.variable} ${recursive.variable} h-full antialiased`}
      suppressHydrationWarning
    >
      <body className="min-h-full flex flex-col">
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
