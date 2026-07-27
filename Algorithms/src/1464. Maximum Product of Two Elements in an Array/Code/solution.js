/**
 * @param {number[]} nums
 * @return {number}
 */
var maxProduct = function(nums) {

    // Store the largest value found so far
    let first = 0;

    // Store the second largest value found so far
    let second = 0;

    // Traverse the array once
    for (const num of nums) {

        // If current number becomes the largest
        if (num >= first) {
            // Old largest becomes second largest
            second = first;

            // Update largest
            first = num;
        }
        // Otherwise update second largest if needed
        else if (num > second) {
            second = num;
        }
    }

    // Return the required product
    return (first - 1) * (second - 1);
};