import Link from "next/link";
import type { PostSummary } from "@/lib/posts";

const arrow = {
  viewBox: "0 0 24 24",
  fill: "none",
  stroke: "currentColor",
  strokeLinecap: "round",
  strokeLinejoin: "round",
  "aria-hidden": true,
  className: "size-3",
} as const;

// Each link fills on hover; the negative margin keeps its text on the column's edge
const link = "flex max-w-[46%] flex-col gap-1 rounded-lg px-3 py-2 hover:bg-hover";
const label = "type-caption inline-flex items-center gap-1 text-muted";
const title = "type-subtitle text-ink";

/**
 * The newer post on the left (上一篇) and the older one on the right (下一篇), the way the archive reads
 * from newest down to oldest.
 */
export function PostPager({ older, newer }: { older?: PostSummary; newer?: PostSummary }) {
  if (!older && !newer) return null;

  return (
    <nav aria-label="相邻文章" className="mt-10 flex justify-between gap-4 pt-4 shadow-[inset_0_1px_0_var(--line)]">
      {newer && (
        <Link href={`/posts/${newer.slug}`} className={`${link} -ml-3`}>
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
        <Link href={`/posts/${older.slug}`} className={`${link} -mr-3 ml-auto items-end text-right`}>
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
