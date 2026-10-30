import type { Toc } from "@stefanprobst/rehype-extract-toc";
import { describe, expect, it } from "vitest";
import { sectionsOf } from "@/lib/headings";

describe("sectionsOf", () => {
  it("lists the ## headings in document order", () => {
    const toc: Toc = [
      { depth: 2, value: "First", id: "first" },
      { depth: 2, value: "Second", id: "second" },
    ];
    expect(sectionsOf(toc)).toEqual([
      { id: "first", text: "First" },
      { id: "second", text: "Second" },
    ]);
  });

  it("finds ## headings nested under other levels and leaves the other levels out", () => {
    const toc: Toc = [
      {
        depth: 1,
        value: "Title",
        id: "title",
        children: [
          { depth: 2, value: "Section", id: "section", children: [{ depth: 3, value: "Detail", id: "detail" }] },
        ],
      },
    ];
    expect(sectionsOf(toc)).toEqual([{ id: "section", text: "Section" }]);
  });

  it("keeps the ids rehype-slug gives repeated titles", () => {
    const toc: Toc = [
      { depth: 2, value: "Notes", id: "notes" },
      { depth: 2, value: "Notes", id: "notes-1" },
    ];
    expect(sectionsOf(toc).map(({ id }) => id)).toEqual(["notes", "notes-1"]);
  });

  it("skips a heading without an id, which the outline could not link to", () => {
    const toc: Toc = [{ depth: 2, value: "No anchor" }];
    expect(sectionsOf(toc)).toEqual([]);
  });
});
