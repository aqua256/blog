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
 * The URL the visitor asked for, as a path under ~: `/posts/typo` → `~/posts/typo`.
 * The 404 page is built once as a static page, so the path is only known in the browser.
 */
export function RequestedPath() {
  const pathname = useSyncExternalStore(noopSubscribe, readablePath, () => "");
  return <>{pathname && `~${pathname.replace(/\/$/, "")}`}</>;
}
