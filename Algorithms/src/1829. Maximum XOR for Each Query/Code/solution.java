class Solution {
    public int[] getMaximumXor(int[] nums, int maximumBit) {
        int n = nums.length;
        int[] answer = new int[n];
        int XORed = 0;

        // Calculate the cumulative XOR of the entire nums array
        for (int num : nums) {
            XORed ^= num;
        }

        // max_k is 2^maximumBit - 1
        int max_k = (1 << maximumBit) - 1;

        // Process each query in reverse
        for (int i = 0; i < n; i++) {
            // Calculate the k that maximizes XOR
            answer[i] = XORed ^ max_k;

            // Update XORed by removing the effect of the last element
            XORed ^= nums[n - 1 - i];
        }

        return answer;
    }
}
