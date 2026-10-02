"use client";

import { useEffect, useState } from "react";

// Once per page load: the first command typed animates, later navigations show it at once
let typedOnce = false;

/**
 * Text that types itself out, one character per step, e.g. the `hello-world.mdx` of `cat hello-world.mdx`.
 * It's always in the HTML; CSS only uncovers it (.typed-text in globals.css, off for reduced motion).
 */
export function TypedText({ text }: { text: string }) {
  const [animate] = useState(() => !typedOnce);
  useEffect(() => {
    typedOnce = true;
  }, []);

  return (
    <span className={animate ? "typed-text" : undefined} style={{ "--chars": text.length } as React.CSSProperties}>
      {text}
    </span>
  );
}
