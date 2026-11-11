"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { isCurrentSection } from "@/lib/nav";
import { site } from "@/lib/site";
import { FluidHover } from "./fluid-hover";

/**
 * 首页 / 归档 / 关于, sharing one hover highlight. The current section sits on a fill and its label turns 550
 * without shifting the others.
 */
export function NavLinks() {
  const pathname = usePathname();

  return (
    <nav aria-label="主导航" className="ml-auto">
      <FluidHover>
        <ul className="flex">
          {site.nav.map(({ label, href }) => {
            const current = isCurrentSection(pathname, href);
            return (
              <li key={href}>
                <Link
                  href={href}
                  aria-current={current ? "page" : undefined}
                  className="type-body flex h-9 items-center rounded-lg px-3 text-muted hover:text-ink aria-[current=page]:bg-active aria-[current=page]:text-ink max-[640px]:px-2.5"
                >
                  <span className="steady" data-text={label}>
                    {label}
                  </span>
                </Link>
              </li>
            );
          })}
        </ul>
      </FluidHover>
    </nav>
  );
}
