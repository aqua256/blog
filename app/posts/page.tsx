import type { Metadata } from "next";
import { groupByYear } from "@/lib/post-order";
import { getPosts } from "@/lib/posts";
import { ArchiveList } from "../_components/archive-list";
import { PageColumns } from "../_components/page-columns";
import { PageTitle } from "../_components/page-title";
import { ProfileCard } from "../_components/profile-card";

export const metadata: Metadata = {
  title: "归档",
};

/**
 * Every post on one line, title on the left and month-day on the right, under a heading for each year,
 * beside the profile card. Above the list, the post count and a control to show one year at a time. The
 * home page shows the latest posts with their descriptions; this is for scanning them all.
 */
export default async function ArchivePage() {
  const posts = await getPosts();
  const years = groupByYear(posts);

  return (
    <PageColumns side={<ProfileCard />}>
      <div className="reveal">
        <PageTitle>归档</PageTitle>
        {years.length > 0 ? (
          <>
            <p className="type-body mt-2 text-muted">{posts.length} 篇文章，按年份排列。</p>
            <ArchiveList
              years={years.map(({ year, posts }) => ({
                year,
                posts: posts.map(({ slug, title, date, draft }) => ({ slug, title, date, draft })),
              }))}
            />
          </>
        ) : (
          <p className="type-body mt-6 text-muted">还没有文章。</p>
        )}
      </div>
    </PageColumns>
  );
}
