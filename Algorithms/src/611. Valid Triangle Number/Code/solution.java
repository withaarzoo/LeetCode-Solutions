import java.util.Arrays;

class Solution {
    public int triangleNumber(int[] nums) {
        int n = nums.length;
        if (n < 3) return 0;
        Arrays.sort(nums);                         // sort ascending
        int count = 0;
        for (int k = n - 1; k >= 2; k--) {
            int i = 0, j = k - 1;                  // two pointers
            while (i < j) {
                if (nums[i] + nums[j] > nums[k]) {
                    count += j - i;               // all between i and j-1 pair with j
                    j--;                          // decrease b
                } else {
                    i++;                          // increase a
                }
            }
        }
        return count;
    }
}
