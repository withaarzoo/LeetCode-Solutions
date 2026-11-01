/**
 * @param {number[]} nums
 * @return {number}
 */
var minMoves = function(nums) {
    let mn = Math.min(...nums);
    let moves = 0;
    for (let x of nums) moves += x - mn;
    return moves;
};