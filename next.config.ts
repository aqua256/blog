import createMDX from "@next/mdx";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  // Allow .md and .mdx files as pages and imports
  pageExtensions: ["js", "jsx", "md", "mdx", "ts", "tsx"],
};

const withMDX = createMDX({
  // Turbopack takes plugins by package name, with JSON-serializable options only
  options: {
    // GitHub Flavored Markdown: tables, ~~strikethrough~~, task lists, footnotes, bare URLs as links
    remarkPlugins: ["remark-gfm"],
    // Syntax highlighting at build time, in Catppuccin. Colors are written as light-dark(latte, mocha),
    // so they follow the page's color-scheme, which next-themes sets from THEME=.
    rehypePlugins: [
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
