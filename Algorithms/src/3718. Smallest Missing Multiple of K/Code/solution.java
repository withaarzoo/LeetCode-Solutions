import java.util.HashSet;
import java.util.Set;

class Solution {
    public int missingMultiple(int[] nums, int k) {
        // Store every number so checking whether a value exists is fast.
        Set<Integer> present = new HashSet<>();

        // Add all values from nums into the hash set.
        for (int num : nums) {
            present.add(num);
        }

        // Start with the smallest positive multiple of k.
        int multiple = k;

        // Continue while the current multiple is already present.
        while (present.contains(multiple)) {
            // Move to the next positive multiple of k.
            multiple += k;
        }

        // Return the first missing multiple.
        return multiple;
    }
}