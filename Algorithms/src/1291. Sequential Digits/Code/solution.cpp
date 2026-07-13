class Solution
{
public:
    vector<int> sequentialDigits(int low, int high)
    {
        // String containing all consecutive digits
        string digits = "123456789";

        // Result array
        vector<int> ans;

        // Number of digits in low and high
        int minLen = to_string(low).size();
        int maxLen = to_string(high).size();

        // Try every possible length
        for (int len = minLen; len <= maxLen; len++)
        {

            // Generate every substring of current length
            for (int start = 0; start + len <= 9; start++)
            {

                // Convert substring into an integer
                int num = stoi(digits.substr(start, len));

                // Keep only numbers inside the range
                if (num >= low && num <= high)
                    ans.push_back(num);
            }
        }

        return ans;
    }
};