/** Things Aqua256 has built, shown in ~/projects and by the agent. Newest first. */
export type Project = {
  /** Shown as a directory: `blog/` */
  name: string;
  /** One line about what it is */
  description: string;
  year: number;
  /** Where the project lives, if anywhere */
  href?: string;
  /** Extra links, e.g. { label: "src", href: "https://github.com/…" } */
  links?: { label: string; href: string }[];
};

// TODO: add the source link once the repository is public
export const projects: Project[] = [
  {
    name: "blog",
    description: "this site: Next.js, MDX and a scripted agent",
    year: 2026,
    href: "/",
  },
];
