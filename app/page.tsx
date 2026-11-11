import Link from "next/link";
import { getPosts } from "@/lib/posts";
import { site } from "@/lib/site";
import { FluidHover } from "./_components/fluid-hover";
import { PageColumns } from "./_components/page-columns";
import { PostCard } from "./_components/post-card";
import { ProfileCard } from "./_components/profile-card";

const LATEST_POSTS = 5;

/** The latest posts as rows, ending in a link to the archive, beside the profile card. */
export default async function HomePage() {
  const posts = (await getPosts()).slice(0, LATEST_POSTS);

  return (
    <PageColumns side={<ProfileCard />}>
      <div className="reveal">
        <h1 className="sr-only">{site.name}</h1>
        <h2 className="sr-only">最新文章</h2>
        {posts.length > 0 ? (
          <FluidHover item="article" className="post-list -mx-3">
            {posts.map((post) => (
              <PostCard key={post.slug} post={post} />
            ))}
          </FluidHover>
        ) : (
          <p className="type-body text-muted">还没有文章。</p>
        )}
        <p className="mt-4 text-right">
          <Link href="/posts" className="button-tertiary type-body">
            全部文章
            <svg
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              aria-hidden="true"
              className="size-4"
            >
              <path d="M5 12h14M12 5l7 7-7 7" />
            </svg>
          </Link>
        </p>
      </div>
    </PageColumns>
  );
}
