/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var targetIndices = function(nums, target) {
    let less = 0, equal = 0;
    for (const x of nums) {
        if (x < target) less++;
        else if (x === target) equal++;
    }
    const ans = [];
    for (let i = 0; i < equal; i++) {
        ans.push(less + i);
    }
    return ans;
};