import { describe, expect, it } from "vitest";
import { isCurrentSection } from "@/lib/nav";

describe("isCurrentSection", () => {
  it("marks the home link only on the home page", () => {
    expect(isCurrentSection("/", "/")).toBe(true);
    expect(isCurrentSection("/posts", "/")).toBe(false);
    expect(isCurrentSection("/posts/hello-world", "/")).toBe(false);
  });

  it("marks the archive link on the archive and on every post", () => {
    expect(isCurrentSection("/posts", "/posts")).toBe(true);
    expect(isCurrentSection("/posts/hello-world", "/posts")).toBe(true);
  });

  it("does not mark a link whose path only starts with the same letters", () => {
    expect(isCurrentSection("/postscript", "/posts")).toBe(false);
  });

  it("does not mark the archive link elsewhere", () => {
    expect(isCurrentSection("/", "/posts")).toBe(false);
    expect(isCurrentSection("/about", "/posts")).toBe(false);
  });
});
