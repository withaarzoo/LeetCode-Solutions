/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var missingMultiple = function(nums, k) {
    // Store all values so membership checks are fast.
    const present = new Set(nums);

    // Start with the smallest positive multiple of k.
    let multiple = k;

    // Keep checking consecutive multiples of k.
    while (present.has(multiple)) {
        // Move to the next positive multiple.
        multiple += k;
    }

    // Return the first multiple that is not present.
    return multiple;
};