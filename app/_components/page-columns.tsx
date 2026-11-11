/**
 * The page's two columns: the content, and a side column with the profile card. The side column wraps
 * under the content when there isn't room for both (below 840px). Without `side`, only the content.
 */
export function PageColumns({ children, side }: { children: React.ReactNode; side?: React.ReactNode }) {
  return (
    <div className="mx-auto flex max-w-252 flex-wrap items-start gap-12 px-6 pt-10 pb-20 max-[640px]:gap-10 max-[640px]:px-4 max-[640px]:pt-6 max-[640px]:pb-14">
      <main className="min-w-0 flex-[999_1_504px]">{children}</main>
      {side && (
        // As tall as the content, so a card in it can stay in view while the content scrolls
        <aside aria-label="侧栏" className="flex flex-[1_1_240px] flex-col self-stretch">
          {side}
        </aside>
      )}
    </div>
  );
}
