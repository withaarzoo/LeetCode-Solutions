import java.util.*;

class Solution {
    public int[] findXSum(int[] nums, int k, int x) {
        int n = nums.length;
        int[] ans = new int[Math.max(0, n - k + 1)];
        Map<Integer, Integer> freq = new HashMap<>();

        // Initial window
        for (int i = 0; i < k; i++) {
            freq.put(nums[i], freq.getOrDefault(nums[i], 0) + 1);
        }

        ans[0] = computeXSum(freq, x);

        // Slide
        for (int i = k; i < n; i++) {
            int add = nums[i];
            int rem = nums[i - k];

            freq.put(add, freq.getOrDefault(add, 0) + 1);
            int fr = freq.get(rem) - 1;
            if (fr == 0)
                freq.remove(rem);
            else
                freq.put(rem, fr);

            ans[i - k + 1] = computeXSum(freq, x);
        }

        return ans;
    }

    private int computeXSum(Map<Integer, Integer> freq, int x) {
        // Build list of (value, freq)
        List<int[]> items = new ArrayList<>();
        for (Map.Entry<Integer, Integer> e : freq.entrySet()) {
            items.add(new int[] { e.getKey(), e.getValue() });
        }
        // Sort by freq desc, value desc
        items.sort((a, b) -> {
            if (a[1] != b[1])
                return b[1] - a[1];
            return b[0] - a[0];
        });
        long sum = 0;
        int take = Math.min(x, items.size());
        for (int i = 0; i < take; i++) {
            sum += 1L * items.get(i)[0] * items.get(i)[1];
        }
        return (int) sum;
    }
}
