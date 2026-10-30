"use client";

import { useSyncExternalStore } from "react";
import type { Heading } from "@/lib/headings";

function subscribe(onChange: () => void) {
  window.addEventListener("scroll", onChange, { passive: true });
  window.addEventListener("resize", onChange);
  return () => {
    window.removeEventListener("scroll", onChange);
    window.removeEventListener("resize", onChange);
  };
}

/** Scrolled all the way down (and the page does scroll). */
function atBottom() {
  const { scrollHeight } = document.documentElement;
  return window.scrollY > 0 && window.scrollY + window.innerHeight >= scrollHeight - 2;
}

/**
 * The last section whose heading has scrolled above 35% of the viewport; none before the first one.
 * At the bottom of the page it's the last section, whose heading may never reach that line if it's short.
 */
function activeSection(headings: Heading[]): string | null {
  if (headings.length > 0 && atBottom()) return headings[headings.length - 1].id;
  let active = null;
  for (const { id } of headings) {
    const el = document.getElementById(id);
    if (el && el.getBoundingClientRect().top < window.innerHeight * 0.35) active = id;
  }
  return active;
}

/**
 * The outline card in the side column, held at the top while the post scrolls. The section being read
 * reads bolder and darker, like the current page in the header.
 */
export function PostOutline({ headings }: { headings: Heading[] }) {
  const active = useSyncExternalStore(
    subscribe,
    () => activeSection(headings),
    () => null,
  );

  return (
    <nav aria-labelledby="outline-title" className="card px-5.5 pt-4.5 pb-4">
      <h2 id="outline-title" className="mb-1.75 text-[13px] leading-normal font-semibold text-ink">
        目录
      </h2>
      <ol>
        {headings.map(({ id, text }) => {
          const current = id === active;
          return (
            <li key={id}>
              <a
                href={`#${id}`}
                aria-current={current ? "location" : undefined}
                className={`block py-1.25 text-[13px] leading-normal ${current ? "font-semibold text-ink" : "text-ink-3 hover:text-ink"}`}
              >
                {text}
              </a>
            </li>
          );
        })}
      </ol>
    </nav>
  );
}

/** On narrow screens there is no side column: the outline folds into a card under the title. */
export function PostOutlineFolded({ headings }: { headings: Heading[] }) {
  return (
    <details className="card group mt-5.5 px-4">
      <summary className="flex min-h-11 cursor-pointer list-none items-center justify-between text-[13px] font-semibold text-ink [&::-webkit-details-marker]:hidden">
        目录
        <svg
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth={1.75}
          strokeLinecap="round"
          strokeLinejoin="round"
          aria-hidden="true"
          className="size-4 text-ink-3 transition-transform duration-200 group-open:rotate-180"
        >
          <path d="m6 9 6 6 6-6" />
        </svg>
      </summary>
      <ol className="mb-3">
        {headings.map(({ id, text }) => (
          <li key={id}>
            <a href={`#${id}`} className="block py-2 text-[13px] leading-normal text-ink-2">
              {text}
            </a>
          </li>
        ))}
      </ol>
    </details>
  );
}
