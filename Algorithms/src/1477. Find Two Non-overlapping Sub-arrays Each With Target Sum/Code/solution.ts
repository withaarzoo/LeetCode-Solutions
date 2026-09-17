function minSumOfLengths(arr: number[], target: number): number {
    const n = arr.length;

    // best[i] stores the shortest target-sum sub-array
    // found completely inside arr[0..i].
    const best: number[] = new Array(n).fill(Infinity);

    let left = 0;       // Left boundary of the sliding window.
    let sum = 0;        // Sum of the current sliding window.
    let answer = Infinity;

    for (let right = 0; right < n; right++) {
        // Expand the window to include arr[right].
        sum += arr[right];

        // Since all elements are positive, removing from the left
        // always decreases the sum.
        while (sum > target) {
            sum -= arr[left];
            left++;
        }

        // We found a valid sub-array with sum equal to target.
        if (sum === target) {
            const currentLength = right - left + 1;

            // best[left - 1] is guaranteed to be non-overlapping
            // with the current window because it ends before left.
            if (left > 0 && best[left - 1] !== Infinity) {
                answer = Math.min(
                    answer,
                    currentLength + best[left - 1]
                );
            }

            // Store the current valid length for this prefix.
            best[right] = currentLength;
        }

        // Keep the shortest valid sub-array seen anywhere in the prefix.
        if (right > 0) {
            best[right] = Math.min(best[right], best[right - 1]);
        }
    }

    // Return -1 when fewer than two non-overlapping valid sub-arrays exist.
    return answer === Infinity ? -1 : answer;
}