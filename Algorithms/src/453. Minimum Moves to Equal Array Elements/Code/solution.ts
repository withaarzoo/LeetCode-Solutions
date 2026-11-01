function minMoves(nums: number[]): number {
  let mn = Math.min(...nums);
  let moves = 0;
  for (let x of nums) moves += x - mn;
  return moves;
}
