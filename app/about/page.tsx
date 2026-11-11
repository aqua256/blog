import type { Metadata } from "next";
import About from "@/content/about.mdx";
import { PageColumns } from "../_components/page-columns";
import { PageTitle } from "../_components/page-title";
import { ProfileCard } from "../_components/profile-card";

export const metadata: Metadata = {
  title: "关于",
};

/** A short introduction, typeset like a post (the text is content/about.mdx), beside the profile card. */
export default function AboutPage() {
  return (
    <PageColumns side={<ProfileCard />}>
      <article className="max-w-160">
        <PageTitle>关于</PageTitle>
        <div className="typeset mt-6">
          <About />
        </div>
      </article>
    </PageColumns>
  );
}
