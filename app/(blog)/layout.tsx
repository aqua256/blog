import { getPosts } from "@/lib/posts";
import { AgentConsole } from "./_components/agent-console";
import { SiteFooter } from "./_components/site-footer";
import { SiteHeader } from "./_components/site-header";

/** Terminal-style chrome for every page of the site. */
export default async function BlogLayout({ children }: { children: React.ReactNode }) {
  // What the agent needs to know about the posts: enough to list them and summarize the current one
  const posts = (await getPosts()).map(({ slug, title, headings }) => ({ slug, title, headings }));

  return (
    <div className="term flex flex-1 flex-col bg-paper font-term text-ui text-ink selection:bg-faint">
      <AgentConsole posts={posts}>
        <SiteHeader />
        <div className="flex-1">{children}</div>
        <SiteFooter />
      </AgentConsole>
    </div>
  );
}
