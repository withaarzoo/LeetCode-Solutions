class Solution {
    public int minimumDistance(int[] nums) {
        Map<Integer, List<Integer>> positions = new HashMap<>();
        
        // Store all indices for each value
        for (int i = 0; i < nums.length; i++) {
            positions.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }
        
        int ans = Integer.MAX_VALUE;
        
        // Check every value's index list
        for (List<Integer> idx : positions.values()) {
            if (idx.size() < 3) continue;
            
            // Check every consecutive group of 3 indices
            for (int i = 0; i + 2 < idx.size(); i++) {
                int distance = 2 * (idx.get(i + 2) - idx.get(i));
                ans = Math.min(ans, distance);
            }
        }
        
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}