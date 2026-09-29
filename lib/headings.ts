import type { Toc } from "@stefanprobst/rehype-extract-toc";

/** A `##` section of a post, for the outline and the agent's summary. */
export type Heading = { id: string; text: string };

/**
 * The `##` sections from a post's table of contents, in order.
 * The MDX pipeline builds it (next.config.ts): rehype-slug gives every heading a unique id
 * (a repeated title becomes `title-1`), and rehype-extract-toc exports the headings as `tableOfContents`.
 */
export function sectionsOf(toc: Toc): Heading[] {
  return toc.flatMap(({ depth, value, id, children }) => [
    ...(depth === 2 && id ? [{ id, text: value }] : []),
    ...sectionsOf(children ?? []),
  ]);
}
