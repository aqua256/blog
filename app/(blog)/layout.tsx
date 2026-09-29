import { AgentConsole } from "./_components/agent-console";
import { SiteFooter } from "./_components/site-footer";
import { SiteHeader } from "./_components/site-header";

/** Terminal-style chrome for the blog pages. The root page keeps its own look. */
export default function BlogLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="term flex flex-1 flex-col bg-paper font-term text-ui text-ink selection:bg-faint">
      <AgentConsole>
        <SiteHeader />
        <div className="flex-1">{children}</div>
        <SiteFooter />
      </AgentConsole>
    </div>
  );
}
