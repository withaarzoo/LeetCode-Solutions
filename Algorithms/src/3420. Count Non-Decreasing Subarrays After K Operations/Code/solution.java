class Solution {
    public long countNonDecreasingSubarrays(int[] nums, int k) {
        long ans = 0;
        long cost = 0;
        Deque<int[]> dq = new ArrayDeque<>();
        int n = nums.length;
        for (int i = n - 1, j = n - 1; i >= 0; --i) {
            int num = nums[i];
            int count = 1;
            while (!dq.isEmpty() && dq.peekLast()[0] < num) {
                int[] back = dq.pollLast();
                count += back[1];
                cost += (long) (num - back[0]) * back[1];
            }
            dq.offerLast(new int[]{num, count});
            while (cost > k) {
                int[] front = dq.pollFirst();
                cost -= (long) (front[0] - nums[j]);
                --j;
                if (front[1] > 1) {
                    dq.offerFirst(new int[]{front[0], front[1] - 1});
                }
            }
            ans += j - i + 1;
        }
        return ans;
    }
}