/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var firstStableIndex = function(nums, k) {
    const n = nums.length; // Store the number of elements in the array.

    const prefixMax = new Array(n); // prefixMax[i] stores max(nums[0..i]).
    const suffixMin = new Array(n); // suffixMin[i] stores min(nums[i..n-1]).

    prefixMax[0] = nums[0]; // For index 0, the prefix contains only nums[0].

    for (let i = 1; i < n; i++) { // Build prefix maximums from left to right.
        prefixMax[i] = Math.max(prefixMax[i - 1], nums[i]); // Keep the largest value seen so far.
    }

    suffixMin[n - 1] = nums[n - 1]; // For the last index, the suffix contains only nums[n-1].

    for (let i = n - 2; i >= 0; i--) { // Build suffix minimums from right to left.
        suffixMin[i] = Math.min(suffixMin[i + 1], nums[i]); // Keep the smallest value in the suffix.
    }

    for (let i = 0; i < n; i++) { // Check every index from smallest to largest.
        const instability = prefixMax[i] - suffixMin[i]; // Calculate the score for index i.

        if (instability <= k) { // The index is stable when its score is at most k.
            return i; // This is the smallest stable index because we scan left to right.
        }
    }

    return -1; // No stable index was found.
};