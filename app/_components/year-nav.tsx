"use client";

import { useActiveSection } from "./use-active-section";

/**
 * `按年份  2026⁵  2025⁴  2024³`: each year jumps to its group in the list below, with its post count. The
 * year being read reads bolder and darker, like the current page in the header; at the top of the page
 * that's the newest year.
 */
export function YearNav({ years }: { years: { year: string; count: number }[] }) {
  const ids = years.map(({ year }) => `year-${year}`);
  // The year whose heading sits just under the header (a jump lands it 80px down). Not the last year at
  // the bottom of the page: on a short archive, a jump to a middle year also ends at the bottom.
  const active = useActiveSection(ids, { before: ids[0] ?? null, line: 96, lastAtBottom: false });

  return (
    <nav aria-labelledby="year-nav-title" className="mt-5.5 flex flex-wrap items-baseline gap-x-5 gap-y-1.5">
      <h2 id="year-nav-title" className="text-[13px] leading-[1.6] font-semibold text-muted">
        按年份
      </h2>
      <ul className="flex flex-wrap gap-x-4.5 gap-y-1.5">
        {years.map(({ year, count }) => {
          const current = active === `year-${year}`;
          return (
            <li key={year}>
              <a
                href={`#year-${year}`}
                aria-current={current ? "location" : undefined}
                className={`text-[14px] leading-[1.6] tabular-nums ${current ? "font-semibold text-ink" : "text-muted hover:text-ink"}`}
              >
                {year}
                <sup className="ml-0.5 text-[10px] font-normal text-muted">
                  {count}
                  <span className="sr-only"> 篇</span>
                </sup>
              </a>
            </li>
          );
        })}
      </ul>
    </nav>
  );
}
