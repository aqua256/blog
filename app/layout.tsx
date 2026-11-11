import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { site, siteUrl } from "@/lib/site";
import { SiteFooter } from "./_components/site-footer";
import { SiteHeader } from "./_components/site-header";
import "./globals.css";
import { Providers } from "./providers";

// Inter Variable, served from this site at build time. The optical size axis comes along so each weight can
// use its own (globals.css); Chinese text falls through to the system sans.
const inter = Inter({
  subsets: ["latin"],
  axes: ["opsz"],
  variable: "--font-inter",
});

export const metadata: Metadata = {
  metadataBase: new URL(siteUrl),
  // "Aqua256" on the home page, "Hello, world · Aqua256" elsewhere
  title: {
    default: site.name,
    template: `%s · ${site.name}`,
  },
  description: site.intro,
  // Lets browsers and feed readers find the RSS feed from any page
  alternates: {
    types: { "application/rss+xml": "/rss.xml" },
  },
};

/** The header and footer around every page of the site, including the 404 page. */
export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html
      lang="zh-CN"
      className={`${inter.variable} h-full antialiased`}
      suppressHydrationWarning
    >
      <body className="flex min-h-full flex-col bg-bg font-sans text-ink selection:bg-selection">
        <Providers>
          <SiteHeader />
          <div className="flex-1">{children}</div>
          <SiteFooter />
        </Providers>
      </body>
    </html>
  );
}
