// Pure helpers shared by the page and the node tests. No DOM here.

// Progress state: { [problemId]: { s: "m" | "r" | "", t: epochMillis } }
// "m" mastered, "r" needs review, "" cleared. Merging two copies keeps, per
// problem, the newer mark; a tie breaks on the status string so that merge is
// commutative and every device converges on the same state.
function pick(a, b) {
  if (!a) return b;
  if (!b) return a;
  if (a.t !== b.t) return a.t > b.t ? a : b;
  return a.s >= b.s ? a : b;
}

function validEntry(e) {
  return (
    e && typeof e === "object" &&
    (e.s === "m" || e.s === "r" || e.s === "") &&
    typeof e.t === "number" && Number.isFinite(e.t)
  );
}

function sanitize(items) {
  const out = {};
  if (!items || typeof items !== "object") return out;
  for (const k of Object.keys(items)) {
    if (/^\d{1,4}$/.test(k) && validEntry(items[k])) out[k] = { s: items[k].s, t: items[k].t };
  }
  return out;
}

function merge(a, b) {
  const out = {};
  const keys = new Set([...Object.keys(a || {}), ...Object.keys(b || {})]);
  for (const k of keys) out[k] = pick(a && a[k], b && b[k]);
  return out;
}

function sameState(a, b) {
  const ka = Object.keys(a), kb = Object.keys(b);
  if (ka.length !== kb.length) return false;
  for (const k of ka) {
    if (!b[k] || b[k].s !== a[k].s || b[k].t !== a[k].t) return false;
  }
  return true;
}

function mark(items, id, status, now) {
  const cur = items[id] ? items[id].s : "";
  const next = cur === status ? "" : status; // pressing the active button clears it
  // Clock skew between devices could make a fresh click look older than the
  // stored mark; never move a problem's timestamp backwards.
  const t = Math.max(now, items[id] ? items[id].t + 1 : 0);
  return { ...items, [id]: { s: next, t } };
}

function statusOf(items, id) {
  return (items[id] && items[id].s) || "";
}

// ---------- Go syntax highlighting ----------

const GO_KEYWORDS = new Set(
  "break case chan const continue default defer else fallthrough for func go goto if import interface map package range return select struct switch type var".split(" ")
);
const GO_TYPES = new Set(
  "bool byte complex64 complex128 error float32 float64 int int8 int16 int32 int64 rune string uint uint8 uint16 uint32 uint64 uintptr any".split(" ")
);
const GO_BUILTINS = new Set(
  "append cap clear close complex copy delete imag len make max min new panic print println real recover".split(" ")
);
const GO_CONSTS = new Set(["true", "false", "nil", "iota"]);

function escapeHtml(s) {
  return s.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");
}

const TOKEN = /(\/\/[^\n]*|\/\*[\s\S]*?\*\/)|("(?:\\.|[^"\\\n])*"|`[^`]*`|'(?:\\.|[^'\\\n])+')|(\b(?:0[xX][0-9a-fA-F_]+|\d[\d_]*(?:\.\d+)?(?:[eE][+-]?\d+)?)\b)|([A-Za-z_][A-Za-z0-9_]*)/g;

function highlightGo(src) {
  let out = "";
  let last = 0;
  TOKEN.lastIndex = 0;
  let m;
  while ((m = TOKEN.exec(src))) {
    out += escapeHtml(src.slice(last, m.index));
    const [tok, comment, str, num, ident] = m;
    let cls = "";
    if (comment) cls = "c";
    else if (str) cls = "s";
    else if (num) cls = "n";
    else if (ident) {
      if (GO_KEYWORDS.has(ident)) cls = "k";
      else if (GO_TYPES.has(ident)) cls = "t";
      else if (GO_CONSTS.has(ident)) cls = "n";
      else if (GO_BUILTINS.has(ident) && src[TOKEN.lastIndex] === "(") cls = "b";
      else if (src[TOKEN.lastIndex] === "(") cls = "f";
    }
    out += cls ? `<span class="${cls}">${escapeHtml(tok)}</span>` : escapeHtml(tok);
    last = TOKEN.lastIndex;
  }
  return out + escapeHtml(src.slice(last));
}

// Inline markup in explanations: 「term」 gets emphasis, `code` becomes code.
function inlineText(s) {
  return escapeHtml(s)
    .replace(/`([^`]+)`/g, "<code>$1</code>")
    .replace(/「([^」]+)」/g, "<em>「$1」</em>");
}

if (typeof module !== "undefined") {
  module.exports = { merge, sanitize, sameState, mark, statusOf, highlightGo, inlineText, escapeHtml };
}
