import { ViewTransition } from "react";

/**
 * Wraps a page's content so navigations animate (React's <ViewTransition>, see Next's view-transitions
 * guide): the old page fades out quickly, the new one prints in with a slight rise. The header is anchored
 * and post titles morph between the list and the post; the CSS is in globals.css.
 *
 * It goes in each page, not the layout: layouts persist, so enter/exit would never fire there.
 */
export function PageTransition({ children }: { children: React.ReactNode }) {
  return (
    <ViewTransition enter="page-in" exit="page-out" default="none">
      {children}
    </ViewTransition>
  );
}

/** The post title, shared between the post lists and the post's own heading so it morphs between them. */
export function PostTitleTransition({ slug, children }: { slug: string; children: React.ReactNode }) {
  return (
    <ViewTransition name={`post-title-${slug}`} share="morph" default="none">
      {children}
    </ViewTransition>
  );
}
