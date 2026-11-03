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
  className: "size-3.5",
} as const;

/**
 * Copies the code of the block it sits in. For 1.5s after a copy it shows a check and announces 已复制;
 * when the browser refuses (permission denied, or no clipboard outside HTTPS) it says 复制失败 instead.
 */
export function CopyButton() {
  const button = useRef<HTMLButtonElement>(null);
  const [status, setStatus] = useState<"idle" | "copied" | "failed">("idle");

  useEffect(() => {
    if (status === "idle") return;
    const timer = setTimeout(() => setStatus("idle"), 1500);
    return () => clearTimeout(timer);
  }, [status]);

  async function copy() {
    const code = button.current?.closest("figure")?.querySelector("pre")?.textContent;
    if (code == null) return;
    try {
      await navigator.clipboard.writeText(code);
      setStatus("copied");
    } catch {
      // The failure itself is the browser's to report; the reader only needs to know to copy by hand
      setStatus("failed");
    }
  }

  return (
    <button
      ref={button}
      type="button"
      onClick={copy}
      aria-label="复制代码"
      title="复制代码"
      className={`grid h-10 min-w-10 cursor-pointer place-items-center rounded-md text-ink-3 hover:bg-ink/5 hover:text-ink ${status === "failed" ? "px-3" : ""}`}
    >
      {status === "copied" ? (
        <svg {...icon}>
          <path d="M20 6 9 17l-5-5" />
        </svg>
      ) : status === "failed" ? (
        <span aria-hidden="true">复制失败</span>
      ) : (
        <svg {...icon}>
          <rect x="9" y="9" width="12" height="12" rx="2" />
          <path d="M5 15V5a2 2 0 0 1 2-2h10" />
        </svg>
      )}
      <output className="sr-only">{status === "copied" ? "已复制" : status === "failed" ? "复制失败" : ""}</output>
    </button>
  );
}
