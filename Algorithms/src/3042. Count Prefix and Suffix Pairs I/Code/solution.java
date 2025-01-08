class Solution {
    public int countPrefixSuffixPairs(String[] words) {
        int n = words.length;
        int count = 0;

        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                String prefix = words[i];
                String word = words[j];
                int len = prefix.length();

                if (word.startsWith(prefix) && word.endsWith(prefix)) {
                    count++;
                }
            }
        }

        return count;
    }
}
