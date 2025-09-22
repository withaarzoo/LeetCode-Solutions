import java.util.HashMap;
import java.util.Map;

class Solution {
    public int maxFrequencyElements(int[] nums) {
        Map<Integer, Integer> freq = new HashMap<>();
        // Count frequency of each number
        for (int x : nums) freq.put(x, freq.getOrDefault(x, 0) + 1);

        // Find the maximum frequency
        int maxFreq = 0;
        for (int v : freq.values()) maxFreq = Math.max(maxFreq, v);

        // Sum frequencies of elements that have frequency == maxFreq
        int result = 0;
        for (int v : freq.values()) if (v == maxFreq) result += v;
        return result;
    }
}
