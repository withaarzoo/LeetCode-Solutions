class Solution {
    public char processStr(String s, long k) {
        int n = s.length();

        // len[i] = length after processing s[i]
        long[] len = new long[n];
        long curLen = 0;

        for (int i = 0; i < n; i++) {
            char c = s.charAt(i);

            if (c >= 'a' && c <= 'z') {
                // Append character
                curLen++;
            } else if (c == '*') {
                // Remove last character if present
                if (curLen > 0)
                    curLen--;
            } else if (c == '#') {
                // Duplicate string
                curLen *= 2;
            } else { // '%'
                // Length unchanged
            }

            len[i] = curLen;
        }

        // Out of bounds
        if (k >= curLen)
            return '.';

        // Undo operations from right to left
        for (int i = n - 1; i >= 0; i--) {
            char c = s.charAt(i);

            long before = (i == 0 ? 0 : len[i - 1]);

            if (c >= 'a' && c <= 'z') {
                // Letter was appended at index "before"
                if (k == before)
                    return c;
            } else if (c == '#') {
                // Undo duplication
                if (before > 0)
                    k %= before;
            } else if (c == '%') {
                // Undo reverse
                k = before - 1 - k;
            } else {
                // '*' needs no index adjustment
            }
        }

        return '.';
    }
}