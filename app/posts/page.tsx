import type { Metadata } from "next";
import { groupByYear } from "@/lib/post-order";
import { getPosts } from "@/lib/posts";
import { PageColumns } from "../_components/page-columns";
import { PageTitle } from "../_components/page-title";
import { PageTransition } from "../_components/page-transition";
import { PostCard } from "../_components/post-card";
import { ProfileCard } from "../_components/profile-card";

export const metadata: Metadata = {
  title: "归档",
};

/** Every post as a card, under a quiet heading for each year, beside the profile card. */
export default async function ArchivePage() {
  const years = groupByYear(await getPosts());

  return (
    <PageTransition>
      <PageColumns side={<ProfileCard />}>
        <PageTitle>归档</PageTitle>
        {years.length > 0 ? (
          years.map(({ year, posts }) => (
            <section key={year} aria-labelledby={`year-${year}`} className="mt-9 flex flex-col gap-4.5">
              <h2 id={`year-${year}`} className="text-[14px] leading-none font-semibold text-ink-3 tabular-nums">
                {year}
              </h2>
              {posts.map((post) => (
                <PostCard key={post.slug} post={post} />
              ))}
            </section>
          ))
        ) : (
          <p className="mt-9 text-ink-3">还没有文章。</p>
        )}
      </PageColumns>
    </PageTransition>
  );
}
