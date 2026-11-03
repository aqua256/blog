"use client";

import type { Heading } from "@/lib/headings";
import { useActiveSection } from "./use-active-section";

/**
 * The outline in the side column, held at the top while the post scrolls. It sits on the page like the post
 * itself rather than in a card, so a long one doesn't float beside the text as a box of its own. The
 * sections hang off a hairline, indented under the larger heading; the one being read reads bolder and
 * darker, like the current page in the header, and marks the line beside it.
 */
export function PostOutline({ headings }: { headings: Heading[] }) {
  // Nothing is current above the first section
  const active = useActiveSection(
    headings.map(({ id }) => id),
    { before: null },
  );

  return (
    <nav aria-labelledby="outline-title">
      <h2 id="outline-title" className="mb-2.25 text-[15px] leading-normal font-semibold tracking-[-0.01em] text-ink">
        目录
      </h2>
      <ol className="ml-px border-l border-line">
        {headings.map(({ id, text }) => {
          const current = id === active;
          return (
            <li key={id}>
              <a
                href={`#${id}`}
                aria-current={current ? "location" : undefined}
                className={`relative block py-1.25 pl-3.25 text-[13px] leading-normal ${current ? "font-semibold text-ink before:absolute before:inset-y-1.5 before:-left-px before:w-0.5 before:rounded-full before:bg-ink" : "text-ink-2 hover:text-ink"}`}
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
      <summary className="flex min-h-11 cursor-pointer list-none items-center justify-between text-[15px] font-semibold tracking-[-0.01em] text-ink [&::-webkit-details-marker]:hidden">
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
      <ol className="mb-3 ml-px border-l border-line">
        {headings.map(({ id, text }) => (
          <li key={id}>
            <a href={`#${id}`} className="block py-2 pl-3.25 text-[13px] leading-normal text-ink-2">
              {text}
            </a>
          </li>
        ))}
      </ol>
    </details>
  );
}
