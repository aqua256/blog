/** A shell line such as `~/writing $ cat hello-world.mdx`. */
export function PromptLine({ cwd, command, args }: { cwd: string; command: string; args?: string }) {
  return (
    <p className="mb-3.5 font-semibold [overflow-wrap:anywhere]">
      <span className="select-none font-normal text-dim">
        <span className="text-prompt">{cwd}</span> ${" "}
      </span>
      <span className="text-cmd">{command}</span>
      {args && <span className="text-arg"> {args}</span>}
    </p>
  );
}
