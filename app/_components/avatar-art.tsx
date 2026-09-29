import art from "@/lib/avatar-art.json";

/**
 * The avatar drawn the way terminal image viewers (chafa, timg) and fastfetch draw pictures: each
 * character cell holds two pixels, one above the other (a `▀` in the upper pixel's color on the lower
 * pixel's background). Here each cell paints its two halves as flat fills (top border, background) on a
 * whole-pixel grid, so cells meet without the hairline gaps that fractional-width glyphs leave between them.
 *
 * lib/avatar-art.json holds the pixel colors (42 × 42), made once from the avatar with sharp; to change
 * the picture, regenerate that file.
 */
export function AvatarArt({ label }: { label: string }) {
  const lines = [];
  for (let y = 0; y < art.height; y += 2) {
    const top = art.pixels[y];
    const bottom = art.pixels[y + 1];
    lines.push(
      <div key={y} className="print-line flex" style={{ "--line": y / 2 } as React.CSSProperties}>
        {top.map((color, x) => (
          <span key={x} style={{ borderTopColor: color, backgroundColor: bottom[x] }} />
        ))}
      </div>,
    );
  }

  return (
    <div role="img" aria-label={label} className="avatar-art">
      {lines}
    </div>
  );
}
