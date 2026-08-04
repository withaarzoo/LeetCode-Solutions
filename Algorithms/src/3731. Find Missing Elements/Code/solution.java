class Solution {
    public List<Integer> findMissingElements(int[] nums) {
        // Find the smallest and largest values
        int mn = Integer.MAX_VALUE;
        int mx = Integer.MIN_VALUE;

        // Store all numbers for fast lookup
        HashSet<Integer> seen = new HashSet<>();

        for (int num : nums) {
            mn = Math.min(mn, num);
            mx = Math.max(mx, num);
            seen.add(num);
        }

        // Store missing numbers
        List<Integer> ans = new ArrayList<>();

        // Check every number in the range
        for (int num = mn; num <= mx; num++) {
            // If the number does not exist, add it
            if (!seen.contains(num)) {
                ans.add(num);
            }
        }

        // Return the result
        return ans;
    }
}