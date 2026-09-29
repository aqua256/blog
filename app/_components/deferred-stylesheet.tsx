"use client";

import { useEffect } from "react";
import { preconnect, preload } from "react-dom";

/**
 * A stylesheet that never holds up the first paint, for font CSS on a third-party CDN.
 * A plain <link rel="stylesheet"> blocks rendering until it arrives, so a slow or unreachable CDN
 * would leave the page blank. Instead the CSS is preloaded (downloads early, blocks nothing) and
 * applied once the page is interactive; until then text shows in the fallback fonts.
 * The font files it points to are CORS requests to the same CDN, so that connection is opened early too.
 */
export function DeferredStylesheet({ href }: { href: string }) {
  preload(href, { as: "style" });
  preconnect(new URL(href).origin, { crossOrigin: "anonymous" });

  useEffect(() => {
    if (document.querySelector(`link[rel="stylesheet"][href="${href}"]`)) return;
    const link = document.createElement("link");
    link.rel = "stylesheet";
    link.href = href;
    document.head.append(link);
  }, [href]);

  return null;
}
