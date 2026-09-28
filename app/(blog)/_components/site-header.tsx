import Link from "next/link";
import { site } from "@/lib/site";
import { NavLinks } from "./nav-links";
import { ThemeToggle } from "./theme-toggle";

/** `aqua256@blog  cd ~ ~/projects ~/writing          THEME=dark` */
export function SiteHeader() {
  return (
    <header className="sticky top-[env(safe-area-inset-top,0px)] z-10 bg-paper/90 px-(--gutter) backdrop-blur-sm max-[560px]:static">
      <div className="page-width mx-auto flex min-h-13 flex-wrap items-center gap-x-5 gap-y-1.5 border-b border-faint py-1.5">
        <Link href="/" className="whitespace-nowrap font-bold">
          <span className="text-prompt">{site.user}</span>@{site.host}
        </Link>
        <NavLinks />
        <div className="ml-auto">
          <ThemeToggle />
        </div>
      </div>
    </header>
  );
}
