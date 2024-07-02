import java.util.*;

class Solution {
    public int[] intersect(int[] nums1, int[] nums2) {
        // Step 1: Create a HashMap to store frequency of each number in nums1
        HashMap<Integer, Integer> countMap = new HashMap<>();
        
        // Step 2: Iterate through nums1 and populate the countMap
        for (int num : nums1) {
            countMap.put(num, countMap.getOrDefault(num, 0) + 1);
        }
        
        // Step 3: Prepare a list to store intersecting elements
        List<Integer> result = new ArrayList<>();
        
        // Step 4: Iterate through nums2 to find intersecting elements
        for (int num : nums2) {
            // Check if num exists in countMap and its count is greater than 0
            if (countMap.getOrDefault(num, 0) > 0) {
                // Add num to result since it's common in both arrays
                result.add(num);
                // Decrease the count in countMap for num
                countMap.put(num, countMap.get(num) - 1);
            }
        }
        
        // Step 5: Convert result list to array and return
        return result.stream().mapToInt(i -> i).toArray();
    }
}
