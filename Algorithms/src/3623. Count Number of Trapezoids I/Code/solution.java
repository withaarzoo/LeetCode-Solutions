import java.util.*;

class Solution {
    public int countTrapezoids(int[][] points) {
        final long MOD = 1_000_000_007L;
        final long INV2 = (MOD + 1) / 2; // modular inverse of 2

        // 1. Count points for each y-coordinate
        HashMap<Integer, Integer> freq = new HashMap<>();
        for (int[] p : points) {
            int y = p[1];
            freq.put(y, freq.getOrDefault(y, 0) + 1);
        }

        long sumF = 0; // S
        long sumF2 = 0; // SQ

        // 2. For each y, compute C(c,2) and accumulate
        for (int c : freq.values()) {
            if (c >= 2) {
                long cc = c;
                long f = (cc * (cc - 1) / 2) % MOD; // C(c,2)
                sumF = (sumF + f) % MOD;
                sumF2 = (sumF2 + f * f % MOD) % MOD;
            }
        }

        // 3. Compute ((S^2 - SQ) / 2) mod MOD
        long ans = (sumF * sumF) % MOD;
        ans = (ans - sumF2 + MOD) % MOD;
        ans = ans * INV2 % MOD;

        return (int) ans;
    }
}
