"use client";

import { THEME_STORAGE_KEY } from "@/lib/theme";

function isDark() {
  const set = document.documentElement.getAttribute("data-theme");
  if (set) return set === "dark";
  return window.matchMedia("(prefers-color-scheme: dark)").matches;
}

/**
 * `THEME=dark` written as a shell assignment: clicking "runs" it.
 * The value shown is the theme you'd switch to; CSS picks it, so server and client markup match.
 */
export function ThemeToggle() {
  function toggle() {
    const next = isDark() ? "light" : "dark";
    document.documentElement.setAttribute("data-theme", next);
    try {
      localStorage.setItem(THEME_STORAGE_KEY, next);
    } catch {}
  }

  return (
    <button type="button" onClick={toggle} className="group cursor-pointer whitespace-nowrap py-1">
      <span className="font-semibold text-cmd">THEME</span>
      <span className="text-dim">=</span>
      <span className="text-str underline-offset-4 group-hover:underline">
        <span className="theme-dark:hidden">dark</span>
        <span className="hidden theme-dark:inline">light</span>
      </span>
      <span className="sr-only">
        <span className="theme-dark:hidden">Switch to dark theme</span>
        <span className="hidden theme-dark:inline">Switch to light theme</span>
      </span>
    </button>
  );
}
