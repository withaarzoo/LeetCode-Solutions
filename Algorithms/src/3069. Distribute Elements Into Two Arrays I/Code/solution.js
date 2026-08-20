/**
 * @param {number[]} nums
 * @return {number[]}
 */
var resultArray = function(nums) {
    // I create two arrays to simulate arr1 and arr2 from the problem.
    const arr1 = [nums[0]];
    const arr2 = [nums[1]];

    // I process every number after the first two required operations.
    for (let i = 2; i < nums.length; i++) {
        // If arr1 has the greater last element, I add nums[i] to arr1.
        if (arr1[arr1.length - 1] > arr2[arr2.length - 1]) {
            arr1.push(nums[i]);
        } else {
            // Otherwise, I add nums[i] to arr2.
            arr2.push(nums[i]);
        }
    }

    // I return arr1 followed by arr2, exactly as the problem requires.
    return [...arr1, ...arr2];
};