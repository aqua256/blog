"use client";

import { useSyncExternalStore } from "react";

const noopSubscribe = () => () => {};

/** `%E6%89%93` → `打`; a malformed escape such as `%E0` is shown as typed. */
function readablePath() {
  const { pathname } = window.location;
  try {
    return decodeURI(pathname);
  } catch {
    return pathname;
  }
}

/**
 * The address the visitor asked for, as code: `/posts/typo`. The 404 page is built once as a static
 * page, so the path is only known in the browser; until then nothing is shown.
 */
export function RequestedPath() {
  const pathname = useSyncExternalStore(noopSubscribe, readablePath, () => "");
  return pathname ? <code className="[overflow-wrap:anywhere]">{pathname.replace(/(.)\/$/, "$1")}</code> : null;
}
