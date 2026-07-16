class Solution {
    // Euclidean algorithm for gcd
    private int gcd(int a, int b) {
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }

    public long gcdSum(int[] nums) {
        int n = nums.length;

        // Store prefix gcd values
        int[] prefixGcd = new int[n];

        // Running prefix maximum
        int prefixMax = 0;

        // Build prefixGcd
        for (int i = 0; i < n; i++) {
            prefixMax = Math.max(prefixMax, nums[i]);
            prefixGcd[i] = gcd(nums[i], prefixMax);
        }

        // Sort the array
        java.util.Arrays.sort(prefixGcd);

        long ans = 0;

        // Pair smallest with largest
        int left = 0;
        int right = n - 1;

        while (left < right) {
            ans += gcd(prefixGcd[left], prefixGcd[right]);
            left++;
            right--;
        }

        return ans;
    }
}