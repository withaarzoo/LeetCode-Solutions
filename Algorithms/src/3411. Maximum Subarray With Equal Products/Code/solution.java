class Solution {
    public int maxLength(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 0; i < n; i++) {
            long prod = 1;
            int g = 0;
            long lc = 1;
            for (int j = i; j < n; j++) {
                int x = nums[j];
                if (g == 0) g = x;
                else {
                    int a = g, b = x;
                    while (b != 0) {
                        int t = b;
                        b = a % b;
                        a = t;
                    }
                    g = a;
                }
                long gg = lc;
                int b = x;
                while (b != 0) {
                    long t = b;
                    b = (int)(gg % b);
                    gg = t;
                }
                if (lc / gg > Long.MAX_VALUE / x) break;
                lc = lc / gg * x;
                if (prod > Long.MAX_VALUE / x) break;
                prod *= x;
                if (prod == lc * g) ans = Math.max(ans, j - i + 1);
            }
        }
        return ans;
    }
}