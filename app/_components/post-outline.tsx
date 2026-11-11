"use client";

import { useRef } from "react";
import type { Heading } from "@/lib/headings";
import { FluidHover } from "./fluid-hover";
import { useActiveSection } from "./use-active-section";

// The height of a row, the compact control size
const ROW = 28;

/**
 * The outline in the side column, held at the top while the post scrolls. It sits on the page like the post
 * itself rather than in a card. The rows share one hover highlight, and the section being read sits on a
 * fill that slides to the next one on the moderate spring as the post scrolls; its label turns 550. A
 * heading too long for its row is cut off with an ellipsis.
 */
export function PostOutline({ headings }: { headings: Heading[] }) {
  // Nothing is current above the first section
  const active = useActiveSection(
    headings.map(({ id }) => id),
    { before: null },
  );
  const index = headings.findIndex(({ id }) => id === active);

  return (
    <nav aria-labelledby="outline-title">
      <h2 id="outline-title" className="type-subtitle mb-2 text-ink">
        目录
      </h2>
      <FluidHover className="-mx-2">
        <span
          aria-hidden="true"
          className={`absolute inset-x-0 top-0 -z-1 h-7 rounded-lg bg-active transition-[transform,opacity] duration-(--duration-moderate) ease-spring motion-reduce:transition-opacity ${index < 0 ? "opacity-0" : ""}`}
          style={{ transform: `translateY(${Math.max(index, 0) * ROW}px)` }}
        />
        <ol>
          {headings.map(({ id, text }) => (
            <li key={id}>
              <a
                href={`#${id}`}
                aria-current={id === active ? "location" : undefined}
                className="type-caption block h-7 truncate rounded-lg px-2 leading-7 text-muted hover:text-ink aria-[current=location]:font-[550] aria-[current=location]:text-ink aria-[current=location]:[--opsz:18]"
              >
                {text}
              </a>
            </li>
          ))}
        </ol>
      </FluidHover>
    </nav>
  );
}

/**
 * On narrow screens there is no side column: the outline folds into a box under the title. Picking a
 * section folds it again, so it doesn't stand open at the top on the way back.
 */
export function PostOutlineFolded({ headings }: { headings: Heading[] }) {
  const details = useRef<HTMLDetailsElement>(null);

  // Runs before the browser follows the link, so it scrolls to the section with the outline already folded
  function fold() {
    if (details.current) details.current.open = false;
  }

  return (
    <details ref={details} className="group mt-4 rounded-xl shadow-[inset_0_0_0_1px_var(--line)]">
      <summary className="type-subtitle flex h-11 cursor-pointer list-none items-center justify-between rounded-xl px-3 text-ink [&::-webkit-details-marker]:hidden">
        目录
        <svg
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          strokeWidth={1.5}
          strokeLinecap="round"
          strokeLinejoin="round"
          aria-hidden="true"
          className="size-4 text-muted transition-transform duration-(--duration-moderate) ease-spring group-open:rotate-180"
        >
          <path d="m6 9 6 6 6-6" />
        </svg>
      </summary>
      <ol className="px-1 pb-1">
        {headings.map(({ id, text }) => (
          <li key={id}>
            <a
              href={`#${id}`}
              onClick={fold}
              className="type-body flex min-h-9 items-center rounded-lg px-2 text-muted hover:bg-hover hover:text-ink"
            >
              {text}
            </a>
          </li>
        ))}
      </ol>
    </details>
  );
}
