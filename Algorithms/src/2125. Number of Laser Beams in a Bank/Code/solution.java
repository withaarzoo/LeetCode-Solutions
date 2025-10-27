class Solution {
    public int numberOfBeams(String[] bank) {
        long ans = 0; // accumulate beams (use long to avoid overflow)
        long prev = 0; // devices in previous non-empty row

        for (String row : bank) {
            long cnt = 0;
            // count '1's in this row
            for (int i = 0; i < row.length(); ++i) {
                if (row.charAt(i) == '1')
                    cnt++;
            }
            if (cnt > 0) {
                ans += prev * cnt;
                prev = cnt;
            }
        }
        return (int) ans;
    }
}
