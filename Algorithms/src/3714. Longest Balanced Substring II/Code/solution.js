/**
 * @param {string} s
 * @return {number}
 *
 * We use Map with string keys "x#y".
 */
var longestBalanced = function (s) {
  const n = s.length;
  let a = 0,
    b = 0,
    c = 0;
  let ans = 0;

  // longest single-char run
  let run = 0,
    prev = "";
  for (let i = 0; i < n; ++i) {
    if (i === 0 || s[i] !== prev) run = 1;
    else run++;
    prev = s[i];
    ans = Math.max(ans, run);
  }

  const map3 = new Map(); // (b-a, c-a)
  const map_ab_c = new Map(); // (b-a, c)
  const map_ac_b = new Map(); // (c-a, b)
  const map_bc_a = new Map(); // (c-b, a)

  function key(x, y) {
    return x + "#" + y;
  }

  map3.set(key(0, 0), 0);
  map_ab_c.set(key(0, 0), 0);
  map_ac_b.set(key(0, 0), 0);
  map_bc_a.set(key(0, 0), 0);

  for (let p = 1; p <= n; ++p) {
    const ch = s[p - 1];
    if (ch === "a") a++;
    else if (ch === "b") b++;
    else c++;

    const k3 = key(b - a, c - a);
    if (map3.has(k3)) ans = Math.max(ans, p - map3.get(k3));
    else map3.set(k3, p);

    const kabc = key(b - a, c);
    if (map_ab_c.has(kabc)) ans = Math.max(ans, p - map_ab_c.get(kabc));
    else map_ab_c.set(kabc, p);

    const kacb = key(c - a, b);
    if (map_ac_b.has(kacb)) ans = Math.max(ans, p - map_ac_b.get(kacb));
    else map_ac_b.set(kacb, p);

    const kbc = key(c - b, a);
    if (map_bc_a.has(kbc)) ans = Math.max(ans, p - map_bc_a.get(kbc));
    else map_bc_a.set(kbc, p);
  }

  return ans;
};
