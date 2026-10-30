import path from "node:path";
import createMDX from "@next/mdx";
import type { NextConfig } from "next";
import flexokiDark from "./lib/shiki/flexoki-dark.json";
import flexokiLight from "./lib/shiki/flexoki-light.json";

type Theme = { tokenColors: { settings: { foreground?: string; fontStyle?: string } }[] };

/** The theme with some token colors swapped, e.g. { "#AD8301": "#896701" }. */
function withColors<T extends Theme>(theme: T, colors: Record<string, string>): T {
  return {
    ...theme,
    tokenColors: theme.tokenColors.map((token) => {
      const replacement = colors[token.settings.foreground ?? ""];
      return replacement ? { ...token, settings: { ...token.settings, foreground: replacement } } : token;
    }),
  };
}

/*
 * Some Flexoki syntax colors fall short of WCAG AA (4.5:1) on the code block background (#f7f9fb in
 * light, #1a1b1e in dark; --surface-code in globals.css). Each is mixed with black (light) or white (dark)
 * by the least that passes.
 */
const flexokiLightReadable = withColors(flexokiLight, {
  "#AD8301": "#906D01", // yellow: 3.30 → 4.55
  "#66800B": "#637C0B", // green: 4.27 → 4.50
  "#24837B": "#237F77", // cyan: 4.32 → 4.55
});
const flexokiDarkReadable = withColors(flexokiDark, {
  "#D14D41": "#D55D52", // red: 3.98 → 4.52
  "#4385BE": "#4989C0", // blue: 4.38 → 4.61
});

const nextConfig: NextConfig = {
  // Every page is built ahead of time, so the site ships as plain files in out/ (served by Cloudflare)
  output: "export",
  // A static export has no server to resize images on request; next/image serves them as they are
  images: { unoptimized: true },
  // Allow .md and .mdx files as pages and imports
  pageExtensions: ["js", "jsx", "md", "mdx", "ts", "tsx"],
};

// Local plugins go by absolute path: the MDX loader resolves names from its own directory
const local = (file: string) => path.resolve("lib", file);

const withMDX = createMDX({
  // Turbopack takes plugins by package name or path, with JSON-serializable options only
  options: {
    // GitHub Flavored Markdown: tables, ~~strikethrough~~, task lists, footnotes, bare URLs as links
    remarkPlugins: ["remark-gfm"],
    rehypePlugins: [
      // Unique ids on every heading (a repeated title becomes `title-1`), then the headings exported
      // from each post as `tableOfContents`, which the outline reads (lib/headings.ts)
      "rehype-slug",
      "@stefanprobst/rehype-extract-toc",
      "@stefanprobst/rehype-extract-toc/mdx",
      // A code block's file name (```ts title="routes.ts"), kept across Shiki's rewrite of the <pre>
      local("rehype-code-title-keep.mjs"),
      // Syntax highlighting at build time, in Flexoki (its official VS Code themes, from kepano/flexoki,
      // MIT). Colors are written as light-dark(light, dark), so they follow the page's color-scheme, which
      // next-themes sets from the theme button.
      [
        "@shikijs/rehype",
        {
          themes: { light: flexokiLightReadable, dark: flexokiDarkReadable },
          defaultColor: "light-dark()",
          colorsRendering: "none",
          addLanguageClass: true,
        },
      ],
      local("rehype-code-title-apply.mjs"),
    ],
  },
});

export default withMDX(nextConfig);
