class Solution {
    public String shortestBeautifulSubstring(String s, int k) {
        String answer = ""; // I store the best beautiful substring found so far.
        int left = 0; // I keep the left boundary of the sliding window.
        int ones = 0; // I count how many '1' characters are inside the window.

        // I expand the window by moving the right pointer through the string.
        for (int right = 0; right < s.length(); right++) {
            // I update the count when the newly added character is '1'.
            if (s.charAt(right) == '1') {
                ones++;
            }

            // If I have too many ones, I shrink the window from the left
            // until it contains at most k ones again.
            while (ones > k) {
                if (s.charAt(left) == '1') {
                    ones--;
                }
                left++;
            }

            // I remove leading zeros because they are unnecessary and only
            // make a valid substring longer.
            while (ones == k && s.charAt(left) == '0') {
                left++;
            }

            // The current window is beautiful when it contains exactly k ones.
            if (ones == k) {
                // I create the current candidate substring.
                String candidate = s.substring(left, right + 1);

                // I update the answer when this candidate is shorter, or when
                // equal-length candidates need lexicographical comparison.
                if (answer.isEmpty() ||
                        candidate.length() < answer.length() ||
                        (candidate.length() == answer.length() &&
                                candidate.compareTo(answer) < 0)) {
                    answer = candidate;
                }
            }
        }

        // I return the best result, or an empty string if no valid substring exists.
        return answer;
    }
}