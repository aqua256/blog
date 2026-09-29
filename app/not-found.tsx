import type { Metadata } from "next";
import Link from "next/link";
import { AskHint } from "./_components/ask-button";
import { PromptLine } from "./_components/prompt-line";
import { RequestedPath } from "./_components/requested-path";

export const metadata: Metadata = {
  title: "404",
};

/**
 * Any URL that isn't a page, as a failed `cd`:
 *
 *   ~ $ cd ~/posts/typo
 *   cd: no such file or directory: ~/posts/typo
 *
 *   try cd ~, or press / to ask
 */
export default function NotFound() {
  return (
    <main className="px-(--gutter) pt-12 pb-18">
      <div className="page-width mx-auto">
        <section aria-label="Page not found" className="max-w-(--col)">
          <PromptLine cwd="~" command="cd" args={<RequestedPath />} />
          <p className="text-err [overflow-wrap:anywhere]">
            cd: no such file or directory: <RequestedPath />
          </p>

          <p className="mt-8 flex flex-wrap items-baseline gap-x-2 text-small text-dim">
            try
            <Link href="/" className="group hover:text-ink">
              <span className="font-semibold text-cmd">cd</span>{" "}
              <span className="underline-offset-4 group-hover:underline">~</span>
            </Link>
            or
            <AskHint>about this site</AskHint>
          </p>
        </section>
      </div>
    </main>
  );
}
