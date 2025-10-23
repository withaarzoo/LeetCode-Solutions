class Solution {
    public boolean hasSameDigits(String s) {
        int n = s.length();
        // convert to int array
        int[] digits = new int[n];
        for (int i = 0; i < n; ++i)
            digits[i] = s.charAt(i) - '0';

        // reduce until length == 2
        while (digits.length > 2) {
            int m = digits.length - 1;
            int[] next = new int[m];
            for (int i = 0; i < m; ++i) {
                next[i] = (digits[i] + digits[i + 1]) % 10;
            }
            digits = next;
        }

        return digits.length == 2 && digits[0] == digits[1];
    }
}
