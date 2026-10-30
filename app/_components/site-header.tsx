import Link from "next/link";
import { site } from "@/lib/site";
import { NavLinks } from "./nav-links";
import { ThemeToggle } from "./theme-toggle";

/**
 * `Aqua256                       首页  归档  ☾`
 *
 * Sticks to the top; clear over the top of the page, it fills in over the first 120px of scrolling
 * (.site-header in globals.css).
 */
export function SiteHeader() {
  return (
    <header style={{ viewTransitionName: "site-header" }} className="site-header sticky top-0 z-10 border-b border-transparent">
      <div className="mx-auto flex min-h-16 max-w-280 items-center gap-2 px-6 max-[640px]:px-4">
        <Link href="/" className="text-[16px] font-semibold tracking-[-0.01em] text-ink">
          {site.name}
        </Link>
        <NavLinks />
        <ThemeToggle />
      </div>
    </header>
  );
}
