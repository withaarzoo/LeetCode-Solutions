/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumOperations = function(nums) {
    let operations = 0;
    
    // Loop through the array
    for (const x of nums) {
        // If remainder when divided by 3 is not zero, need 1 operation
        if (x % 3 !== 0) {
            operations++;
        }
    }
    
    return operations;
};
