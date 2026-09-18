class Solution {
    public List<String> maxNumOfSubstrings(String s) {
        int n = s.length();
        // first and last occurrence of each letter
        int[] left = new int[26];
        int[] right = new int[26];
        Arrays.fill(left, -1);
        Arrays.fill(right, -1);
        for (int i = 0; i < n; i++) {
            int c = s.charAt(i) - 'a';
            if (left[c] == -1) left[c] = i;   // record first time we see it
            right[c] = i;                     // always update last time
        }

        // collect every valid closed interval
        List<int[]> intervals = new ArrayList<>();
        for (int i = 0; i < 26; i++) {
            if (left[i] == -1) continue;      // letter never appeared
            int L = left[i], R = right[i];
            boolean valid = true;
            // expand right end while scanning the current range
            for (int j = L; j <= R; j++) {
                int c = s.charAt(j) - 'a';
                if (left[c] < L) {            // needs to start earlier -> not closed
                    valid = false;
                    break;
                }
                R = Math.max(R, right[c]);    // push right end if needed
            }
            if (valid) intervals.add(new int[]{L, R});
        }

        // sort by ending position so greedy can pick earliest-ending first
        intervals.sort((a, b) -> Integer.compare(a[1], b[1]));

        // greedy selection of non-overlapping intervals
        List<String> ans = new ArrayList<>();
        int lastEnd = -1;
        for (int[] iv : intervals) {
            int L = iv[0], R = iv[1];
            if (L > lastEnd) {                // completely after previous piece
                ans.add(s.substring(L, R + 1));
                lastEnd = R;
            }
        }
        return ans;
    }
}