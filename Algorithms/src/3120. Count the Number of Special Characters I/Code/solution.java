class Solution {
    public int numberOfSpecialChars(String word) {

        // Store all characters inside a HashSet
        HashSet<Character> set = new HashSet<>();

        // Add every character from the string
        for (char ch : word.toCharArray()) {
            set.add(ch);
        }

        // Variable to store answer
        int count = 0;

        // Check all lowercase English letters
        for (char ch = 'a'; ch <= 'z'; ch++) {

            // Check if both lowercase and uppercase exist
            if (set.contains(ch) && set.contains((char) (ch - 'a' + 'A'))) {
                count++;
            }
        }

        // Return final count
        return count;
    }
}