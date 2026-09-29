import Link from "next/link";
import { getPosts } from "@/lib/posts";
import { site } from "@/lib/site";
import { AskHint } from "./_components/ask-button";
import { PostList } from "./_components/post-list";
import { PromptLine } from "./_components/prompt-line";

const RECENT_POSTS = 5;

/**
 * The home directory:
 *
 *   ~ $ whoami
 *      aqua256
 *      <a few lines about Aqua256>
 *
 *   ~ $ ls -t ~/posts | head -5
 *      2026-09-29   Hello, world   2 min
 *      cd ~/posts                        (only when there are more)
 *
 *   ~ $ ask
 *      press / to ask anything about Aqua256
 */
export default async function HomePage() {
  const posts = await getPosts();

  return (
    <main className="px-(--gutter) pt-12 pb-18">
      <div className="page-width mx-auto">
        <div className="flex max-w-(--col) flex-col gap-12">
          {/* Command output is indented one gutter, like the post page */}
          <section aria-label="About">
            <PromptLine cwd="~" command="whoami" />
            <div className="ml-12 max-[560px]:ml-6.5">
              <h1 className="font-bold">{site.user}</h1>
              <p className="mt-1.5 font-[380] text-body leading-[1.7] [font-variation-settings:'MONO'_0,'CASL'_0.3]">
                {site.intro}
              </p>
            </div>
          </section>

          <section aria-label="Recent posts">
            <PromptLine cwd="~" command="ls" args={`-t ~/posts | head -${RECENT_POSTS}`} />
            <div className="ml-12 max-[560px]:ml-6.5">
              <PostList posts={posts.slice(0, RECENT_POSTS)} />
              {posts.length > RECENT_POSTS && (
                <Link href="/posts" className="group mt-4 inline-block text-dim hover:text-ink">
                  <span className="font-semibold text-cmd">cd</span>{" "}
                  <span className="underline-offset-4 group-hover:underline">~/posts</span> for all {posts.length}
                </Link>
              )}
            </div>
          </section>

          <section aria-label="Ask">
            <PromptLine cwd="~" command="ask" />
            <div className="ml-12 max-[560px]:ml-6.5">
              <AskHint>anything about Aqua256</AskHint>
            </div>
          </section>
        </div>
      </div>
    </main>
  );
}
