/**
 * @param {string[]} patterns
 * @param {string} word
 * @return {number}
 */
var numOfStrings = function(patterns, word) {

    // Store the number of matching patterns
    let count = 0;

    // Check every pattern
    for (const pattern of patterns) {

        // includes() returns true if pattern is a substring
        if (word.includes(pattern)) {
            count++;
        }
    }

    // Return the total count
    return count;
};