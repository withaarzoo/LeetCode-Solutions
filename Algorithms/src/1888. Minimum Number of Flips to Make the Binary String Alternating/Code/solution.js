var minFlips = function (s) {
  const n = s.length;
  const ss = s + s;

  let diff1 = 0,
    diff2 = 0;
  let ans = Infinity;

  for (let i = 0; i < ss.length; i++) {
    let expected1 = i % 2 === 0 ? "0" : "1";
    let expected2 = i % 2 === 0 ? "1" : "0";

    if (ss[i] !== expected1) diff1++;
    if (ss[i] !== expected2) diff2++;

    if (i >= n) {
      let prev = ss[i - n];

      let prevExp1 = (i - n) % 2 === 0 ? "0" : "1";
      let prevExp2 = (i - n) % 2 === 0 ? "1" : "0";

      if (prev !== prevExp1) diff1--;
      if (prev !== prevExp2) diff2--;
    }

    if (i >= n - 1) {
      ans = Math.min(ans, diff1, diff2);
    }
  }

  return ans;
};
