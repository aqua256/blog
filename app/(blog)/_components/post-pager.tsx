import Link from "next/link";
import type { PostSummary } from "@/lib/posts";

/** Older post on the left, newer on the right; stacked on phones. */
export function PostPager({ older, newer }: { older?: PostSummary; newer?: PostSummary }) {
  if (!older && !newer) return null;

  return (
    <nav
      aria-label="More writing"
      className="mt-5 grid grid-cols-2 gap-x-8 gap-y-4 border-t border-dashed border-faint pt-4.5 text-small max-[560px]:grid-cols-1"
    >
      {older ? (
        <Link href={`/posts/${older.slug}`} className="group flex flex-col gap-1">
          <span className="text-dim">← older</span>
          <span className="leading-snug text-link underline-offset-3 group-hover:underline">{older.title}</span>
        </Link>
      ) : (
        <span />
      )}
      {newer && (
        <Link
          href={`/posts/${newer.slug}`}
          className="group flex flex-col items-end gap-1 text-right max-[560px]:items-start max-[560px]:text-left"
        >
          <span className="text-dim">newer →</span>
          <span className="leading-snug text-link underline-offset-3 group-hover:underline">{newer.title}</span>
        </Link>
      )}
    </nav>
  );
}
