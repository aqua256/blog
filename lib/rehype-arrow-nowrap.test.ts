import { describe, expect, it } from "vitest";
import rehypeArrowNowrap from "@/lib/rehype-arrow-nowrap.mjs";

type Node = { type: string; tagName?: string; value?: string; children?: Node[] };

const text = (value: string): Node => ({ type: "text", value });
const element = (tagName: string, children: Node[]): Node => ({ type: "element", tagName, children });

function run(tree: Node) {
  rehypeArrowNowrap()(tree);
  return tree;
}

describe("rehype-arrow-nowrap", () => {
  it("joins each arrow to the word before it", () => {
    const tree = run({ type: "root", children: [element("li", [text("Notes → Settings → Body")])] });
    expect(tree.children?.[0].children?.[0].value).toBe("Notes → Settings → Body");
  });

  it("leaves code as written", () => {
    const tree = run({ type: "root", children: [element("p", [element("code", [text("a → b")])]), element("pre", [text("a → b")])] });
    expect(tree.children?.[0].children?.[0].children?.[0].value).toBe("a → b");
    expect(tree.children?.[1].children?.[0].value).toBe("a → b");
  });
});
