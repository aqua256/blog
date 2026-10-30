import Link from "next/link";
import { getPosts } from "@/lib/posts";
import { site } from "@/lib/site";
import { PageColumns } from "./_components/page-columns";
import { PageTransition } from "./_components/page-transition";
import { PostCard } from "./_components/post-card";
import { ProfileCard } from "./_components/profile-card";

const LATEST_POSTS = 5;

/** The latest posts as cards, ending in a link to the archive, beside the profile card. */
export default async function HomePage() {
  const posts = (await getPosts()).slice(0, LATEST_POSTS);

  return (
    <PageTransition>
      <PageColumns side={<ProfileCard />}>
        {/* The leaf shadow reaches further down the home page (.leaves in globals.css) */}
        <div data-leaves="long" className="flex flex-col gap-4.5">
          <h1 className="sr-only">{site.name}</h1>
          <h2 className="sr-only">最新文章</h2>
          {posts.length > 0 ? (
            posts.map((post) => <PostCard key={post.slug} post={post} />)
          ) : (
            <p className="text-ink-3">还没有文章。</p>
          )}
          <p className="mt-2.75 text-right">
            <Link href="/posts" className="inline-flex min-h-11 items-center gap-1.25 text-[13px] text-ink-2 hover:text-ink">
              全部文章
              <svg
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth={1.75}
                strokeLinecap="round"
                strokeLinejoin="round"
                aria-hidden="true"
                className="size-3.5"
              >
                <path d="M5 12h14M12 5l7 7-7 7" />
              </svg>
            </Link>
          </p>
        </div>
      </PageColumns>
    </PageTransition>
  );
}
