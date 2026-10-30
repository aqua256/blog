import { site } from "@/lib/site";

// The site is built ahead of time, so this is the year of the build
const year = new Date().getFullYear();

/** Only the copyright: the links are in the profile card on every page. */
export function SiteFooter() {
  return (
    <footer className="border-t border-line">
      <p className="mx-auto max-w-292 px-6 py-6 text-[14px] leading-normal text-ink-3 max-[640px]:px-4">
        © {year} {site.name}
      </p>
    </footer>
  );
}
