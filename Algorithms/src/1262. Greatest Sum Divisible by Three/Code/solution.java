class Solution {
    public int maxSumDivThree(int[] nums) {
        long sum = 0;
        int INF = (int) 1e9;
        int r1_min1 = INF, r1_min2 = INF; // two smallest remainder-1 numbers
        int r2_min1 = INF, r2_min2 = INF; // two smallest remainder-2 numbers

        for (int x : nums) {
            sum += x;
            int r = x % 3;
            if (r == 1) {
                if (x < r1_min1) {
                    r1_min2 = r1_min1;
                    r1_min1 = x;
                } else if (x < r1_min2) {
                    r1_min2 = x;
                }
            } else if (r == 2) {
                if (x < r2_min1) {
                    r2_min2 = r2_min1;
                    r2_min1 = x;
                } else if (x < r2_min2) {
                    r2_min2 = x;
                }
            }
        }

        int mod = (int) (sum % 3);
        if (mod == 0)
            return (int) sum;

        long removeCost = (long) 1e18;

        if (mod == 1) {
            if (r1_min1 != INF)
                removeCost = Math.min(removeCost, (long) r1_min1);
            if (r2_min2 != INF)
                removeCost = Math.min(removeCost, (long) r2_min1 + r2_min2);
        } else { // mod == 2
            if (r2_min1 != INF)
                removeCost = Math.min(removeCost, (long) r2_min1);
            if (r1_min2 != INF)
                removeCost = Math.min(removeCost, (long) r1_min1 + r1_min2);
        }

        if (removeCost >= (long) 1e18)
            return 0;
        return (int) (sum - removeCost);
    }
}
