/**
 * @param {number} n
 * @return {number}
 */
var nextBeautifulNumber = function (n) {
  const isBalanced = (x) => {
    let cnt = new Array(10).fill(0);
    let t = x;
    while (t > 0) {
      cnt[t % 10]++;
      t = Math.floor(t / 10);
    }
    if (cnt[0] > 0) return false; // 0 cannot appear
    for (let d = 1; d <= 9; ++d) {
      if (cnt[d] !== 0 && cnt[d] !== d) return false;
    }
    return true;
  };

  let x = n + 1;
  while (true) {
    if (isBalanced(x)) return x;
    x++;
  }
};
