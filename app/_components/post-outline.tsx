"use client";

import type { Heading } from "@/lib/headings";
import { useActiveSection } from "./use-active-section";

/**
 * The outline card in the side column, held at the top while the post scrolls. The section being read
 * reads bolder and darker, like the current page in the header.
 */
export function PostOutline({ headings }: { headings: Heading[] }) {
  // Nothing is current above the first section
  const active = useActiveSection(
    headings.map(({ id }) => id),
    { before: null },
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
