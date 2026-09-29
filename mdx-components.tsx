import type { MDXComponents } from "mdx/types";
import { headingId, nodeText } from "@/lib/headings";

const components: MDXComponents = {
  // `##` sections get an id so the outline can link to them
  h2: ({ children, ...props }) => (
    <h2 id={headingId(nodeText(children))} {...props}>
      {children}
    </h2>
  ),
};

export function useMDXComponents(): MDXComponents {
  return components;
}
