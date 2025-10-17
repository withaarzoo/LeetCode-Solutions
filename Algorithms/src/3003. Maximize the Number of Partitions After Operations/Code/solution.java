import java.util.*;

class Solution {
    private Map<Long, Integer> memo = new HashMap<>();
    private String s;
    private int k;

    private int dp(int i, long mask, boolean canChange) {
        if (i == s.length()) return 0;
        long key = ((long)i << 27) | (mask << 1) | (canChange ? 1 : 0);
        if (memo.containsKey(key)) return memo.get(key);

        int ch = s.charAt(i) - 'a';
        long newMask = mask | (1L << ch);
        int res;

        if (Long.bitCount(newMask) > k)
            res = 1 + dp(i + 1, 1L << ch, canChange);
        else
            res = dp(i + 1, newMask, canChange);

        if (canChange) {
            for (int j = 0; j < 26; j++) {
                long changeMask = mask | (1L << j);
                if (Long.bitCount(changeMask) > k)
                    res = Math.max(res, 1 + dp(i + 1, 1L << j, false));
                else
                    res = Math.max(res, dp(i + 1, changeMask, false));
            }
        }

        memo.put(key, res);
        return res;
    }

    public int maxPartitionsAfterOperations(String s, int k) {
        this.s = s;
        this.k = k;
        memo.clear();
        return dp(0, 0, true) + 1;
    }
}
