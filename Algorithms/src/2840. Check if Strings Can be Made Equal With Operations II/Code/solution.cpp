class Solution {
public:
    bool checkStrings(string s1, string s2) {
        // Frequency arrays for even and odd positions
        vector<int> even(26, 0), odd(26, 0);

        for (int i = 0; i < s1.size(); i++) {
            if (i % 2 == 0) {
                // Count characters at even indexes
                even[s1[i] - 'a']++;
                even[s2[i] - 'a']--;
            } else {
                // Count characters at odd indexes
                odd[s1[i] - 'a']++;
                odd[s2[i] - 'a']--;
            }
        }

        // If any frequency is not zero, strings cannot be made equal
        for (int i = 0; i < 26; i++) {
            if (even[i] != 0 || odd[i] != 0) {
                return false;
            }
        }

        return true;
    }
};