class Solution {
    public String lexGreaterPermutation(String s, String target) {
        // Store the frequency of every character available in s.
        int[] count = new int[26];
        for (char ch : s.toCharArray()) {
            count[ch - 'a']++;
        }

        int n = s.length();
        int matched = 0;

        // Try to keep the answer exactly equal to target for as long as possible.
        while (matched < n && count[target.charAt(matched) - 'a'] > 0) {
            // Use target[matched] because matching it keeps the prefix smallest.
            count[target.charAt(matched) - 'a']--;
            matched++;
        }

        // Start where matching failed, or at the last position if all matched.
        int start = matched < n ? matched : n - 1;

        // Move backward because changing a later position gives a smaller answer.
        for (int i = start; i >= 0; i--) {
            // Restore this character if it was previously used to match target.
            if (i < matched) {
                count[target.charAt(i) - 'a']++;
            }

            // Find the smallest available character greater than target[i].
            int bigger = -1;
            for (int ch = target.charAt(i) - 'a' + 1; ch < 26; ch++) {
                if (count[ch] > 0) {
                    bigger = ch;
                    break;
                }
            }

            // Build the answer when an increasing character is found.
            if (bigger != -1) {
                // Consume the character used to make the answer strictly greater.
                count[bigger]--;

                // Keep the prefix unchanged.
                StringBuilder answer = new StringBuilder(target.substring(0, i));

                // Add the smallest possible character that is greater here.
                answer.append((char) ('a' + bigger));

                // Add every remaining character in sorted order.
                for (int ch = 0; ch < 26; ch++) {
                    while (count[ch]-- > 0) {
                        answer.append((char) ('a' + ch));
                    }
                }

                return answer.toString();
            }
        }

        // No permutation is strictly greater than target.
        return "";
    }
}