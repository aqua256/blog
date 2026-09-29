import createMDX from "@next/mdx";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // Every page is built ahead of time, so the site ships as plain files in out/ (served by Cloudflare)
  output: "export",
  // Allow .md and .mdx files as pages and imports
  pageExtensions: ["js", "jsx", "md", "mdx", "ts", "tsx"],
};

const withMDX = createMDX({
  // Turbopack takes plugins by package name, with JSON-serializable options only
  options: {
    // GitHub Flavored Markdown: tables, ~~strikethrough~~, task lists, footnotes, bare URLs as links
    remarkPlugins: ["remark-gfm"],
    rehypePlugins: [
      // Unique ids on every heading (a repeated title becomes `title-1`), then the headings exported
      // from each post as `tableOfContents`, which the outline and the agent read (lib/headings.ts)
      "rehype-slug",
      "@stefanprobst/rehype-extract-toc",
      "@stefanprobst/rehype-extract-toc/mdx",
      // Syntax highlighting at build time, in Catppuccin. Colors are written as light-dark(latte, mocha),
      // so they follow the page's color-scheme, which next-themes sets from THEME=.
      [
        "@shikijs/rehype",
        {
          themes: { light: "catppuccin-latte", dark: "catppuccin-mocha" },
          defaultColor: "light-dark()",
          colorsRendering: "none",
          addLanguageClass: true,
        },
      ],
    ],
  },
});

export default withMDX(nextConfig);
