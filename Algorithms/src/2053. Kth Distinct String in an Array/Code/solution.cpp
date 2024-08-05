#include <vector>
#include <string>
#include <unordered_map>

class Solution
{
public:
    string kthDistinct(std::vector<std::string> &arr, int k)
    {
        // Create an unordered map to count occurrences of each string
        std::unordered_map<std::string, int> count;
        // Create a vector to store distinct strings
        std::vector<std::string> distinct;

        // Count occurrences of each string in the array
        for (const std::string &str : arr)
        {
            count[str]++;
        }

        // Collect distinct strings (strings that appear exactly once) in order
        for (const std::string &str : arr)
        {
            if (count[str] == 1)
            {
                distinct.push_back(str);
            }
        }

        // If k is within the range of distinct strings, return the k-th distinct string
        // Otherwise, return an empty string
        if (k <= distinct.size())
        {
            return distinct[k - 1];
        }
        else
        {
            return "";
        }
    }
};
