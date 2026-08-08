class Solution {
    public int[] validSequence(String word1, String word2) {
        int n = word1.length(); // I store the length of word1 for the two scans.
        int m = word2.length(); // I store the length of word2 because the answer has m indices.

        int[] last = new int[m]; // I store a usable position for every suffix of word2.

        // I initialize every position as -1 to mean that no matching position was found.
        java.util.Arrays.fill(last, -1);

        int i = n - 1; // I start from the end of word1.
        int j = m - 1; // I start from the end of word2.

        // I greedily match word2 from right to left.
        // This gives me the positions needed to safely match the remaining suffix.
        while (i >= 0 && j >= 0) {
            // If the current characters match, this position can represent word2[j].
            if (word1.charAt(i) == word2.charAt(j)) {
                last[j] = i; // I remember the position for this suffix.
                --j;         // I now search for the previous character of word2.
            }

            --i; // I continue moving left through word1.
        }

        int[] ans = new int[m]; // I allocate space for the final sequence.
        int size = 0;           // I track how many indices I have inserted.

        boolean canSkip = true; // I have not used the one allowed mismatch yet.
        j = 0;                  // I start matching word2 from the beginning.

        // I scan word1 from left to right to get the lexicographically smallest indices.
        for (i = 0; i < n && j < m; ++i) {
            // A matching character can always be selected without using the mismatch.
            if (word1.charAt(i) == word2.charAt(j)) {
                ans[size++] = i; // I take the earliest valid index.
                ++j;             // I move to the next character of word2.
            }
            // I can use a mismatch only when the remaining suffix can still be matched.
            else if (canSkip &&
                     (j == m - 1 || i < last[j + 1])) {
                canSkip = false; // I spend the one allowed mismatch.
                ans[size++] = i; // I take this index as the mismatching character.
                ++j;             // I move to the next character of word2.
            }
        }

        // If every character of word2 was matched, I return the filled answer.
        if (j == m) {
            return ans;
        }

        // Otherwise, no valid sequence exists.
        return new int[0];
    }
}