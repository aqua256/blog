export const site = {
  user: "aqua256",
  host: "blog",
  // TODO: Aqua256's own introduction, shown by `whoami` on the home page and by the agent
  intro: "This introduction is a placeholder until Aqua256 writes one.",
  nav: [
    { path: "~", href: "/" },
    { path: "~/projects", href: "/projects" },
    { path: "~/posts", href: "/posts" },
  ],
  social: [
    // TODO: fill in real profile URLs
    { label: "linkedin", href: "#" },
    { label: "github", href: "#" },
    { label: "twitter", href: "#" },
    { label: "rss", href: "/rss.xml" },
  ],
} as const;

/**
 * The site's public address, for absolute links (RSS, metadata). Read at build time on the server:
 * SITE_URL if set, else Vercel's production domain, else the local dev server.
 */
export const siteUrl =
  process.env.SITE_URL ??
  (process.env.VERCEL_PROJECT_PRODUCTION_URL
    ? `https://${process.env.VERCEL_PROJECT_PRODUCTION_URL}`
    : "http://localhost:3000");
