import type { Metadata } from "next";
import { getPosts } from "@/lib/posts";
import { site, siteUrl } from "@/lib/site";
import { AgentConsole } from "./_components/agent-console";
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

/** Terminal-style chrome for every page of the site, including the 404 page. */
export default async function RootLayout({ children }: LayoutProps<"/">) {
  // What the agent needs to know about the posts: enough to list them and summarize the current one
  const posts = (await getPosts()).map(({ slug, title, headings }) => ({ slug, title, headings }));

  return (
    <html
      lang="zh-CN"
      className="h-full antialiased"
      suppressHydrationWarning
    >
      <body className="flex min-h-full flex-col bg-bg font-sans text-ui text-ink selection:bg-selection">
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
