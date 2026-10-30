import type { Metadata } from "next";
import About from "@/content/about.mdx";
import { PageColumns } from "../_components/page-columns";
import { PageTitle } from "../_components/page-title";
import { PageTransition } from "../_components/page-transition";
import { ProfileCard } from "../_components/profile-card";

export const metadata: Metadata = {
  title: "关于",
};

/** A short introduction, typeset like a post (the text is content/about.mdx), beside the profile card. */
export default function AboutPage() {
  return (
    <PageTransition>
      <PageColumns side={<ProfileCard />}>
        <article className="max-w-[630px]">
          <PageTitle>关于</PageTitle>
          <div className="post-body mt-9 max-[640px]:mt-7">
            <About />
          </div>
        </article>
      </PageColumns>
    </PageTransition>
  );
}
