class Solution {
    public boolean canConstruct(String s, int k) {
        if (k > s.length())
            return false; // More palindromes than characters
        int[] freq = new int[26]; // Frequency array for lowercase letters
        for (char c : s.toCharArray()) {
            freq[c - 'a']++;
        }
        int oddCount = 0;
        for (int count : freq) {
            if (count % 2 != 0) {
                oddCount++;
            }
        }
        return oddCount <= k;
    }
}
