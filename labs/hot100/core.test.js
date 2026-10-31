// node --test labs/hot100/core.test.js
const test = require("node:test");
const assert = require("node:assert/strict");
const { merge, sanitize, sameState, mark, highlightGo, inlineText } = require("./core.js");

// Small seeded PRNG so failures reproduce.
function rng(seed) {
  let x = seed >>> 0 || 1;
  return () => ((x ^= x << 13), (x ^= x >>> 17), (x ^= x << 5), (x >>> 0) / 2 ** 32);
}
const R = rng(20261031);
const int = (n) => Math.floor(R() * n);

function randState() {
  const s = {};
  const n = int(6);
  for (let i = 0; i < n; i++) {
    s[String(int(5) + 1)] = { s: ["m", "r", ""][int(3)], t: int(4) };
  }
  return s;
}

const N = 2000;

test("merge is commutative, associative, idempotent, with {} as identity", () => {
  for (let i = 0; i < N; i++) {
    const a = randState(), b = randState(), c = randState();
    assert.ok(sameState(merge(a, b), merge(b, a)), "commutative");
    assert.ok(sameState(merge(merge(a, b), c), merge(a, merge(b, c))), "associative");
    assert.ok(sameState(merge(a, a), a), "idempotent");
    assert.ok(sameState(merge(a, {}), a), "identity");
  }
});

test("merge never loses a newer mark", () => {
  for (let i = 0; i < N; i++) {
    const a = randState(), b = randState();
    const m = merge(a, b);
    for (const k of Object.keys(m)) {
      const ts = [a[k], b[k]].filter(Boolean).map((e) => e.t);
      assert.equal(m[k].t, Math.max(...ts));
    }
  }
});

test("devices converge whatever order they exchange state in", () => {
  for (let i = 0; i < 500; i++) {
    const devices = [randState(), randState(), randState()];
    // random gossip, then one full round
    for (let j = 0; j < 6; j++) {
      const x = int(3), y = int(3);
      devices[x] = merge(devices[x], devices[y]);
    }
    const all = devices.reduce(merge, {});
    for (const d of devices) assert.ok(sameState(merge(d, all), all));
  }
});

test("mark toggles and its timestamp only moves forward", () => {
  for (let i = 0; i < N; i++) {
    const st = randState();
    const id = String(int(5) + 1);
    const status = ["m", "r"][int(2)];
    const now = int(6);
    const next = mark(st, id, status, now);
    const before = st[id] ? st[id].s : "";
    assert.equal(next[id].s, before === status ? "" : status);
    if (st[id]) assert.ok(next[id].t > st[id].t, "a fresh click wins over the stored mark");
    // The fresh click survives merging with the old state from any device.
    assert.equal(merge(st, next)[id].s, next[id].s);
    assert.notEqual(next, st, "returns a new object");
  }
});

test("sanitize drops anything that is not a valid entry", () => {
  const dirty = {
    "42": { s: "m", t: 5 },
    "43": { s: "x", t: 5 },
    "44": { s: "r", t: "5" },
    "__proto__": { s: "m", t: 1 },
    "abc": { s: "m", t: 1 },
    "45": null,
    "46": { s: "", t: Infinity },
  };
  assert.deepEqual(sanitize(dirty), { "42": { s: "m", t: 5 } });
  assert.deepEqual(sanitize(null), {});
  assert.deepEqual(sanitize("nope"), {});
});

function unhighlight(html) {
  return html
    .replace(/<[^>]+>/g, "")
    .replace(/&lt;/g, "<").replace(/&gt;/g, ">").replace(/&quot;/g, '"').replace(/&amp;/g, "&");
}

test("highlightGo round-trips: stripping the markup gives back the source", () => {
  const alphabet = ['a', 'z', 'f', 'u', 'n', 'c', ' ', '\n', '(', ')', '"', '`', "'", '/', '*', '<', '>', '&', '0', '9', '.', 'x', '\\', '_', '{', '}', 'func ', 'return ', 'len(', 'int', '// c\n', '"<b>"'];
  for (let i = 0; i < N; i++) {
    let src = "";
    const n = int(40);
    for (let j = 0; j < n; j++) src += alphabet[int(alphabet.length)];
    const html = highlightGo(src);
    assert.equal(unhighlight(html), src);
    assert.ok(!/<(?!\/?span)/.test(html), "only span tags are emitted: " + JSON.stringify(src));
  }
});

test("highlightGo classifies the basics", () => {
  const h = highlightGo('func trap(h []int) int { return len(h) } // x');
  assert.match(h, /<span class="k">func<\/span>/);
  assert.match(h, /<span class="f">trap<\/span>/);
  assert.match(h, /<span class="t">int<\/span>/);
  assert.match(h, /<span class="b">len<\/span>/);
  assert.match(h, /<span class="c">\/\/ x<\/span>/);
});

test("inlineText escapes before adding markup", () => {
  assert.equal(inlineText("a<b 「c」 `d&e`"), "a&lt;b <em>「c」</em> <code>d&amp;e</code>");
});
