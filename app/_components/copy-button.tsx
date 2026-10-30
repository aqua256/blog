"use client";

import { useEffect, useRef, useState } from "react";

const icon = {
  viewBox: "0 0 24 24",
  fill: "none",
  stroke: "currentColor",
  strokeWidth: 1.75,
  strokeLinecap: "round",
  strokeLinejoin: "round",
  "aria-hidden": true,
  className: "size-4",
} as const;

/** Copies the code of the block it sits in. For 1.5s after a copy it shows a check and announces 已复制. */
export function CopyButton() {
  const button = useRef<HTMLButtonElement>(null);
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    if (!copied) return;
    const timer = setTimeout(() => setCopied(false), 1500);
    return () => clearTimeout(timer);
  }, [copied]);

  async function copy() {
    const code = button.current?.closest("figure")?.querySelector("pre")?.textContent;
    if (code == null) return;
    await navigator.clipboard.writeText(code);
    setCopied(true);
  }

  return (
    <button
      ref={button}
      type="button"
      onClick={copy}
      aria-label="复制代码"
      title="复制代码"
      className="grid size-10 cursor-pointer place-items-center rounded-md text-ink-3 hover:bg-ink/5 hover:text-ink"
    >
      {copied ? (
        <svg {...icon}>
          <path d="M20 6 9 17l-5-5" />
        </svg>
      ) : (
        <svg {...icon}>
          <rect x="9" y="9" width="12" height="12" rx="2" />
          <path d="M5 15V5a2 2 0 0 1 2-2h10" />
        </svg>
      )}
      <output className="sr-only">{copied ? "已复制" : ""}</output>
    </button>
  );
}
