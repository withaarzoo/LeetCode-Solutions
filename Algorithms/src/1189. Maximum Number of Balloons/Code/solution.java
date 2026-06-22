class Solution {
    public int maxNumberOfBalloons(String text) {
        // Store frequency of all lowercase letters
        int[] freq = new int[26];

        // Count each character
        for (char ch : text.toCharArray()) {
            freq[ch - 'a']++;
        }

        // Return the smallest possible complete balloon count
        return Math.min(
                Math.min(freq['b' - 'a'], freq['a' - 'a']),
                Math.min(
                        Math.min(freq['l' - 'a'] / 2, freq['o' - 'a'] / 2),
                        freq['n' - 'a']));
    }
}