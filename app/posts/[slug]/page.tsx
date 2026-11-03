import type { Metadata } from "next";
import { notFound } from "next/navigation";
import { getAdjacentPosts, getPost, getPosts } from "@/lib/posts";
import { PageColumns } from "../../_components/page-columns";
import { PageTitle } from "../../_components/page-title";
import { PostOutline, PostOutlineFolded } from "../../_components/post-outline";
import { PostPager } from "../../_components/post-pager";
import { ProfileCard } from "../../_components/profile-card";

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
 * The post beside the profile card and, on wide screens, an outline card that stays in view. On narrow
 * screens the outline folds into a card under the title. The newer and older posts close the page.
 */
export default async function PostPage(props: PageProps<"/posts/[slug]">) {
  const { slug } = await props.params;
  const post = await getPost(slug);
  if (!post) notFound();

  const { older, newer } = await getAdjacentPosts(slug);
  const { Content, headings } = post;

  return (
    <PageColumns
      side={
        <>
          <ProfileCard />
          {headings.length > 0 && (
            <div className="sticky top-21.5 mt-9 max-[869px]:hidden">
              <PostOutline headings={headings} />
            </div>
          )}
        </>
      }
    >
      <article className="max-w-[630px]">
        <header>
          <PageTitle>{post.title}</PageTitle>
          <div className="mt-3.5 flex flex-wrap items-center gap-x-4.5 gap-y-1.75 text-[13px] leading-normal">
            <time dateTime={post.date} className="text-ink-3 tabular-nums">
              {post.date}
            </time>
            {post.tags && post.tags.length > 0 && (
              <ul aria-label="标签" className="flex gap-3.5 text-ink-2">
                {post.tags.map((tag) => (
                  <li key={tag}>#{tag}</li>
                ))}
              </ul>
            )}
          </div>
          {headings.length > 0 && (
            <div className="min-[870px]:hidden">
              <PostOutlineFolded headings={headings} />
            </div>
          )}
        </header>
        <div className="post-body mt-9 max-[640px]:mt-7">
          <Content />
        </div>
        <PostPager older={older} newer={newer} />
      </article>
    </PageColumns>
  );
}
