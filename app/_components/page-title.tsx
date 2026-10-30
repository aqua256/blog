/** The large title at the top of the archive, a post and the about page. */
export function PageTitle({ children }: { children: React.ReactNode }) {
  return (
    <h1 className="text-[28px] leading-[1.3] font-semibold tracking-[-0.02em] text-balance text-ink max-[640px]:text-[24px]">
      {children}
    </h1>
  );
}
