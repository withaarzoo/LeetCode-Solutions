/**
 * @param {string} s
 * @param {number} a
 * @param {number} b
 * @return {string}
 */
var findLexSmallestString = function (s, a, b) {
  const n = s.length;
  const seen = new Set();
  const q = [];
  seen.add(s);
  q.push(s);
  let ans = s;

  while (q.length) {
    const cur = q.shift();
    if (cur < ans) ans = cur;

    // add a to odd indices
    let ch = cur.split("");
    for (let i = 1; i < n; i += 2) {
      ch[i] = String((parseInt(ch[i]) + a) % 10);
    }
    const addOp = ch.join("");
    if (!seen.has(addOp)) {
      seen.add(addOp);
      q.push(addOp);
    }

    // rotate right by b
    const rotOp = cur.slice(n - b) + cur.slice(0, n - b);
    if (!seen.has(rotOp)) {
      seen.add(rotOp);
      q.push(rotOp);
    }
  }

  return ans;
};
