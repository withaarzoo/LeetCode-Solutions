class Solution {
    public int closestTarget(String[] words, String target, int startIndex) {
        int n = words.length;
        int ans = Integer.MAX_VALUE;

        // Check every index in the array
        for (int i = 0; i < n; i++) {
            // If current word matches target
            if (words[i].equals(target)) {
                // Normal distance between indices
                int diff = Math.abs(i - startIndex);

                // Circular distance
                int circularDist = n - diff;

                // Update minimum answer
                ans = Math.min(ans, Math.min(diff, circularDist));
            }
        }

        // If target does not exist
        return ans == Integer.MAX_VALUE ? -1 : ans;
    }
}