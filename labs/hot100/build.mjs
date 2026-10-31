// node labs/hot100/build.mjs  ->  labs/hot100/dist/index.html
// Assembles chapters/*.json + go/**/solution.go into one self-contained page,
// and refuses to build when the content breaks an invariant.
import { readFileSync, readdirSync, existsSync, mkdirSync, writeFileSync } from "node:fs";
import { join, dirname } from "node:path";
import { fileURLToPath } from "node:url";
import { createRequire } from "node:module";

const root = dirname(fileURLToPath(import.meta.url));
const require = createRequire(import.meta.url);
const { highlightGo, inlineText, escapeHtml } = require("./core.js");

const errors = [];
const fail = (msg) => errors.push(msg);

// ---------- expected assignment ----------
const plan = readFileSync(join(root, "CHAPTERS.md"), "utf8")
  .split("\n")
  .filter((l) => /^\d\d /.test(l))
  .map((l) => {
    const [id, slug, title, day, , ...ids] = l.trim().split(/\s+/);
    return { id, slug, title, day: Number(day.slice(1)), ids: ids.map(Number) };
  });

// ---------- load ----------
function displayCode(file) {
  const src = readFileSync(file, "utf8").replace(/\r\n/g, "\n");
  const lines = src.split("\n");
  const out = [];
  let inImport = false;
  for (const line of lines) {
    if (/^package\s/.test(line)) continue;
    if (/^import\s*\($/.test(line)) { inImport = true; continue; }
    if (inImport) { if (/^\)$/.test(line)) inImport = false; continue; }
    if (/^import\s+"/.test(line)) continue;
    out.push(line);
  }
  return out.join("\n").replace(/^\n+/, "").replace(/\s+$/, "") + "\n";
}

const goDir = join(root, "go");
const chapters = [];
const seen = new Set();
for (const p of plan) {
  const file = join(root, "chapters", p.id + ".json");
  if (!existsSync(file)) { fail(`missing chapters/${p.id}.json (${p.title})`); continue; }
  let ch;
  try { ch = JSON.parse(readFileSync(file, "utf8")); } catch (e) { fail(`${file}: ${e.message}`); continue; }
  for (const k of ["id", "slug", "title", "day"]) {
    if (String(ch[k]) !== String(p[k])) fail(`${p.id}: ${k} is ${ch[k]}, expected ${p[k]}`);
  }
  if (!ch.when) fail(`${p.id}: missing "when"`);
  const got = (ch.problems || []).map((x) => x.id);
  const missing = p.ids.filter((i) => !got.includes(i));
  const extra = got.filter((i) => !p.ids.includes(i));
  if (missing.length) fail(`${p.id}: missing problems ${missing}`);
  if (extra.length) fail(`${p.id}: unexpected problems ${extra}`);
  for (const pr of ch.problems || []) {
    if (seen.has(pr.id)) fail(`duplicate problem ${pr.id}`);
    seen.add(pr.id);
    if (!["简单", "中等", "困难"].includes(pr.difficulty)) fail(`${pr.id}: difficulty ${pr.difficulty}`);
    for (const k of ["title", "insight", "pitfall", "time", "space"]) if (!pr[k]) fail(`${pr.id}: missing ${k}`);
    const dir = join(goDir, "p" + String(pr.id).padStart(4, "0"));
    if (!existsSync(join(dir, "solution.go"))) { fail(`${pr.id}: missing ${dir}/solution.go`); continue; }
    if (!existsSync(join(dir, "solution_test.go"))) fail(`${pr.id}: missing solution_test.go`);
    pr.code = displayCode(join(dir, "solution.go"));
  }
  if (ch.template) {
    const f = join(goDir, ch.template.dir, "template.go");
    if (!existsSync(f)) fail(`${p.id}: missing ${f}`);
    else ch.template.code = displayCode(f);
  }
  chapters.push(ch);
}
const total = chapters.reduce((n, c) => n + (c.problems || []).length, 0);
if (total !== 100) fail(`expected 100 problems, got ${total}`);

if (errors.length && !process.argv.includes("--partial")) {
  console.error("build refused:\n  " + errors.join("\n  "));
  process.exit(1);
}
if (errors.length) console.warn(`building partial page (${errors.length} problems):\n  ` + errors.slice(0, 20).join("\n  "));

// ---------- render ----------
const DIFF = { 简单: "d-easy", 中等: "d-mid", 困难: "d-hard" };
const lc = (id) => `https://leetcode.cn/problemset/?search=${id}`;
let seq = 0;

function codeBlock(id, code) {
  return `<pre id="${id}"><code>${highlightGo(code)}</code></pre>`;
}

function renderProblem(pr) {
  seq++;
  const lines = pr.code.trimEnd().split("\n").length;
  const codeId = `code-${pr.id}`;
  return `
<article class="card" id="p${pr.id}" data-id="${pr.id}" data-title="${escapeHtml(pr.title)}" data-hot="${pr.hot ? 1 : 0}">
  <div class="p-head">
    <a class="p-id" href="${lc(pr.id)}" title="在力扣打开">${pr.id}</a>
    <h3 class="p-title">${escapeHtml(pr.title)}</h3>
    <span class="tag ${DIFF[pr.difficulty]}">${pr.difficulty}</span>
    ${pr.hot ? `<span class="tag hot">高频</span>` : ""}
    <span class="seq">第 ${seq} 题</span>
  </div>
  <div class="acts">
    <button class="act m" aria-pressed="false">✓ 掌握</button>
    <button class="act r" aria-pressed="false">↻ 要复习</button>
  </div>
  <p class="insight">${inlineText(pr.insight)}</p>
  <div class="trap"><span class="tag hot">易错</span><span>${inlineText(pr.pitfall)}</span></div>
  <div class="cx"><b>复杂度</b>　时间 ${escapeHtml(pr.time)}，空间 ${escapeHtml(pr.space)}</div>
  <details>
    <summary>Go 代码 <span class="ln">${lines} 行</span><button class="copy" data-for="${codeId}">复制</button></summary>
    ${codeBlock(codeId, pr.code)}
  </details>
</article>`;
}

function renderChapter(ch) {
  const n = ch.problems.length;
  const hot = ch.problems.filter((p) => p.hot).length;
  const tpl = ch.template
    ? `
  <div class="tpl">
    <div class="tpl-head"><span class="lab">模板</span>${escapeHtml(ch.template.title)}<button class="copy" data-for="tpl-${ch.id}">复制</button></div>
    <div class="lang">GO</div>
    ${codeBlock("tpl-" + ch.id, ch.template.code)}
  </div>`
    : "";
  return `
<section class="chapter" id="c${ch.id}" data-day="${ch.day}">
  <div class="ch-head">
    <div class="ch-no">${ch.id}</div>
    <div>
      <h2 class="ch-title">${escapeHtml(ch.title)}</h2>
      <div class="ch-meta"><span class="tag-day">Day ${ch.day}</span><span>${n} 题</span><span title="高频题数">★ ${hot}</span><span>已掌握 <span class="ch-done">0</span>/${n}</span></div>
    </div>
  </div>
  <p class="when">${inlineText(ch.when)}</p>
  ${ch.pitfalls && ch.pitfalls.length ? `<ul class="pitfalls">${ch.pitfalls.map((x) => `<li>${inlineText(x)}</li>`).join("")}</ul>` : ""}
  ${tpl}
  <div class="cards">${ch.problems.map(renderProblem).join("")}</div>
</section>`;
}

const ready = chapters.filter((c) => c.problems && c.problems.every((p) => p.code));
const days = [...new Set(ready.map((c) => c.day))].sort((a, b) => a - b);
let strip = "";
let lastDay = 0;
for (const c of ready) {
  if (c.day !== lastDay) { strip += `<span class="day">D${c.day}</span>`; lastDay = c.day; }
  strip += `<a href="#c${c.id}"><span class="no">${c.id}</span>${escapeHtml(c.title)} <span class="pg">0/${c.problems.length}</span></a>`;
}

const tests = readdirSync(goDir).filter((d) => existsSync(join(goDir, d, "solution_test.go")) || existsSync(join(goDir, d, "template_test.go"))).length;
const html = readFileSync(join(root, "page.html"), "utf8")
  .replace("/*__CORE__*/", () => readFileSync(join(root, "core.js"), "utf8"))
  .replace("<!--DAY_OPTIONS-->", days.map((d) => `<option value="${d}">Day ${d}</option>`).join(""))
  .replace("<!--STRIP-->", () => strip)
  .replace("<!--INTRO-->", `力扣热题 100，按套路分成 ${ready.length} 章、${days.length} 天刷完。每章先读一句「什么时候用」和模板，再逐题过思路、易错点和代码；做完点「掌握」，没把握点「要复习」，之后用顶部筛选只刷要复习的题。`)
  .replace("<!--CHAPTERS-->", () => ready.map(renderChapter).join("\n"))
  .replace("<!--FOOT-->", `全部 Go 代码都在真实的 Go 包里编译并通过测试（${tests} 个测试包，含与暴力解的随机对拍）。进度登录后跨设备同步，只有你自己看得到。`);

mkdirSync(join(root, "dist"), { recursive: true });
writeFileSync(join(root, "dist", "index.html"), html);
console.log(`built dist/index.html: ${ready.length} chapters, ${seq} problems, ${(html.length / 1024).toFixed(0)} KiB`);
