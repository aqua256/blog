export const site = {
  /** The author's name: page titles, the RSS feed, anywhere people should recognize the author */
  name: "Aqua256",
  // Shown in the profile card, in the RSS feed and as the page description
  intro:
    "一块自留地，放些做过的东西和碎碎念。",
  nav: [
    { label: "首页", href: "/" },
    { label: "归档", href: "/posts" },
    { label: "关于", href: "/about" },
  ],
  // The profile card shows these as icons, with the label for screen readers and the title as a tooltip
  social: [
    { icon: "github", label: "GitHub", title: "GitHub", href: "https://github.com/aqua256" },
    { icon: "mail", label: "邮件", title: "maix9527@gmail.com", href: "mailto:maix9527@gmail.com" },
    { icon: "rss", label: "RSS", title: "RSS", href: "/rss.xml" },
  ],
} as const;

/**
 * The site's public address, for absolute links (RSS, metadata). Read at build time on the server:
 * SITE_URL (set as a build variable on Cloudflare), else the local dev server.
 */
export const siteUrl = (process.env.SITE_URL ?? "http://localhost:3000").replace(/\/+$/, ""); // "https://example.com/" → "https://example.com", so paths can be appended
