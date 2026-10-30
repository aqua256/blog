/**
 * The leaf shadow: a drooping branch reaching in from the top-right corner, swaying slowly. It lies on
 * the page background under the content, so cards and text cover it. The drawing is public/leaves.svg;
 * the masks, sway and dark-mode version are .leaves in globals.css.
 */
export function Leaves() {
  return (
    <div aria-hidden="true" className="leaves">
      <span />
    </div>
  );
}
