"use client";

import { useSyncExternalStore } from "react";
import type { Heading } from "@/lib/headings";

const TOP = "top";

function subscribe(onChange: () => void) {
  window.addEventListener("scroll", onChange, { passive: true });
  window.addEventListener("resize", onChange);
  return () => {
    window.removeEventListener("scroll", onChange);
    window.removeEventListener("resize", onChange);
  };
}

/** The last section whose heading has scrolled above 35% of the viewport; "top" before the first one. */
function activeSection(headings: Heading[]): string {
  let active = TOP;
  for (const { id } of headings) {
    const el = document.getElementById(id);
    if (el && el.getBoundingClientRect().top < window.innerHeight * 0.35) active = id;
  }
  return active;
}

/**
 * Outline for wide screens. The section being read gets a `>` marker.
 *
 *   outline
 *   > top
 *     为什么又做了一个博客
 */
export function PostOutline({ headings }: { headings: Heading[] }) {
  const active = useSyncExternalStore(
    subscribe,
    () => activeSection(headings),
    () => TOP,
  );

  const items = [{ id: TOP, text: "top" }, ...headings];

  return (
    <nav aria-label="Outline" className="text-small">
      <h2 className="mb-3 font-semibold text-dim">outline</h2>
      <ol className="flex flex-col gap-2">
        {items.map(({ id, text }) => {
          const current = id === active;
          return (
            <li key={id}>
              <a
                href={`#${id}`}
                aria-current={current ? "location" : undefined}
                className={current ? "block leading-snug font-bold text-ink" : "block leading-snug text-dim hover:text-ink"}
              >
                {current && <span className="text-prompt">&gt; </span>}
                {text}
              </a>
            </li>
          );
        })}
      </ol>
    </nav>
  );
}
