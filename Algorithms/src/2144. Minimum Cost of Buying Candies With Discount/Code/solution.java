class Solution {
    public int minimumCost(int[] cost) {
        // Sort in ascending order first
        Arrays.sort(cost);

        int total = 0;

        // Traverse from largest to smallest
        int position = 0;

        for (int i = cost.length - 1; i >= 0; i--) {
            // Every third candy is free
            if (position % 3 != 2) {
                total += cost[i];
            }

            position++;
        }

        return total;
    }
}