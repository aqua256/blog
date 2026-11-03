/**
 * Keeps an arrow on the line of the word before it, so a menu path such as `Notes → Settings → Body`
 * never starts a wrapped line with `→`: the space before each arrow becomes a no-break space. Text in
 * code is left as written. A plain .mjs file with no dependencies, like the other local plugins:
 * Turbopack loads MDX plugins by path.
 */
const SPACE_BEFORE_ARROW = /[ \t]+→/g;

export default function rehypeArrowNowrap() {
  return walk;
}

function walk(node) {
  node.children?.forEach((child) => {
    if (child.type === "text") {
      child.value = child.value.replace(SPACE_BEFORE_ARROW, " →");
    } else if (child.type !== "element" || (child.tagName !== "code" && child.tagName !== "pre")) {
      walk(child);
    }
  });
}
