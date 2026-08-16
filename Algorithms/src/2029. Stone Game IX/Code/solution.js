/**
 * @param {number[]} stones
 * @return {boolean}
 */
var stoneGameIX = function (stones) {
  // cnt[r] stores how many stones have remainder r modulo 3.
  const cnt = [0, 0, 0];

  // Count the stones in each remainder group.
  for (const stone of stones) {
    // Only the remainder matters for the game.
    cnt[stone % 3]++;
  }

  // With an even number of remainder-0 stones,
  // Alice needs at least one stone from both useful groups.
  if (cnt[0] % 2 === 0) {
    return cnt[1] > 0 && cnt[2] > 0;
  }

  // With an odd number of remainder-0 stones,
  // a difference greater than 2 lets Alice force a win.
  return Math.abs(cnt[1] - cnt[2]) > 2;
};
