# Aqua256's blog

A personal blog: posts as cards beside a profile card, with a leaf shadow reaching in from the top-right
corner of every page. The first page of a visit opens with that shadow coming into focus.

Built with Next.js 16 (App Router), Tailwind CSS v4 and MDX. Colors are neutral grays with one blue accent,
light and dark, and code is highlighted in [Flexoki](https://stephango.com/flexoki). Text uses the system's
own fonts (PingFang SC, Microsoft YaHei or Noto Sans CJK for Chinese), so no font files are downloaded.

## Development

```bash
pnpm install
pnpm dev        # http://localhost:3000
pnpm typecheck  # route types, then TypeScript 7
pnpm lint       # oxlint
pnpm test       # Vitest
pnpm build      # production build; drafts are left out
```

TypeScript is version 7, which has no JavaScript API yet: keep your editor on its bundled TypeScript rather
than the workspace version, and rely on `pnpm typecheck`.

## Writing a post

Each post is one MDX file in `content/posts/`. The file name is the URL: `content/posts/hello-world.mdx`
is served at `/posts/hello-world`.

A post starts with its metadata:

```mdx
export const metadata = {
  title: "Hello, world",
  date: "2026-09-29", // YYYY-MM-DD; posts are sorted by it
  tags: ["meta"], // optional
  description: "One or two sentences.", // optional; for cards, RSS and shared links, defaults to the first paragraph
  draft: true, // optional; drafts only show up in `pnpm dev`
};

The first paragraph.

## A section
```

- `##` headings become the outline beside the post on wide screens (folded under the title on phones).
  Every heading gets an id to link to (`#title`); a repeated title becomes `#title-1`, `#title-2`, as on
  GitHub.
- GitHub Flavored Markdown works: tables, `~~strikethrough~~`, task lists (`- [x]`) and footnotes (`[^1]`).
  Footnotes are listed at the end under 参考.
- Fenced code blocks with a language (```` ```ts ````) are highlighted at build time. Add a file name to
  show it above the code: ```` ```ts title="routes.ts" ````. Every block has a copy button.
- Images fill the column; for a caption, write `<figure>` with an `<img>` and a `<figcaption>`.
- Each post gets a share image (title, date and tags on white, with the leaf shadow), made at build time.

## Site content

| What | Where |
| --- | --- |
| Name, introduction, navigation, social links | `lib/site.ts` |
| The about page's text | `content/about.mdx` |
| The avatar in the profile card (also the Apple touch icon) | `app/apple-icon.png` |

The introduction in `lib/site.ts` is shared by the profile card, the RSS feed and the page description.

## Deployment

The site is a static export (`output: "export"`): `pnpm build` writes every page to `out/`, and Cloudflare
Workers serves those files as static assets (`wrangler.jsonc`), with `out/404.html` for unknown URLs.
There is no server code. `public/_headers` and `public/_redirects` are copied into `out/` and applied by
Cloudflare; `/projects`, a page that no longer exists, redirects to the home page.

On Cloudflare, the Worker is connected to this GitHub repository and deploys on every push to `main`:

- Build command: `pnpm build`
- Deploy command: `npx wrangler deploy`
- Build variable: `SITE_URL`

To try the production build locally: `pnpm build && npx wrangler dev`.

At least one published (non-draft) post is needed: a static export can't build `/posts/[slug]` with no
posts.

Set `SITE_URL` to the site's public address, e.g. `https://example.com`. It is used for absolute links in
the RSS feed, the sitemap and page metadata. Without it the build falls back to `http://localhost:3000`.

These are generated at build time, so changing `SITE_URL` needs a rebuild:

- `/rss.xml`: RSS 2.0 feed of the published posts
- `/sitemap.xml` and `/robots.txt`

## Layout

```
app/
  layout.tsx          header, leaf shadow, footer, site metadata
  page.tsx            home: the latest posts
  posts/              the archive (归档) and each post, with its share image
  about/              the about page
  not-found.tsx       404
  rss.xml/, sitemap.ts, robots.ts
  _components/        shared components
content/posts/        the posts (MDX)
content/about.mdx     the about page's text
lib/                  site data, posts and their order, headings, code block titles, Flexoki code themes
mdx-components.tsx    how MDX elements render (code blocks)
public/               the leaf shadow drawing, Cloudflare headers and redirects
wrangler.jsonc        Cloudflare: serve out/ as static assets
```
