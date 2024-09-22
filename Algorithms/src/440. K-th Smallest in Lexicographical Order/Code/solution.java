class Solution {

    // Helper function to calculate the number of steps between 'curr' and 'n'
    // when traversing the numbers in lexicographical order.
    // 'curr' is the current prefix we are considering, and 'n' is the upper limit.
    private long countSteps(long curr, long n) {
        long steps = 0; // Initialize the number of steps.
        long first = curr; // Starting number for the current range.
        long last = curr; // Ending number for the current range.

        // Continue while the starting number is within the limit 'n'.
        while (first <= n) {
            // Calculate the number of valid numbers between 'first' and 'last'.
            // We use 'n + 1' because we are considering all numbers from 'first' to 'min(n,
            // last)'.
            steps += Math.min(n + 1, last + 1) - first;

            // Move to the next level by multiplying by 10.
            // This expands the range by one more digit.
            first *= 10;
            last = last * 10 + 9;
        }

        return steps; // Return the total number of steps calculated.
    }

    // Main function to find the k-th smallest number in lexicographical order.
    public int findKthNumber(int n, int k) {
        long curr = 1; // Start from the prefix 1 as the smallest lexicographical number.
        k--; // Since we are considering the first number as 1, we decrement 'k' by 1.

        // While we still need to find the k-th number (i.e., k > 0),
        // we continue adjusting the current prefix.
        while (k > 0) {
            // Calculate the number of steps between 'curr' and 'n' for the current prefix.
            long steps = countSteps(curr, n);

            // If the number of steps for the current prefix is less than or equal to 'k',
            // it means the k-th number is beyond this prefix, so we move to the next
            // prefix.
            if (steps <= k) {
                curr++; // Move to the next prefix by incrementing 'curr'.
                k -= steps; // Reduce 'k' by the number of steps skipped.
            } else {
                // If the number of steps is greater than 'k', the k-th number lies within this
                // prefix.
                // So, we go deeper into the current prefix by multiplying 'curr' by 10,
                // effectively exploring the next level in the lexicographical order.
                curr *= 10;
                k--; // Decrement 'k' because we've accounted for one step.
            }
        }

        return (int) curr; // Return the k-th number as an integer.
    }
}
