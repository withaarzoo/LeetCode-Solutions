class Solution {
    public int maximumLengthSubstring(String s) {
        // Store the number of times each lowercase letter appears
        // inside the current sliding window.
        int[] freq = new int[26];

        // left marks the start of the current window.
        int left = 0;

        // ans stores the maximum valid window length found so far.
        int ans = 0;

        // Expand the window one character at a time.
        for (int right = 0; right < s.length(); right++) {
            // Convert the current character into an index from 0 to 25
            // and increase its frequency in the window.
            freq[s.charAt(right) - 'a']++;

            // If the current character appears more than two times,
            // shrink the window until the condition becomes valid again.
            while (freq[s.charAt(right) - 'a'] > 2) {
                // Remove the character leaving the window.
                freq[s.charAt(left) - 'a']--;

                // Move the left pointer forward.
                left++;
            }

            // The current window is valid, so update the maximum length.
            ans = Math.max(ans, right - left + 1);
        }

        // Return the longest valid substring length.
        return ans;
    }
}