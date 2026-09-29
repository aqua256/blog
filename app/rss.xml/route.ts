import { Feed } from "feed";
import { getPosts } from "@/lib/posts";
import { site, siteUrl } from "@/lib/site";

// Built once at build time, like the posts themselves
export const dynamic = "force-static";

/** RSS 2.0 feed of the published posts, newest first, with each post's summary. */
export async function GET() {
  const posts = await getPosts();

  const feed = new Feed({
    title: `${site.user}@${site.host}`,
    description: site.intro,
    id: `${siteUrl}/`,
    link: `${siteUrl}/`,
    language: "zh-CN",
    copyright: `© ${site.user}`,
    updated: posts[0] ? new Date(posts[0].date) : undefined,
    feedLinks: { rss: `${siteUrl}/rss.xml` },
    author: { name: site.user, link: `${siteUrl}/` },
  });

  for (const post of posts) {
    const url = `${siteUrl}/posts/${post.slug}`;
    feed.addItem({
      title: post.title,
      id: url,
      link: url,
      description: post.description,
      date: new Date(post.date),
      category: post.tags?.map((tag) => ({ name: tag })),
    });
  }

  return new Response(feed.rss2(), {
    headers: { "Content-Type": "application/rss+xml; charset=utf-8" },
  });
}
