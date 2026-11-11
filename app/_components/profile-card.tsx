import Image from "next/image";
import { site } from "@/lib/site";
import { FluidHover } from "./fluid-hover";
// The same picture as the Apple touch icon
import avatar from "../apple-icon.png";

const icon = {
  viewBox: "0 0 24 24",
  fill: "none",
  stroke: "currentColor",
  strokeLinecap: "round",
  strokeLinejoin: "round",
  "aria-hidden": true,
  className: "size-4",
} as const;

const icons = {
  github: (
    <svg {...icon}>
      <path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.4 5.4 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4" />
      <path d="M9 18c-4.51 2-5-2-7-2" />
    </svg>
  ),
  mail: (
    <svg {...icon}>
      <rect x="2" y="4" width="20" height="16" rx="2" />
      <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
    </svg>
  ),
  rss: (
    <svg {...icon}>
      <path d="M4 11a9 9 0 0 1 9 9" />
      <path d="M4 4a16 16 0 0 1 16 16" />
      <circle cx="5" cy="19" r="1" />
    </svg>
  ),
};

/**
 * Avatar, name, the one-line introduction and the links, at the top of the side column. It sits on the page
 * without a card: it only shows who writes here, so nothing but the links reacts to the pointer.
 */
export function ProfileCard() {
  return (
    <section aria-label="关于作者" className="flex flex-col items-center pt-2 text-center">
      {/* A hairline drawn inside the picture's edge, so a light avatar keeps its outline on the light page */}
      <Image
        src={avatar}
        alt={`${site.name} 的头像`}
        width={112}
        height={112}
        className="rounded-xs outline-1 -outline-offset-1 outline-black/10 dark:outline-white/10"
      />
      <h2 className="type-subtitle mt-4 text-ink">{site.name}</h2>
      <p className="type-body mt-1 break-keep text-muted">{site.intro}</p>
      <nav aria-label="外部链接" className="mt-3">
        <FluidHover className="flex">
          {site.social.map(({ icon, label, title, href }) => (
            <a key={href} href={href} aria-label={label} title={title} className="icon-link grid size-9 place-items-center rounded-lg text-muted hover:text-ink">
              {icons[icon]}
            </a>
          ))}
        </FluidHover>
      </nav>
    </section>
  );
}
