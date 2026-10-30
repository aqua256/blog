import { HOLDER } from "./rehype-code-title-keep.mjs";

/**
 * Runs after @shikijs/rehype: moves the title remembered by rehype-code-title-keep.mjs onto the highlighted
 * <pre> as data-title, and removes the placeholder.
 */
export default function rehypeCodeTitleApply() {
  return apply;
}

function apply(node) {
  node.children?.forEach((child, i) => {
    if (child.type !== "element") return;
    if (child.tagName === HOLDER) {
      // Shiki puts its output in as a fragment (a root node) holding the <pre>
      const pre = child.children
        .flatMap((c) => (c.type === "root" ? c.children : [c]))
        .find((c) => c.type === "element" && c.tagName === "pre");
      if (pre) {
        pre.properties = { ...pre.properties, dataTitle: child.properties.title };
        node.children[i] = pre;
      }
      return;
    }
    apply(child);
  });
}
