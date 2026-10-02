import type { Metadata } from "next";
import { getPosts } from "@/lib/posts";
import { PageTransition } from "../_components/page-transition";
import { PostList } from "../_components/post-list";
import { PromptLine } from "../_components/prompt-line";

export const metadata: Metadata = {
  title: "posts",
};

/**
 * The posts section as an `ls -t` of ~/posts:
 *
 *   ~/posts $ ls -t
 *      2026-09-29   Hello, world
 *      2026-09-01   Designing APIs That Read Like Prose
 */
export default async function PostsPage() {
  const posts = await getPosts();

  return (
    <PageTransition>
      <main className="px-(--gutter) pt-12 pb-18">
        <div className="page-width mx-auto grid grid-cols-[minmax(0,var(--col))] gap-x-18 min-[1180px]:grid-cols-[minmax(0,var(--col))_220px]">
          <section aria-label="Posts">
            <PromptLine cwd="~/posts" command="ls" args="-t" />

            {/* Command output is indented one gutter, like the post page */}
            <div className="ml-12 max-[560px]:ml-6.5">
              <PostList posts={posts} />
            </div>
          </section>
        </div>
      </main>
    </PageTransition>
  );
}
