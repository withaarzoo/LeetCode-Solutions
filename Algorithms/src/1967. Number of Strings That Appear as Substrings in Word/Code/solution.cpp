class Solution {
public:
    int numOfStrings(vector<string>& patterns, string word) {
        // Store the number of matching patterns
        int count = 0;

        // Check every pattern one by one
        for (string &pattern : patterns) {
            // If the pattern exists inside word, increase the answer
            if (word.find(pattern) != string::npos) {
                count++;
            }
        }

        // Return the total number of matching patterns
        return count;
    }
};