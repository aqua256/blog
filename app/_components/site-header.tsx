import Link from "next/link";
import { site } from "@/lib/site";
import { NavLinks } from "./nav-links";
import { ThemeToggle } from "./theme-toggle";

/**
 * `Aqua256                  首页  归档  关于  ☾`
 *
 * Sticks to the top; clear over the top of the page, it fills in over the first 120px of scrolling
 * (.site-header in globals.css).
 */
export function SiteHeader() {
  return (
    <header className="site-header sticky top-0 z-10 border-b border-transparent">
      <div className="mx-auto flex min-h-14 max-w-264 items-center gap-2 px-6 max-[640px]:px-4">
        <Link href="/" className="text-[15px] font-semibold tracking-[-0.01em] text-ink">
          {site.name}
        </Link>
        <NavLinks />
        <ThemeToggle />
      </div>
    </header>
  );
}
