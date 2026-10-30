/**
 * Lets a fenced code block carry a file name: ```ts title="routes.ts".
 *
 * @shikijs/rehype replaces each <pre> with its highlighted version and drops the fence's meta string, so
 * the title is carried across in two steps around it (next.config.ts). This one runs before Shiki and
 * wraps a titled <pre> in a placeholder that remembers the title; rehype-code-title-apply.mjs runs after
 * it. Two plugins rather than one with a "step" option, because unified merges repeated uses of the same
 * plugin into one. Plain .mjs files with no dependencies: Turbopack loads MDX plugins by path, with JSON
 * options only.
 */
export const HOLDER = "code-title-holder";

const TITLE = /(?:^|\s)title="([^"]*)"/;

export default function rehypeCodeTitleKeep() {
  return keep;
}

function keep(node) {
  node.children?.forEach((child, i) => {
    if (child.type !== "element") return;
    if (child.tagName === "pre") {
      const title = TITLE.exec(child.children?.[0]?.data?.meta ?? "")?.[1];
      if (title) {
        node.children[i] = { type: "element", tagName: HOLDER, properties: { title }, children: [child] };
      }
      return;
    }
    keep(child);
  });
}
