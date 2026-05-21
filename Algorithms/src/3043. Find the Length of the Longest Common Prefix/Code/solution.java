class Solution {
    public int longestCommonPrefix(int[] arr1, int[] arr2) {
        
        // Hash set to store all prefixes from arr1
        HashSet<Integer> prefixes = new HashSet<>();

        // Generate all prefixes
        for (int num : arr1) {

            int x = num;

            // Keep removing last digit
            while (x > 0) {

                // Store current prefix
                prefixes.add(x);

                // Remove last digit
                x /= 10;
            }
        }

        int ans = 0;

        // Process arr2
        for (int num : arr2) {

            int x = num;

            // Keep checking prefixes
            while (x > 0) {

                // Prefix found
                if (prefixes.contains(x)) {

                    // Update maximum length
                    ans = Math.max(ans, String.valueOf(x).length());

                    // Stop because larger prefix already found
                    break;
                }

                // Remove last digit
                x /= 10;
            }
        }

        return ans;
    }
}