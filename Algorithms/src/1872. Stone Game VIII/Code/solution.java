class Solution {
    public int stoneGameVIII(int[] stones) {
        // Convert stones into prefix sums without using another array.
        for (int i = 1; i < stones.length; i++) {
            stones[i] += stones[i - 1];
        }

        // Taking all stones gives the total sum as the starting state.
        int best = stones[stones.length - 1];

        // Process every possible prefix from right to left.
        for (int i = stones.length - 2; i >= 1; i--) {
            // Choose between keeping the current best result and
            // taking this prefix, then subtracting the opponent's result.
            best = Math.max(best, stones[i] - best);
        }

        // Return the best possible score difference.
        return best;
    }
}