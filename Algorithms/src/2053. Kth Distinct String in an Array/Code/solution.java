import java.util.*;

class Solution {
    public String kthDistinct(String[] arr, int k) {
        // Create a HashMap to count occurrences of each string
        Map<String, Integer> count = new HashMap<>();

        // Create a List to store distinct strings
        List<String> distinct = new ArrayList<>();

        // Loop through each string in the array
        for (String str : arr) {
            // Increment the count for each string
            count.put(str, count.getOrDefault(str, 0) + 1);
        }

        // Loop through each string in the array again to collect distinct strings
        for (String str : arr) {
            // If the string appears only once, add it to the distinct list
            if (count.get(str) == 1) {
                distinct.add(str);
            }
        }

        // Check if the k-th distinct string exists
        if (k <= distinct.size()) {
            // Return the k-th distinct string (1-based index)
            return distinct.get(k - 1);
        } else {
            // If there are fewer than k distinct strings, return an empty string
            return "";
        }
    }
}
