import java.util.*;

class Solution {
    public int specialTriplets(int[] nums) {
        final int MOD = 1_000_000_007;
        Map<Integer, Long> right = new HashMap<>();
        Map<Integer, Long> left = new HashMap<>();

        // Count all elements into 'right'
        for (int x : nums) {
            right.put(x, right.getOrDefault(x, 0L) + 1);
        }

        long ans = 0;

        for (int x : nums) {
            // Current x becomes the middle index, remove from right
            right.put(x, right.get(x) - 1);

            long target = (long) x * 2L; // value 2x

            long cntLeft = left.getOrDefault((int) target, 0L);
            long cntRight = right.getOrDefault((int) target, 0L);

            long add = (cntLeft * cntRight) % MOD;
            ans = (ans + add) % MOD;

            // Move x to left
            left.put(x, left.getOrDefault(x, 0L) + 1);
        }

        return (int) (ans % MOD);
    }
}
