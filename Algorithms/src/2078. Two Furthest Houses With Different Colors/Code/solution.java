class Solution {
    public int maxDistance(int[] colors) {
        int n = colors.length;
        int ans = 0;

        // Check farthest house from the first house
        for (int i = n - 1; i >= 0; i--) {
            if (colors[i] != colors[0]) {
                ans = Math.max(ans, i);
                break;
            }
        }

        // Check farthest house from the last house
        for (int i = 0; i < n; i++) {
            if (colors[i] != colors[n - 1]) {
                ans = Math.max(ans, n - 1 - i);
                break;
            }
        }

        return ans;
    }
}