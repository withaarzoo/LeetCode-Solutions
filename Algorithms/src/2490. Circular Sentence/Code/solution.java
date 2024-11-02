class Solution {
    public boolean isCircularSentence(String sentence) {
        // Step 1: Split the sentence into words
        String[] words = sentence.split(" ");

        // Step 2: Check adjacent pairs and the circular condition
        for (int i = 0; i < words.length; i++) {
            char lastChar = words[i].charAt(words[i].length() - 1);
            char firstChar = words[(i + 1) % words.length].charAt(0);
            if (lastChar != firstChar) {
                return false;
            }
        }

        return true;
    }
}
