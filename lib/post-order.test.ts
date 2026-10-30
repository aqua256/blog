import { describe, expect, it } from "vitest";
import { newestFirst } from "@/lib/post-order";

describe("newestFirst", () => {
  it("puts the newest post first, across months and years", () => {
    const posts = [{ date: "2026-09-29" }, { date: "2027-01-03" }, { date: "2026-10-14" }];
    expect(newestFirst(posts).map(({ date }) => date)).toEqual(["2027-01-03", "2026-10-14", "2026-09-29"]);
  });

  it("keeps posts from the same day in their original order", () => {
    const posts = [
      { slug: "a", date: "2026-10-14" },
      { slug: "b", date: "2026-10-14" },
    ];
    expect(newestFirst(posts).map(({ slug }) => slug)).toEqual(["a", "b"]);
  });

  it("leaves the array it was given untouched", () => {
    const posts = [{ date: "2026-09-29" }, { date: "2026-10-14" }];
    newestFirst(posts);
    expect(posts.map(({ date }) => date)).toEqual(["2026-09-29", "2026-10-14"]);
  });
});
