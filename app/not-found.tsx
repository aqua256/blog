import type { Metadata } from "next";
import Link from "next/link";
import { PageColumns } from "./_components/page-columns";
import { PageTitle } from "./_components/page-title";
import { PageTransition } from "./_components/page-transition";
import { RequestedPath } from "./_components/requested-path";

export const metadata: Metadata = {
  title: "找不到这个页面",
};

/** Any address that isn't a page: what was asked for, and the way back. One column, without the profile card. */
export default function NotFound() {
  return (
    <PageTransition>
      <PageColumns>
        <article className="max-w-[630px]">
          <PageTitle>找不到这个页面</PageTitle>
          <div className="post-body mt-9 max-[640px]:mt-7">
            <p>
              <RequestedPath /> 这个地址不存在，可能已经改名或删掉了。
            </p>
            <p>
              <Link href="/">回到首页</Link>，或者<Link href="/posts">去归档看看</Link>。
            </p>
          </div>
        </article>
      </PageColumns>
    </PageTransition>
  );
}
