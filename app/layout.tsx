import type { Metadata } from "next";
import { site, siteUrl } from "@/lib/site";
import { Leaves } from "./_components/leaves";
import { introScript, LeavesIntro } from "./_components/leaves-intro";
import { SiteFooter } from "./_components/site-footer";
import { SiteHeader } from "./_components/site-header";
import "./globals.css";
import { Providers } from "./providers";

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

/** The header, leaf shadow and footer around every page of the site, including the 404 page. */
export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html
      lang="zh-CN"
      className="h-full antialiased"
      suppressHydrationWarning
    >
      <head>
        {/* Before the first paint: whether this page opens with the leaf shadow coming into focus */}
        <script dangerouslySetInnerHTML={{ __html: introScript }} />
      </head>
      <body className="relative flex min-h-full flex-col bg-bg font-sans text-ink selection:bg-selection">
        <Providers>
          <Leaves />
          <LeavesIntro />
          <SiteHeader />
          {/* Positioned above the leaf shadow, which lies on the page background; both fade in with the intro */}
          <div data-intro="fade" className="relative z-1 flex-1">
            {children}
          </div>
          <div data-intro="fade" className="relative z-1">
            <SiteFooter />
          </div>
        </Providers>
      </body>
    </html>
  );
}
