/**
 * @param {number[]} skill
 * @param {number[]} mana
 * @return {number}
 */
var minTime = function (skill, mana) {
  const n = skill.length;
  const m = mana.length;
  if (m === 0) return 0;

  // Use BigInt inside to remain safe with large products,
  // finally convert to Number (problem constraints fit in JS Number here).
  const pref = new Array(n);
  for (let i = 0; i < n; ++i) {
    pref[i] = BigInt(skill[i]) + (i ? pref[i - 1] : 0n);
  }

  let S = 0n;
  for (let j = 1; j < m; ++j) {
    const prev = BigInt(mana[j - 1]);
    const cur = BigInt(mana[j]);
    let best = null;
    for (let i = 0; i < n; ++i) {
      const prev_pref = i ? pref[i - 1] : 0n;
      const cand = pref[i] * prev - prev_pref * cur;
      if (best === null || cand > best) best = cand;
    }
    S += best;
  }

  const ans = S + pref[n - 1] * BigInt(mana[m - 1]);
  // Convert back to Number (safe under constraints). If you prefer BigInt, return ans.
  return Number(ans);
};
