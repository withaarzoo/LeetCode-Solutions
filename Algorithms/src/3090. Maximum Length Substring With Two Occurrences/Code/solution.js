/**
 * @param {string} s
 * @return {number}
 */
var maximumLengthSubstring = function(s) {
    // Store the number of times each lowercase letter appears
    // inside the current sliding window.
    const freq = new Array(26).fill(0);

    // left marks the start of the current window.
    let left = 0;

    // ans stores the maximum valid window length found so far.
    let ans = 0;

    // Expand the window one character at a time.
    for (let right = 0; right < s.length; right++) {
        // Convert the current character into an index from 0 to 25
        // and increase its frequency in the window.
        const index = s.charCodeAt(right) - 97;
        freq[index]++;

        // If this character appears more than two times,
        // shrink the window until it appears at most twice.
        while (freq[index] > 2) {
            // Remove the character leaving the window.
            freq[s.charCodeAt(left) - 97]--;

            // Move the left pointer forward.
            left++;
        }

        // The current window is valid, so update the maximum length.
        ans = Math.max(ans, right - left + 1);
    }

    // Return the longest valid substring length.
    return ans;
};