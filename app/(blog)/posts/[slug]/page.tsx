import type { Metadata } from "next";
import { notFound } from "next/navigation";
import { getAdjacentPosts, getPost, getPosts } from "@/lib/posts";
import { PostOutline } from "../../_components/post-outline";
import { PostPager } from "../../_components/post-pager";
import { PromptLine } from "../../_components/prompt-line";

export async function generateStaticParams() {
  const posts = await getPosts();
  return posts.map(({ slug }) => ({ slug }));
}

export const dynamicParams = false;

export async function generateMetadata(props: PageProps<"/posts/[slug]">): Promise<Metadata> {
  const { slug } = await props.params;
  const post = await getPost(slug);
  return post ? { title: post.title } : {};
}

/**
 * The post as a `cat` of its .mdx file:
 *
 *   ~/posts $ cat hello-world.mdx
 *      # Title
 *        2026-09-29 · 2 min read · #tag
 *        body…
 *   (END)
 *   ← older                     newer →
 */
export default async function PostPage(props: PageProps<"/posts/[slug]">) {
  const { slug } = await props.params;
  const post = await getPost(slug);
  if (!post) notFound();

  const { older, newer } = await getAdjacentPosts(slug);
  const { Content } = post;

  return (
    <main className="px-(--gutter) pt-12 pb-18">
      {/* Same width as the header; the second column is kept free for the outline on wide screens. */}
      <div className="page-width mx-auto grid grid-cols-[minmax(0,var(--col))] gap-x-18 min-[1180px]:grid-cols-[minmax(0,var(--col))_220px]">
        <article id="top" className="scroll-mt-20">
          <PromptLine cwd="~/posts" command="cat" args={`${slug}.mdx`} />

          <div className="post-file">
            <h1 className="post-title">{post.title}</h1>
            <p className="mb-9 text-small text-dim">
              <time dateTime={post.date}>{post.date}</time> · {post.readingMinutes} min read
              {post.tags && post.tags.length > 0 && (
                <>
                  {" · "}
                  {post.tags.map((tag) => (
                    <span key={tag} className="whitespace-nowrap text-tag">
                      #{tag}{" "}
                    </span>
                  ))}
                </>
              )}
            </p>
            <div className="post-body">
              <Content />
            </div>
          </div>

          <p className="mt-12 text-small text-dim">(END)</p>
          <PostPager older={older} newer={newer} />
        </article>

        {post.headings.length > 0 && (
          <aside className="hidden min-[1180px]:block">
            <PostOutline headings={post.headings} />
          </aside>
        )}
      </div>
    </main>
  );
}
