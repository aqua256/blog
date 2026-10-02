import Link from "next/link";
import { site } from "@/lib/site";
import { AskButton, AskCommand } from "./ask-button";
import { NavLinks } from "./nav-links";
import { ThemeToggle } from "./theme-toggle";

/**
 * `aqua256@blog  cd ~ ~/projects ~/posts          ask  THEME=dark`
 *
 * Phones get two deliberate rows instead of an arbitrary wrap:
 *
 *   aqua256@blog         ask  THEME=dark
 *   cd ~ ~/projects ~/posts
 */
export function SiteHeader() {
  return (
    <header className="sticky top-[env(safe-area-inset-top,0px)] z-10 bg-paper/90 px-(--gutter) backdrop-blur-sm max-[560px]:static">
      <div className="page-width mx-auto flex min-h-13 flex-wrap items-center gap-x-5 gap-y-1.5 border-b border-faint py-1.5">
        <Link href="/" className="whitespace-nowrap font-bold">
          <span className="text-prompt">{site.user}</span>@{site.host}
        </Link>
        <NavLinks className="max-[560px]:order-last max-[560px]:basis-full" />
        <div className="ml-auto flex items-center gap-x-5">
          <AskButton
            aria-label={`Ask ${site.givenName}'s agent`}
            title={`Ask ${site.givenName}'s agent (press /)`}
            className="py-1"
          >
            <AskCommand />
          </AskButton>
          <ThemeToggle />
        </div>
      </div>
    </header>
  );
}
