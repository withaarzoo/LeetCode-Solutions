class Solution {
    public int minFlips(String s) {
        int n = s.length();
        String ss = s + s;

        int diff1 = 0, diff2 = 0;
        int ans = Integer.MAX_VALUE;

        for (int i = 0; i < ss.length(); i++) {
            char c = ss.charAt(i);

            char expected1 = (i % 2 == 0) ? '0' : '1';
            char expected2 = (i % 2 == 0) ? '1' : '0';

            if (c != expected1)
                diff1++;
            if (c != expected2)
                diff2++;

            if (i >= n) {
                char prev = ss.charAt(i - n);

                char prevExp1 = ((i - n) % 2 == 0) ? '0' : '1';
                char prevExp2 = ((i - n) % 2 == 0) ? '1' : '0';

                if (prev != prevExp1)
                    diff1--;
                if (prev != prevExp2)
                    diff2--;
            }

            if (i >= n - 1) {
                ans = Math.min(ans, Math.min(diff1, diff2));
            }
        }

        return ans;
    }
}