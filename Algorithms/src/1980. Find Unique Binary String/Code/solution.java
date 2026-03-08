class Solution {
    public String findDifferentBinaryString(String[] nums) {
        int n = nums.length;
        StringBuilder result = new StringBuilder();

        // Flip the diagonal bits
        for (int i = 0; i < n; i++) {
            if (nums[i].charAt(i) == '0')
                result.append('1');
            else
                result.append('0');
        }

        return result.toString();
    }
}