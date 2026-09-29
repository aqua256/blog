"use client";

import { usePathname } from "next/navigation";
import { createContext, useContext, useEffect, useEffectEvent, useRef } from "react";
import { AGENT_COMMAND, type AgentPost } from "@/lib/agent-replies";
import { AgentChat } from "./agent-chat";

const OpenAgentContext = createContext<(() => void) | null>(null);

/** Opens the agent console. Any component under <AgentConsole> can use it. */
export function useOpenAgent() {
  const open = useContext(OpenAgentContext);
  if (!open) throw new Error("useOpenAgent must be used inside <AgentConsole>");
  return open;
}

/** `/` opens the console, except while the visitor is typing somewhere. */
function isTyping(target: EventTarget | null) {
  return target instanceof HTMLElement && (target.isContentEditable || /^(INPUT|TEXTAREA|SELECT)$/.test(target.tagName));
}

/** The shell's working directory for a URL: `/posts/hello-world` → `~/posts`. */
function cwdOf(pathname: string) {
  const section = pathname.split("/")[1];
  return section ? `~/${section}` : "~";
}

/**
 * A quake-style console that drops down from the top of the page.
 * A modal <dialog> gives focus trapping, Esc to close and focus return for free.
 *
 *   ~/posts $ aqua256-agent                                  [exit]
 *   demo · scripted replies
 *   <the conversation, see AgentChat>
 */
export function AgentConsole({ posts, children }: { posts: AgentPost[]; children: React.ReactNode }) {
  const dialogRef = useRef<HTMLDialogElement>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const cwd = cwdOf(usePathname());

  function open() {
    const dialog = dialogRef.current;
    if (dialog && !dialog.open) dialog.showModal();
    inputRef.current?.focus();
  }

  function close() {
    dialogRef.current?.close();
  }

  const onKeyDown = useEffectEvent((event: KeyboardEvent) => {
    if (event.key !== "/" || event.metaKey || event.ctrlKey || event.altKey || isTyping(event.target)) return;
    event.preventDefault();
    open();
  });

  useEffect(() => {
    window.addEventListener("keydown", onKeyDown);
    return () => window.removeEventListener("keydown", onKeyDown);
  }, []);

  return (
    <OpenAgentContext value={open}>
      {children}

      <dialog
        ref={dialogRef}
        aria-labelledby="agent-title"
        // The content fills the whole dialog, so a click on the dialog itself is a click on the backdrop.
        onClick={(event) => event.target === event.currentTarget && close()}
        className="fixed inset-x-0 top-0 bottom-auto m-0 max-h-[min(78vh,640px)] w-full max-w-full overflow-auto overscroll-contain border-b border-faint bg-paper-2 text-ink shadow-[0_18px_40px_-18px_rgb(0_0_0/0.35)] backdrop:bg-crust/45 open:animate-drop motion-reduce:open:animate-none max-[560px]:max-h-[88dvh]"
      >
        <div className="px-(--gutter) pt-[calc(18px+env(safe-area-inset-top,0px))] pb-5.5">
          <div className="page-width mx-auto">
            <div className="max-w-(--col)">
              <div className="flex items-baseline justify-between gap-4">
                <p id="agent-title" className="mb-1.5 font-semibold">
                  <span className="font-normal text-dim">
                    <span className="text-prompt">{cwd}</span> ${" "}
                  </span>
                  <span className="text-cmd">{AGENT_COMMAND}</span>
                </p>
                <button
                  type="button"
                  onClick={close}
                  className="cursor-pointer rounded border border-faint px-2.5 py-0.5 text-dim hover:border-prompt hover:text-ink"
                >
                  exit
                </button>
              </div>
              <p className="mb-4 text-small text-dim">demo · scripted replies</p>

              <AgentChat posts={posts} inputRef={inputRef} onExit={close} />
            </div>
          </div>
        </div>
      </dialog>
    </OpenAgentContext>
  );
}
