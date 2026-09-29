import type { Heading } from "./headings";
import { site } from "./site";

/**
 * Scripted replies for the agent console demo: a question is matched by keywords (English or Chinese)
 * to one topic, and each topic has a fixed answer. No model is involved yet.
 */

/** A piece of an answer: plain text, inline code, or a link. */
export type Segment = string | { code: string } | { href: string; text: string };

/** What the agent knows about the posts, passed down from the layout. */
export type AgentPost = { slug: string; title: string; headings: Heading[] };

type Context = {
  /** Newest first */
  posts: AgentPost[];
  /** The post being read, if any */
  current?: AgentPost;
};

type Topic = "help" | "summary" | "building" | "writing" | "contact" | "about" | "fallback";

/** Checked in order; the first match wins. */
const TOPICS: [Topic, RegExp][] = [
  ["help", /^(help|\?|man|帮助)$/],
  ["summary", /summar|tl;?dr|this post|this article|recap|总结|这篇|概括/],
  ["building", /build|project|working on|making|work on|portfolio|项目|作品|在做/],
  ["writing", /writ|post|article|blog|read next|what else|文章|博客/],
  ["contact", /contact|touch|reach|email|linkedin|github|hire|联系|邮箱/],
  ["about", /who|about|introduce|yourself|aqua256|关于|介绍|你是谁|是谁/],
];

function topicOf(question: string): Topic {
  const text = question.toLowerCase();
  return TOPICS.find(([, pattern]) => pattern.test(text))?.[0] ?? "fallback";
}

const lines = (items: Segment[][]): Segment[] => items.flatMap((item, i) => (i === 0 ? item : ["\n", ...item]));

function summary({ current }: Context): Segment[] {
  if (!current) return ["Open a post first, then ask me to summarize it."];
  if (current.headings.length === 0) return [`“${current.title}” is short enough that it has no sections to list.`];
  return lines([
    [`“${current.title}” has ${current.headings.length} sections:`],
    ...current.headings.map(({ id, text }) => [{ href: `#${id}`, text }]),
  ]);
}

function writing({ posts, current }: Context): Segment[] {
  if (posts.length === 0) return ["Nothing in ", { code: "~/posts" }, " yet."];
  return lines([
    ["In ", { code: "~/posts" }, ", newest first:"],
    ...posts.map((post): Segment[] =>
      post.slug === current?.slug
        ? [{ href: "#top", text: post.title }, " (this one)"]
        : [{ href: `/posts/${post.slug}`, text: post.title }],
    ),
  ]);
}

function contact(): Segment[] {
  return lines([["The best places to reach Aqua256:"], ...site.social.map(({ label, href }) => [{ href, text: label }])]);
}

const REPLIES: Record<Topic, (context: Context) => Segment[]> = {
  // TODO: Aqua256's own introduction
  about: () => ["Aqua256 hasn't written this part yet. For now, the posts in ", { code: "~/posts" }, " say more than I can."],
  building: () => [
    "Right now Aqua256 is building this blog, one small step at a time. Other work will live in ",
    { code: "~/projects" },
    ".",
  ],
  summary,
  writing,
  contact,
  help: () =>
    lines([
      ["Things I can answer:"],
      [{ code: 'ask "who is Aqua256?"' }],
      [{ code: 'ask "what are you building?"' }],
      [{ code: 'ask "summarize this post"' }],
      [{ code: 'ask "recent writing"' }],
      [{ code: 'ask "how do I get in touch?"' }],
      ["Type ", { code: "clear" }, " to wipe the log, or ", { code: "exit" }, " to close."],
    ]),
  fallback: () => [
    "I only know a few things so far. Try asking who Aqua256 is, what Aqua256 is building, for a summary of this post, for recent writing, or how to get in touch. 你也可以用中文问：介绍、项目、总结、文章、联系。",
  ],
};

export function reply(question: string, context: Context): Segment[] {
  return REPLIES[topicOf(question)](context);
}

/** `$ ask "who is Aqua256?"` → `who is Aqua256?`, so visitors can type the command the chips show. */
export function cleanQuestion(raw: string): string {
  return raw
    .trim()
    .replace(/^(\$\s*)?(ask|aqua256-agent)\s+/i, "")
    .replace(/^["'“”‘’]+|["'“”‘’]+$/g, "")
    .trim();
}
