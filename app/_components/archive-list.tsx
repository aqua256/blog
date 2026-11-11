"use client";

import Link from "next/link";
import { useState } from "react";
import type { PostSummary } from "@/lib/posts";
import { FluidHover } from "./fluid-hover";

type ArchivePost = Pick<PostSummary, "slug" | "title" | "date" | "draft">;

// The width of one segment
const SEGMENT = 64;

/**
 * The archive's posts under a heading for each year, with a segmented control above them: 全部 shows every
 * year, a year shows only its own posts. The selected segment is a raised surface that slides across on the
 * moderate spring. Each post is a 36px row, title on the left and month-day on the right, and a year's rows
 * share one hover highlight. The page is built with 全部 selected, so without JavaScript every post is listed.
 */
export function ArchiveList({ years }: { years: { year: string; posts: ArchivePost[] }[] }) {
  const [selected, setSelected] = useState<string | null>(null);
  const segments = [{ label: "全部", year: null }, ...years.map(({ year }) => ({ label: year, year }))];
  const index = segments.findIndex(({ year }) => year === selected);

  return (
    <>
      <fieldset className="relative mt-6 inline-grid grid-flow-col rounded-xl bg-hover p-1" style={{ gridAutoColumns: SEGMENT }}>
        <legend className="sr-only">按年份筛选</legend>
        <span
          aria-hidden="true"
          className="absolute top-1 left-1 h-7 rounded-lg bg-surface shadow-(--raised) transition-transform duration-(--duration-moderate) ease-spring motion-reduce:transition-none"
          style={{ width: SEGMENT, transform: `translateX(${index * SEGMENT}px)` }}
        />
        {segments.map(({ label, year }) => (
          <button
            key={label}
            type="button"
            aria-pressed={year === selected}
            onClick={() => setSelected(year)}
            className="type-body relative h-7 cursor-pointer rounded-lg text-muted tabular-nums hover:text-ink aria-pressed:text-ink"
          >
            <span className="steady" data-text={label}>
              {label}
            </span>
          </button>
        ))}
      </fieldset>
      {years
        .filter(({ year }) => selected === null || year === selected)
        .map(({ year, posts }) => (
          <section key={year} aria-labelledby={`year-${year}-title`}>
            <h2 id={`year-${year}-title`} className="type-subtitle mt-6 mb-1 text-ink tabular-nums">
              {year}
            </h2>
            <FluidHover className="-mx-3">
              <ol className="row-rules">
                {posts.map((post) => (
                  <li key={post.slug} className="relative">
                    <Link href={`/posts/${post.slug}`} className="flex h-9 items-center gap-4 rounded-lg px-3">
                      <span className="type-body min-w-0 flex-1 truncate text-ink">
                        {post.title}
                        {post.draft && <span className="text-muted"> · 草稿</span>}
                      </span>
                      <time dateTime={post.date} className="type-caption text-muted tabular-nums">
                        {post.date.slice(5)}
                      </time>
                    </Link>
                  </li>
                ))}
              </ol>
            </FluidHover>
          </section>
        ))}
    </>
  );
}
