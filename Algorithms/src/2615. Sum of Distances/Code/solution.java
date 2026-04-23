class Solution {
    public long[] distance(int[] nums) {
        int n = nums.length;
        Map<Integer, List<Integer>> map = new HashMap<>();

        // Group indices
        for (int i = 0; i < n; i++) {
            map.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }

        long[] res = new long[n];

        for (List<Integer> idx : map.values()) {
            int k = idx.size();

            long prefixSum = 0;
            long totalSum = 0;

            for (int x : idx)
                totalSum += x;

            for (int i = 0; i < k; i++) {
                long curr = idx.get(i);

                long left = curr * i - prefixSum;
                long right = (totalSum - prefixSum - curr) - curr * (k - i - 1);

                res[(int) curr] = left + right;

                prefixSum += curr;
            }
        }

        return res;
    }
}