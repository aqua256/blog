/** The large title at the top of the archive, a post and the about page. */
export function PageTitle({ children }: { children: React.ReactNode }) {
  return (
    <h1 className="text-[32px] leading-[1.3] font-semibold tracking-[-0.02em] text-balance text-ink max-[640px]:text-[26px]">
      {children}
    </h1>
  );
}
