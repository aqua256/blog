import { ImageResponse } from "next/og";
import { getPost, getPosts } from "@/lib/posts";
import { site } from "@/lib/site";

/**
 * The card shown when a post is shared: the post page itself, as a dark (Flexoki) terminal.
 *
 *   ~/posts $ cat hello-world.mdx
 *   # Hello, world
 *   2026-09-29 · #meta #blog
 *                                   aqua256@blog
 */

export const alt = "Post title on a dark terminal";
export const size = { width: 1200, height: 630 };
export const contentType = "image/png";

export async function generateStaticParams() {
  const posts = await getPosts();
  return posts.map(({ slug }) => ({ slug }));
}

// Flexoki dark, as in globals.css
const dark = {
  bg: "#100f0f",
  text: "#cecdc3",
  dim: "#878580",
  faint: "#343331",
  purple: "#8b7ec8",
  green: "#879a39",
  orange: "#da702c",
};

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
  const prompt = `~/posts $ cat ${slug}.mdx`;
  const tags = (post.tags ?? []).map((tag) => `#${tag}`).join(" ");
  const meta = tags ? `${post.date} · ${tags}` : post.date;
  // Long titles step down a size so three lines still fit
  const titleSize = post.title.length > 16 ? 68 : 88;
  const signature = `${site.user}@${site.host}`;
  const all = prompt + post.title + meta + signature + "#";

  // Latin in Recursive (mono, casual), as on the site; Chinese in LXGW WenKai. Google only has the
  // TC cut of WenKai, which also covers simplified characters.
  const [recursive, recursiveBold, wenkai] = await Promise.all([
    googleFont("Recursive:slnt,wght,CASL,CRSV,MONO@0,400,1,0.5,1", all),
    googleFont("Recursive:slnt,wght,CASL,CRSV,MONO@0,800,1,0.5,0.5", all),
    googleFont("LXGW+WenKai+TC:wght@700", post.title),
  ]);

  return new ImageResponse(
    <div
      style={{
        width: "100%",
        height: "100%",
        display: "flex",
        flexDirection: "column",
        padding: "72px 88px",
        background: dark.bg,
        color: dark.text,
        fontFamily: "Recursive, WenKai",
        fontSize: 32,
      }}
    >
      <div style={{ display: "flex", gap: 16 }}>
        <span style={{ color: dark.purple }}>~/posts</span>
        <span style={{ color: dark.dim }}>$</span>
        <span style={{ color: dark.green, fontWeight: 800 }}>cat</span>
        <span style={{ color: dark.orange }}>{`${slug}.mdx`}</span>
      </div>

      <div style={{ display: "flex", flex: 1, flexDirection: "column", justifyContent: "center" }}>
        {/* The # sits on the first line, sharing its line height */}
        <div style={{ display: "flex", gap: 32, fontWeight: 800 }}>
          <span style={{ color: dark.faint, fontSize: titleSize * 0.72, lineHeight: 1.2 / 0.72 }}>#</span>
          <span style={{ fontSize: titleSize, lineHeight: 1.2, letterSpacing: -1 }}>{post.title}</span>
        </div>
        <div style={{ display: "flex", marginTop: 32, marginLeft: 70, color: dark.dim }}>{meta}</div>
      </div>

      <div style={{ display: "flex", justifyContent: "flex-end", color: dark.purple, fontWeight: 800 }}>
        {signature}
      </div>
    </div>,
    {
      ...size,
      fonts: [
        { name: "Recursive", data: recursive, weight: 400 },
        { name: "Recursive", data: recursiveBold, weight: 800 },
        { name: "WenKai", data: wenkai, weight: 800 },
      ],
    },
  );
}
