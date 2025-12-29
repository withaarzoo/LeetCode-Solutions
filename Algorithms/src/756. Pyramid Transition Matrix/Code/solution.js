var pyramidTransition = function (bottom, allowed) {
  const rules = {};
  for (const s of allowed) {
    const key = s[0] + s[1];
    if (!rules[key]) rules[key] = new Set();
    rules[key].add(s[2]);
  }

  const bad = new Set();

  function dfs(row, idx, next) {
    if (row.length === 1) return true;

    if (idx === row.length - 1) {
      if (bad.has(next)) return false;
      const ok = dfs(next, 0, "");
      if (!ok) bad.add(next);
      return ok;
    }

    const key = row.substring(idx, idx + 2);
    if (!rules[key]) return false;

    for (const c of rules[key]) {
      if (dfs(row, idx + 1, next + c)) return true;
    }
    return false;
  }

  return dfs(bottom, 0, "");
};
