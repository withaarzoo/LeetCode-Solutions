/**
 * @param {number[]} arr
 * @param {number} target
 * @return {number}
 */
var minSumOfLengths = function(arr, target) {
    const n = arr.length;

    // best[i] stores the shortest target-sum sub-array
    // found completely inside arr[0..i].
    const best = new Array(n).fill(Infinity);

    let left = 0;       // Left boundary of the sliding window.
    let sum = 0;        // Sum of the current window.
    let answer = Infinity;

    for (let right = 0; right < n; right++) {
        // Expand the window by adding arr[right].
        sum += arr[right];

        // All values are positive, so shrinking from the left
        // is guaranteed to reduce the sum.
        while (sum > target) {
            sum -= arr[left];
            left++;
        }

        // The current window has exactly the target sum.
        if (sum === target) {
            const currentLength = right - left + 1;

            // best[left - 1] represents a valid sub-array
            // that ends before the current one starts.
            if (left > 0 && best[left - 1] !== Infinity) {
                answer = Math.min(
                    answer,
                    currentLength + best[left - 1]
                );
            }

            // Save the current valid sub-array length for this prefix.
            best[right] = currentLength;
        }

        // Keep the shortest valid sub-array seen anywhere up to right.
        if (right > 0) {
            best[right] = Math.min(best[right], best[right - 1]);
        }
    }

    // Return -1 if we could not build two non-overlapping sub-arrays.
    return answer === Infinity ? -1 : answer;
};