class Solution {
    public int getCommon(int[] nums1, int[] nums2) {
        
        // Pointer for nums1
        int i = 0;

        // Pointer for nums2
        int j = 0;

        // Traverse both arrays
        while (i < nums1.length && j < nums2.length) {

            // Common value found
            if (nums1[i] == nums2[j]) {
                return nums1[i];
            }

            // Move the smaller value forward
            if (nums1[i] < nums2[j]) {
                i++;
            } else {
                j++;
            }
        }

        // No common element exists
        return -1;
    }
}