"use client";

import { useTheme } from "next-themes";
import { useSyncExternalStore } from "react";
import { flushSync } from "react-dom";

const noopSubscribe = () => () => {};

/** Crossfade from the old colors to the new ones, unless the browser can't or the reader prefers no motion. */
function withCrossfade(update: () => void) {
  if (!document.startViewTransition || matchMedia("(prefers-reduced-motion: reduce)").matches) {
    update();
    return;
  }
  // flushSync makes next-themes set data-theme before the browser takes the "after" snapshot
  document.startViewTransition({ update: () => flushSync(update), types: ["theme"] });
}

const icon = {
  viewBox: "0 0 24 24",
  fill: "none",
  stroke: "currentColor",
  strokeWidth: 1.75,
  strokeLinecap: "round",
  strokeLinejoin: "round",
  "aria-hidden": true,
} as const;

/**
 * A moon in light mode, a sun in dark mode. A first visit follows the system setting; a click switches to
 * the other theme and remembers it. The icon is picked by CSS, so it is right before hydration too.
 */
export function ThemeToggle() {
  const { resolvedTheme, setTheme } = useTheme();

  // The resolved theme is only known in the browser, so the label is generic during server rendering.
  // useSyncExternalStore gives false on the server and during hydration, true afterwards — no effect needed.
  const mounted = useSyncExternalStore(
    noopSubscribe,
    () => true,
    () => false,
  );
  const label = !mounted ? "切换深浅色" : resolvedTheme === "dark" ? "切换到浅色" : "切换到深色";

  return (
    <button
      type="button"
      aria-label={label}
      title={label}
      onClick={() => withCrossfade(() => setTheme(resolvedTheme === "dark" ? "light" : "dark"))}
      className="-mr-3.25 grid size-11 cursor-pointer place-items-center rounded-md text-ink-3 hover:bg-ink/5 hover:text-ink"
    >
      <svg {...icon} className="size-4.5 dark:hidden">
        <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z" />
      </svg>
      <svg {...icon} className="hidden size-4.5 dark:block">
        <circle cx="12" cy="12" r="4" />
        <path d="M12 2v2M12 20v2M4.93 4.93l1.41 1.41M17.66 17.66l1.41 1.41M2 12h2M20 12h2M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41" />
      </svg>
    </button>
  );
}
