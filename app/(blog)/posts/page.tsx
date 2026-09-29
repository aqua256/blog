import type { Metadata } from "next";
import Link from "next/link";
import { getPosts } from "@/lib/posts";
import { PromptLine } from "../_components/prompt-line";

export const metadata: Metadata = {
  title: "posts",
};

/**
 * The posts section as an `ls -t` of ~/posts:
 *
 *   ~/posts $ ls -t
 *      2026-09-29   Hello, world                          2 min
 *      2026-09-01   Designing APIs That Read Like Prose   6 min
 */
export default async function PostsPage() {
  const posts = await getPosts();

  return (
    <main className="px-(--gutter) pt-12 pb-18">
      <div className="page-width mx-auto grid grid-cols-[minmax(0,var(--col))] gap-x-18 min-[1180px]:grid-cols-[minmax(0,var(--col))_220px]">
        <section aria-label="Posts">
          <PromptLine cwd="~/posts" command="ls" args="-t" />

          {/* Command output is indented one gutter, like the post page */}
          <div className="ml-12 max-[560px]:ml-6.5">
            {posts.length === 0 ? (
              <p className="text-dim">total 0</p>
            ) : (
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
            )}
          </div>
        </section>
      </div>
    </main>
  );
}
