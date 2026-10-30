import type { MDXComponents } from "mdx/types";
import { isValidElement, type ReactNode } from "react";
import { CopyButton } from "./app/_components/copy-button";

/** `ts` for a <code class="language-ts"> inside a code block (the class comes from @shikijs/rehype). */
function languageOf(children: ReactNode) {
  if (!isValidElement<{ className?: string }>(children)) return undefined;
  return /language-(\S+)/.exec(children.props.className ?? "")?.[1];
}

const components: MDXComponents = {
  // A code block gets a bar above the code: its file name (```ts title="routes.ts", see
  // lib/rehype-code-title-keep.mjs) or else its language, and a copy button. The bar stays put while the
  // <pre> scrolls sideways.
  pre: ({ "data-title": title, ...props }: React.ComponentProps<"pre"> & { "data-title"?: string }) => (
    <figure className="code-block">
      <figcaption>
        <span>{title ?? languageOf(props.children)}</span>
        <CopyButton />
      </figcaption>
      <pre {...props} />
    </figure>
  ),
};

export function useMDXComponents(): MDXComponents {
  return components;
}
