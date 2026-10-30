import { configDefaults, defineConfig } from "vitest/config";

export default defineConfig({
  // Resolve the `@/` alias from tsconfig.json, as Next.js does
  resolve: { tsconfigPaths: true },
  test: {
    // Claude Code's git worktrees hold full copies of the repository, tests included
    exclude: [...configDefaults.exclude, ".claude/**"],
  },
});
