/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @return {number}
 */
var longestCommonPrefix = function(arr1, arr2) {
    
    // Set to store all prefixes from arr1
    const prefixes = new Set();

    // Generate prefixes
    for (let num of arr1) {

        let x = num;

        // Keep removing last digit
        while (x > 0) {

            // Store current prefix
            prefixes.add(x);

            // Remove last digit
            x = Math.floor(x / 10);
        }
    }

    let ans = 0;

    // Check numbers from arr2
    for (let num of arr2) {

        let x = num;

        // Try all prefixes
        while (x > 0) {

            // Prefix found
            if (prefixes.has(x)) {

                // Update answer using digit count
                ans = Math.max(ans, x.toString().length);

                // Stop because this is the longest for current number
                break;
            }

            // Remove last digit
            x = Math.floor(x / 10);
        }
    }

    return ans;
};