#include <bits/stdc++.h>
using namespace std;

class Solution
{
public:
    // Function to find the length of the longest common prefix between two arrays of integers
    int longestCommonPrefix(vector<int> &arr1, vector<int> &arr2)
    {
        // Create an unordered map to store prefixes of numbers from arr1.
        unordered_map<string, int> prefixMap;

        // Step 1: Build the prefix map for all numbers in arr1
        // Iterate over each number in arr1
        for (int num : arr1)
        {
            // Convert the number into a string to handle its digits
            string strNum = to_string(num);
            string prefix = ""; // Initialize an empty string to build the prefix

            // Iterate over each character (digit) of the stringified number
            for (char ch : strNum)
            {
                // Append the current character (digit) to the prefix
                prefix += ch;

                // Insert the current prefix into the prefix map and increment its count
                prefixMap[prefix]++;
            }
        }

        // Variable to keep track of the maximum length of common prefix found
        int maxLength = 0;

        // Step 2: Check for common prefixes in arr2
        // Iterate over each number in arr2
        for (int num : arr2)
        {
            // Convert the number into a string to handle its digits
            string strNum = to_string(num);
            string prefix = ""; // Initialize an empty string to build the prefix

            // Iterate over each character (digit) of the stringified number
            for (char ch : strNum)
            {
                // Append the current character (digit) to the prefix
                prefix += ch;

                // Check if this prefix exists in the prefix map (i.e., it's a common prefix)
                if (prefixMap.find(prefix) != prefixMap.end())
                {
                    // Update maxLength to the maximum length of the common prefix found so far
                    maxLength = max(maxLength, static_cast<int>(prefix.length()));
                }
            }
        }

        // Return the length of the longest common prefix found
        return maxLength;
    }
};
