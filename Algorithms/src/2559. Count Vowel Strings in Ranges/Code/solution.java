import java.util.*;

class Solution {
    public int[] vowelStrings(String[] words, int[][] queries) {
        Set<Character> vowels = Set.of('a', 'e', 'i', 'o', 'u');
        int n = words.length;
        int[] prefix = new int[n];

        // Precompute the prefix sum
        for (int i = 0; i < n; i++) {
            if (vowels.contains(words[i].charAt(0)) && vowels.contains(words[i].charAt(words[i].length() - 1))) {
                prefix[i] = 1;
            }
            if (i > 0) {
                prefix[i] += prefix[i - 1];
            }
        }

        // Answer the queries
        int[] result = new int[queries.length];
        for (int i = 0; i < queries.length; i++) {
            int l = queries[i][0], r = queries[i][1];
            result[i] = prefix[r] - (l > 0 ? prefix[l - 1] : 0);
        }
        return result;
    }
}