import Link from "next/link";
import { getPosts } from "@/lib/posts";
import { site } from "@/lib/site";
import { AskHint } from "./_components/ask-button";
import { AvatarArt } from "./_components/avatar-art";
import { PageTransition } from "./_components/page-transition";
import { PostList } from "./_components/post-list";
import { PromptLine } from "./_components/prompt-line";

const RECENT_POSTS = 5;
/** The color swatches fastfetch prints last, in Catppuccin's accents */
const SWATCHES = ["red", "peach", "yellow", "green", "teal", "blue", "mauve", "pink"];

/** Style for the nth line of output, so lines print in turn (see .print-line in globals.css) */
const nth = (line: number) => ({ "--line": line }) as React.CSSProperties;

const link = "text-link underline-offset-3 hover:underline";

/**
 * The home directory:
 *
 *   ~ $ fastfetch
 *      ▀▀▀▀▀▀▀▀   aqua256
 *      ▀▀▀▀▀▀▀▀   ───────
 *      ▀▀▀▀▀▀▀▀   Role     学生
 *      ▀▀▀▀▀▀▀▀   Skills   TypeScript
 *      (avatar)   Uptime   since 2026-09-29
 *                 Links    linkedin · github · rss
 *                 ███ ███ ███ ███ ███ ███ ███ ███
 *      <a few lines about Aqua256>
 *
 *   ~ $ ls -t ~/posts | head -5
 *      2026-09-29   Hello, world
 *      cd ~/posts                        (only when there are more)
 *
 *   ~ $ ask
 *      press / to ask anything about Aqua256
 *
 *   ~ $ █
 */
export default async function HomePage() {
  const posts = await getPosts();
  const first = posts.at(-1);

  const info: { key: string; value: React.ReactNode }[] = [
    ...site.profile,
    ...(first ? [{ key: "Uptime", value: `since ${first.date}` }] : []),
    {
      key: "Links",
      value: site.social.map(({ label, href }, i) => (
        <span key={label}>
          {i > 0 && <span className="text-dim"> · </span>}
          <a href={href} className={link}>
            {label}
          </a>
        </span>
      )),
    },
  ];

  return (
    <PageTransition>
      <main className="px-(--gutter) pt-12 pb-18">
        <div className="page-width mx-auto">
          <div className="flex max-w-(--col) flex-col gap-12">
            {/* Command output is indented one gutter, like the post page */}
            <section aria-label="About">
              <PromptLine cwd="~" command="fastfetch" />
              <div className="ml-12 max-[560px]:ml-6.5">
                <div className="flex items-start gap-8 max-[560px]:flex-col max-[560px]:gap-5">
                  <AvatarArt label={`${site.name} 的头像`} />
                  <div className="min-w-0 max-w-full">
                    <h1 className="print-line font-bold text-prompt" style={nth(0)}>
                      {site.user}
                    </h1>
                    <p aria-hidden className="print-line text-dim" style={nth(1)}>
                      {"─".repeat(site.user.length)}
                    </p>
                    <dl className="grid grid-cols-[auto_minmax(0,1fr)] gap-x-4">
                      {info.map(({ key, value }, i) => (
                        <div key={key} className="print-line col-span-2 grid grid-cols-subgrid" style={nth(i + 2)}>
                          <dt className="font-semibold text-prompt">{key}</dt>
                          <dd>{value}</dd>
                        </div>
                      ))}
                    </dl>
                    <p aria-hidden className="print-line mt-3 flex" style={nth(info.length + 2)}>
                      {/* Swatches may shrink (clipping their blocks) so the row fits narrow screens */}
                      {SWATCHES.map((name) => (
                        <span key={name} className="min-w-0 overflow-clip" style={{ color: `var(--ctp-${name})` }}>
                          ███
                        </span>
                      ))}
                    </p>
                  </div>
                </div>
                <p className="mt-8 font-[380] text-body leading-[1.7] text-pretty [font-variation-settings:'MONO'_0,'CASL'_0.3]">
                  {site.intro}
                </p>
              </div>
            </section>

            <section aria-label="Recent posts">
              <PromptLine cwd="~" command="ls" args={`-t ~/posts | head -${RECENT_POSTS}`} />
              <div className="ml-12 max-[560px]:ml-6.5">
                <PostList posts={posts.slice(0, RECENT_POSTS)} />
                {posts.length > RECENT_POSTS && (
                  <Link href="/posts" className="group mt-4 inline-block text-dim hover:text-ink">
                    <span className="font-semibold text-cmd">cd</span>{" "}
                    <span className="underline-offset-4 group-hover:underline">~/posts</span> for all {posts.length}
                  </Link>
                )}
              </div>
            </section>

            <section aria-label="Ask">
              <PromptLine cwd="~" command="ask" />
              <div className="ml-12 max-[560px]:ml-6.5">
                <AskHint>anything about {site.givenName}</AskHint>
              </div>
            </section>

            <PromptLine cwd="~" cursor />
          </div>
        </div>
      </main>
    </PageTransition>
  );
}
