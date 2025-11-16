class Solution {
    public int numSub(String s) {
        final long MOD = 1_000_000_007L;
        long res = 0L;
        long cnt = 0L; // current consecutive '1's

        for (int i = 0; i < s.length(); ++i) {
            if (s.charAt(i) == '1') {
                cnt++;
            } else {
                res = (res + (cnt * (cnt + 1) / 2) % MOD) % MOD;
                cnt = 0L;
            }
        }
        // last block
        res = (res + (cnt * (cnt + 1) / 2) % MOD) % MOD;
        return (int) res;
    }
}
