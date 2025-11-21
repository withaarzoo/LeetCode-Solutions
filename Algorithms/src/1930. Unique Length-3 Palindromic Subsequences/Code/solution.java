class Solution {
    public int countPalindromicSubsequence(String s) {
        int n = s.length();
        int A = 26;
        int[] first = new int[A];
        int[] last = new int[A];
        Arrays.fill(first, Integer.MAX_VALUE);
        Arrays.fill(last, -1);

        // record first and last occurrence for every letter
        for (int i = 0; i < n; ++i) {
            int c = s.charAt(i) - 'a';
            first[c] = Math.min(first[c], i);
            last[c] = Math.max(last[c], i);
        }

        int ans = 0;
        // for each outer letter, count distinct middle letters between first and last
        for (int c = 0; c < A; ++c) {
            if (first[c] < last[c]) {
                boolean[] seen = new boolean[A];
                for (int i = first[c] + 1; i < last[c]; ++i) {
                    seen[s.charAt(i) - 'a'] = true;
                }
                for (int j = 0; j < A; ++j)
                    if (seen[j])
                        ans++;
            }
        }
        return ans;
    }
}
