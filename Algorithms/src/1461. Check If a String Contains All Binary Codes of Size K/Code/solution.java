class Solution {
    public boolean hasAllCodes(String s, int k) {
        int n = s.length();

        if (n < k)
            return false;

        int total = 1 << k;
        if (n - k + 1 < total)
            return false;

        boolean[] seen = new boolean[total];
        int mask = total - 1;
        int curr = 0;
        int count = 0;

        // First window
        for (int i = 0; i < k; i++) {
            curr = (curr << 1) | (s.charAt(i) - '0');
        }

        if (!seen[curr]) {
            seen[curr] = true;
            count++;
        }

        // Sliding window
        for (int i = k; i < n; i++) {
            curr = ((curr << 1) & mask) | (s.charAt(i) - '0');

            if (!seen[curr]) {
                seen[curr] = true;
                count++;
                if (count == total)
                    return true;
            }
        }

        return count == total;
    }
}