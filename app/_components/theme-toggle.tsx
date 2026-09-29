"use client";

import { useTheme } from "next-themes";
import { useSyncExternalStore } from "react";

const next: Record<string, string> = { system: "light", light: "dark", dark: "system" };

const noopSubscribe = () => () => {};

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
      onClick={() => setTheme((current) => next[current] ?? "system")}
      className="group cursor-pointer whitespace-nowrap py-1"
    >
      <span className="font-semibold text-cmd">THEME</span>
      <span className="text-dim">=</span>
      <span className="text-str underline-offset-4 group-hover:underline">{mounted ? theme : "…"}</span>
    </button>
  );
}
