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
