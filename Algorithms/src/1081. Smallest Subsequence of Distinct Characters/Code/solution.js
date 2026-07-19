/**
 * @param {string} s
 * @return {string}
 */
var smallestSubsequence = function(s) {

    // Count remaining occurrences
    const freq = new Array(26).fill(0);
    for (const ch of s) {
        freq[ch.charCodeAt(0) - 97]++;
    }

    // Track characters already inside the stack
    const inStack = new Array(26).fill(false);

    // Array works as a stack
    const stack = [];

    for (const ch of s) {

        // Current occurrence has been processed
        freq[ch.charCodeAt(0) - 97]--;

        // Skip duplicates
        if (inStack[ch.charCodeAt(0) - 97]) {
            continue;
        }

        // Remove larger characters if they appear again later
        while (
            stack.length > 0 &&
            stack[stack.length - 1] > ch &&
            freq[stack[stack.length - 1].charCodeAt(0) - 97] > 0
        ) {
            inStack[stack.pop().charCodeAt(0) - 97] = false;
        }

        // Push current character
        stack.push(ch);
        inStack[ch.charCodeAt(0) - 97] = true;
    }

    // Join stack into the answer
    return stack.join("");
};