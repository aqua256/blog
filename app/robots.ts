import type { MetadataRoute } from "next";
import { siteUrl } from "@/lib/site";

/** /robots.txt: everything may be crawled; points crawlers at the sitemap. */
export default function robots(): MetadataRoute.Robots {
  return {
    rules: { userAgent: "*", allow: "/" },
    sitemap: `${siteUrl}/sitemap.xml`,
  };
}
