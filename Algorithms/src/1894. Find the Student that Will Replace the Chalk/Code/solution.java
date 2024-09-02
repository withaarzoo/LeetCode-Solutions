class Solution {
    public int chalkReplacer(int[] chalk, int k) {
        // Step 1: Calculate the total amount of chalk used in one complete round.
        // We use a long to avoid potential overflow issues since the sum can be large.
        long totalChalk = 0;
        for (int c : chalk) {
            totalChalk += c; // Add each student's chalk usage to the total.
        }

        // Step 2: Use the modulo operation to reduce k.
        // The idea is that after consuming chalk in multiple rounds, we only care about
        // the remainder,
        // as full rounds won't affect the result (we only care about the first
        // incomplete round).
        k %= totalChalk;

        // Step 3: Determine which student will run out of chalk and thus need to
        // replace it.
        // Iterate over each student's chalk usage.
        for (int i = 0; i < chalk.length; i++) {
            // If k is less than the current student's chalk usage, that student can't
            // complete their turn.
            if (k < chalk[i]) {
                return i; // Return the index of the student who will replace the chalk.
            }
            // Otherwise, reduce k by the amount of chalk the current student uses.
            k -= chalk[i];
        }

        // Step 4: Safety return (though logically, we should never reach this point in
        // the code).
        // The loop will always find the student who will replace the chalk.
        return -1;
    }
}
