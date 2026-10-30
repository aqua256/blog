"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { isCurrentSection } from "@/lib/nav";
import { site } from "@/lib/site";

/** 首页 / 归档 / 关于: the current section only reads bolder and darker. */
export function NavLinks() {
  const pathname = usePathname();

  return (
    <nav aria-label="主导航" className="ml-auto">
      <ul className="flex gap-1">
        {site.nav.map(({ label, href }) => {
          const current = isCurrentSection(pathname, href);
          return (
            <li key={href}>
              <Link
                href={href}
                aria-current={current ? "page" : undefined}
                className={`flex min-h-14 items-center px-3.25 text-[13px] max-[640px]:px-2.5 ${
                  current ? "font-semibold text-ink" : "text-ink-3 hover:text-ink"
                }`}
              >
                {label}
              </Link>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}
