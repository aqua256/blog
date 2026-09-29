import { site } from "@/lib/site";
import { PromptLine } from "./_components/prompt-line";

/**
 * The home directory:
 *
 *   ~ $ whoami
 *      aqua256
 *      <a few lines about Aqua256>
 */
export default function HomePage() {
  return (
    <main className="px-(--gutter) pt-12 pb-18">
      <div className="page-width mx-auto">
        <section aria-label="About" className="max-w-(--col)">
          <PromptLine cwd="~" command="whoami" />
          {/* Command output is indented one gutter, like the post page */}
          <div className="ml-12 max-[560px]:ml-6.5">
            <h1 className="font-bold">{site.user}</h1>
            <p className="mt-1.5 font-[380] text-body leading-[1.7] [font-variation-settings:'MONO'_0,'CASL'_0.3]">
              {site.intro}
            </p>
          </div>
        </section>
      </div>
    </main>
  );
}
