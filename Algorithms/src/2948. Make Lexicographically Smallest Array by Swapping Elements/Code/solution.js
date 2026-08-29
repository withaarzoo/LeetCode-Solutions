/**
 * @param {number[]} nums
 * @param {number} limit
 * @return {number[]}
 */
var lexicographicallySmallestArray = function(nums, limit) {
    const n = nums.length; // Store the array size.
    
    // Store every value with its original index.
    const elements = nums.map((value, index) => [value, index]);
    
    // Sort by value so connected values become consecutive.
    elements.sort((a, b) => a[0] - b[0]);
    
    // Store the lexicographically smallest result.
    const answer = new Array(n);
    
    let start = 0; // First element of the current connected group.
    
    while (start < n) {
        let end = start; // Expand the current group.
        
        // Consecutive values stay in the same group when their difference
        // is at most limit, so they are connected through valid swaps.
        while (
            end + 1 < n &&
            elements[end + 1][0] - elements[end][0] <= limit
        ) {
            end++;
        }
        
        // Collect all original indices of the current group.
        const indices = [];
        
        for (let i = start; i <= end; i++) {
            indices.push(elements[i][1]);
        }
        
        // Sort positions so smaller values are placed earlier.
        indices.sort((a, b) => a - b);
        
        // Values are already sorted inside the current group.
        for (let i = 0; i < indices.length; i++) {
            answer[indices[i]] = elements[start + i][0];
        }
        
        // Move to the next connected group.
        start = end + 1;
    }
    
    return answer; // Return the lexicographically smallest array.
};