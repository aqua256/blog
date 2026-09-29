import type { MDXComponents } from "mdx/types";
import { isValidElement, type ReactNode } from "react";
import { headingId, nodeText } from "@/lib/headings";

/** `ts` for a <code class="language-ts"> inside a code block (the class comes from @shikijs/rehype). */
function languageOf(children: ReactNode) {
  if (!isValidElement<{ className?: string }>(children)) return undefined;
  return /language-(\S+)/.exec(children.props.className ?? "")?.[1];
}

const components: MDXComponents = {
  // `##` sections get an id so the outline can link to them
  h2: ({ children, ...props }) => (
    <h2 id={headingId(nodeText(children))} {...props}>
      {children}
    </h2>
  ),
  // Code blocks get a wrapper that stays put while the <pre> scrolls sideways,
  // so the ``` mark and the language label don't scroll with the code
  pre: (props) => (
    <div className="code-block" data-language={languageOf(props.children)}>
      <pre {...props} />
    </div>
  ),
};

export function useMDXComponents(): MDXComponents {
  return components;
}
