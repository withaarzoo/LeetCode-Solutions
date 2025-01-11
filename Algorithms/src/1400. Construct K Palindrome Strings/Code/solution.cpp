class Solution {
public:
    bool canConstruct(string s, int k) {
        if (k > s.length()) return false; // More palindromes than characters
        vector<int> freq(26, 0); // Frequency array for lowercase letters
        for (char c : s) {
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
};
