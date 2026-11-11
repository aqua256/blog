"use client";

import { useState } from "react";

type Highlight = { x: number; y: number; width: number; height: number; on: boolean; glide: boolean };

/**
 * One hover highlight shared by a group of links (Fluid Functionalism's fluid hover). Coming in from outside
 * the group, it fades in where the pointer lands; moving from one item to the next, it glides over on the
 * fast spring; leaving the group, it fades out. `item` picks what it covers, by default the link or button
 * under the pointer. Touch and pen get no highlight, as there is no hover to follow. The look and motion are
 * .fluid-highlight in globals.css.
 */
export function FluidHover({
  item = "a, button",
  className = "",
  children,
}: {
  item?: string;
  className?: string;
  children: React.ReactNode;
}) {
  const [highlight, setHighlight] = useState<Highlight | null>(null);

  function onPointerOver(event: React.PointerEvent<HTMLDivElement>) {
    if (event.pointerType !== "mouse" || !(event.target instanceof Element)) return;
    const group = event.currentTarget;
    const target = event.target.closest(item);
    // Between items (in a gap, or on the group's own padding) the highlight stays where it was
    if (!target || !group.contains(target)) return;

    const from = group.getBoundingClientRect();
    const to = target.getBoundingClientRect();
    setHighlight((previous) => ({
      x: to.left - from.left,
      y: to.top - from.top,
      width: to.width,
      height: to.height,
      on: true,
      glide: previous?.on ?? false,
    }));
  }

  function onPointerLeave() {
    setHighlight((previous) => previous && { ...previous, on: false, glide: false });
  }

  return (
    <div className={`relative isolate ${className}`} onPointerOver={onPointerOver} onPointerLeave={onPointerLeave}>
      <span
        aria-hidden="true"
        className="fluid-highlight"
        data-on={highlight?.on || undefined}
        data-glide={highlight?.glide || undefined}
        style={
          highlight
            ? {
                transform: `translate(${highlight.x}px, ${highlight.y}px)`,
                width: highlight.width,
                height: highlight.height,
              }
            : undefined
        }
      />
      {children}
    </div>
  );
}
