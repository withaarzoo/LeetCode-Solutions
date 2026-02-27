class Solution {
    public int minOperations(String s, int k) {
        int n = s.length();
        int zero = 0;

        for (char c : s.toCharArray())
            if (c == '0')
                zero++;

        if (zero == 0)
            return 0;

        if (n == k) {
            if (zero == n)
                return 1;
            if (zero == 0)
                return 0;
            return -1;
        }

        int one = n - zero;
        int base = n - k;

        long ans = Long.MAX_VALUE;

        // Odd operations
        if ((k % 2) == (zero % 2)) {
            long m = Math.max(
                    (zero + k - 1) / k,
                    (one + base - 1) / base);

            if (m % 2 == 0)
                m++;

            ans = Math.min(ans, m);
        }

        // Even operations
        if (zero % 2 == 0) {
            long m = Math.max(
                    (zero + k - 1) / k,
                    (zero + base - 1) / base);

            if (m % 2 == 1)
                m++;

            ans = Math.min(ans, m);
        }

        return ans == Long.MAX_VALUE ? -1 : (int) ans;
    }
}