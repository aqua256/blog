/**
 * Whether a nav link's section is the one being shown: the home link only on the home page, any other
 * link on its own page and everything under it (the archive link stays current on each post).
 */
export function isCurrentSection(pathname: string, href: string): boolean {
  return href === "/" ? pathname === "/" : pathname === href || pathname.startsWith(`${href}/`);
}
