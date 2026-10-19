class Solution {
    public List<Integer> targetIndices(int[] nums, int target) {
        int less = 0, equal = 0;
        for (int x : nums) {
            if (x < target) less++;
            else if (x == target) equal++;
        }
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < equal; i++) {
            ans.add(less + i);
        }
        return ans;
    }
}