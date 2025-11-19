import java.util.HashSet;
import java.util.Set;

class Solution {
    public int findFinalValue(int[] nums, int original) {
        // Put all numbers in a HashSet for O(1) average lookup
        Set<Integer> set = new HashSet<>();
        for (int x : nums)
            set.add(x);
        // Keep doubling while original is present
        while (set.contains(original)) {
            original *= 2;
        }
        return original;
    }
}
