import java.util.TreeMap;

class Solution {
    // TreeMap to simulate a Multiset (Value -> Count)
    private void add(TreeMap<Integer, Integer> map, int val) {
        map.put(val, map.getOrDefault(val, 0) + 1);
    }

    private void remove(TreeMap<Integer, Integer> map, int val) {
        if (map.get(val) == 1)
            map.remove(val);
        else
            map.put(val, map.get(val) - 1);
    }

    // Helper variables
    TreeMap<Integer, Integer> L = new TreeMap<>();
    TreeMap<Integer, Integer> R = new TreeMap<>();
    long currentLSum = 0;
    int sizeL = 0;

    public long minimumCost(int[] nums, int k, int dist) {
        int target = k - 2;
        int n = nums.length;

        // Initial window for pivot i = 1: [2, dist + 1]
        for (int j = 2; j <= Math.min(n - 1, dist + 1); j++) {
            add(L, nums[j]);
            currentLSum += nums[j];
            sizeL++;
        }

        balance(target);

        long minCost = (long) nums[0] + nums[1] + currentLSum;

        // Slide the window
        for (int i = 2; i <= n - (k - 1); i++) {
            int outVal = nums[i];
            // Remove outgoing element
            if (L.containsKey(outVal)) {
                remove(L, outVal);
                currentLSum -= outVal;
                sizeL--;
            } else {
                remove(R, outVal);
            }

            // Add incoming element
            if (i + dist < n) {
                int inVal = nums[i + dist];
                add(L, inVal);
                currentLSum += inVal;
                sizeL++;
            }

            balance(target);

            minCost = Math.min(minCost, (long) nums[0] + nums[i] + currentLSum);
        }

        return minCost;
    }

    private void balance(int target) {
        // 1. Move largest from L to R if L is too big
        while (sizeL > target) {
            int maxL = L.lastKey();
            remove(L, maxL);
            currentLSum -= maxL;
            sizeL--;
            add(R, maxL);
        }

        // 2. Move smallest from R to L if L is too small
        while (sizeL < target && !R.isEmpty()) {
            int minR = R.firstKey();
            remove(R, minR);
            add(L, minR);
            currentLSum += minR;
            sizeL++;
        }

        // 3. Ensure all in L <= all in R
        while (!L.isEmpty() && !R.isEmpty() && L.lastKey() > R.firstKey()) {
            int maxL = L.lastKey();
            int minR = R.firstKey();

            remove(L, maxL);
            currentLSum -= maxL;
            sizeL--;
            add(R, maxL);

            remove(R, minR);
            add(L, minR);
            currentLSum += minR;
            sizeL++;
        }
    }
}