export const site = {
  /** The author's name: page titles, the RSS feed, anywhere people should recognize the author */
  name: "Aqua256",
  /** How the agent refers to the author in a sentence ("What Aqua256 is building") */
  givenName: "Aqua256",
  /** The shell prompt in the header, `user@host` */
  user: "aqua256",
  host: "blog",
  // Shown under `fastfetch` on the home page, by the agent, in the RSS feed and as the page description
  intro:
    "这个博客主要写我做东西的过程：一个想法怎么一步步变成能跑起来的代码，中间走过哪些弯路、做过哪些取舍；其余的是读文档、踩坑时顺手记下的技术笔记。",
  // Lines of `fastfetch` on the home page, before Uptime and Links
  profile: [
    { key: "Role", value: "学生" },
    { key: "Skills", value: "TypeScript" },
  ],
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
 * SITE_URL (set as a build variable on Cloudflare), else the local dev server.
 */
export const siteUrl = (process.env.SITE_URL ?? "http://localhost:3000").replace(/\/+$/, ""); // "https://example.com/" → "https://example.com", so paths can be appended
