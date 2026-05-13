class Solution {
    public int minMoves(int[] nums, int limit) {
        
        int n = nums.length;

        // Difference array
        int[] diff = new int[2 * limit + 2];

        // Process all pairs
        for (int i = 0; i < n / 2; i++) {

            int a = Math.min(nums[i], nums[n - 1 - i]);
            int b = Math.max(nums[i], nums[n - 1 - i]);

            // 1 move range
            diff[a + 1] -= 1;
            diff[b + limit + 1] += 1;

            // Exact sum needing 0 moves
            diff[a + b] -= 1;
            diff[a + b + 1] += 1;
        }

        int pairs = n / 2;

        // Initially assume 2 moves for every pair
        int current = pairs * 2;

        int answer = Integer.MAX_VALUE;

        // Prefix sum traversal
        for (int sum = 2; sum <= 2 * limit; sum++) {

            current += diff[sum];

            answer = Math.min(answer, current);
        }

        return answer;
    }
}