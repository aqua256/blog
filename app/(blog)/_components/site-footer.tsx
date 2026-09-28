import { site } from "@/lib/site";

export function SiteFooter() {
  return (
    <footer aria-label="Elsewhere" className="px-(--gutter) pb-18">
      <div className="page-width mx-auto flex flex-wrap gap-x-4 gap-y-1.5 border-t border-faint pt-5 text-dim max-[560px]:gap-x-3">
        {site.social.map(({ label, href }) => (
          <a key={label} href={href} className="hover:text-ink">
            {label}
          </a>
        ))}
      </div>
    </footer>
  );
}
