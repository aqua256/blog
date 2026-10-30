import fs from "node:fs/promises";
import path from "node:path";
import { ImageResponse } from "next/og";
import { getPost, getPosts } from "@/lib/posts";
import { site } from "@/lib/site";

/**
 * The card shown when a post is shared, in the site's light theme: the title, the date and tags, and the
 * author's name, with the leaf shadow reaching in from the top-right corner.
 *
 *                                         (leaves)
 *   TypeScript 的 satisfies
 *   到底解决了什么
 *   2026-10-14   #typescript
 *
 *   Aqua256
 */

export const alt = "文章标题，右上角是树影";
export const size = { width: 1200, height: 630 };
export const contentType = "image/png";

export async function generateStaticParams() {
  const posts = await getPosts();
  return posts.map(({ slug }) => ({ slug }));
}

// The light theme's colors, as in globals.css
const colors = { bg: "#ffffff", ink: "#18212c", ink3: "#5d6b7b" };

/**
 * A font from Google Fonts cut down to just `text`, the way next/font fetches fonts at build time.
 * With `text=`, Google serves a small TrueType file, which is what ImageResponse can read.
 */
async function googleFont(family: string, text: string): Promise<ArrayBuffer> {
  const css = await fetch(
    `https://fonts.googleapis.com/css2?family=${family}&text=${encodeURIComponent(text)}`,
  ).then((res) => res.text());
  const url = css.match(/src: url\((.+?)\) format\('(?:truetype|opentype)'\)/)?.[1];
  if (!url) throw new Error(`No TrueType font for ${family}`);
  return fetch(url).then((res) => res.arrayBuffer());
}

export default async function Image({ params }: { params: Promise<{ slug: string }> }) {
  const { slug } = await params;
  const post = (await getPost(slug))!;
  const tags = (post.tags ?? []).map((tag) => `#${tag}`).join("  ");
  // Long titles step down a size so three lines still fit
  const titleSize = post.title.length > 16 ? 56 : 68;

  // The site itself uses the reader's system fonts, which an image can't; Noto Sans SC is closest to them
  const [bold, regular, leaves] = await Promise.all([
    googleFont("Noto+Sans+SC:wght@600", post.title + site.name),
    googleFont("Noto+Sans+SC:wght@400", post.date + tags),
    fs.readFile(path.join(process.cwd(), "public", "leaves.svg"), "base64"),
  ]);

  return new ImageResponse(
    <div
      style={{
        width: "100%",
        height: "100%",
        display: "flex",
        flexDirection: "column",
        padding: "72px 88px",
        background: colors.bg,
        color: colors.ink,
        fontFamily: "Noto Sans SC",
      }}
    >
      {/* The leaf shadow as on the site: faint, fading out downwards. ImageResponse draws plain <img> only. */}
      {/* oxlint-disable-next-line nextjs/no-img-element -- next/image does not exist inside ImageResponse */}
      <img
        alt=""
        src={`data:image/svg+xml;base64,${leaves}`}
        width={720}
        height={750}
        style={{
          position: "absolute",
          top: 0,
          right: 0,
          opacity: 0.14,
          maskImage: "linear-gradient(to bottom, black 45%, transparent 85%)",
        }}
      />

      <div style={{ display: "flex", flex: 1, flexDirection: "column", justifyContent: "center", maxWidth: 820 }}>
        <div style={{ fontSize: titleSize, fontWeight: 600, lineHeight: 1.3, letterSpacing: -1 }}>{post.title}</div>
        <div style={{ display: "flex", gap: 32, marginTop: 28, fontSize: 28, color: colors.ink3 }}>
          <span>{post.date}</span>
          {tags && <span>{tags}</span>}
        </div>
      </div>

      <div style={{ display: "flex", fontSize: 30, fontWeight: 600 }}>{site.name}</div>
    </div>,
    {
      ...size,
      fonts: [
        { name: "Noto Sans SC", data: bold, weight: 600 },
        { name: "Noto Sans SC", data: regular, weight: 400 },
      ],
    },
  );
}
