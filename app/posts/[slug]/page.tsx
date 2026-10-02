import type { Metadata } from "next";
import { notFound } from "next/navigation";
import { getAdjacentPosts, getPost, getPosts } from "@/lib/posts";
import { AskButton, AskCommand, AskHint, Key } from "../../_components/ask-button";
import { PageTransition, PostTitleTransition } from "../../_components/page-transition";
import { PostOutline } from "../../_components/post-outline";
import { PostPager } from "../../_components/post-pager";
import { PromptLine } from "../../_components/prompt-line";
import { TypedText } from "../../_components/typed-text";

export async function generateStaticParams() {
  const posts = await getPosts();
  return posts.map(({ slug }) => ({ slug }));
}

export const dynamicParams = false;

export async function generateMetadata(props: PageProps<"/posts/[slug]">): Promise<Metadata> {
  const { slug } = await props.params;
  const post = await getPost(slug);
  if (!post) return {};
  // Shared links show the title and description; the image comes from opengraph-image.tsx
  return {
    title: post.title,
    description: post.description,
    openGraph: {
      type: "article",
      title: post.title,
      description: post.description,
      url: `/posts/${slug}`,
      publishedTime: post.date,
      tags: post.tags,
    },
    twitter: { card: "summary_large_image" },
  };
}

/**
 * The post as a `cat` of its .mdx file:
 *
 *   ~/posts $ cat hello-world.mdx
 *      # Title
 *        2026-09-29 · #tag
 *        body…
 *   (END)  press / to ask about this post
 *   ← older                     newer →
 */
export default async function PostPage(props: PageProps<"/posts/[slug]">) {
  const { slug } = await props.params;
  const post = await getPost(slug);
  if (!post) notFound();

  const { older, newer } = await getAdjacentPosts(slug);
  const { Content } = post;

  // Keyed by slug so moving from one post to the next animates too
  return (
    <PageTransition key={slug}>
      <main className="px-(--gutter) pt-12 pb-18">
        {/* Same width as the header; the second column is kept free for the outline on wide screens. */}
        <div className="page-width mx-auto grid grid-cols-[minmax(0,var(--col))] gap-x-18 min-[1180px]:grid-cols-[minmax(0,var(--col))_220px]">
          <article id="top" className="scroll-mt-20">
            <PromptLine cwd="~/posts" command="cat" args={<TypedText text={`${slug}.mdx`} />} />

            <div className="post-file">
              <h1 className="post-title">
                {/* inline-block: a title that wraps is still one box to morph */}
                <PostTitleTransition slug={slug}>
                  <span className="inline-block">{post.title}</span>
                </PostTitleTransition>
              </h1>
              <p className="mb-9 text-small text-dim">
                <time dateTime={post.date}>{post.date}</time>
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

            <p className="mt-12 flex flex-wrap items-baseline gap-x-4 gap-y-1 text-small text-dim">
              (END)
              <AskHint>about this post</AskHint>
            </p>
            <PostPager older={older} newer={newer} />
          </article>

          {post.headings.length > 0 && (
            <aside className="hidden min-[1180px]:block">
              <div className="sticky top-24">
                <PostOutline headings={post.headings} />
                <AskButton className="mt-5 text-small">
                  <AskCommand /> the agent <Key>/</Key>
                </AskButton>
              </div>
            </aside>
          )}
        </div>
      </main>
    </PageTransition>
  );
}
