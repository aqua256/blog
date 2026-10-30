/**
 * The page's two columns: the content, and a side column with the profile card. The side column wraps
 * under the content when there isn't room for both (about 900px).
 */
export function PageColumns({ children, side }: { children: React.ReactNode; side: React.ReactNode }) {
  return (
    <div className="mx-auto flex max-w-292 flex-wrap items-start gap-14 px-6 pt-14 pb-24 max-[640px]:gap-12 max-[640px]:px-4 max-[640px]:pt-8 max-[640px]:pb-16">
      <main className="min-w-0 flex-[999_1_560px]">{children}</main>
      {/* As tall as the content, so a card in it can stay in view while the content scrolls */}
      <aside aria-label="侧栏" className="flex flex-[1_1_300px] flex-col self-stretch">
        {side}
      </aside>
    </div>
  );
}
