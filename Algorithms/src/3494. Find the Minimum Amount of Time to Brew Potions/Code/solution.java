class Solution {
    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        int m = mana.length;
        if (m == 0) return 0L;

        // prefix sums of skills (use long)
        long[] pref = new long[n];
        for (int i = 0; i < n; ++i) {
            pref[i] = skill[i] + (i > 0 ? pref[i-1] : 0L);
        }

        long S = 0L; // S_j (start time of current potion on wizard 0)
        for (int j = 1; j < m; ++j) {
            long prev = (long) mana[j-1];
            long cur  = (long) mana[j];
            long best = Long.MIN_VALUE;
            for (int i = 0; i < n; ++i) {
                long prev_pref = (i > 0 ? pref[i-1] : 0L);
                long cand = pref[i] * prev - prev_pref * cur;
                if (cand > best) best = cand;
            }
            S += best;
        }

        long ans = S + pref[n-1] * (long) mana[m-1];
        return ans;
    }
}
