class Solution {

    // Fenwick Tree for frequency counting
    static class Fenwick {
        int[] bit;

        Fenwick(int n) {
            bit = new int[n + 1];
        }

        // Add one occurrence
        void update(int idx) {
            while (idx < bit.length) {
                bit[idx]++;
                idx += idx & -idx;
            }
        }

        // Prefix frequency
        long query(int idx) {
            long sum = 0;
            while (idx > 0) {
                sum += bit[idx];
                idx -= idx & -idx;
            }
            return sum;
        }
    }

    public long countMajoritySubarrays(int[] nums, int target) {

        int n = nums.length;

        // Prefix sums after transformation
        int[] pref = new int[n + 1];
        for (int i = 0; i < n; i++) {
            pref[i + 1] = pref[i] + (nums[i] == target ? 1 : -1);
        }

        // Coordinate compression
        int[] values = pref.clone();
        java.util.Arrays.sort(values);

        int m = 0;
        for (int x : values) {
            if (m == 0 || values[m - 1] != x) {
                values[m++] = x;
            }
        }

        Fenwick ft = new Fenwick(m);

        long ans = 0;

        for (int x : pref) {

            int idx = java.util.Arrays.binarySearch(values, 0, m, x) + 1;

            ans += ft.query(idx - 1);

            ft.update(idx);
        }

        return ans;
    }
}