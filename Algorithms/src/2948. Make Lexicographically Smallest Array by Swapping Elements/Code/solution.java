import java.util.*;

class Solution {
    public int[] lexicographicallySmallestArray(int[] nums, int limit) {
        int n = nums.length; // Store the array size for easier use.
        
        // Each pair stores {value, originalIndex}.
        int[][] elements = new int[n][2];
        
        for (int i = 0; i < n; i++) {
            elements[i][0] = nums[i]; // Store the value.
            elements[i][1] = i;       // Store its original position.
        }
        
        // Sort all pairs by their values.
        Arrays.sort(elements, (a, b) -> Integer.compare(a[0], b[0]));
        
        // Store the final lexicographically smallest array.
        int[] answer = new int[n];
        
        int start = 0; // First position of the current connected group.
        
        while (start < n) {
            int end = start; // Expand the current group.
            
            // Keep consecutive values in the same group while they can
            // be connected through valid swaps.
            while (end + 1 < n &&
                   (long) elements[end + 1][0] - elements[end][0] <= limit) {
                end++;
            }
            
            // Store all original indices belonging to this group.
            int[] indices = new int[end - start + 1];
            
            for (int i = start; i <= end; i++) {
                indices[i - start] = elements[i][1];
            }
            
            // Sort positions so smaller values go to earlier indices.
            Arrays.sort(indices);
            
            // The values in the current group are already sorted because
            // the complete elements array was sorted by value.
            for (int i = 0; i < indices.length; i++) {
                answer[indices[i]] = elements[start + i][0];
            }
            
            // Move to the next connected group.
            start = end + 1;
        }
        
        return answer; // Return the final array.
    }
}