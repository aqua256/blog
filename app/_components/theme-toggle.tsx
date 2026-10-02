"use client";

import { useTheme } from "next-themes";
import { useSyncExternalStore } from "react";
import { flushSync } from "react-dom";

const next: Record<string, string> = { system: "light", light: "dark", dark: "system" };

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

/** `THEME=system` — the current setting, written as a shell variable. Click to cycle system → light → dark. */
export function ThemeToggle() {
  const { theme, setTheme } = useTheme();

  // The saved theme is only known in the browser, so show a placeholder during server rendering.
  // useSyncExternalStore gives false on the server and during hydration, true afterwards — no effect needed.
  const mounted = useSyncExternalStore(
    noopSubscribe,
    () => true,
    () => false,
  );

  return (
    <button
      type="button"
      onClick={() => withCrossfade(() => setTheme((current) => next[current] ?? "system"))}
      className="group cursor-pointer whitespace-nowrap py-1"
    >
      <span className="font-semibold text-cmd">THEME</span>
      <span className="text-dim">=</span>
      <span className="text-str underline-offset-4 group-hover:underline">{mounted ? theme : "…"}</span>
    </button>
  );
}
