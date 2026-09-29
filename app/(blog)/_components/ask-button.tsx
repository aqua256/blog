"use client";

import { useOpenAgent } from "./agent-console";

/** A button that opens the agent console. The word `ask` inside it reads as a shell command. */
export function AskButton({ className = "", children, ...props }: React.ComponentProps<"button">) {
  const open = useOpenAgent();
  return (
    <button type="button" onClick={open} className={`group cursor-pointer text-dim hover:text-ink ${className}`} {...props}>
      {children}
    </button>
  );
}

export function AskCommand() {
  return <span className="font-semibold text-cmd underline-offset-4 group-hover:underline">ask</span>;
}

/** A keycap, e.g. the `/` that opens the console. */
export function Key({ children }: { children: React.ReactNode }) {
  return (
    <kbd className="rounded-[3px] border border-b-2 border-faint px-1.25 text-key leading-normal text-dim">{children}</kbd>
  );
}
