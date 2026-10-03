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
 * Some Flexoki syntax colors fall short of WCAG AA (4.5:1) on the code block background (base-50 in
 * light, base-950 in dark). Each is mixed with black (light) or white (dark) by the least that passes,
 * as the text colors in globals.css are.
 */
const flexokiLightReadable = withColors(flexokiLight, {
  "#6F6E69": "#6E6D68", // base-600, comments and punctuation: 4.47 → 4.54
  "#BC5215": "#B44F14", // orange: 4.22 → 4.51
  "#AD8301": "#896701", // yellow: 3.05 → 4.59
  "#66800B": "#5E760A", // green: 3.94 → 4.51
  "#24837B": "#217971", // cyan: 3.99 → 4.55
});
const flexokiDarkReadable = withColors(flexokiDark, {
  "#D14D41": "#D55D52", // red: 3.97 → 4.52
  "#4385BE": "#4989C0", // blue: 4.37 → 4.60
});

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
      // Syntax highlighting at build time, in Flexoki (its official VS Code themes, from kepano/flexoki,
      // MIT). Colors are written as light-dark(light, dark), so they follow the page's color-scheme, which
      // next-themes sets from THEME=.
      [
        "@shikijs/rehype",
        {
          themes: { light: flexokiLightReadable, dark: flexokiDarkReadable },
          defaultColor: "light-dark()",
          colorsRendering: "none",
          addLanguageClass: true,
        },
      ],
    ],
  },
});

export default withMDX(nextConfig);
