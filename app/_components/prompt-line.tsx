/**
 * A shell line such as `~/posts $ cat hello-world.mdx`. With `cursor`, the line ends in a blinking block
 * cursor; with only a cwd it's the idle prompt `~ $ █`.
 */
export function PromptLine({
  cwd,
  command,
  args,
  cursor = false,
}: {
  cwd: string;
  command?: string;
  args?: React.ReactNode;
  cursor?: boolean;
}) {
  return (
    <p className="mb-3.5 font-semibold [overflow-wrap:anywhere]">
      <span className="select-none font-normal text-dim">
        <span className="text-prompt">{cwd}</span> ${" "}
      </span>
      {command && <span className="text-cmd">{command}</span>}
      {args && <span className="text-arg"> {args}</span>}
      {cursor && <span aria-hidden className="block-cursor" />}
    </p>
  );
}
