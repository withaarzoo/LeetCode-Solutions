import java.util.*;

class Solution {
    public int intersectionSizeTwo(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> {
            if (a[1] != b[1])
                return Integer.compare(a[1], b[1]);
            return Integer.compare(b[0], a[0]); // start descending if ends equal
        });

        int a = Integer.MIN_VALUE / 2, b = Integer.MIN_VALUE / 2; // last two chosen
        int ans = 0;
        for (int[] iv : intervals) {
            int l = iv[0], r = iv[1];
            if (l > b) {
                ans += 2;
                a = r - 1;
                b = r;
            } else if (l > a) {
                ans += 1;
                a = b;
                b = r;
            } else {
                // already have two numbers inside
            }
        }
        return ans;
    }
}
