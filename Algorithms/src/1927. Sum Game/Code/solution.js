/**
 * @param {string} num
 * @return {boolean}
 */
var sumGame = function(num) {
    // Find where the right half starts.
    const mid = num.length / 2;

    // Store sums of known digits and counts of '?' characters.
    let leftSum = 0;
    let rightSum = 0;
    let leftQuestion = 0;
    let rightQuestion = 0;

    // Scan every character once.
    for (let i = 0; i < num.length; i++) {
        if (i < mid) {
            // Update information for the left half.
            if (num[i] === '?') {
                leftQuestion++;
            } else {
                leftSum += Number(num[i]);
            }
        } else {
            // Update information for the right half.
            if (num[i] === '?') {
                rightQuestion++;
            } else {
                rightSum += Number(num[i]);
            }
        }
    }

    // If Bob cannot satisfy this equality, Alice can force a win.
    return 2 * (leftSum - rightSum) !==
           9 * (rightQuestion - leftQuestion);
};