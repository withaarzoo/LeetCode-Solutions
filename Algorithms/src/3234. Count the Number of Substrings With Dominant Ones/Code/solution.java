import java.util.*;

class Solution {
    public int numberOfSubstrings(String s) {
        int n = s.length();
        long ans = 0;
        // count all-ones substrings
        long run = 0;
        for (int i = 0; i < n; ++i) {
            if (s.charAt(i) == '1')
                run++;
            else {
                ans += run * (run + 1) / 2;
                run = 0;
            }
        }
        ans += run * (run + 1) / 2;

        // collect zero positions
        ArrayList<Integer> zeroPos = new ArrayList<>();
        for (int i = 0; i < n; ++i)
            if (s.charAt(i) == '0')
                zeroPos.add(i);
        int m = zeroPos.size();
        if (m == 0)
            return (int) ans;

        int K = (int) Math.floor(Math.sqrt(n));
        for (int k = 1; k <= K && k <= m; ++k) {
            for (int i = 0; i + k - 1 < m; ++i) {
                int leftPrev = (i == 0 ? -1 : zeroPos.get(i - 1));
                int rightNext = (i + k - 1 == m - 1 ? n : zeroPos.get(i + k));
                int leftOnes = zeroPos.get(i) - leftPrev - 1;
                int rightOnes = rightNext - zeroPos.get(i + k - 1) - 1;
                int baseLen = zeroPos.get(i + k - 1) - zeroPos.get(i) + 1;
                long needLen = 1L * k * k + k;
                long t = needLen - baseLen; // require x+y >= t
                long totalPairs = 1L * (leftOnes + 1) * (rightOnes + 1);
                if (t <= 0) {
                    ans += totalPairs;
                    continue;
                }

                // Count pairs with x+y < t using closed-form arithmetic
                long pairs_lt = 0;
                long s0 = t - 1;
                if (s0 >= 0) {
                    long L = leftOnes;
                    long R = rightOnes;
                    long x_max = Math.min(L, s0);
                    if (x_max >= 0) {
                        long x0 = Math.max(0L, s0 - R);
                        if (x0 > x_max) {
                            pairs_lt = (x_max + 1) * (R + 1);
                        } else {
                            long part1 = x0 * (R + 1);
                            long n2 = x_max - x0 + 1;
                            long sum_x = (x0 + x_max) * n2 / 2;
                            long part2 = n2 * (s0 + 1) - sum_x;
                            pairs_lt = part1 + part2;
                        }
                    } else
                        pairs_lt = 0;
                } else
                    pairs_lt = 0;

                long valid = totalPairs - pairs_lt;
                if (valid > 0)
                    ans += valid;
            }
        }

        return (int) ans; // cast safe under constraints
    }
}
