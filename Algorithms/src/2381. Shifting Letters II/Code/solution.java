class Solution {
    public String shiftingLetters(String s, int[][] shifts) {
        int n = s.length();
        int[] diff = new int[n + 1];

        // Build the difference array
        for (int[] shift : shifts) {
            int start = shift[0], end = shift[1], direction = shift[2];
            int delta = (direction == 1) ? 1 : -1;
            diff[start] += delta;
            if (end + 1 < n)
                diff[end + 1] -= delta;
        }

        // Calculate cumulative shifts
        int shift = 0;
        char[] result = s.toCharArray();
        for (int i = 0; i < n; i++) {
            shift += diff[i];
            shift = (shift % 26 + 26) % 26; // Normalize shift to [0, 25]
            result[i] = (char) ('a' + (result[i] - 'a' + shift) % 26);
        }

        return new String(result);
    }
}
