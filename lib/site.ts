export const site = {
  /** Full name: page titles, the RSS feed, anywhere people should recognize the author */
  name: "Aqua256",
  /** How the agent refers to the author in a sentence ("What Aqua256 is building") */
  givenName: "Aqua256",
  /** The shell prompt in the header, `user@host` */
  user: "aqua256",
  host: "blog",
  // Shown by `whoami` on the home page, by the agent, in the RSS feed and as the page description
  intro:
    "我是一名后端开发工程师，同时也还是个学生。这里写下的，大多是我一边学一边做东西的过程：为什么想做它，怎么一点一点把它做出来，做完之后回头又想明白了什么；偶尔也会留下一些零散的技术笔记，算是给以后的自己翻看。",
  nav: [
    { path: "~", href: "/" },
    { path: "~/projects", href: "/projects" },
    { path: "~/posts", href: "/posts" },
  ],
  social: [
    { label: "linkedin", href: "https://www.linkedin.com/in/sheng-xiao-b20846367" },
    { label: "github", href: "https://github.com/aqua256" },
    { label: "rss", href: "/rss.xml" },
  ],
} as const;

/**
 * The site's public address, for absolute links (RSS, metadata). Read at build time on the server:
 * SITE_URL if set, else Vercel's production domain, else the local dev server.
 */
export const siteUrl = (
  process.env.SITE_URL ??
  (process.env.VERCEL_PROJECT_PRODUCTION_URL
    ? `https://${process.env.VERCEL_PROJECT_PRODUCTION_URL}`
    : "http://localhost:3000")
).replace(/\/+$/, ""); // "https://example.com/" → "https://example.com", so paths can be appended
