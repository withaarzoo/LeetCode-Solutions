class Solution {
    public int countPermutations(int[] complexity) {
        final int MOD = 1_000_000_007;
        int n = complexity.length;

        // 1. Find global minimum and its frequency
        int minVal = complexity[0];
        int cntMin = 0;
        for (int x : complexity) {
            if (x < minVal) {
                minVal = x;
                cntMin = 1;
            } else if (x == minVal) {
                cntMin++;
            }
        }

        // 2. Check if index 0 has unique minimum
        if (complexity[0] != minVal || cntMin != 1) {
            return 0;
        }

        // 3. Compute (n - 1)! % MOD using long
        long ans = 1;
        for (int i = 2; i <= n - 1; i++) {
            ans = (ans * i) % MOD;
        }

        return (int) ans;
    }
}
