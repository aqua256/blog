import Link from "next/link";
import type { PostSummary } from "@/lib/posts";
import { PostTitleTransition } from "./page-transition";

/**
 * A post as a card: title, description, then the date on the left and the tags on the right. The whole
 * card opens the post (the title link stretches over it); the description, date and tags sit above that
 * layer so their text can still be selected.
 */
export function PostCard({ post }: { post: PostSummary }) {
  return (
    <article className="card post-card relative px-7 py-6 max-[640px]:p-5">
      <h3 className="text-[18px] leading-[1.4] font-semibold tracking-[-0.015em]">
        <PostTitleTransition slug={post.slug}>
          <Link href={`/posts/${post.slug}`} className="text-ink before:absolute before:inset-0 hover:text-accent">
            {post.title}
          </Link>
        </PostTitleTransition>
      </h3>
      <p className="relative z-1 mt-2.25 text-[15px] leading-[1.8] text-pretty text-ink-2">{post.description}</p>
      <div className="mt-3.5 flex flex-wrap items-center gap-x-4.5 gap-y-1.75 text-[13px] leading-normal">
        <p className="relative z-1 text-ink-3 tabular-nums">
          <time dateTime={post.date}>{post.date}</time>
          {post.draft && " · 草稿"}
        </p>
        {post.tags && post.tags.length > 0 && (
          <ul aria-label="标签" className="relative z-1 ml-auto flex gap-3.5 text-ink-2">
            {post.tags.map((tag) => (
              <li key={tag}>#{tag}</li>
            ))}
          </ul>
        )}
      </div>
    </article>
  );
}
