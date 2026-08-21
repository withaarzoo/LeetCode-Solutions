import java.util.*;

class Solution {
    public long findKthSmallest(int[] coins, int k) {
        // Sort coins so smaller denominations are processed first.
        Arrays.sort(coins);

        // Store only denominations that add new valid multiples.
        List<Long> usefulList = new ArrayList<>();

        for (int coin : coins) {
            boolean redundant = false;

            // Check whether a smaller kept coin already covers this coin.
            for (long prev : usefulList) {
                if (coin % prev == 0) {
                    redundant = true;
                    break;
                }
            }

            // Keep only non-redundant coins.
            if (!redundant) {
                usefulList.add((long) coin);
            }
        }

        int m = usefulList.size();

        // Convert the list into an array for faster repeated access.
        long[] useful = new long[m];
        for (int i = 0; i < m; i++) {
            useful[i] = usefulList.get(i);
        }

        // The kth multiple of the smallest coin is always a valid upper bound.
        long low = 1;
        long high = useful[0] * k;

        int totalMasks = 1 << m;

        // Store the LCM for every subset.
        long[] lcms = new long[totalMasks];

        // Store +1 for odd subsets and -1 for even subsets.
        int[] signs = new int[totalMasks];

        // Precompute LCM values and inclusion-exclusion signs.
        for (int mask = 1; mask < totalMasks; mask++) {
            long currentLCM = 1;
            int bits = 0;

            for (int i = 0; i < m; i++) {
                // Process this coin only when its bit belongs to the subset.
                if ((mask & (1 << i)) != 0) {
                    long g = gcd(currentLCM, useful[i]);

                    // Divide first to keep multiplication smaller.
                    currentLCM /= g;

                    // Cap the LCM when it becomes too large to matter.
                    if (currentLCM > high / useful[i]) {
                        currentLCM = high + 1;
                        break;
                    }

                    currentLCM *= useful[i];
                    bits++;
                }
            }

            // Save the final LCM for this subset.
            lcms[mask] = currentLCM;

            // Apply inclusion-exclusion based on subset size.
            signs[mask] = (bits % 2 == 1) ? 1 : -1;
        }

        // Binary search for the smallest value with at least k valid amounts.
        while (low < high) {
            long mid = low + (high - low) / 2;
            long count = 0;

            // Count numbers up to mid using inclusion-exclusion.
            for (int mask = 1; mask < totalMasks; mask++) {
                // This subset contributes only when its LCM is not larger than mid.
                if (lcms[mask] <= mid) {
                    count += signs[mask] * (mid / lcms[mask]);
                }
            }

            // Move left when mid already contains at least k valid amounts.
            if (count >= k) {
                high = mid;
            } else {
                // Otherwise, move right.
                low = mid + 1;
            }
        }

        // low is the kth smallest valid amount.
        return low;
    }

    private long gcd(long a, long b) {
        // Use the Euclidean algorithm to calculate the greatest common divisor.
        while (b != 0) {
            long temp = a % b;
            a = b;
            b = temp;
        }

        return a;
    }
}