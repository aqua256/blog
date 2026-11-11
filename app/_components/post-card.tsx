import Link from "next/link";
import type { PostSummary } from "@/lib/posts";

/**
 * A post as a row on the home page: title, description, then the date on the left and the tags on the
 * right. The description is cut off with an ellipsis after two lines, and tags that don't fit beside the
 * date are left off. The whole row opens the post: the title link stretches over it, and its keyboard focus
 * ring goes around the row too. Rows sit on the page without a card; the list's hover highlight and the
 * hairlines between rows set them apart (.post-list in globals.css).
 */
export function PostCard({ post }: { post: PostSummary }) {
  return (
    <article className="relative px-3 py-4">
      <h3 className="type-subtitle text-balance text-ink">
        <Link
          href={`/posts/${post.slug}`}
          className="outline-none before:absolute before:inset-0 before:rounded-lg focus-visible:before:outline-2 focus-visible:before:outline-ring"
        >
          {post.title}
        </Link>
      </h3>
      <p className="type-body mt-1 line-clamp-2 text-pretty text-muted">{post.description}</p>
      <div className="type-caption mt-2 flex items-center gap-3 text-muted">
        <p className="shrink-0 tabular-nums">
          <time dateTime={post.date}>{post.date}</time>
          {post.draft && " · 草稿"}
        </p>
        {post.tags && post.tags.length > 0 && (
          // Tags that don't fit wrap onto a second row, which the one-line height hides
          <ul aria-label="标签" className="ml-auto flex h-4 min-w-0 flex-wrap justify-end gap-x-3 overflow-hidden">
            {post.tags.map((tag) => (
              <li key={tag}>#{tag}</li>
            ))}
          </ul>
        )}
      </div>
    </article>
  );
}
