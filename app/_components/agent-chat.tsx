"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useRef, useState } from "react";
import { cleanQuestion, reply, type AgentPost, type Segment } from "@/lib/agent-replies";
import { site } from "@/lib/site";

/** Delay between words while an answer is being typed out. */
const WORD_DELAY_MS = 34;

type Exchange = { question: string; answer: Segment[] };

/** Plain text split into words (each keeping its trailing space), so answers can appear word by word. */
function toWords(answer: Segment[]): Segment[] {
  return answer.flatMap((segment): Segment[] =>
    typeof segment === "string" ? (segment.match(/\S+\s*|\s+/g) ?? []) : [segment],
  );
}

function prefersReducedMotion() {
  return window.matchMedia("(prefers-reduced-motion: reduce)").matches;
}

function SegmentView({ segment }: { segment: Segment }) {
  if (typeof segment === "string") return segment;
  if ("code" in segment) {
    return <code className="term text-[0.92em] text-code">{segment.code}</code>;
  }
  const className = "text-link underline underline-offset-3";
  return segment.href.startsWith("/") ? (
    <Link href={segment.href} className={className}>
      {segment.text}
    </Link>
  ) : (
    <a href={segment.href} className={className}>
      {segment.text}
    </a>
  );
}

const CHIPS = [
  { label: `who is ${site.givenName}?`, question: `who is ${site.givenName}?` },
  { label: "summarize this post", question: "summarize this post", onPostOnly: true },
  { label: "what else?", question: "what else have you written?" },
  { label: "get in touch", question: "how do I get in touch?" },
];

/**
 * The conversation inside the agent console:
 *
 *   $ ask "who is Aqua256?"
 *   <answer, typed out word by word>
 *   [ask "who is Aqua256?"] [ask "summarize this post"] …
 *   $ _
 */
export function AgentChat({
  posts,
  inputRef,
  onExit,
}: {
  posts: AgentPost[];
  inputRef: React.RefObject<HTMLInputElement | null>;
  onExit: () => void;
}) {
  const pathname = usePathname();
  const current = posts.find((post) => pathname === `/posts/${post.slug}`);

  const [log, setLog] = useState<Exchange[]>([]);
  // How many words of the newest answer are visible so far
  const [shownWords, setShownWords] = useState(0);
  const lastQuestion = useRef("");
  const formRef = useRef<HTMLFormElement>(null);

  const typing = log.length > 0 && shownWords < log[log.length - 1].answer.length;

  // Type the newest answer out, one word per tick
  useEffect(() => {
    if (!typing) return;
    const timer = setTimeout(() => setShownWords((n) => n + 1), WORD_DELAY_MS);
    return () => clearTimeout(timer);
  }, [typing, shownWords]);

  // Keep the prompt in view as the log grows
  useEffect(() => {
    formRef.current?.scrollIntoView({ block: "nearest" });
  }, [log, shownWords]);

  function ask(raw: string) {
    const question = cleanQuestion(raw);
    if (typing || !question) return;
    if (/^(exit|quit)$/i.test(question)) return onExit();
    if (/^(clear|cls)$/i.test(question)) return setLog([]);

    lastQuestion.current = raw.trim();
    const answer = toWords(reply(question, { posts, current }));
    setLog((log) => [...log, { question, answer }]);
    setShownWords(prefersReducedMotion() ? answer.length : 0);
  }

  function onSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const input = inputRef.current;
    if (!input || typing) return;
    ask(input.value);
    input.value = "";
  }

  // ↑ on an empty prompt brings back the last question, like shell history
  function onKeyDown(event: React.KeyboardEvent<HTMLInputElement>) {
    const input = event.currentTarget;
    if (event.key === "ArrowUp" && !input.value && lastQuestion.current) {
      event.preventDefault();
      input.value = lastQuestion.current;
    }
  }

  // Links in an answer lead somewhere on the page or site, so the console gets out of the way first
  function onLogClick(event: React.MouseEvent) {
    if (event.target instanceof Element && event.target.closest("a")) onExit();
  }

  return (
    <>
      <div aria-live="polite" onClick={onLogClick} className="flex flex-col gap-4.5 empty:hidden">
        {log.map(({ question, answer }, i) => {
          const visible = i === log.length - 1 ? answer.slice(0, shownWords) : answer;
          return (
            <div key={i}>
              <p className="mb-1.5 font-semibold [overflow-wrap:anywhere]">
                <span className="font-normal text-dim">$ </span>
                <span className="text-cmd">ask</span> <span className="text-str">&quot;{question}&quot;</span>
              </p>
              <p className="font-[380] text-body leading-[1.7] whitespace-pre-line [font-variation-settings:'MONO'_0,'CASL'_0.3]">
                {visible.map((segment, j) => (
                  <SegmentView key={j} segment={segment} />
                ))}
              </p>
            </div>
          );
        })}
      </div>

      <ul aria-label="Suggested questions" className="mt-4.5 mb-3.5 flex flex-wrap gap-x-2.5 gap-y-2">
        {CHIPS.filter((chip) => !chip.onPostOnly || current).map(({ label, question }) => (
          <li key={label}>
            <button
              type="button"
              onClick={() => ask(question)}
              className="cursor-pointer rounded border border-faint px-2.5 py-1 text-small text-dim hover:border-prompt hover:text-ink"
            >
              ask &quot;{label}&quot;
            </button>
          </li>
        ))}
      </ul>

      <form
        ref={formRef}
        onSubmit={onSubmit}
        autoComplete="off"
        aria-busy={typing}
        className="flex items-center gap-x-2.5 border-y border-faint py-2.5 font-semibold focus-within:border-prompt"
      >
        <span aria-hidden className="text-dim">
          $
        </span>
        <input
          ref={inputRef}
          type="text"
          aria-label={`Ask ${site.givenName}'s agent`}
          maxLength={200}
          enterKeyHint="send"
          placeholder={`ask anything about ${site.givenName}`}
          onKeyDown={onKeyDown}
          className="min-w-[10ch] flex-1 bg-transparent py-1 font-medium text-ink caret-cursor outline-none placeholder:font-normal placeholder:text-dim/80"
        />
      </form>
    </>
  );
}
