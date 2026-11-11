import { site } from "@/lib/site";

// The site is built ahead of time, so this is the year of the build
const year = new Date().getFullYear();

/** Only the copyright: the links are in the profile card on every page. */
export function SiteFooter() {
  return (
    <footer className="border-t border-line">
      <p className="type-caption mx-auto max-w-252 px-6 py-5 text-muted max-[640px]:px-4">
        © {year} {site.name}
      </p>
    </footer>
  );
}
