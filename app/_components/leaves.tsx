/**
 * The leaf shadow: a drooping branch reaching in from the top-right corner, swaying slowly. It lies on
 * the page background under the content, so cards and text cover it. The drawing is public/leaves.svg;
 * the masks, sway and dark-mode version are .leaves in globals.css.
 *
 * The outer layer is as tall as the page and clips the shadow to it: the drawing is about 1250px tall,
 * and left alone it would make a page that fits on one screen scroll on into the empty space below.
 */
export function Leaves() {
  return (
    <div aria-hidden="true" className="pointer-events-none absolute inset-0 overflow-hidden">
      <div className="leaves">
        <span />
      </div>
    </div>
  );
}
