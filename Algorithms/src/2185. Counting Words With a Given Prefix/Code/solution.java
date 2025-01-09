class Solution {
    public int prefixCount(String[] words, String pref) {
        int count = 0;
        for (String word : words) {
            // Check if the word starts with the prefix
            if (word.startsWith(pref)) {
                count++;
            }
        }
        return count;
    }
}
