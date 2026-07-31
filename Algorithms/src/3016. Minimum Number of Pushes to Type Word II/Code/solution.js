/**
 * @param {string} word
 * @return {number}
 */
var minimumPushes = function(word) {
    // Store the frequency of every letter
    const freq = new Array(26).fill(0);

    // Count each character
    for (const ch of word) {
        freq[ch.charCodeAt(0) - 97]++;
    }

    // Sort from largest to smallest
    freq.sort((a, b) => b - a);

    let ans = 0;

    // Assign push cost based on sorted position
    for (let i = 0; i < 26; i++) {
        // Stop when no more letters exist
        if (freq[i] === 0) break;

        // Every group of 8 letters gets one more push
        const pushes = Math.floor(i / 8) + 1;

        // Add contribution
        ans += freq[i] * pushes;
    }

    return ans;
};