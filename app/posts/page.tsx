import type { Metadata } from "next";
import Link from "next/link";
import { groupByYear } from "@/lib/post-order";
import { getPosts } from "@/lib/posts";
import { PageColumns } from "../_components/page-columns";
import { PageTitle } from "../_components/page-title";
import { PageTransition, PostTitleTransition } from "../_components/page-transition";
import { ProfileCard } from "../_components/profile-card";

export const metadata: Metadata = {
  title: "归档",
};

/**
 * Every post on one line, title on the left and month-day on the right, under a quiet heading for each
 * year, beside the profile card. The home page shows the latest posts as cards with their descriptions;
 * this is for scanning them all.
 */
export default async function ArchivePage() {
  const years = groupByYear(await getPosts());

  return (
    <PageTransition>
      <PageColumns side={<ProfileCard />}>
        <PageTitle>归档</PageTitle>
        {years.length > 0 ? (
          years.map(({ year, posts }, i) => (
            <section key={year} aria-labelledby={`year-${year}`} className={i === 0 ? "mt-9" : "mt-12.5"}>
              <h2 id={`year-${year}`} className="mb-2 text-[14px] leading-none font-semibold text-ink-3 tabular-nums">
                {year}
              </h2>
              <ol className="border-b border-line">
                {posts.map((post) => (
                  <li key={post.slug} className="flex items-baseline gap-5.5 border-t border-line py-3.25 max-[640px]:gap-4">
                    <PostTitleTransition slug={post.slug}>
                      <Link
                        href={`/posts/${post.slug}`}
                        className="min-w-0 flex-1 text-[15px] leading-[1.6] text-pretty text-ink hover:text-accent"
                      >
                        {post.title}
                        {post.draft && <span className="text-ink-3"> · 草稿</span>}
                      </Link>
                    </PostTitleTransition>
                    <time dateTime={post.date} className="ml-auto flex-none text-[13px] leading-[1.6] text-ink-3 tabular-nums">
                      {post.date.slice(5)}
                    </time>
                  </li>
                ))}
              </ol>
            </section>
          ))
        ) : (
          <p className="mt-9 text-ink-3">还没有文章。</p>
        )}
      </PageColumns>
    </PageTransition>
  );
}
