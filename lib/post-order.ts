/** Posts sorted newest first by their YYYY-MM-DD date, as a new array. Posts on the same day keep their order. */
export function newestFirst<T extends { date: string }>(posts: readonly T[]): T[] {
  return posts.toSorted((a, b) => b.date.localeCompare(a.date));
}
