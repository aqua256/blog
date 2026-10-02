"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { site } from "@/lib/site";

function isCurrent(pathname: string, href: string) {
  return href === "/" ? pathname === "/" : pathname === href || pathname.startsWith(`${href}/`);
}

/** `cd ~ ~/projects ~/posts` — one `cd` for every path, the current one in the prompt color. */
export function NavLinks({ className = "" }: { className?: string }) {
  const pathname = usePathname();

  return (
    <nav aria-label="Site" className={`flex items-baseline gap-3 ${className}`}>
      <span aria-hidden="true" className="font-semibold text-cmd">
        cd
      </span>
      <ul className="flex flex-wrap gap-x-4 gap-y-0.5 max-[560px]:gap-x-3">
        {site.nav.map(({ path, href }) => {
          const current = isCurrent(pathname, href);
          return (
            <li key={href}>
              <Link
                href={href}
                aria-current={current ? "page" : undefined}
                className={
                  current
                    ? "whitespace-nowrap py-1 font-bold text-prompt"
                    : "whitespace-nowrap py-1 text-dim hover:text-ink"
                }
              >
                {path}
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}
