/** A shell line such as `~/posts $ cat hello-world.mdx`. The idle prompt you can type into is PromptInput. */
export function PromptLine({
  cwd,
  command,
  args,
}: {
  cwd: string;
  command?: string;
  args?: React.ReactNode;
}) {
  return (
    <p className="mb-3.5 font-semibold [overflow-wrap:anywhere]">
      <span className="select-none font-normal text-dim">
        <span className="text-prompt">{cwd}</span> ${" "}
      </span>
      {command && <span className="text-cmd">{command}</span>}
      {args && <span className="text-arg"> {args}</span>}
    </p>
  );
}
