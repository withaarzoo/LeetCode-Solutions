/**
 * @param {number[]} nums
 * @return {number}
 */
var maxWidthRamp = function(nums) {
    let n = nums.length;
    let stack = [];
    
    // Step 1: Build a decreasing stack of indices
    for (let i = 0; i < n; i++) {
        if (stack.length === 0 || nums[stack[stack.length - 1]] > nums[i]) {
            stack.push(i);
        }
    }
    
    let maxWidth = 0;
    
    // Step 2: Traverse from the end and find maximum width ramp
    for (let j = n - 1; j >= 0; j--) {
        while (stack.length > 0 && nums[stack[stack.length - 1]] <= nums[j]) {
            maxWidth = Math.max(maxWidth, j - stack.pop());
        }
    }
    
    return maxWidth;
};