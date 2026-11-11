/** The large title at the top of the archive, a post and the about page. */
export function PageTitle({ children }: { children: React.ReactNode }) {
  return (
    <h1 className="type-display text-balance text-ink">
      {children}
    </h1>
  );
}
