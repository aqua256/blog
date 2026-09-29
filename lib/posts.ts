import fs from "node:fs/promises";
import path from "node:path";
import type { MDXContent } from "mdx/types";

/** What each post declares with `export const metadata = { ... }` at the top of its .mdx file. */
export type PostMetadata = {
  title: string;
  /** YYYY-MM-DD */
  date: string;
  tags?: string[];
  /** Drafts are only listed in development. */
  draft?: boolean;
};

export type PostSummary = PostMetadata & {
  slug: string;
  readingMinutes: number;
};

export type Post = PostSummary & {
  Content: MDXContent;
};

const POSTS_DIR = path.join(process.cwd(), "content", "posts");
// Reading speeds: English words vs. CJK characters (which have no spaces between words)
const WORDS_PER_MINUTE = 200;
const CJK_CHARS_PER_MINUTE = 350;
const CJK_CHAR = /[぀-ヿ㐀-鿿豈-﫿가-힯]/g;
const LATIN_WORD = /[A-Za-z0-9]+(?:['’][A-Za-z]+)*/g;

async function listSlugs(): Promise<string[]> {
  const files = await fs.readdir(POSTS_DIR).catch(() => [] as string[]);
  return files.filter((file) => file.endsWith(".mdx")).map((file) => file.slice(0, -".mdx".length));
}

/**
 * Rough reading time from the raw source, skipping the metadata export and import/export lines.
 * CJK characters and English words are counted separately, then added up.
 */
async function readingMinutes(slug: string): Promise<number> {
  const source = await fs.readFile(path.join(POSTS_DIR, `${slug}.mdx`), "utf8");
  const prose = source
    .replace(/^export const metadata = \{[\s\S]*?\};?\s*$/m, "")
    .replace(/^(import|export) .*$/gm, "");
  const cjkChars = prose.match(CJK_CHAR)?.length ?? 0;
  const latinWords = prose.replace(CJK_CHAR, " ").match(LATIN_WORD)?.length ?? 0;
  const minutes = cjkChars / CJK_CHARS_PER_MINUTE + latinWords / WORDS_PER_MINUTE;
  return Math.max(1, Math.round(minutes));
}

async function loadPost(slug: string): Promise<Post> {
  const mod = (await import(`@/content/posts/${slug}.mdx`)) as {
    default: MDXContent;
    metadata: PostMetadata;
  };
  return {
    ...mod.metadata,
    slug,
    readingMinutes: await readingMinutes(slug),
    Content: mod.default,
  };
}

function isVisible(post: PostMetadata) {
  return !post.draft || process.env.NODE_ENV === "development";
}

/** Visible posts, newest first. */
export async function getPosts(): Promise<PostSummary[]> {
  const posts = await Promise.all((await listSlugs()).map(loadPost));
  return posts
    .filter(isVisible)
    .sort((a, b) => b.date.localeCompare(a.date))
    .map(({ slug, title, date, tags, draft, readingMinutes }) => ({ slug, title, date, tags, draft, readingMinutes }));
}

export async function getPost(slug: string): Promise<Post | null> {
  if (!(await listSlugs()).includes(slug)) return null;
  const post = await loadPost(slug);
  return isVisible(post) ? post : null;
}
