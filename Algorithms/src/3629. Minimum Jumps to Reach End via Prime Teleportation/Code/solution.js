/**
 * @param {number[]} nums
 * @return {number}
 */
var minJumps = function (nums) {
  const n = nums.length;

  // Already at destination
  if (n === 1) return 0;

  // Find maximum value
  const mx = Math.max(...nums);

  // Smallest prime factor array
  const spf = Array(mx + 1).fill(0);

  // Initialize SPF
  for (let i = 0; i <= mx; i++) {
    spf[i] = i;
  }

  // Sieve preprocessing
  for (let i = 2; i * i <= mx; i++) {
    if (spf[i] === i) {
      for (let j = i * i; j <= mx; j += i) {
        if (spf[j] === j) {
          spf[j] = i;
        }
      }
    }
  }

  // Prime factor -> indices
  const mp = new Map();

  for (let i = 0; i < n; i++) {
    let x = nums[i];

    const used = new Set();

    // Extract unique prime factors
    while (x > 1) {
      const p = spf[x];

      if (!used.has(p)) {
        if (!mp.has(p)) {
          mp.set(p, []);
        }

        mp.get(p).push(i);

        used.add(p);
      }

      x = Math.floor(x / p);
    }
  }

  // BFS queue
  const q = [0];

  // Distance array
  const dist = Array(n).fill(-1);

  dist[0] = 0;

  let front = 0;

  while (front < q.length) {
    const i = q[front++];

    const steps = dist[i];

    // Reached destination
    if (i === n - 1) {
      return steps;
    }

    // Move left
    if (i - 1 >= 0 && dist[i - 1] === -1) {
      dist[i - 1] = steps + 1;
      q.push(i - 1);
    }

    // Move right
    if (i + 1 < n && dist[i + 1] === -1) {
      dist[i + 1] = steps + 1;
      q.push(i + 1);
    }

    const val = nums[i];

    // Current value must be prime
    if (val > 1 && spf[val] === val) {
      const list = mp.get(val) || [];

      // Teleport moves
      for (const nxt of list) {
        if (dist[nxt] === -1) {
          dist[nxt] = steps + 1;
          q.push(nxt);
        }
      }

      // Prevent repeated processing
      mp.set(val, []);
    }
  }

  return -1;
};
