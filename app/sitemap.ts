import type { MetadataRoute } from "next";
import { getPosts } from "@/lib/posts";
import { siteUrl } from "@/lib/site";

// Built once at build time; required with output: "export"
export const dynamic = "force-static";

/**
 * /sitemap.xml: the section pages and every published post, built at build time.
 * Only url and lastModified; search engines ignore changeFrequency and priority.
 */
export default async function sitemap(): Promise<MetadataRoute.Sitemap> {
  const posts = await getPosts();
  const latest = posts[0]?.date;

  return [
    { url: `${siteUrl}/`, lastModified: latest },
    { url: `${siteUrl}/posts`, lastModified: latest },
    { url: `${siteUrl}/projects` },
    ...posts.map((post) => ({ url: `${siteUrl}/posts/${post.slug}`, lastModified: post.date })),
  ];
}
