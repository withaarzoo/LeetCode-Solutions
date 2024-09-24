import java.util.HashMap;

class Solution {

    // Method to find the length of the longest common prefix between two integer
    // arrays.
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        // Step 1: Create a HashMap to store prefixes from the first array (arr1)
        HashMap<String, Integer> prefixMap = new HashMap<>();

        // Step 2: Build the prefix map using elements from arr1
        // For each number in arr1, we will extract its prefixes and store them in the
        // map.
        for (int num : arr1) {
            // Convert the number to a string so we can handle its individual digits.
            String strNum = Integer.toString(num);
            // This string will store the growing prefix of the number as we process each
            // digit.
            String prefix = "";

            // Loop through each character (digit) in the string representation of the
            // number.
            for (char ch : strNum.toCharArray()) {
                // Append the current character to the prefix.
                prefix += ch;
                // Put the prefix in the map, or increment its count if it's already there.
                // getOrDefault(prefix, 0) ensures that if the prefix doesn't exist, we start
                // with 0.
                prefixMap.put(prefix, prefixMap.getOrDefault(prefix, 0) + 1);
            }
        }

        // Step 3: Initialize maxLength to store the length of the longest common
        // prefix.
        int maxLength = 0;

        // Step 4: Check the prefixes of the numbers in arr2 against the ones stored in
        // the map.
        for (int num : arr2) {
            // Convert the number to a string to process its digits one by one.
            String strNum = Integer.toString(num);
            // Initialize an empty string to accumulate the prefix as we go through the
            // digits.
            String prefix = "";

            // Loop through each character (digit) in the string representation of the
            // number.
            for (char ch : strNum.toCharArray()) {
                // Append the current character to the prefix.
                prefix += ch;

                // If this prefix exists in the map (i.e., it was found in arr1),
                // we consider it as a potential common prefix.
                if (prefixMap.containsKey(prefix)) {
                    // Update maxLength with the maximum of the current maxLength and the length of
                    // the common prefix.
                    maxLength = Math.max(maxLength, prefix.length());
                }
            }
        }

        // Step 5: Return the length of the longest common prefix found.
        return maxLength;
    }
}
