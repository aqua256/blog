# Aqua256's blog

A personal blog that reads like a friendly terminal: every page is a shell command and its output
(`cd ~/posts`, `ls -t`, `cat hello-world.mdx`), with a drop-down agent console you can ask about the site.

Built with Next.js 16 (App Router), Tailwind CSS v4 and MDX. Colors are [Catppuccin](https://catppuccin.com)
Latte (light) and Mocha (dark); type is [Recursive](https://www.recursive.design) for Latin text and
[LXGW WenKai](https://github.com/lxgw/LxgwWenKai) for Chinese.

## Development

```bash
pnpm install
pnpm dev      # http://localhost:3000
pnpm lint
pnpm build    # production build; drafts are left out
```

## Writing a post

Each post is one MDX file in `content/posts/`. The file name is the URL: `content/posts/hello-world.mdx`
is served at `/posts/hello-world`.

A post starts with its metadata:

```mdx
export const metadata = {
  title: "Hello, world",
  date: "2026-09-29", // YYYY-MM-DD; posts are sorted by it
  tags: ["meta"], // optional
  description: "One or two sentences.", // optional; RSS summary, defaults to the first paragraph
  draft: true, // optional; drafts only show up in `pnpm dev`
};

The first paragraph is shown larger, as the lede.

## A section
```

- `##` headings become the outline beside the post on wide screens, and the agent lists them when asked
  to summarize the post. Two sections with the same title share one link.
- GitHub Flavored Markdown works: tables, `~~strikethrough~~`, task lists (`- [x]`) and footnotes (`[^1]`).
- Fenced code blocks with a language (```` ```ts ````) are highlighted at build time.
- Reading time is estimated from the prose (code blocks are skipped).

## Site content

| What | Where |
| --- | --- |
| Name, introduction, navigation, social links | `lib/site.ts` |
| Projects shown in `~/projects` | `lib/projects.ts` |
| The agent's scripted replies | `lib/agent-replies.ts` |

The introduction in `lib/site.ts` is shared by the home page (`whoami`), the agent, the RSS feed and
the page description.

## Deployment

The site is a static export (`output: "export"`): `pnpm build` writes every page to `out/`, and Cloudflare
Workers serves those files as static assets (`wrangler.jsonc`), with `out/404.html` for unknown URLs.
There is no server code.

On Cloudflare, the Worker is connected to this GitHub repository and deploys on every push to `main`:

- Build command: `pnpm build`
- Deploy command: `npx wrangler deploy`
- Build variable: `SITE_URL`

To try the production build locally: `pnpm build && npx wrangler dev`.

At least one published (non-draft) post is needed: a static export can't build `/posts/[slug]` with no
posts.

Set `SITE_URL` to the site's public address, e.g. `https://example.com`. It is used for absolute links in
the RSS feed, the sitemap and page metadata. Without it the build falls back to Vercel's production domain
(`VERCEL_PROJECT_PRODUCTION_URL`), then to `http://localhost:3000`.

These are generated at build time, so changing `SITE_URL` needs a rebuild:

- `/rss.xml`: RSS 2.0 feed of the published posts
- `/sitemap.xml` and `/robots.txt`

LXGW WenKai is loaded from jsDelivr (`@callmebill/lxgw-wenkai-web`) without blocking the first paint.
Its version is pinned in two places, `app/layout.tsx` and `app/globals.css`; update both together, and
check the slice file names used for punctuation in `globals.css`, which can change between versions.

## Layout

```
app/
  layout.tsx          header, footer, agent console, fonts, site metadata
  page.tsx            ~ (home)
  posts/              ~/posts and each post
  projects/           ~/projects
  not-found.tsx       404, shown as a failed cd
  rss.xml/, sitemap.ts, robots.ts
wrangler.jsonc        Cloudflare: serve out/ as static assets
  _components/        shared components
content/posts/        the posts (MDX)
lib/                  site data, posts, headings, agent replies
mdx-components.tsx    how MDX elements render (heading ids, code blocks)
```
