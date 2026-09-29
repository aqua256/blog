import Link from "next/link";
import type { PostSummary } from "@/lib/posts";

/**
 * Posts as `ls` output, one per line:
 *
 *   2026-09-29   Hello, world                          2 min
 *   2026-09-01   Designing APIs That Read Like Prose   6 min
 */
export function PostList({ posts }: { posts: PostSummary[] }) {
  if (posts.length === 0) return <p className="text-dim">total 0</p>;

  return (
    <ul className="flex flex-col gap-3 max-[560px]:gap-5">
      {posts.map((post) => (
        <li
          key={post.slug}
          className="grid grid-cols-[auto_minmax(0,1fr)_auto] items-baseline gap-x-6 max-[560px]:grid-cols-[1fr_auto] max-[560px]:gap-y-1"
        >
          <time dateTime={post.date} className="text-dim tabular-nums">
            {post.date}
          </time>
          <Link
            href={`/posts/${post.slug}`}
            className="text-link underline-offset-3 hover:underline max-[560px]:col-span-2 max-[560px]:row-start-2"
          >
            {post.title}
          </Link>
          <span className="whitespace-nowrap text-right text-dim max-[560px]:col-start-2 max-[560px]:row-start-1">
            {post.readingMinutes} min
            {post.draft && " · draft"}
          </span>
        </li>
      ))}
    </ul>
  );
}
