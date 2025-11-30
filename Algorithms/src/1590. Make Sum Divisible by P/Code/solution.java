import java.util.*;

class Solution {
    public int minSubarray(int[] nums, int p) {
        long total = 0;
        for (int x : nums) {
            total = (total + x) % p; // keep modulo to avoid overflow
        }

        int need = (int) total;
        if (need == 0)
            return 0; // already divisible

        int n = nums.length;
        Map<Integer, Integer> lastIndex = new HashMap<>();
        lastIndex.put(0, -1); // prefix before any element

        int ans = n;
        long prefix = 0;

        for (int i = 0; i < n; i++) {
            prefix = (prefix + nums[i]) % p;
            int prefMod = (int) prefix;

            int target = prefMod - need;
            if (target < 0)
                target += p; // (prefMod - need + p) % p

            if (lastIndex.containsKey(target)) {
                ans = Math.min(ans, i - lastIndex.get(target));
            }

            // store latest index for this prefix remainder
            lastIndex.put(prefMod, i);
        }

        return ans == n ? -1 : ans;
    }
}
