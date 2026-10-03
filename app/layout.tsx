import type { Metadata } from "next";
import localFont from "next/font/local";
import { getPosts } from "@/lib/posts";
import { site, siteUrl } from "@/lib/site";
import { AgentConsole } from "./_components/agent-console";
import { DeferredStylesheet } from "./_components/deferred-stylesheet";
import { SiteFooter } from "./_components/site-footer";
import { SiteHeader } from "./_components/site-header";
import "./globals.css";
import { Providers } from "./providers";

/**
 * Latin text: Recursive, self-hosted by next/font. It has no Chinese glyphs, so Chinese characters fall
 * through to LXGW WenKai (loaded below), then to the usual system fonts while it downloads.
 * The file is Google's Latin subset with each axis narrowed to what the site uses, about half the
 * size of the full font; scripts/subset-recursive.sh rebuilds it.
 */
const recursive = localFont({
  src: "./_fonts/recursive.woff2",
  weight: "380 800",
  variable: "--font-recursive",
  fallback: ["LXGW WenKai", "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "Noto Sans CJK SC", "sans-serif"],
});

/**
 * Chinese text: LXGW WenKai (simplified), cut into ~220 unicode-range slices so a page only downloads
 * the characters it uses. Loaded without blocking the first paint, since it comes from a CDN.
 * Pinned version; globals.css points at the same version for punctuation.
 */
const LXGW_WENKAI_CSS =
  "https://cdn.jsdelivr.net/npm/@callmebill/lxgw-wenkai-web@1.522.0/lxgwwenkai-regular/result.css";

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

/** Terminal-style chrome for every page of the site, including the 404 page. */
export default async function RootLayout({ children }: LayoutProps<"/">) {
  // What the agent needs to know about the posts: enough to list them and summarize the current one
  const posts = (await getPosts()).map(({ slug, title, headings }) => ({ slug, title, headings }));

  return (
    <html
      lang="zh-CN"
      className={`${recursive.variable} h-full antialiased`}
      suppressHydrationWarning
    >
      <body className="term flex min-h-full flex-col bg-paper font-term text-ui text-ink selection:bg-faint">
        <DeferredStylesheet href={LXGW_WENKAI_CSS} />
        <Providers>
          <AgentConsole posts={posts}>
            <SiteHeader />
            <div className="flex-1">{children}</div>
            <SiteFooter />
          </AgentConsole>
        </Providers>
      </body>
    </html>
  );
}
