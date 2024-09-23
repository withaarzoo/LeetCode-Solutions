#include <iostream>
#include <vector>
#include <string>
#include <unordered_set>
using namespace std;

class Solution
{
public:
    // Function to calculate the minimum number of extra characters in the string
    int minExtraChar(string s, vector<string> &dictionary)
    {
        // Convert the vector of dictionary words into an unordered set for O(1) lookup
        unordered_set<string> dict(dictionary.begin(), dictionary.end());

        int n = s.size(); // Length of the input string

        // DP array where dp[i] represents the minimum number of extra characters for the substring s[0:i]
        // Initialized to the maximum number of extra characters possible, which is the length of the string (n)
        vector<int> dp(n + 1, n);

        // Base case: No extra characters are needed for an empty string
        dp[0] = 0;

        // Iterate through each index of the string from 1 to n (i is the end index of a substring)
        for (int i = 1; i <= n; i++)
        {
            // Try all possible substrings ending at position i
            for (int j = 0; j < i; j++)
            {
                // Extract the substring s[j:i] (substring from index j to i-1)
                string sub = s.substr(j, i - j);

                // Check if the substring exists in the dictionary
                if (dict.find(sub) != dict.end())
                {
                    // If the substring is found in the dictionary, update dp[i]
                    // We take the minimum between the current dp[i] and dp[j], since dp[j] represents
                    // the state before adding the valid substring, meaning no extra characters are needed for this substring
                    dp[i] = min(dp[i], dp[j]);
                }
            }

            // If no valid substring was found, consider the current character at index i-1 as an extra character
            // This updates dp[i] by taking dp[i-1] (previous state) + 1 (since the current character is considered extra)
            dp[i] = min(dp[i], dp[i - 1] + 1);
        }

        // The result is stored in dp[n], which represents the minimum number of extra characters for the entire string
        return dp[n];
    }
};
