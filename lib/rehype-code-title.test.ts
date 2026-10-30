import { describe, expect, it } from "vitest";
import rehypeCodeTitleApply from "@/lib/rehype-code-title-apply.mjs";
import rehypeCodeTitleKeep from "@/lib/rehype-code-title-keep.mjs";

type Node = { type: string; tagName?: string; properties?: Record<string, unknown>; data?: { meta?: string }; children?: Node[] };

/** A fenced code block as remark-rehype hands it over: <pre><code> with the fence's meta string. */
function block(meta?: string): Node {
  return {
    type: "element",
    tagName: "pre",
    properties: {},
    children: [{ type: "element", tagName: "code", properties: {}, data: meta ? { meta } : {}, children: [] }],
  };
}

/** What @shikijs/rehype does to each <pre>: puts a new one in its place, as a fragment, without the meta. */
function highlight(node: Node) {
  node.children?.forEach((child, i) => {
    if (child.tagName === "pre") {
      node.children![i] = { type: "root", children: [{ type: "element", tagName: "pre", properties: { class: "shiki" }, children: [] }] };
    } else highlight(child);
  });
}

function run(tree: Node) {
  rehypeCodeTitleKeep()(tree);
  highlight(tree);
  rehypeCodeTitleApply()(tree);
  return tree;
}

describe("rehype-code-title", () => {
  it("puts the fence's title on the highlighted <pre>", () => {
    const tree = run({ type: "root", children: [block('title="routes.ts"')] });
    expect(tree.children).toEqual([
      { type: "element", tagName: "pre", properties: { class: "shiki", dataTitle: "routes.ts" }, children: [] },
    ]);
  });

  it("finds a title among other meta and in nested blocks", () => {
    const tree = run({ type: "root", children: [{ type: "element", tagName: "li", children: [block('{1,3} title="a b.ts"')] }] });
    expect(tree.children?.[0].children?.[0].properties?.dataTitle).toBe("a b.ts");
  });

  it("leaves blocks without a title alone", () => {
    const tree = run({ type: "root", children: [block(), block("showLineNumbers")] });
    expect(tree.children?.map((c) => c.type)).toEqual(["root", "root"]);
  });
});
