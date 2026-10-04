"use client";

import { site } from "@/lib/site";
import { useOpenAgent } from "./agent-console";

/**
 * The idle prompt `~ $ █` as a real input. Typing there opens the agent console with what was typed,
 * so the last line of the page works like the terminal it looks like.
 */
export function PromptInput({ cwd }: { cwd: string }) {
  const open = useOpenAgent();

  function handOver(input: HTMLInputElement) {
    const text = input.value;
    input.value = "";
    // `/` opens the console everywhere else, so here it does too, without being typed
    if (text) open(text === "/" ? "" : text);
  }

  return (
    // The label makes a click anywhere on the line focus the input; the focus ring hugs just the prompt
    <label className="mb-3.5 block cursor-text font-semibold">
      <span className="has-focus-visible:outline-2 has-focus-visible:outline-offset-3 has-focus-visible:outline-prompt">
        <span className="select-none font-normal text-dim">
          <span className="text-prompt">{cwd}</span> ${" "}
        </span>
        <input
          type="text"
          aria-label={`Ask ${site.givenName}'s agent`}
          autoComplete="off"
          size={1}
          // While an IME is composing (pinyin, kana), wait for the finished characters
          onInput={(event) => !event.nativeEvent.isComposing && handOver(event.currentTarget)}
          onCompositionEnd={(event) => handOver(event.currentTarget)}
          className="max-w-full bg-transparent p-0 font-semibold text-ink caret-transparent outline-none field-sizing-content"
        />
        <span aria-hidden className="block-cursor" />
      </span>
    </label>
  );
}
