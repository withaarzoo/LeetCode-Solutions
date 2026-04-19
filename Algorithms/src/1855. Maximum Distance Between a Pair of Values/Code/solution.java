class Solution {
    public int maxDistance(int[] nums1, int[] nums2) {
        int i = 0, j = 0;
        int ans = 0;

        while (i < nums1.length && j < nums2.length) {

            // Ensure i <= j
            if (i > j) {
                j++;
                continue;
            }

            // Valid pair
            if (nums1[i] <= nums2[j]) {
                ans = Math.max(ans, j - i);
                j++; // Try for a bigger distance
            } else {
                // Invalid pair
                i++;
            }
        }

        return ans;
    }
}