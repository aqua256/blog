import { slug } from "github-slugger";
import { isValidElement, type ReactNode } from "react";

/** A `##` section of a post, for the outline. */
export type Heading = { id: string; text: string };

/**
 * One slug function for both sides, so outline links always match the rendered <h2> ids:
 * the h2 in mdx-components.tsx and extractHeadings() below both call it.
 * Two sections with the exact same title get the same id (the link goes to the first).
 */
export const headingId = (text: string) => slug(text);

/** Plain text of rendered MDX children, e.g. `Why <code>ls</code>` → "Why ls". */
export function nodeText(node: ReactNode): string {
  if (typeof node === "string" || typeof node === "number") return String(node);
  if (Array.isArray(node)) return node.map(nodeText).join("");
  if (isValidElement<{ children?: ReactNode }>(node)) return nodeText(node.props.children);
  return "";
}

/** Markdown inline syntax → the text it renders as. */
function stripInline(markdown: string): string {
  return markdown
    .replace(/!?\[([^\]]*)\]\([^)]*\)/g, "$1") // links and images → their text
    .replace(/`([^`]*)`/g, "$1") // inline code
    .replace(/(\*\*|__|\*|_|~~)(.+?)\1/g, "$2") // bold, italic, strikethrough
    .trim();
}

/** The `##` headings in an MDX source, in order. Lines inside fenced code blocks are ignored. */
export function extractHeadings(source: string): Heading[] {
  const headings: Heading[] = [];
  let inFence = false;
  for (const line of source.split("\n")) {
    if (/^\s*(```|~~~)/.test(line)) inFence = !inFence;
    if (inFence) continue;
    const match = /^##\s+(.+?)\s*#*\s*$/.exec(line);
    if (match) {
      const text = stripInline(match[1]);
      headings.push({ id: headingId(text), text });
    }
  }
  return headings;
}
