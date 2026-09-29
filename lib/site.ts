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
  // TODO: fill in real profile URLs
  social: [
    { label: "linkedin", href: "#" },
    { label: "github", href: "#" },
    { label: "twitter", href: "#" },
    { label: "rss", href: "#" },
  ],
} as const;
