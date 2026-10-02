"use client";

import { ThemeProvider } from "next-themes";

/**
 * Sets data-theme on <html> ("light" | "dark"), remembers the choice, and follows the system by default.
 * CSS transitions are paused while the theme switches, so hover fades don't lag behind the new colors.
 */
export function Providers({ children }: { children: React.ReactNode }) {
  return <ThemeProvider attribute="data-theme" disableTransitionOnChange>{children}</ThemeProvider>;
}
