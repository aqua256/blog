"use client";

import { useEffect } from "react";

const PLAYED = "intro-played";
const LENGTH = 1800;
const SKIP_ON = ["pointerdown", "wheel", "keydown", "touchstart"] as const;

/**
 * Runs in <head> before the first paint (app/layout.tsx): on the first page of a visit (a browser tab's
 * session), unless the reader prefers reduced motion, it marks <html> with the `intro` class that starts
 * the opening in globals.css. If storage is unavailable (blocked, or some private modes) it throws, and
 * the page simply opens without the intro.
 */
export const introScript = `try {
  if (!sessionStorage.getItem("${PLAYED}") && !matchMedia("(prefers-reduced-motion: reduce)").matches) {
    sessionStorage.setItem("${PLAYED}", "1");
    document.documentElement.classList.add("intro");
  }
} catch (e) {}`;

/**
 * Ends the opening: removes the `intro` class when it is over, so later pages of the visit don't fade in
 * again, or at once when the reader clicks, scrolls or presses a key.
 */
export function LeavesIntro() {
  useEffect(() => {
    const root = document.documentElement;
    if (!root.classList.contains("intro")) return;

    const end = () => {
      root.classList.remove("intro");
      clearTimeout(timer);
      SKIP_ON.forEach((type) => removeEventListener(type, end));
    };
    // The animations started with the first paint, at about the start of the navigation
    const timer = setTimeout(end, Math.max(0, LENGTH - performance.now()));
    SKIP_ON.forEach((type) => addEventListener(type, end, { passive: true }));
    return end;
  }, []);

  return null;
}
