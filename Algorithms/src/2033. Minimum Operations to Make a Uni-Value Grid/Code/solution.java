class Solution {
    public int minOperations(int[][] grid, int x) {
        List<Integer> nums = new ArrayList<>();

        // Flatten grid
        for (int[] row : grid) {
            for (int val : row) {
                nums.add(val);
            }
        }

        // Check feasibility
        int rem = nums.get(0) % x;
        for (int num : nums) {
            if (num % x != rem)
                return -1;
        }

        // Sort
        Collections.sort(nums);

        // Median
        int median = nums.get(nums.size() / 2);

        // Count operations
        int ops = 0;
        for (int num : nums) {
            ops += Math.abs(num - median) / x;
        }

        return ops;
    }
}