/** Posts sorted newest first by their YYYY-MM-DD date, as a new array. Posts on the same day keep their order. */
export function newestFirst<T extends { date: string }>(posts: readonly T[]): T[] {
  return posts.toSorted((a, b) => b.date.localeCompare(a.date));
}

/** Posts grouped by the year of their date, newest year first and newest post first within each year. */
export function groupByYear<T extends { date: string }>(posts: readonly T[]): { year: string; posts: T[] }[] {
  const groups: { year: string; posts: T[] }[] = [];
  for (const post of newestFirst(posts)) {
    const year = post.date.slice(0, 4);
    const last = groups.at(-1);
    if (last?.year === year) last.posts.push(post);
    else groups.push({ year, posts: [post] });
  }
  return groups;
}
