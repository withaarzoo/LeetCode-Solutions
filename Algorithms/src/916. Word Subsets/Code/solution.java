class Solution {
    public List<String> wordSubsets(String[] words1, String[] words2) {
        int[] maxFreq = new int[26];

        // Precompute the maximum frequency for each character in words2
        for (String word : words2) {
            int[] freq = new int[26];
            for (char c : word.toCharArray())
                freq[c - 'a']++;
            for (int i = 0; i < 26; i++) {
                maxFreq[i] = Math.max(maxFreq[i], freq[i]);
            }
        }

        List<String> result = new ArrayList<>();
        // Check each word in words1
        for (String word : words1) {
            int[] freq = new int[26];
            for (char c : word.toCharArray())
                freq[c - 'a']++;
            boolean isUniversal = true;
            for (int i = 0; i < 26; i++) {
                if (freq[i] < maxFreq[i]) {
                    isUniversal = false;
                    break;
                }
            }
            if (isUniversal)
                result.add(word);
        }

        return result;
    }
}
