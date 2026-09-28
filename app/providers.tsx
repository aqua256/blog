"use client";

import { ThemeProvider } from "next-themes";

/** Sets data-theme on <html> ("light" | "dark"), remembers the choice, and follows the system by default. */
export function Providers({ children }: { children: React.ReactNode }) {
  return <ThemeProvider attribute="data-theme">{children}</ThemeProvider>;
}
