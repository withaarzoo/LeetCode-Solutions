class Solution {
    public int numberOfSubstrings(String s) {

        // Store the frequency of 'a', 'b', and 'c'
        int[] freq = new int[3];

        int left = 0;
        int ans = 0;
        int n = s.length();

        // Expand the window
        for (int right = 0; right < n; right++) {

            // Include the current character
            freq[s.charAt(right) - 'a']++;

            // Shrink while all three characters exist
            while (freq[0] > 0 && freq[1] > 0 && freq[2] > 0) {

                // Count every possible ending position
                ans += (n - right);

                // Remove the leftmost character
                freq[s.charAt(left) - 'a']--;

                // Move left forward
                left++;
            }
        }

        // Return the final answer
        return ans;
    }
}