import type { Metadata } from "next";
import Link from "next/link";
import { projects } from "@/lib/projects";
import { PromptLine } from "../_components/prompt-line";

export const metadata: Metadata = {
  title: "projects",
};

/**
 * The projects section as an `ls` of ~/projects:
 *
 *   ~/projects $ ls
 *      blog/   this site: Next.js, MDX and a scripted agent   2026   src ↗
 */
export default function ProjectsPage() {
  return (
    <main className="px-(--gutter) pt-12 pb-18">
      <div className="page-width mx-auto grid grid-cols-[minmax(0,var(--col))] gap-x-18 min-[1180px]:grid-cols-[minmax(0,var(--col))_220px]">
        <section aria-label="Projects">
          <PromptLine cwd="~/projects" command="ls" />

          {/* Command output is indented one gutter, like the post page */}
          <div className="ml-12 max-[560px]:ml-6.5">
            {projects.length === 0 ? (
              <p className="text-dim">total 0</p>
            ) : (
              <ul className="flex flex-col gap-3 max-[560px]:gap-5">
                {projects.map((project) => (
                  <li
                    key={project.name}
                    className="grid grid-cols-[auto_minmax(0,1fr)_auto] items-baseline gap-x-6 max-[560px]:grid-cols-[1fr_auto] max-[560px]:gap-y-1"
                  >
                    {project.href ? (
                      <Link href={project.href} className="font-semibold text-link underline-offset-3 hover:underline">
                        {project.name}/
                      </Link>
                    ) : (
                      <span className="font-semibold">{project.name}/</span>
                    )}
                    <span className="text-dim max-[560px]:col-span-2 max-[560px]:row-start-2">
                      {project.description}
                      {project.links?.map(({ label, href }) => (
                        <a key={href} href={href} className="ml-3 whitespace-nowrap text-link hover:underline">
                          {label} ↗
                        </a>
                      ))}
                    </span>
                    <span className="text-right text-dim tabular-nums">{project.year}</span>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </section>
      </div>
    </main>
  );
}
