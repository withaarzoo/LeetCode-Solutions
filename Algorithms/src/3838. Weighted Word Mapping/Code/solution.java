class Solution {
    public String mapWordWeights(String[] words, int[] weights) {
        StringBuilder result = new StringBuilder();

        // Process every word
        for (String word : words) {
            int sumWeight = 0;

            // Add weights of all characters
            for (char ch : word.toCharArray()) {
                sumWeight += weights[ch - 'a'];
            }

            // Take modulo 26
            int value = sumWeight % 26;

            // Convert to reverse alphabetical character
            result.append((char) ('z' - value));
        }

        return result.toString();
    }
}