function minMoves(nums: number[]): number {
  let mx = Math.max(...nums);
  let ans = 0;
  for (let x of nums) ans += mx - x;
  return ans;
}
