class Solution {
    public int minimumDistance(int[] nums) {
        Map<Integer, List<Integer>> map = new HashMap<>();
        
        // Store all indices for each value
        for (int i = 0; i < nums.length; i++) {
            map.computeIfAbsent(nums[i], k -> new ArrayList<>()).add(i);
        }
        
        int ans = Integer.MAX_VALUE;
        
        // Process each value's indices
        for (List<Integer> indices : map.values()) {
            if (indices.size() < 3) continue;
            
            // Check every consecutive group of 3 indices
            for (int i = 0; i + 2 < indices.size(); i++) {
                int distance = 2 * (indices.get(i + 2) - indices.get(i));
                ans = Math.min(ans, distance);
            }
        }
        
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}