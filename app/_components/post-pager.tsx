import Link from "next/link";
import type { PostSummary } from "@/lib/posts";

const arrow = {
  viewBox: "0 0 24 24",
  fill: "none",
  stroke: "currentColor",
  strokeWidth: 1.75,
  strokeLinecap: "round",
  strokeLinejoin: "round",
  "aria-hidden": true,
  className: "size-3.5",
} as const;

const link = "flex max-w-[46%] flex-col gap-1.25 py-1 group";
const label = "inline-flex items-center gap-1.25 text-[13px] leading-normal text-muted";
const title = "text-[15px] leading-normal font-semibold text-muted group-hover:text-accent";

/**
 * The newer post on the left (上一篇) and the older one on the right (下一篇), the way the archive reads
 * from newest down to oldest.
 */
export function PostPager({ older, newer }: { older?: PostSummary; newer?: PostSummary }) {
  if (!older && !newer) return null;

  return (
    <nav aria-label="相邻文章" className="mt-14.5 flex justify-between gap-5.5 border-t border-line pt-5.5">
      {newer && (
        <Link href={`/posts/${newer.slug}`} className={link}>
          <span className={label}>
            <svg {...arrow}>
              <path d="M19 12H5M12 19l-7-7 7-7" />
            </svg>
            上一篇
          </span>
          <span className={title}>{newer.title}</span>
        </Link>
      )}
      {older && (
        <Link href={`/posts/${older.slug}`} className={`${link} ml-auto items-end text-right`}>
          <span className={label}>
            下一篇
            <svg {...arrow}>
              <path d="M5 12h14M12 5l7 7-7 7" />
            </svg>
          </span>
          <span className={title}>{older.title}</span>
        </Link>
      )}
    </nav>
  );
}
