import Link from "next/link";
import type { PostSummary } from "@/lib/posts";

/**
 * Posts as `ls` output, one per line:
 *
 *   2026-09-29   Hello, world
 *   2026-09-01   Designing APIs That Read Like Prose
 */
export function PostList({ posts }: { posts: PostSummary[] }) {
  if (posts.length === 0) return <p className="text-dim">total 0</p>;

  return (
    <ul className="flex flex-col gap-3 max-[560px]:gap-5">
      {posts.map((post) => (
        <li
          key={post.slug}
          className="grid grid-cols-[auto_minmax(0,1fr)] items-baseline gap-x-6 max-[560px]:grid-cols-1 max-[560px]:gap-y-1"
        >
          <time dateTime={post.date} className="text-dim tabular-nums">
            {post.date}
          </time>
          <span>
            <Link href={`/posts/${post.slug}`} className="text-link underline-offset-3 hover:underline">
              {post.title}
            </Link>
            {post.draft && <span className="text-dim"> · draft</span>}
          </span>
        </li>
      ))}
    </ul>
  );
}
