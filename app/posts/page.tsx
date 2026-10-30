import type { Metadata } from "next";
import Link from "next/link";
import { groupByYear } from "@/lib/post-order";
import { getPosts } from "@/lib/posts";
import { PageColumns } from "../_components/page-columns";
import { PageTitle } from "../_components/page-title";
import { ProfileCard } from "../_components/profile-card";
import { YearNav } from "../_components/year-nav";

export const metadata: Metadata = {
  title: "归档",
};

/**
 * Every post on one line, title on the left and month-day on the right, under a quiet heading for each
 * year, beside the profile card. Above the list, the post count and a link to each year. The home page
 * shows the latest posts as cards with their descriptions; this is for scanning them all.
 */
export default async function ArchivePage() {
  const posts = await getPosts();
  const years = groupByYear(posts);

  return (
    <PageColumns side={<ProfileCard />}>
      <PageTitle>归档</PageTitle>
      {years.length > 0 ? (
        <>
          <p className="mt-2.5 text-[14px] leading-[1.6] text-ink-3">{posts.length} 篇文章，按年份排列。</p>
          <YearNav years={years.map(({ year, posts }) => ({ year, count: posts.length }))} />
          {years.map(({ year, posts }) => (
            <section key={year} id={`year-${year}`} aria-labelledby={`year-${year}-title`} className="mt-8 scroll-mt-20">
              <h2 id={`year-${year}-title`} className="mb-2 text-[14px] leading-none font-semibold text-ink-3 tabular-nums">
                {year}
              </h2>
              <ol className="border-b border-line">
                {posts.map((post) => (
                  <li key={post.slug} className="flex items-baseline gap-5.5 border-t border-line py-3.25 max-[640px]:gap-4">
                    <Link
                      href={`/posts/${post.slug}`}
                      className="min-w-0 flex-1 text-[15px] leading-[1.6] text-pretty text-ink hover:text-accent"
                    >
                      {post.title}
                      {post.draft && <span className="text-ink-3"> · 草稿</span>}
                    </Link>
                    <time dateTime={post.date} className="ml-auto flex-none text-[13px] leading-[1.6] text-ink-3 tabular-nums">
                      {post.date.slice(5)}
                    </time>
                  </li>
                ))}
              </ol>
            </section>
          ))}
        </>
      ) : (
        <p className="mt-9 text-ink-3">还没有文章。</p>
      )}
    </PageColumns>
  );
}
