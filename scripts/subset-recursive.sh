#!/usr/bin/env bash
# Rebuilds app/_fonts/recursive.woff2: Recursive's Latin subset from Google Fonts, with each axis
# narrowed to the range the site uses, about half the size of the full font.
#
#   weight 380–800 · slant −8–0 · MONO 0–1 · CASL 0–1
#
# Widen a range here before using a value outside it in globals.css or a font-* class.
# Needs fonttools with brotli: pip install fonttools brotli
set -euo pipefail

out="$(dirname "$0")/../app/_fonts/recursive.woff2"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

# A modern browser's user agent gets woff2 files split by script; take the "latin" one
ua="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36"
css="$(curl -fsS -A "$ua" "https://fonts.googleapis.com/css2?family=Recursive:slnt,wght,CASL,MONO@-15..0,300..1000,0..1,0..1")"
url="$(awk '/\/\* latin \*\//{found=1} found && /src:/{print; exit}' <<<"$css" | grep -o 'https://[^)]*')"
curl -fsS "$url" -o "$tmp/full.woff2"

fonttools varLib.instancer "$tmp/full.woff2" wght=380:800 slnt=-8:0 -o "$tmp/narrow.ttf" -q
pyftsubset "$tmp/narrow.ttf" --unicodes='*' --layout-features='*' --flavor=woff2 --output-file="$out"
echo "$out: $(wc -c <"$out" | tr -d ' ') bytes"
