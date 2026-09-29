import type { Metadata } from "next";
import { notFound } from "next/navigation";
import { getPost, getPosts } from "@/lib/posts";

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

// Page chrome (command line, title, meta, end-of-post navigation) comes in the next step.
export default async function PostPage(props: PageProps<"/posts/[slug]">) {
  const { slug } = await props.params;
  const post = await getPost(slug);
  if (!post) notFound();

  const { Content } = post;
  return (
    <main className="px-(--gutter) py-12">
      <article className="page-width mx-auto">
        <h1>{post.title}</h1>
        <div className="post-body">
          <Content />
        </div>
      </article>
    </main>
  );
}
