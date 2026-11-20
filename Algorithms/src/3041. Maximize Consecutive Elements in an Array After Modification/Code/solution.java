class Solution {
    public int maxSelectedElements(int[] nums) {
        Arrays.sort(nums);
        Map<Integer, Integer> dp = new HashMap<>();
        int ans = 0;
        for (int num : nums) {
            dp.put(num + 1, dp.getOrDefault(num, 0) + 1);
            dp.put(num, dp.getOrDefault(num - 1, 0) + 1);
            ans = Math.max(ans, Math.max(dp.get(num), dp.get(num + 1)));
        }
        return ans;
    }
}