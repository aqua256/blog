import { useSyncExternalStore } from "react";

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

type Options = {
  /** The answer while no section has reached the line, and during server rendering */
  before: string | null;
};

/**
 * The id of the section being read: the last of `ids` whose element has scrolled above a line 35% down the
 * viewport. At the bottom of the page it is the last section, which may never reach the line if it's short.
 */
export function useActiveSection(ids: readonly string[], { before }: Options): string | null {
  return useSyncExternalStore(
    subscribe,
    () => {
      if (ids.length > 0 && atBottom()) return ids[ids.length - 1];
      const y = window.innerHeight * 0.35;
      let active = before;
      for (const id of ids) {
        const el = document.getElementById(id);
        if (el && el.getBoundingClientRect().top < y) active = id;
      }
      return active;
    },
    () => before,
  );
}
